package model

import (
	"time"
)

type Category struct {
	ID              uint `gorm:"primarykey"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
	Name            string `gorm:"type:varchar(255);not null;unique" json:"name"`
	Description     string `gorm:"type:text" json:"description"`
	Image           string `gorm:"type:varchar(255)" json:"image"`
	HeaderImage     string `gorm:"type:varchar(255)" json:"headerImage"`
	Slug            string `gorm:"type:varchar(255);not null;unique" json:"slug"`
	ParentID        *uint  `gorm:"index" json:"parentId"`
	SortIndex       int    `gorm:"default:100" json:"sortIndex"`
	SeoTitle        string `gorm:"type:varchar(255)" json:"seoTitle"`
	SeoDescription  string `gorm:"type:text" json:"seoText"`
	SeoKeys         string `gorm:"type:text" json:"seoKeys"`
	IsActive        bool   `gorm:"default:true" json:"isActive"`
	IsShade         bool   `gorm:"default:false" json:"isShade"`
	IsVisibleInMenu bool   `gorm:"default:true" json:"isVisibleInMenu"`
	IsVisibleOnMain bool   `gorm:"default:false" json:"isVisibleOnMain"`
}
