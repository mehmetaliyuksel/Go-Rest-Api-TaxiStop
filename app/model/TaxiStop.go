package model

type TaxiStop struct {
	Model              `gorm:"embedded"`
	Name               string   `gorm:"not null" validate:"required" json:"name"`
	YearBuilt          int      `gorm:"not null" validate:"required" json:"YearBuilt"`
	PhoneNumber        string   `gorm:"not null" validate:"required" json:"PhoneNumber"`
	Address            string   `gorm:"not null" validate:"required" json:"address"`
	TaxRegistrationNum string   `gorm:"not null" validate:"required" json:"TaxRegistrationNum"`
	NumberOfDrivers    int      `gorm:"not null" validate:"required" json:"NumberOfDrivers"`
	Driver             []Driver `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"Drivers"`
	User               []User   `gorm:"many2many:TaxiStop_User;" json:"Users"`
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
func (ts *TaxiStop) FindAssociatedUsers(taxiStop TaxiStop) ([]User, error) {
	var users []User
	if err := db.Model(&taxiStop).Association("User").Find(&users); err != nil {
		return nil, err
	}

	return users, nil
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
