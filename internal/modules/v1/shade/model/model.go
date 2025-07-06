package model

import "time"

type Shade struct {
	ID        uint `gorm:"primarykey" json:"id"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string `gorm:"type:varchar(255);not null" json:"name"`
	Image     string `gorm:"type:varchar(255);not null" json:"image"`
	SortIndex int    `gorm:"not null;default:0" json:"sort_index"`
}
