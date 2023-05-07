package model

import (
	"github.com/jinzhu/gorm"
	"github.com/wejectchen/ginblog/utils/errmsg"
)

type Article struct {
	Category Category
	gorm.Model
	Title   string `gorm:"type:varchar(100);not null" json:"title"` //标题
	Cid     int    `gorm:"type:int;not null" json:"cid"`            //分类
	Desc    string `gorm:"type:varchar(200)" json:"desc"`           //描述
	Content string `gorm:"type:longtext" json:"content"`            //内容
	Img     string `gorm:"type:varchar(100)" json:"img"`            //图片
}

func CreateArt(data *Article) int {
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR // 500
	}
	return errmsg.SUCCSE
}
