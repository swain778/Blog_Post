package controller

import (
	"net/http"
	"v0/models"
	"v0/service"
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
