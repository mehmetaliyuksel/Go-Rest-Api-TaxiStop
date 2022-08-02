package model

import "golang.org/x/crypto/bcrypt"

type User struct {
	Model    `gorm:"embedded"`
	Username string `gorm:"not null" validate:"required" json:"username"`
	Name     string `gorm:"not null" validate:"required" json:"name"`
	Surname  string `gorm:"not null" validate:"required" json:"surname"`
	Age      int    `json:"age"`
	Email    string `gorm:"not null" validate:"required" json:"email"`
	Password string `gorm:"not null" validate:"required" json:"password"`
}

func (u *User) Create(user User) (User, error) {
	if err := db.Create(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (u *User) GetAll() ([]User, error) {
	var users []User
	if err := db.Where("deleted = ?", false).Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func (u *User) FindById(id uint) (User, error) {
	var user User
	if err := db.Where("deleted = ?", false).First(&user, id).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (u *User) FindByEmail(email string) (User, error) {
	var user User
	if err := db.Where("deleted = ? and email = ?", false, email).First(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

// ????
func (u *User) FindByTaxiStop() ([]User, error) {
	var users []User
	if err := db.Model(&User{}).Preload("TaxiStop").Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func (u *User) Update(user User) error {
	if err := db.Model(&user).Updates(user).Error; err != nil {
		return err
	}
	return nil
}

// DeleteBy soft delete
func (u *User) DeleteBy(id uint) error {
	if err := db.Model(&User{}).Where("id = ?", id).Update("deleted", true).Error; err != nil {
		return err
	}
	return nil
}

func (u *User) HashPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)

	if err != nil {
		return err
	}

	u.Password = string(bytes)
	return nil
}
func (u *User) CheckPassword(providedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(providedPassword))

	if err != nil {
		return err
	}
	return nil
}

func (u *User) IsExist(email string) (bool, error) {
	var user User

	if err := db.Where("email = ?", email).First(&user).Error; err != nil {
		return false, err
	}

	return true, nil
}
