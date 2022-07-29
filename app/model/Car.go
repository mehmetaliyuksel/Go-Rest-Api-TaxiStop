package model

type Car struct {
	Model            `gorm:"embedded"`
	CarModelName     string `gorm:"not null" validate:"required" json:"car_model_name"`
	PlateNumber      string `gorm:"not null" validate:"required" json:"plate_number"`
	Color            string `gorm:"not null" validate:"required" json:"color"`
	YearManufactured int    `gorm:"not null" validate:"required" json:"year_manufactured"`
	DriverID         uint   `json:"driver_id"`
}
