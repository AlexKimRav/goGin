package model

import (
	"time"

	"gorm.io/gorm"
)

// column->:false : disabled read from db
// json:”-” : ignore a field

type BaseModel struct {
	CreatedAt time.Time      `gorm:"->:false;column:created_at" json:"-"`
	UpdatedAt time.Time      `gorm:"->:false;column:updated_at" json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"->:false;column:deleted_at" json:"-"`
}

type Album struct {
	Id     int     `gorm:"column:id; primary_key; not null" json:"id"`
	Title  string  `gorm:"column:title" json:"title"`
	Artist string  `gorm:"column:artist" json:"artist"`
	Price  float64 `gorm:"column:price" json:"price"`
	BaseModel
}

type User struct {
	Id       uint   `gorm:"column:id; primary_key; not null" json:"id"`
	Name     string `gorm:"column:name" json:"name"`
	Email    string `gorm:"column:email" json:"email"`
	Nickname string `gorm:"column:nickname" json:"nickname"`
	Age      int    `gorm:"column:age" json:"age"`
	BaseModel
}
