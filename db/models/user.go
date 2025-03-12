package models

import (
	"errors"
	"time"
	"gorm.io/gorm"
)

type User struct {
	Id                 uint               `gorm:"primaryKey"`
	FirstName          string             `json:"first_name" gorm:"not null"`
	LastName           string             `json:"last_name" gorm:"not null"`
	Email              string             `json:"email" gorm:"unique;not null"`
	Password           string             `json:"password" gorm:"not null"`
	LastModified       time.Time          `json:"last_modified" gorm:"autoUpdateTime"`
	Gender             string             `json:"gender"`
	MaritalStatus      string             `json:"marital_status"`
	ResidentialDetails *ResidentialDetails `json:"residential_details" gorm:"type:json"`
	OfficeDetails      *OfficeDetails      `json:"office_details" gorm:"type:json"`
}

type ResidentialDetails struct {
	Address string `json:"address"`
	City    string `json:"city"`
	State   string `json:"state"`
	Country string `json:"country"`
	Contact string `json:"contact"`
}

type OfficeDetails struct {
	Name         string `json:"name"`
	Email        string `json:"email"`
	EmployeeCode string `json:"employee_code"`
	Address      string `json:"address"`
	City         string `json:"city"`
	State        string `json:"state"`
	Country      string `json:"country"`
	Contact      string `json:"contact"`
}

func (user *User) GetUser(db *gorm.DB, queries [][]string) error {
	q := db.Model(User{})
	for _, query := range queries {
		q = q.Where(query[0], query[1])
	}
	result := q.First(&user)

	if result.Error != nil {
		return errors.New("User not found")
	}
	return nil
}
