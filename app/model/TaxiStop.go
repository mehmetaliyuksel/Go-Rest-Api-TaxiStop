package model

type TaxiStop struct {
	Model            `gorm:"embedded"`
	Name             string   `gorm:"not null" validate:"required" json:"name"`
	YearBuilt        int      `gorm:"not null" validate:"required" json:"year_built"`
	PhoneNumber      string   `gorm:"not null" validate:"required" json:"phone_number"`
	Address          string   `gorm:"not null" validate:"required" json:"address"`
	LicenceSerialNum string   `gorm:"not null" validate:"required" json:"licence_serial_num"`
	NumberOfDrivers  int      `gorm:"not null" validate:"required" json:"number_of_drivers"`
	Driver           []Driver `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"drivers"`
	User             []User   `gorm:"many2many:TaxiStop_User;" json:"users"`
}
