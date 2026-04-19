package http

import (
	"github.com/go-chi/chi/v5"
	ver1 "github.com/twenty99h/wishes-service/internal/controller/http/v1"
	"github.com/twenty99h/wishes-service/internal/usecase"
)

func WishesRouter(r *chi.Mux, uc *usecase.UseCase) {
	v1 := ver1.New(uc)

	r.Route("/api", func(r chi.Router) {
		r.Route("/v1", func(r chi.Router) {
			r.Post("/wishes", v1.CreateWish)
			r.Get("/wishes/{id}", v1.GetWish)
			r.Put("/wishes/{id}", v1.UpdateWish)
			r.Delete("/wishes/{id}", v1.DeleteWish)
			r.Get("/wishes", v1.GetWishes)
		})
	})
}
