// Package article 应用的文章模型
package article

import (
	"goblog/app/models"
	"goblog/app/models/category"
	"goblog/app/models/user"
	"goblog/pkg/route"
	"strconv"
)

// Article 文章模型
type Article struct {
	models.BaseModel

	Title string `gorm:"type:varchar(255);not null" valid:"title"`
	Body  string `gorm:"type:longtext;not null" valid:"body"`

	UserID uint64 `gorm:"not null;index"`
	User   user.User

	CategoryID uint64 `gorm:"not null;default:4;index"`
	Category   category.Category
}

// Link 方法用来生成文章链接
func (a Article) Link() string {
	return route.Name2URL("articles.show", "id", strconv.FormatUint(a.ID, 10))
}

func (a Article) CreatedAtDate() string {
	return a.CreatedAt.Format("2006-01-02")
}
