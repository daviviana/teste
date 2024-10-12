package models

import (
	"log"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	ID       uint    `gorm:"primaryKey" json:"id"`
	Name     string  `gorm:"size:255" json:"name"`
	Email    string  `json:"email"`
	Password string  `json:"-"`
	Address  Address `gorm:"embedded" json:"address"`
}

type Address struct {
	Street       string `json:"street"`
	Neighborhood string `json:"neighborhood"`
	Number       uint   `json:"number"`
	City         string `json:"city"`
	State        string `json:"state"`
	ZipCode      string `json:"zipcode"`
}

func (user *User) SetPassword(password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	return nil
}

func (user *User) CheckPassword(password string) bool {
	log.Println(password)
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	return err == nil
}

func (user *User) BeforeSave(tx *gorm.DB) (err error) {
	if user.Password != "" {
		err = user.SetPassword(user.Password)
	}
	return
}
