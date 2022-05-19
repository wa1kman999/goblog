package model

import "gorm.io/gorm"

// Category 分类
type Category struct {
	gorm.Model
	//ID   uint   `gorm:"primaryKey;auto_increment" json:"id"`
	Name string `gorm:"type:varchar(20);not null" json:"name"`
}
