package entity

import (
	"gorm.io/datatypes"
)

type Record struct {
	Base
	Table         string         `json:"userId" orm:"comment:table name"`
	RecordId      uint           `json:"recordId" orm:"comment:record id"`
	Field         string         `json:"field" orm:"comment:fields separated by comma"`
	Before        string         `json:"before" orm:"comment:old values separated by comma"`
	After         string         `json:"after" orm:"comment:new values separated by comma"`
	Status        uint           `json:"status" orm:"comment:record status"`
	OperatorId    uint           `json:"operatorId" orm:"comment:user id"`
	OperationTime datatypes.Date `json:"operationTime" orm:"comment:operation time"`
}

// Record Status
const (
	CREATED uint = iota
	UPDATED
	DELETED
)
