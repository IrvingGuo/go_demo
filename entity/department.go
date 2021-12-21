package entity

import (
	"resource-plan-improvement/util"
	"strings"
)

type Department struct {
	Base
	Name     string `json:"name" orm:"comment:department name"`
	UserId   uint   `json:"userId" orm:"comment:user id"`
	Level    string `json:"level" orm:"comment:department level structure"`
	ParentId uint   `json:"parentId" orm:"comment:parent department id"`
}

func (dept *Department) Save() error {
	return db.Save(dept).Error
}

func FindAllDepartments() (departments []Department, err error) {
	err = db.Find(&departments).Error
	return
}

func FindDeptById(deptId uint) (dept Department, err error) {
	err = db.First(&dept, "id = ?", deptId).Error
	return
}

func FindDeptByUserId(userId uint) (deptList []Department, err error) {
	err = db.Find(&deptList, "user_id = ?", userId).Error
	return
}

func FindChildDeptByLevel(level string) (deptList []Department, err error) {
	err = db.Find(&deptList, "level like ?", level+"%").Error
	return
}

func FindParentDeptByLevel(level string) (deptList []Department, err error) {
	strIds := strings.Split(level, ".")[1:]
	var ids []uint
	if ids, err = util.ConvertStringArrToUintArr(strIds); err != nil {
		return
	}
	var dept Department
	for _, id := range ids {
		if dept, err = FindDeptById(id); err != nil {
			return
		} else {
			deptList = append(deptList, dept)
		}
	}
	return
}
