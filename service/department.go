package service

import (
	"fmt"
	"resource-plan-improvement/entity"
)

// current departments and all relative departments including parent and child
func GetRelativeDepartmentsUnderUser(userId uint, includeParent bool) (departments []entity.Department, err error) {
	// use map to prune duplicates
	deptMap := make(map[uint](entity.Department))
	// get direct departments whose user_id is userId
	var directDeptList []entity.Department
	if directDeptList, err = entity.FindDeptByUserId(userId); err != nil {
		return
	}
	// get indirect departments from direct departments, i.e, get its parent and child dept
	for _, directDept := range directDeptList {
		deptMap[directDept.Id] = directDept
		var tmpList []entity.Department
		// get child dept
		if tmpList, err = entity.FindChildDeptByLevel(fmt.Sprintf("%s.%d", directDept.Level, directDept.Id)); err != nil {
			return
		} else {
			for _, tmpDept := range tmpList {
				deptMap[tmpDept.Id] = tmpDept
			}
		}
		// if includeParent, get parent dept
		if includeParent {
			if tmpList, err = entity.FindParentDeptByLevel(directDept.Level); err != nil {
				return
			} else {
				for _, tmpDept := range tmpList {
					deptMap[tmpDept.Id] = tmpDept
				}
			}
		}
	}
	// get value list
	for _, v := range deptMap {
		departments = append(departments, v)
	}
	return
}
