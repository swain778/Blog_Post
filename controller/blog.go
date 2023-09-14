package controller

import (
	"net/http"
	"strconv"
	"v0/models"
	"v0/service"

	"github.com/go-chi/chi"
)

// CreateBlog creates a blog
func CreatePost(w http.ResponseWriter, r *http.Request) {
	service := service.NewBlogService()

	blog, err := service.CreatePost(&models.Post{
		Title:   r.FormValue("title"),
		Content: r.FormValue("content"),
	})

	if err != nil {
		ApiResponse(w, &Res{
			Code:    900,
			Message: "error",
			Data:    err.Error(),
		})
		return
	}
	ApiResponse(w, &Res{
		Code:    200,
		Message: "blog created",
		Data:    blog,
	})
}

// GetsPosts give the posts list
func GetsPosts(w http.ResponseWriter, r *http.Request) {
	service := service.NewBlogService()
	blog, err := service.GetsPosts()
	if err != nil {
		ApiResponse(w, &Res{
			Code:    900,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}
	ApiResponse(w, &Res{
		Code:    200,
		Message: "success",
		Data:    blog,
	})
}

// GetBlog gets a blog with specific ID
func GetPost(w http.ResponseWriter, r *http.Request) {
	service := service.NewBlogService()

	blog, err := service.GetPost(chi.URLParam(r, "id"))
	if err != nil {
		ApiResponse(w, &Res{
			Code:    900,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}
	ApiResponse(w, &Res{
		Code:    200,
		Message: "success",
		Data:    blog,
	})
}

// DeletePost deletes a post with specific ID
func DeletePost(w http.ResponseWriter, r *http.Request) {
	service := service.NewBlogService()

	blog, err := service.DeletePost(chi.URLParam(r, "id"))
	if err != nil {
		ApiResponse(w, &Res{
			Code:    900,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}
	ApiResponse(w, &Res{
		Code:    200,
		Message: "deleted successfully",
		Data:    blog,
	})
}

// Update post update the specific post by ID
func UpdatePost(w http.ResponseWriter, r *http.Request) {
	service := service.NewBlogService()

	title := r.FormValue("title")
	content := r.FormValue("content")

	if title == "" && content == "" {
		ApiResponse(w, &Res{
			Code:    903,
			Message: "can't update blank values",
			Data:    nil,
		})
		return
	}

	post := &models.Post{}
	if title != "" {
		post.Title = title
	}

	if content != "" {
		post.Content = content
	}
	postID, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		ApiResponse(w, &Res{
			Code:    901,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}
	post.ID = uint(postID)
	blog, err := service.UpdatePost(post)
	if err != nil {
		ApiResponse(w, &Res{
			Code:    900,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}
	ApiResponse(w, &Res{
		Code:    200,
		Message: "success",
		Data:    blog,
	})
}
