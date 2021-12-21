package entity

import (
	"gorm.io/datatypes"
)

type Program struct {
	Base
	Code        string         `json:"code" orm:"comment:program code"`
	Type        string         `json:"type" orm:"comment:program type"`
	Name        string         `json:"name" orm:"comment:program name"`
	UserId      uint           `json:"userId" orm:"comment:user id aka tpm id"`
	CloseDate   datatypes.Date `json:"closeDate" orm:"comment:program close date"`
	Status      uint           `json:"status" orm:"comment:program status"`
	Subprograms string         `json:"subprograms" orm:"comment:subprogram id list joined with ','"`
	Activities  string         `json:"activities" orm:"comment:activity id list joined with ','"`
}

// Program Status
const (
	ACTIVE uint = iota
	CLOSE
	SUSPENDING
)

func (program *Program) Save() error {
	return db.Save(program).Error
}

func (program *Program) Delete() error {
	return db.Delete(program).Error
}

func FindAllPrograms() (programs []Program, err error) {
	db.Find(&programs)
	return
}

func FindProgramById(id uint) (program Program, err error) {
	err = db.First(&program, "id = ?", id).Error
	return
}

func FindProgramByUserId(userId uint) (programs []Program, err error) {
	err = db.Find(&programs, "user_id = ?", userId).Error
	return
}

func DeleteProgramById(id uint) error {
	return db.Where("id = ?", id).Delete(&Program{}).Error
}
