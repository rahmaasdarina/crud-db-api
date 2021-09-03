package models

import "gorm.io/gorm"

type Movies struct {
	gorm.Model
	Title       string `gorm:"type:varchar(255);not null; index" json:"title"`
	Slug        string `gorm:"type:varchar(255);unique;not null; index" json:"slug"`
	Description string `gorm:"type:text;not null" json:"description"`
	Duration    uint   `gorm:"type:int(5);not null" json:"duration"`
	Image       string `gorm:"type:varchar(255);not null" json:"image"`
}

// type MoviesData struct {
// 	ID          int    `gorm:"column:id; PRIMARY_KEY" json:"id"`
// 	Title       string `gorm:"column:title" json:"title"`
// 	Slug        string `gorm:"column:slug" json:"slug"`
// 	Description string `gorm:"column:description" json:"description"`
// 	Duration    uint   `gorm:"column:duration" json:"duration"`
// 	Image       string `gorm:"column:image" json:"image"`
// }
