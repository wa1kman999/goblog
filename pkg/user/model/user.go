package model

import "gorm.io/gorm"

// User 用户
type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(20);not null " json:"username" binding:"required,max=12" label:"用户名"`
	Password string `gorm:"type:varchar(500);not null" json:"password" binding:"required,min=6,max=120" label:"密码"`
	Role     int    `gorm:"type:int;DEFAULT:2" json:"role" binding:"required,gte=2" label:"角色码"`
}
