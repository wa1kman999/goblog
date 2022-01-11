package initialize

import (
	"os"

	"github.com/sirupsen/logrus"
	articleModel "github.com/wa1kman999/goblog/pkg/article/model"
	categoryModel "github.com/wa1kman999/goblog/pkg/category/model"
	commentModel "github.com/wa1kman999/goblog/pkg/comment/model"
	profileModel "github.com/wa1kman999/goblog/pkg/profile/model"
	userModel "github.com/wa1kman999/goblog/pkg/user/model"
	"gorm.io/gorm"
)

// RegisterTables 注册数据库表专用
func RegisterTables(db *gorm.DB) {
	err := db.AutoMigrate(
		&userModel.User{},
		&articleModel.Article{},
		&categoryModel.Category{},
		&commentModel.Comment{},
		&profileModel.Profile{},
	)

	if err != nil {
		logrus.Errorf("register table failed: %s", err.Error())
		os.Exit(0)
	}
	logrus.Info("register table success")
}
