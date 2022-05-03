package dao

import (
	"github.com/wa1kman999/goblog/global"
	"github.com/wa1kman999/goblog/pkg/user/model"
	"gorm.io/gorm"
)

const (
	maxLimit     int64 = 10000
	defaultLimit int64 = 10
)

type User interface {
	// Create 新建一个用户
	Create(data *model.User) error
	// FindOne 查询一个
	FindOne(fields string, query interface{}, args ...interface{}) (model.User, error)
	// FindManyByPage 分页查询
	FindManyByPage(fields string, query *model.User, pageIndex, pageSize int64) ([]*model.User, int64, error)
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

// FindOne 通过名字查询
func (entity *UserEntity) FindOne(fields string, query interface{}, args ...interface{}) (model.User, error) {
	var user model.User
	err := entity.dao.Select(fields).Where(query, args...).First(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

// FindManyByPage 分页查询
func (entity *UserEntity) FindManyByPage(fields string, query *model.User, pageIndex, pageSize int64) ([]*model.User, int64, error) {
	var users []*model.User
	var count int64
	if pageSize == 0 {
		pageSize = defaultLimit
	}
	if pageIndex == 0 {
		pageIndex = 1
	}
	if pageSize > maxLimit {
		pageSize = maxLimit
	}
	offset := (pageIndex - 1) * pageSize
	db := entity.dao.Select(fields)
	if query.Username != "" {
		db = db.Where("username LIKE ?", "%"+query.Username+"%")
	}
	if query.Role != 0 {
		db = db.Where("role = ?", query.Role)
	}
	err := db.Model(&model.User{}).Count(&count).Error
	if err != nil {
		return nil, 0, err
	}
	err = db.Offset(int(offset)).Limit(int(pageSize)).Find(&users).Error
	if err != nil {
		return nil, 0, err
	}

	return users, count, nil
}
