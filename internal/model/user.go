package model

import "time"

type User struct {
	ID          int64     `gorm:"primaryKey;autoIncrement" json:"id"`
	CompanyName string    `gorm:"type:varchar(255);not null" json:"company_name"`
	Name        string    `gorm:"type:varchar(255);not null" json:"name"`
	Email       string    `gorm:"type:varchar(255);uniqueIndex;not null" json:"email"`
	Password    string    `gorm:"type:varchar(255);not null" json:"password"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

