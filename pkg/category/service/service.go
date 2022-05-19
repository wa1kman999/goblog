package service

import (
	"github.com/wa1kman999/goblog/pkg/category/dao"
	"github.com/wa1kman999/goblog/pkg/category/model"
)

type DomainCategory interface {
	// Create 创建一个用户
	Create(user model.Category) error
	// FindOne 查询一个
	FindOne(fields string, query interface{}, args ...interface{}) (model.Category, error)
	// FindManyByPage 分页查询
	FindManyByPage(fields string, query *model.Category, pageIndex, pageSize int64) ([]*model.Category, int64, error)
	// Update 更新
	Update(value map[string]interface{}, query interface{}, args ...interface{}) error
	// Delete 删除
	Delete(query interface{}, args ...interface{}) error
}

// DomainCategoryService 分类
type DomainCategoryService struct {
}

func NewDomainCategoryService() DomainCategory {
	return new(DomainCategoryService)
}

// Create 新建
func (domain *DomainCategoryService) Create(category model.Category) error {
	entity, err := dao.NewCategoryEntity()
	if err != nil {
		return err
	}
	return entity.Create(category)
}

// FindOne 查询一个
func (domain *DomainCategoryService) FindOne(fields string, query interface{}, args ...interface{}) (model.Category, error) {
	entity, err := dao.NewCategoryEntity()
	if err != nil {
		return model.Category{}, err
	}
	return entity.FindOne(fields, query, args)
}

// FindManyByPage 分页查询
func (domain *DomainCategoryService) FindManyByPage(fields string, query *model.Category, pageIndex, pageSize int64) ([]*model.Category, int64, error) {
	entity, err := dao.NewCategoryEntity()
	if err != nil {
		return nil, 0, err
	}
	return entity.FindManyByPage(fields, query, pageIndex, pageSize)
}

// Update 更新
func (domain *DomainCategoryService) Update(value map[string]interface{}, query interface{}, args ...interface{}) error {
	entity, err := dao.NewCategoryEntity()
	if err != nil {
		return err
	}
	return entity.Update(value, query, args)
}

// Delete 删除
func (domain *DomainCategoryService) Delete(query interface{}, args ...interface{}) error {
	entity, err := dao.NewCategoryEntity()
	if err != nil {
		return err
	}
	return entity.Delete(query, args)
}
