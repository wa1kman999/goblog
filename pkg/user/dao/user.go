package dao

import (
	"github.com/wa1kman999/goblog/global"
	"github.com/wa1kman999/goblog/pkg/user/model"
	"gorm.io/gorm"
)

type User interface {
	// Create 新建一个用户
	Create(data *model.User) error
}

type UserEntity struct {
	dao *gorm.DB
}

func NewUserEntity() (User, error) {
	return &UserEntity{
		dao: global.GBMysql,
	}, nil
}

// Create 新建一个用户
func (entity *UserEntity) Create(user *model.User) error {
	return entity.dao.Create(user).Error
}
