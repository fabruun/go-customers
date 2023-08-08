package application

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/fabruun/go-customers/repository"
)

func loadRoutes() *chi.Mux {
	router := chi.NewRouter()

	router.Use(middleware.Logger)

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	router.Route("/customers", loadCustomerRoutes)

	return router
}

func loadCustomerRoutes(router chi.Router) {
	customerHandler := &repository.Customer{}

	router.Get("/", customerHandler.List)
}
