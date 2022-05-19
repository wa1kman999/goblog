package model

import (
	"github.com/wa1kman999/goblog/pkg/category/model"
	"gorm.io/gorm"
)

// Article 文章
type Article struct {
	Category model.Category `gorm:"foreignkey:Cid" json:"category"`
	gorm.Model
	Title        string `gorm:"type:varchar(100);not null" json:"title" binding:"required"`
	Cid          int    `gorm:"type:int;not null" json:"cid" binding:"required"`
	Desc         string `gorm:"type:varchar(200)" json:"desc" binding:"required"`
	Content      string `gorm:"type:longtext" json:"content" binding:"required"`
	Img          string `gorm:"type:varchar(100)" json:"img" binding:"required"`
	CommentCount int    `gorm:"type:int;not null;default:0" json:"comment_count"`
	ReadCount    int    `gorm:"type:int;not null;default:0" json:"read_count"`
}
