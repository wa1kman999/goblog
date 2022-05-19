package user

import (
	"github.com/go-errors/errors"
	"github.com/wa1kman999/goblog/pkg/category/model"
	"github.com/wa1kman999/goblog/pkg/category/service"
)

type ICategoryService interface {
	CreateCategory(param model.Category) error
	GetCategoryInfo(id int) (model.Category, error)
	GetCategoryList(categoryName string, pageIndex, pageSize int64) ([]*model.Category, int64, error)
	EditCategory(param model.Category) error
	DeleteCategory(id int) error
}

type ServiceCategory struct{}

func NewCategoryService() ICategoryService {
	return &ServiceCategory{}
}

// CreateCategory 创建分类
func (app *ServiceCategory) CreateCategory(param model.Category) error {
	categoryService := service.NewDomainCategoryService()
	// 查询是否有同名的
	category, _ := categoryService.FindOne("id", "name = ?", param.Name)
	// 大于0 表示有同名的
	if category.ID > 0 {
		return errors.New("名字重复")
	}
	if err := categoryService.Create(param); err != nil {
		return err
	}
	return nil
}

// GetCategoryInfo 获取单个信息
func (app *ServiceCategory) GetCategoryInfo(id int) (model.Category, error) {
	categoryService := service.NewDomainCategoryService()
	categoryInfo, err := categoryService.FindOne("id, name", "id = ?", id)
	if err != nil {
		return model.Category{}, err
	}
	return categoryInfo, nil
}

// GetCategoryList 获取category列表
func (app *ServiceCategory) GetCategoryList(categoryName string, pageIndex, pageSize int64) ([]*model.Category, int64, error) {
	categoryService := service.NewDomainCategoryService()
	query := &model.Category{
		Name: categoryName,
	}
	categoryList, count, err := categoryService.FindManyByPage("id, name", query, pageIndex, pageSize)
	if err != nil {
		return nil, 0, err
	}
	return categoryList, count, nil
}

// EditCategory 更新user信息
func (app *ServiceCategory) EditCategory(param model.Category) error {
	categoryService := service.NewDomainCategoryService()
	return categoryService.Update(map[string]interface{}{
		"name": param.Name,
	}, "id = ?", param.ID)
}

func (app *ServiceCategory) DeleteCategory(id int) error {
	return service.NewDomainCategoryService().Delete("id = ?", id)
}
