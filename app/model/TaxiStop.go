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

func (ts *TaxiStop) Create(taxiStop TaxiStop) (TaxiStop, error) {
	if err := db.Create(&taxiStop).Error; err != nil {
		return taxiStop, err
	}
	return taxiStop, nil
}

func (ts *TaxiStop) GetAll() ([]TaxiStop, error) {
	var taxiStops []TaxiStop
	if err := db.Where("deleted = ?", false).Find(&taxiStops).Error; err != nil {
		return nil, err
	}

	return taxiStops, nil
}

func (ts *TaxiStop) FindBy(id uint) (TaxiStop, error) {
	var taxiStop TaxiStop
	if err := db.Where("deleted = ?", false).First(&taxiStop, id).Error; err != nil {
		return taxiStop, err
	}

	return taxiStop, nil
}

// ????
func (ts *TaxiStop) FindByUser() ([]TaxiStop, error) {
	var taxiStops []TaxiStop
	if err := db.Model(&TaxiStop{}).Preload("User").Find(&taxiStops).Error; err != nil {
		return nil, err
	}

	return taxiStops, nil
}

func (ts *TaxiStop) Update(taxiStop TaxiStop) error {
	if err := db.Model(&taxiStop).Updates(taxiStop).Error; err != nil {
		return err
	}
	return nil
}

// DeleteBy soft delete
func (ts *TaxiStop) DeleteBy(id uint) error {
	if err := db.Model(&TaxiStop{}).Where("id = ?", id).Update("deleted", true).Error; err != nil {
		return err
	}
	return nil
}
