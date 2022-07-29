package model

import (
	"gorm.io/gorm"
	"strconv"
)

type Driver struct {
	Model            `gorm:"embedded"`
	Name             string `gorm:"not null" validate:"required" json:"name"`
	Surname          string `gorm:"not null" validate:"required" json:"surname"`
	Age              int    `json:"age"`
	PhoneNumber      string `gorm:"not null" validate:"required" json:"phone_number"`
	LicenceSerialNum string `gorm:"not null" validate:"required" json:"licence_serial_num"`
	Car              []Car  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"cars"`
	TaxiStopID       uint   `json:"taxiStopID"`
}

func getQueryToIncreaseNumOfDriversBy(val int) string {
	return "UPDATE TaxiStop SET NumberOfDrivers = NumberOfDrivers + " + strconv.Itoa(val) + " WHERE id = ?"
}

func (d *Driver) AfterCreate(tx *gorm.DB) (err error) {
	tx.Exec(getQueryToIncreaseNumOfDriversBy(1), d.TaxiStopID)

	return nil
}

func (d *Driver) BeforeUpdate(tx *gorm.DB) (err error) {
	if tx.Statement.Changed("TaxiStopID") {
		tx.Exec(getQueryToIncreaseNumOfDriversBy(-1), d.TaxiStopID)
	}

	return nil
}

func (d *Driver) AfterUpdate(tx *gorm.DB) (err error) {
	if tx.Statement.Changed("TaxiStopID") {
		tx.Exec(getQueryToIncreaseNumOfDriversBy(1), d.TaxiStopID)
	}
	return nil
}
