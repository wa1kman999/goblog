package model

// Category 分类
type Category struct {
	ID   uint   `cond:"primary_key;auto_increment" json:"id"`
	Name string `cond:"type:varchar(20);not null" json:"name"`
}
