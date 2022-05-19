package dao

import (
	"github.com/wa1kman999/goblog/global"
	"github.com/wa1kman999/goblog/pkg/article/model"
	"gorm.io/gorm"
)

const (
	maxLimit     int64 = 10000
	defaultLimit int64 = 10
)

type Article interface {
	// Create 添加文章
	Create(data model.Article) error
	// FindOne 查询一个
	FindOne(fields string, query interface{}, args ...interface{}) (model.Article, error)
	// FindManyByPage 分页查询
	FindManyByPage(fields string, query *model.Article, page, pageSize int64) ([]*model.Article, int64, error)
	// Update 更新操作
	Update(value map[string]interface{}, query interface{}, args ...interface{}) error
	// Delete 删除
	Delete(query interface{}, args ...interface{}) error
}

type ArticleEntity struct {
	dao *gorm.DB
}

func NewArticleEntity() (Article, error) {
	return &ArticleEntity{
		dao: global.GBMysql,
	}, nil
}

// Create 新建一个用户
func (entity *ArticleEntity) Create(article model.Article) error {
	return entity.dao.Create(&article).Error
}

// FindOne 通过名字查询
func (entity *ArticleEntity) FindOne(fields string, query interface{}, args ...interface{}) (model.Article, error) {
	var article model.Article
	err := entity.dao.Select(fields).Where(query, args...).First(&article).Error
	if err != nil {
		return article, err
	}
	return article, nil
}

// FindManyByPage 分页查询
func (entity *ArticleEntity) FindManyByPage(fields string, query *model.Article, page, pageSize int64) ([]*model.Article, int64, error) {
	var articles []*model.Article
	var count int64
	if pageSize == 0 {
		pageSize = defaultLimit
	}
	if page == 0 {
		page = 1
	}
	if pageSize > maxLimit {
		pageSize = maxLimit
	}
	offset := (page - 1) * pageSize
	db := entity.dao.Select(fields)
	if query.Title != "" {
		db = db.Where("title LIKE ?", "%"+query.Title+"%")
	}
	err := db.Model(&model.Article{}).Count(&count).Error
	if err != nil {
		return nil, 0, err
	}
	err = db.Order("article.created_at DESC").Joins("Category").Offset(int(offset)).Limit(int(pageSize)).Find(&articles).Error
	if err != nil {
		return nil, 0, err
	}

	return articles, count, nil
}

// Update 更新操作
func (entity *ArticleEntity) Update(value map[string]interface{}, query interface{}, args ...interface{}) error {
	var article model.Article
	return entity.dao.Model(&article).Where(query, args).Updates(value).Error
}

// Delete 删除
func (entity *ArticleEntity) Delete(query interface{}, args ...interface{}) error {
	return entity.dao.Where(query, args...).Delete(&model.Article{}).Error
}
