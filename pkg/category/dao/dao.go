package dao

import (
	"github.com/wa1kman999/goblog/global"
	"github.com/wa1kman999/goblog/pkg/category/model"
	"gorm.io/gorm"
)

const (
	maxLimit     int64 = 10000
	defaultLimit int64 = 10
)

type Category interface {
	// Create 新建一个用户
	Create(data model.Category) error
	// FindOne 查询一个
	FindOne(fields string, query interface{}, args ...interface{}) (model.Category, error)
	// FindManyByPage 分页查询
	FindManyByPage(fields string, query *model.Category, pageIndex, pageSize int64) ([]*model.Category, int64, error)
	// Update 更新操作
	Update(value map[string]interface{}, query interface{}, args ...interface{}) error
	// Delete 删除
	Delete(query interface{}, args ...interface{}) error
}

type CategoryEntity struct {
	dao *gorm.DB
}

func NewCategoryEntity() (Category, error) {
	return &CategoryEntity{
		dao: global.GBMysql,
	}, nil
}

// Create 新建一个用户
func (entity *CategoryEntity) Create(user model.Category) error {
	return entity.dao.Create(&user).Error
}

// FindOne 通过名字查询
func (entity *CategoryEntity) FindOne(fields string, query interface{}, args ...interface{}) (model.Category, error) {
	var category model.Category
	err := entity.dao.Select(fields).Where(query, args...).First(&category).Error
	if err != nil {
		return category, err
	}
	return category, nil
}

// FindManyByPage 分页查询
func (entity *CategoryEntity) FindManyByPage(fields string, query *model.Category, pageIndex, pageSize int64) ([]*model.Category, int64, error) {
	var categories []*model.Category
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
	if query.Name != "" {
		db = db.Where("name LIKE ?", "%"+query.Name+"%")
	}
	err := db.Model(&model.Category{}).Count(&count).Error
	if err != nil {
		return nil, 0, err
	}
	err = db.Offset(int(offset)).Limit(int(pageSize)).Find(&categories).Error
	if err != nil {
		return nil, 0, err
	}

	return categories, count, nil
}

// Update 更新操作
func (entity *CategoryEntity) Update(value map[string]interface{}, query interface{}, args ...interface{}) error {
	var category model.Category
	return entity.dao.Model(&category).Where(query, args).Updates(value).Error
}

// Delete 删除
func (entity *CategoryEntity) Delete(query interface{}, args ...interface{}) error {
	return entity.dao.Where(query, args...).Delete(&model.Category{}).Error
}
