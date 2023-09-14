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

// DeletePost deletes a post with specific ID
func (b *BlogService) DeletePost(blogID string) (bool, error) {
	err := b.db.Where("id=?", blogID).Delete(&models.Post{}).Error
	if err != nil {
		return false, errors.New("can't delete blog")
	}
	return true, nil
}

// Update post update the specific post by ID
func (b *BlogService) UpdatePost(post *models.Post) (bool, error) {
	err := b.db.Where("id=?", post.ID).Updates(&post).Error
	if err != nil {
		return false, errors.New("can't update blogs")
	}
	return true, nil
}
