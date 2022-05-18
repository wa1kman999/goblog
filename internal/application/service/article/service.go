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
}

type ServiceArticle struct{}

func NewArticleService() IArticleService {
	return &ServiceArticle{}
}

// CreateArticle 新建文章
func (app *ServiceArticle) CreateArticle(param model.Article) error {
	userService := service.NewDomainUserService()
	if err := userService.Create(param); err != nil {
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
	return p, nil
}
