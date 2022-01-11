package model

import (
	"github.com/wa1kman999/goblog/pkg/category/model"
	"gorm.io/gorm"
)

// Article 文章
type Article struct {
	Category model.Category `cond:"foreignkey:Cid"`
	gorm.Model
	Title        string `cond:"type:varchar(100);not null" json:"title"`
	Cid          int    `cond:"type:int;not null" json:"cid"`
	Desc         string `cond:"type:varchar(200)" json:"desc"`
	Content      string `cond:"type:longtext" json:"content"`
	Img          string `cond:"type:varchar(100)" json:"img"`
	CommentCount int    `cond:"type:int;not null;default:0" json:"comment_count"`
	ReadCount    int    `cond:"type:int;not null;default:0" json:"read_count"`
}
