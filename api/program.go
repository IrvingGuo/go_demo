package api

import (
	"resource-plan-improvement/config"
	"resource-plan-improvement/entity"
	"resource-plan-improvement/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

func SaveProgram(c *gin.Context) {
	var err error
	var program entity.Program
	if err = c.BindJSON(&program); err != nil {
		FailWithErr(c, err)
		return
	}
	if err = program.Save(); err != nil {
		FailWithErr(c, err)
		return
	}
	OkWithData(c, program)
}

func GetAllPrograms(c *gin.Context) {
	if programs, err := entity.FindAllPrograms(); err != nil {
		AbortFindFailed(c, err)
	} else {
		OkWithData(c, programs)
	}
}

func GetProgramById(c *gin.Context) {
	var programId int
	var err error
	if programId, err = strconv.Atoi(c.Param("id")); err != nil {
		AbortBadRequest(c, err)
		return
	}
	var program entity.Program
	if program, err = entity.FindProgramById(uint(programId)); err != nil {
		AbortFindFailed(c, err)
		return
	}
	OkWithData(c, program)
}

func DeleteProgramById(c *gin.Context) {
	var programId int
	var err error
	if programId, err = strconv.Atoi(c.Param("id")); err != nil {
		FailWithErr(c, err)
		return
	}

	if err = service.DeleteProgram(uint(programId)); err != nil {
		FailWithErr(c, err)
		return
	}
	Ok(c)
}

func GetProgramsUnderUser(c *gin.Context) {
	userId := c.GetUint(config.CTX_KEY_USER_ID)
	var err error
	var programs []entity.Program
	if programs, err = entity.FindProgramByUserId(uint(userId)); err != nil {
		FailWithMsg(c, err.Error())
		return
	}
	OkWithData(c, programs)
}
