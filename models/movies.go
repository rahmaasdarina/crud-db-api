package models

import "gorm.io/gorm"

type Movies struct {
	gorm.Model
	Title       string `gorm:"type:varchar(255);not null; index"`
	Slug        string `gorm:"type:varchar(255);unique;not null; index"`
	Description string `gorm:"type:text;not null"`
	Duration    uint   `gorm:"type:int(5);not null"`
	Image       string `gorm:"type:varchar(255);not null"`
}

type MoviesData struct {
	ID          int    `gorm:"column:id; PRIMARY_KEY" json:"id"`
	Title       string `gorm:"column:title" json:"title"`
	Slug        string `gorm:"column:slug" json:"slug"`
	Description string `gorm:"column:description" json:"description"`
	Duration    uint   `gorm:"column:duration" json:"duration"`
	Image       string `gorm:"column:image" json:"image"`
}
