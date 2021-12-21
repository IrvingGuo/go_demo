package api

import (
	"resource-plan-improvement/entity"
	"strconv"

	"github.com/gin-gonic/gin"
)

func SaveSubprogram(c *gin.Context) {
	var err error
	var subprogram entity.Subprogram
	if err = c.BindJSON(&subprogram); err != nil {
		FailWithMsg(c, err.Error())
		return
	}
	if err = subprogram.Save(); err != nil {
		FailWithMsg(c, err.Error())
		return
	}
	OkWithData(c, subprogram)
}

func GetAllSubprograms(c *gin.Context) {
	if subprograms, err := entity.FindAllSubprograms(); err != nil {
		AbortFindFailed(c, err)
	} else {
		OkWithData(c, subprograms)
	}
}

func DeleteSubprogramById(c *gin.Context) {
	var subprogramId int
	var err error
	if subprogramId, err = strconv.Atoi(c.Param("id")); err != nil {
		FailWithErr(c, err)
		return
	}
	if err = entity.DeleteSubprogramById(uint(subprogramId)); err != nil {
		FailWithErr(c, err)
		return
	}
	Ok(c)
}
