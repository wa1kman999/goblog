package model

import "gorm.io/gorm"

// Comment 评论
type Comment struct {
	gorm.Model
	UserId    uint   `json:"user_id"`
	ArticleId uint   `json:"article_id"`
	Title     string `json:"article_title"`
	Username  string `json:"username"`
	Content   string `cond:"type:varchar(500);not null;" json:"content"`
	Status    int8   `cond:"type:tinyint;default:2" json:"status"`
}
