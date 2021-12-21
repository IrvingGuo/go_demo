package service

import (
	"resource-plan-improvement/entity"
)

// divsion leader can get all subordinates from all relatives departments and groups
func GetUsersUnderLeader(leaderId uint) (users []entity.User, err error) {
	var deptList []entity.Department
	if deptList, err = GetRelativeDepartmentsUnderUser(leaderId, false); err != nil {
		return
	}
	deptIds := make([]uint, len(deptList))
	for i, dept := range deptList {
		deptIds[i] = dept.Id
	}
	return entity.FindUsersByDeptIds(deptIds)
}

// delete user and relative assignments
func DeleteUser(userId uint) (err error) {
	if err = entity.DeleteUserById(uint(userId)); err != nil {
		return
	}
	if err = entity.DeleteAssignmentsByUserId(userId); err != nil {
		return
	}
	return
}
