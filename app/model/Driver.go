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
	return "UPDATE taxi_stops SET number_of_drivers = number_of_drivers " + "+ " + strconv.Itoa(val) + " WHERE id = ?"
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

func (d *Driver) Create(driver Driver) (Driver, error) {
	if err := db.Create(&driver).Error; err != nil {
		return driver, err
	}
	return driver, nil
}

func (d *Driver) GetAll() ([]Driver, error) {
	var drivers []Driver
	if err := db.Where("deleted = ?", false).Find(&drivers).Error; err != nil {
		return nil, err
	}

	return drivers, nil
}

func (d *Driver) FindBy(id uint) (Driver, error) {
	var driver Driver
	if err := db.Where("deleted = ?", false).First(&driver, id).Error; err != nil {
		return driver, err
	}

	return driver, nil
}

func (d *Driver) Update(driver Driver) error {
	if err := db.Model(&driver).Updates(driver).Error; err != nil {
		return err
	}
	return nil
}

// DeleteBy soft delete
func (d *Driver) DeleteBy(id uint) error {
	if err := db.Model(&Driver{}).Where("id = ?", id).Update("deleted", true).Error; err != nil {
		return err
	}
	return nil
}
