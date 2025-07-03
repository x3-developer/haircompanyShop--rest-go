package model

import (
	"time"
)

type DashboardUser struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Email     string `gorm:"type:varchar(255);not null;unique" json:"email"`
	Password  string `gorm:"type:varchar(255);not null" json:"password"`
	Role      string `gorm:"type:ENUM('admin','manager');not null" json:"role"`
}
