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

// Router returns a go-chi mux that can be passed to http.ListenAndServe()
func (s *server) Router() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Server is running"))
	})
	r.Route("/v1", func(r chi.Router) {
		r.Post("/location", s.fetchLoc)
	})

	return r
}
