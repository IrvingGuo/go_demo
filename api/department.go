package api

import (
	"resource-plan-improvement/config"
	"resource-plan-improvement/entity"
	"resource-plan-improvement/service"

	"github.com/gin-gonic/gin"
)

func SaveDepartment(c *gin.Context) {
	var err error
	var dept entity.Department
	if err = c.BindJSON(&dept); err != nil {
		FailWithErr(c, err)
		return
	}
	if err = dept.Save(); err != nil {
		FailWithErr(c, err)
		return
	}
	OkWithData(c, dept)
}

func GetAllDepartments(c *gin.Context) {
	if departments, err := entity.FindAllDepartments(); err != nil {
		AbortFindFailed(c, err)
	} else {
		OkWithData(c, departments)
	}
}

func GetDepartmentById(c *gin.Context) {

}

func DeleteDepartmentById(c *gin.Context) {

}

func GetDepartmentsUnderLoginUser(c *gin.Context) {
	userId := c.GetUint(config.CTX_KEY_USER_ID)
	var err error
	var departments []entity.Department
	if departments, err = service.GetRelativeDepartmentsUnderUser(uint(userId), true); err != nil {
		FailWithMsg(c, err.Error())
		return
	}
	OkWithData(c, departments)
}
