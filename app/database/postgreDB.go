package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "123456"
	dbname   = "TaxiStop"
)

func InitDBConnection() *gorm.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	gormDB, err := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})
	sqlDB, err := gormDB.DB()

	if err != nil {
		panic(err)
	}
	//defer sqlDB.Close()

	err = sqlDB.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")

	return gormDB
}
