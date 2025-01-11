package postgres

import "gorm.io/gorm"

type CloudResource struct {
	gorm.Model
	ID         int64
	Name       string `gorm:"unique"`
	Type       string
	Region     string
	CustomerID int64
}

type Customer struct {
	gorm.Model
	ID    int64
	Name  string
	Email string `gorm:"unique"`
}

type Notification struct {
	gorm.Model
	ID      int64
	UserID  string
	Message string
}
