package router

import (
	"github.com/go-chi/chi"

	"myapp/app/app"
)

// New return a pointer to a new router
func New() *chi.Mux {
	r := chi.NewRouter()

	r.MethodFunc("GET", "/", app.HandleIndex)
	return r
}
