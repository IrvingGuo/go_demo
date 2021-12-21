package service

import "resource-plan-improvement/entity"

func GetAssignmentsUnderLeader(leaderId uint) (assignments []entity.Assignment, err error) {
	var departments []entity.Department
	if departments, err = GetRelativeDepartmentsUnderUser(leaderId, false); err != nil {
		return
	}
	var deptIds = make([]uint, len(departments))
	for _, dept := range departments {
		deptIds = append(deptIds, dept.Id)
	}
	return entity.FindAssignmentsByDeptIds(deptIds)
}

func GetAssignmentsUnderTpm(tpmId uint) (assignmentWithCn []entity.AssignmentWithCn, err error) {
	var programs []entity.Program
	if programs, err = entity.FindProgramByUserId(tpmId); err != nil {
		return
	}
	var programIds = make([]uint, len(programs))
	for _, program := range programs {
		programIds = append(programIds, program.Id)
	}
	return entity.FindAssignmentsWithCnByProgramIds(programIds)
}
