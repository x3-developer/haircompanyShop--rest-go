package model

import (
	"time"
)

type ClientUser struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Phone     string `gorm:"type:varchar(255);not null;unique" json:"phone"`
}
