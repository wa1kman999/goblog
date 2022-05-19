package article

import (
	"github.com/wa1kman999/goblog/pkg/article/model"
	"github.com/wa1kman999/goblog/pkg/article/service"
	"github.com/wa1kman999/goblog/pkg/common/utils"
	"io"
	"mime/multipart"
	"os"
	"path"
	"strings"
	"time"
)

type IArticleService interface {
	CreateArticle(param model.Article) error
	Upload(file *multipart.FileHeader) (string, error)
	GetImg(path string) (*os.File, error)
	GetArticleList(title string, pageIndex, pageSize int64) ([]*model.Article, int64, error)
}

type ServiceArticle struct{}

func NewArticleService() IArticleService {
	return &ServiceArticle{}
}

// CreateArticle 新建文章
func (app *ServiceArticle) CreateArticle(param model.Article) error {
	articleService := service.NewDomainArticleService()
	if err := articleService.Create(param); err != nil {
		return err
	}
	return nil
}

// Upload 上传文件
func (app *ServiceArticle) Upload(file *multipart.FileHeader) (string, error) {
	// 读取文件后缀
	ext := path.Ext(file.Filename)
	// 读取文件名并加密
	name := strings.TrimSuffix(file.Filename, ext)
	name = utils.MD5V([]byte(name))
	// 拼接新文件名
	filename := name + "_" + time.Now().Format("20060102150405") + ext
	// 尝试创建此路径
	err := os.MkdirAll("fileDir", os.ModePerm)
	if err != nil {
		return "", err
	}
	// 拼接路径和文件名
	p := "fileDir" + "/" + filename

	f, err := file.Open() // 读取文件
	if err != nil {
		return "", err
	}
	defer f.Close() // 创建文件 defer 关闭

	out, err := os.Create(p)
	if err != nil {
		return "", err
	}
	defer out.Close() // 创建文件 defer 关闭

	_, err = io.Copy(out, f) // 传输（拷贝）文件
	if err != nil {
		return "", err
	}
	return filename, nil
}

// GetImg 获取图片
func (app *ServiceArticle) GetImg(path string) (*os.File, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	return f, nil
}

// GetArticleList 获取文章列表
func (app *ServiceArticle) GetArticleList(title string, pageIndex, pageSize int64) ([]*model.Article, int64, error) {
	categoryService := service.NewDomainArticleService()
	query := &model.Article{
		Title: title,
	}
	categoryList, count, err := categoryService.FindManyByPage("article.id, title,`desc`,comment_count,read_count, content, img, article.created_at,article.updated_at", query, pageIndex, pageSize)
	if err != nil {
		return nil, 0, err
	}
	return categoryList, count, nil
}
