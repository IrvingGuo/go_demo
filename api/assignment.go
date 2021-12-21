package api

import (
	"resource-plan-improvement/config"
	"resource-plan-improvement/entity"
	"resource-plan-improvement/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

func SaveAssignments(c *gin.Context) {
	var err error
	var assignments []entity.Assignment
	if err = c.BindJSON(&assignments); err != nil {
		FailWithMsg(c, err.Error())
		return
	}
	if err = entity.SaveInBatches(assignments); err != nil {
		FailWithMsg(c, err.Error())
		return
	}
	OkWithData(c, assignments)
}

func UpdateAssignments(c *gin.Context) {
	var err error
	var assignments []entity.Assignment
	if err = c.BindJSON(&assignments); err != nil {
		FailWithMsg(c, err.Error())
		return
	}
	if err = entity.UpdateAssignmentsInBatches(assignments); err != nil {
		FailWithMsg(c, err.Error())
		return
	}
	OkWithData(c, assignments)
}

func UpdateAssignmentsStatus(c *gin.Context) {
	var err error
	var statusPayloads []entity.StatusPayload
	if err = c.BindJSON(&statusPayloads); err != nil {
		FailWithMsg(c, err.Error())
		return
	}
	if err = entity.UpdateAssignmentsStatus(statusPayloads); err != nil {
		FailWithMsg(c, err.Error())
		return
	}
	OkWithData(c, statusPayloads)
}

func GetAllAssignments(c *gin.Context) {
	if assignments, err := entity.FindAllAssignments(); err != nil {
		FailWithMsg(c, err.Error())
	} else {
		OkWithData(c, assignments)
	}
}

func GetAssignmentById(c *gin.Context) {
	var assignmentId int
	var err error
	if assignmentId, err = strconv.Atoi(c.Param("id")); err != nil {
		FailWithMsg(c, err.Error())
		return
	}
	var assignment entity.Assignment
	if assignment, err = entity.FindAssignmentById(uint(assignmentId)); err != nil {
		FailWithMsg(c, err.Error())
		return
	}
	OkWithData(c, assignment)
}

func DeleteAssignmentsByIds(c *gin.Context) {
	var assignments []entity.Assignment
	var err error
	if err = c.BindJSON(&assignments); err != nil {
		FailWithMsg(c, err.Error())
		return
	}
	entity.PrintAssignments(assignments)
	// todo: use generic
	var assignmentIds = make([]uint, len(assignments))
	for _, assignment := range assignments {
		if assignment.Id <= 0 {
			FailWithMsg(c, entity.ErrUpdateId.Error())
			return
		}
		assignmentIds = append(assignmentIds, assignment.Id)
	}
	if err = entity.DeleteAssignmentsByIds(assignmentIds); err != nil {
		FailWithMsg(c, err.Error())
		return
	}
	Ok(c)
}

func GetAssignmentsUnderLeader(c *gin.Context) {
	userId := c.GetUint(config.CTX_KEY_USER_ID)
	if assignments, err := service.GetAssignmentsUnderLeader(userId); err != nil {
		FailWithMsg(c, err.Error())
	} else {
		OkWithData(c, assignments)
	}
}

func GetAssignmentsUnderTpm(c *gin.Context) {
	userId := c.GetUint(config.CTX_KEY_USER_ID)
	if assignmentsWithCn, err := service.GetAssignmentsUnderTpm(userId); err != nil {
		FailWithMsg(c, err.Error())
	} else {
		OkWithData(c, assignmentsWithCn)
	}
}

func GetAssignmentsByProgramId(c *gin.Context) {
	var err error
	var programId int
	if programId, err = strconv.Atoi(c.Param("id")); err != nil {
		FailWithMsg(c, err.Error())
		return
	}
	if assignmentsWithCn, err := entity.FindAssignmentsWithCnByProgramIds([]uint{uint(programId)}); err != nil {
		FailWithMsg(c, err.Error())
	} else {
		OkWithData(c, assignmentsWithCn)
	}
}

func GetMasterResPlanExcelFilename(c *gin.Context) {
	var filename string
	var err error
	if filename, err = service.GenExcel(); err != nil {
		FailWithMsg(c, err.Error())
		return
	}
	OkWithData(c, filename)
}
