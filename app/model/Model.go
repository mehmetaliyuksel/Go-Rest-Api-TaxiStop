package model

import (
	"TaxiStop/app/database"
	"gorm.io/gorm"
	"reflect"
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

func IsExist(entity interface{}) bool {
	var exists bool

	identifierParam := getIdentifierParam(entity)
	identifierParamValue := getValueOfField(entity, identifierParam)

	err := db.Model(&entity).
		Select("count(*) > 0").
		Where(identifierParam+" = ?", identifierParamValue).
		Find(&exists).
		Error

	if err != nil {
		panic(err)
	}

	return exists

}

func getValueOfField(v interface{}, field string) string {
	r := reflect.ValueOf(v).FieldByName(field)
	return r.String()
}

func getIdentifierParam(entity interface{}) string {
	var identifierParam string

	switch reflect.TypeOf(entity).String() {
	case "model.User":
		identifierParam = "Email"
	case "model.TaxiStop":
		identifierParam = "TaxRegistrationNum"
	case "model.Driver":
		identifierParam = "LicenceSerialNum"
	case "model.Car":
		identifierParam = "PlateNumber"
	}

	return identifierParam
}
