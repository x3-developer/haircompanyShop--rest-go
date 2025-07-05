package model

import "time"

type Line struct {
	ID        uint `gorm:"primarykey" json:"id"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string `gorm:"type:varchar(255);not null;unique" json:"name"`
	Color     string `gorm:"type:varchar(7);not null" json:"color"` // Hex color code
}
