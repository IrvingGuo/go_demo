package entity

type Subprogram struct {
	Base
	Name   string `json:"name" orm:"comment:sub program name"`
	Status bool   `json:"status" orm:"comment:sub program status"`
}

func (subprogram *Subprogram) Save() error {
	return db.Save(subprogram).Error
}

func FindAllSubprograms() (subprograms []Subprogram, err error) {
	err = db.Find(&subprograms).Error
	return
}

func DeleteSubprogramById(subprogramId uint) error {
	return db.Delete(Subprogram{}, "id = ?", subprogramId).Error
}
