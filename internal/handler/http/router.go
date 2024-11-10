package http

import (
	"github.com/go-chi/chi"
)

func NewRouters(handler *Service) *chi.Mux {
	r := chi.NewRouter()

	r.Route("/v1", func(r chi.Router) {
		r.Route("/user", func(r chi.Router) {
			r.Post("/", handler.CreateUser)
			r.Get("/", handler.SearchUser)
			r.Get("/{email}", handler.GetUser)
			r.Put("/{email}", handler.UpdateUser)
			r.Delete("/{email}", handler.DeleteUser)
		})
	})

	return r
}
