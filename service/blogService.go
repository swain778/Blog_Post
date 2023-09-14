package service

import (
	"v0/database"
	"v0/models"

	"gorm.io/gorm"
)

type BlogService struct {
	db *gorm.DB
}

func NewBlogService() *BlogService {
	return &BlogService{
		db: database.GetDB(),
	}
}

// CreateBlog creates blog
func (b *BlogService) CreatePost(post *models.Post) (*models.Post, error) {

	err := b.db.Create(&post).Error
	if err != nil {
		return nil, err
	}
	return post, nil
}
