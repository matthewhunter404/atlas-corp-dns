package server

import (
	"net/http"

	"go-atlas-corp/domain/sector"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

type Server interface {
	fetchLoc(w http.ResponseWriter, r *http.Request)
	Router() *chi.Mux
}

type server struct {
	Sector sector.Sector
}

func New(sector sector.Sector) Server {
	return &server{
		Sector: sector,
	}
}

func (s *server) Router() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Server is running"))
	})
	r.Post("/location", s.fetchLoc)

	return r
}
