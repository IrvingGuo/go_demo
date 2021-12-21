package entity

import (
	"errors"
	"resource-plan-improvement/config"
	"time"
)

type Base struct {
	Id uint `orm:"primarykey" json:"id" gorm:"primaryKey;autoIncrement:true"`
}

var (
	db  = config.Db
	log = config.Logger
)

var (
	ErrUpdateId = errors.New("id in updated or deleted entity must be greater than 0")
)

var entities = map[string]interface{}{
	"user":       &User{},
	"program":    &Program{},
	"subprogram": &Subprogram{},
	"activity":   &Activity{},
	"assignment": &Assignment{},
	"record":     &Record{},
	"department": &Department{},
}

func migrate() {
	for _, entity := range entities {
		if err := db.AutoMigrate(entity); err != nil {
			log.Warnf("Fail to migrate table %d, try again, err: %d", entity, err)
			time.Sleep(time.Second)
			if err := db.AutoMigrate(entity); err != nil {
				panic(err)
			}
		}
	}
}

func init() {
	migrate()
}
