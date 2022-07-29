package model

import (
	"TaxiStop/app/database"
	"gorm.io/gorm"
	"time"
)

var db *gorm.DB = database.InitDBConnection()

type Model struct {
	ID        uint      `gorm:"primary_key;auto_increment;unique;not null;" json:"id"`
	Deleted   bool      `gorm:"not null"  json:"isDeleted"`
	CreatedAt time.Time `gorm:"not null"  json:"created_at"`
	UpdatedAt time.Time `gorm:"not null"  json:"updated_at"`
	DeletedAt time.Time `gorm:"not null"  json:"deleted_at"`
}

//type crud interface {
//	Create(uint) (interface{}, error)
//	FindBy(interface{}) (interface{}, error)
//	GetAll() (interface{}, error)
//	Update(interface{}) (interface{}, error)
//	Delete(interface{}) (interface{}, error)
//}

var models = []interface{}{
	&User{},
	&Car{},
	&Driver{},
	&TaxiStop{},
}

func GetModels() []interface{} {
	return models
}

func CreateTables() {
	for _, model := range GetModels() {
		err := db.AutoMigrate(model)

		if err != nil {
			panic(err)
		}
	}
}
