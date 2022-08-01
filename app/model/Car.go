package model

type Car struct {
	Model            `gorm:"embedded"`
	CarModelName     string `gorm:"not null" validate:"required" json:"car_model_name"`
	PlateNumber      string `gorm:"not null" validate:"required" json:"plate_number"`
	Color            string `gorm:"not null" validate:"required" json:"color"`
	YearManufactured int    `gorm:"not null" validate:"required" json:"year_manufactured"`
	DriverID         uint   `json:"driver_id"`
}

func (c *Car) Create(car Car) (Car, error) {
	if err := db.Create(&car).Error; err != nil {
		return car, err
	}
	return car, nil
}

func (c *Car) GetAll() ([]Car, error) {
	var cars []Car
	if err := db.Where("deleted = ?", false).Find(&cars).Error; err != nil {
		return nil, err
	}

	return cars, nil
}

func (c *Car) FindBy(id uint) (Car, error) {
	var car Car
	if err := db.Where("deleted = ?", false).First(&car, id).Error; err != nil {
		return car, err
	}

	return car, nil
}

func (c *Car) Update(car Car) error {
	if err := db.Model(&car).Updates(car).Error; err != nil {
		return err
	}
	return nil
}

// DeleteBy soft delete
func (c *Car) DeleteBy(id uint) error {
	if err := db.Model(&Car{}).Where("id = ?", id).Update("deleted", true).Error; err != nil {
		return err
	}
	return nil
}
