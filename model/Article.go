package model

import (
	"github.com/jinzhu/gorm"
	"github.com/wejectchen/ginblog/utils/errmsg"
)

type Article struct {
	Category Category `gorm:"foreignkey:Cid"`
	gorm.Model
	Title        string `gorm:"type:varchar(100);not null" json:"title"`
	Cid          int    `gorm:"type:int;not null" json:"cid"`
	Desc         string `gorm:"type:varchar(200)" json:"desc"`
	Content      string `gorm:"type:longtext" json:"content"`
	Img          string `gorm:"type:varchar(100)" json:"img"`
	CommentCount int    `gorm:"type:int;not null;default:0" json:"comment_count"`
	ReadCount    int    `gorm:"type:int;not null;default:0" json:"read_count"`
}

// 新增文章
func CreateArt(data *Article) int {
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR // 500
	}
	return errmsg.SUCCSE
}

// 查询分类下的所有文章
func GetCateArt(id int, pageSize int, pageNum int) ([]Article, int, int64) {
	var cateArtList []Article
	var total int64

	err = db.Preload("Category").Limit(pageSize).Offset((pageNum-1)*pageSize).Where(
		"cid =?", id).Find(&cateArtList).Error
	db.Model(&cateArtList).Where("cid =?", id).Count(&total)
	if err != nil {
		return nil, errmsg.ERROR_CATE_NOT_EXIST, 0
	}
	return cateArtList, errmsg.SUCCSE, total

}

// 查询单个文章
func GetArtInfo(id int) (Article, int) {
	var art Article
	err := db.Preload("Category").Where("id = ?", id).First(&art).Error
	if err != nil {
		return art, errmsg.ERROR_ART_NOT_EXIST
	}
	return art, errmsg.SUCCSE
}

// 查询文章列表
func GetArt(pageSize int, pageNum int) ([]Article, int, int64) {
	var articleList []Article
	var err error
	var total int64
	err = db.Preload("Category").Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&articleList).Count(&total).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, errmsg.ERROR, 0
	}
	return articleList, errmsg.SUCCSE, total

}

// SearchArticle 搜索文章标题
func SearchArticle(title string, pageSize int, pageNum int) ([]Article, int, int64) {
	var articleList []Article
	var err error
	var total int64
	err = db.Select("article.id,title, img, created_at, updated_at, `desc`, comment_count, read_count, Category.name").Order("Created_At DESC").Joins("Category").Where("title LIKE ?",
		title+"%",
	).Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&articleList).Error
	//单独计数
	db.Model(&articleList).Where("title LIKE ?",
		title+"%",
	).Count(&total)

	if err != nil {
		return nil, errmsg.ERROR, 0
	}
	return articleList, errmsg.SUCCSE, total
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
