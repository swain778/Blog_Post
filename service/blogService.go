package service

import (
	"errors"
	"v0/database"
	"v0/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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

// GetsPosts give the posts list
func (b *BlogService) GetsPosts() (*[]models.Post, error) {
	posts := &[]models.Post{}
	err := b.db.Model(&[]models.Post{}).Find(posts).Error
	if err != nil {
		return nil, errors.New("can't gets blog posts")
	}
	return posts, nil
}

// GetBlog gets a blog with specific ID
func (b *BlogService) GetPost(blogID string) (*models.Post, error) {
	post := &models.Post{}
	err := b.db.Model(&models.Post{}).Preload(clause.Associations).
		First(post, "id=?", blogID).Error
	if err != nil {
		return nil, errors.New("can't get post")
	}
	return post, nil
}
