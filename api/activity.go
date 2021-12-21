package api

import (
	"resource-plan-improvement/entity"
	"strconv"

	"github.com/gin-gonic/gin"
)

func SaveActivity(c *gin.Context) {
	var err error
	var activity entity.Activity
	if err = c.BindJSON(&activity); err != nil {
		FailWithMsg(c, err.Error())
		return
	}
	if err = activity.Save(); err != nil {
		FailWithMsg(c, err.Error())
		return
	}
	OkWithData(c, activity)
}

func GetAllActivities(c *gin.Context) {
	if activities, err := entity.FindAllActivities(); err != nil {
		AbortFindFailed(c, err)
	} else {
		OkWithData(c, activities)
	}
}

func DeleteActivityById(c *gin.Context) {
	var activityId int
	var err error
	if activityId, err = strconv.Atoi(c.Param("id")); err != nil {
		FailWithErr(c, err)
		return
	}
	if err = entity.DeleteActivityById(uint(activityId)); err != nil {
		FailWithErr(c, err)
		return
	}
	Ok(c)
}
