package service

import "resource-plan-improvement/entity"

// delete user and relative assignments
func DeleteProgram(progId uint) (err error) {
	if err = entity.DeleteProgramById(uint(progId)); err != nil {
		return
	}
	if err = entity.DeleteAssignmentsByProgId(progId); err != nil {
		return
	}
	return
}
