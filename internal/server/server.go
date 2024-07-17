package server

import (
	"log"
	"net/http"

	"app/internal/post/handler"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Server struct {
	router      *chi.Mux
	postHandler *handler.PostHandler
}

func NewServer(postHandler *handler.PostHandler) *Server {
	s := &Server{
		router:      chi.NewRouter(),
		postHandler: postHandler,
	}
	s.setupRoutes()
	return s
}

func (s *Server) setupRoutes() {
	s.router.Use(middleware.Logger)
	s.router.Use(middleware.Recoverer)

	s.router.Route("/posts", func(r chi.Router) {
		r.Get("/", s.postHandler.GetAllPosts)
		r.Post("/", s.postHandler.CreatePost)
		r.Get("/{id}", s.postHandler.GetPost)
		r.Delete("/delete", s.postHandler.DeletePosts)
	})
}

func (s *Server) Run(addr string) {
	log.Printf("Server is running on %s", addr)
	http.ListenAndServe(addr, s.router)
}
