package model

import "time"

type DesiredResult struct {
	ID        uint `gorm:"primarykey" json:"id"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string `gorm:"type:varchar(255);not null;unique" json:"name"`
}
