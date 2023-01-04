package http

import (
	"fmt"
	"log"
	"net/http"

	parser "github.com/acool-kaz/api-gateway-service/pkg/parser/client"
	post_crud "github.com/acool-kaz/api-gateway-service/pkg/post_crud/client"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Handler struct {
	parserClient *parser.ParserClient
	postClient   *post_crud.PostCRUDClient
}

func InitHandler(parserClient *parser.ParserClient, postClient *post_crud.PostCRUDClient) *Handler {
	log.Println("init http handler")
	return &Handler{
		parserClient: parserClient,
		postClient:   postClient,
	}
}

func (h *Handler) InitRoutes() http.Handler {
	log.Println("init routes")

	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			h.ServeHTTP(w, r)
		})
	})

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "hello this is root of api gateway service")
	})

	router.Route("/post", func(post chi.Router) {
		post.Get("/", h.getAllPosts)
		post.Post("/", h.createPost)
		post.Get("/{post_id}", h.getPostById)
		post.Delete("/{post_id}", h.deletePostById)
		post.Patch("/{post_id}", h.updatePostById)
	})

	router.Route("/parser", func(parser chi.Router) {
		parser.Get("/", h.parserHandler)
	})

	return router
}
