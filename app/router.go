package app

import (
	"v0/controller"

	"github.com/go-chi/chi"
)

func (ctn *Container) LoadRoutes() {
	r := ctn.Mux

	r.Route("/api", func(r chi.Router) {
		r.Post("/post", controller.CreatePost)
		r.Get("/posts", controller.GetsPosts)
		r.Get("/post/{id}", controller.GetPost)
		r.Delete("/post/{id}", controller.DeletePost)
		r.Put("/post/{id}", controller.UpdatePost)
	})
}
