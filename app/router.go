package app

import (
	"v0/controller"

	"github.com/go-chi/chi"
)

func (ctn *Container) LoadRoutes() {
	r := ctn.Mux

	r.Route("/api", func(r chi.Router) {
		r.Post("/post", controller.CreatePost)
	})
}
