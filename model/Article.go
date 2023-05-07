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

// 新增文章
func CreateArt(data *Article) int {
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR // 500
	}
	return errmsg.SUCCSE
}

// todo 查询分类下的所有文章

// todo 查询单个文章

// 查询文章列表
func GetArt(pageSize int, pageNum int) []Category {
	var cate []Category
	err = db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&cate).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil
	}
	return cate
}

// 编辑文章
func EditArt(id int, data *Article) int {
	var art Article
	var maps = make(map[string]interface{})
	maps["title"] = data.Title
	maps["cid"] = data.Cid
	maps["desc"] = data.Desc
	maps["content"] = data.Content
	maps["img"] = data.Img

	err = db.Model(&art).Where("id = ? ", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

// 删除文章
func DeleteArt(id int) int {
	var art Article
	err = db.Where("id = ? ", id).Delete(&art).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}
