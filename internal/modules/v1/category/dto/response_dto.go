package dto

import (
	"time"
)

type ResponseDTO struct {
	Id              uint      `json:"id"`
	CreatedAt       time.Time `json:"createdAt"`
	UpdatedAt       time.Time `json:"updatedAt"`
	Name            string    `json:"name"`
	Description     string    `json:"description"`
	Image           string    `json:"image"`
	HeaderImage     string    `json:"headerImage"`
	Slug            string    `json:"slug"`
	ParentID        *uint     `json:"parentId"`
	SortIndex       int       `json:"sortIndex"`
	SeoTitle        string    `json:"seoTitle"`
	SeoDescription  string    `json:"seoDescription"`
	SeoKeys         string    `json:"seoKeys"`
	IsActive        bool      `json:"isActive"`
	IsShade         bool      `json:"isShade"`
	IsVisibleInMenu bool      `json:"isVisibleInMenu"`
	IsVisibleOnMain bool      `json:"isVisibleOnMain"`
}
