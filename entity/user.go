package entity

import (
	"gorm.io/datatypes"
)

type User struct {
	Base
	Cn             string         `json:"cn" orm:"comment:e.g. Chen Ruonan"`
	Title          string         `json:"title" orm:"comment:e.g. Engineer,Grey Services"`
	DeptId         uint           `json:"deptId" orm:"comment:department id"`
	SAMAccountName string         `json:"sAMAccount" orm:"comment:chen_r1"`
	EntryDate      datatypes.Date `json:"entryDate" orm:"comment:chen_r1"`
	ResignDate     datatypes.Date `json:"resignDate" orm:"comment:chen_r1"`
	Gender         bool           `json:"gender" orm:"comment:gender"`
	Location       string         `json:"location" orm:"comment:location"`
	ResourceType   string         `json:"resourceType" orm:"comment:internal or external"`
	Privilege      uint           `json:"privilege" orm:"comment:Group Leader or Technical Product Manager"`
	Status         uint           `json:"status" orm:"comment:User Status"`
	Token          string         `json:"token" gorm:"-"`
}

// Privilege
const (
	PRIVILEGE_DEPARTMENT uint = 1 << iota
	PRIVILEGE_TPM
	ADMIN
)

// Status
const (
	ON_THE_JOB uint = iota
	RESIGN
)

func (user *User) Save() error {
	return db.Save(user).Error
}

func FindAllUsers() (users []User, err error) {
	err = db.Find(&users).Error
	return
}

func FindUserById(id uint) (user User, err error) {
	err = db.First(&user, "id = ?", id).Error
	return
}

func FindUserByDistinguishedName(distinguishedName string) (user User, err error) {
	err = db.First(&user, "distinguished_name = ?", distinguishedName).Error
	return
}

func FindUserBySAMAccount(sAMAccountName string) (user User, err error) {
	err = db.First(&user, "sam_account_name = ?", sAMAccountName).Error
	return
}

func FindUsersByDeptIds(deptIds []uint) (users []User, err error) {
	err = db.Find(&users, "dept_id in ?", deptIds).Error
	return
}

func DeleteUserById(userId uint) error {
	return db.Delete(User{}, "id = ?", userId).Error
}
