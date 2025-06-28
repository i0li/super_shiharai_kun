package model

import "time"

type User struct {
	ID          int64     `gorm:"primaryKey" json:"id"`
	CompanyName string    `gorm:"not null" json:"company_name"`
	Name        string    `gorm:"not null" json:"name"`
	Email       string    `gorm:"uniqueIndex;not null" json:"email"`
	Password    string    `gorm:"not null" json:"password"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
