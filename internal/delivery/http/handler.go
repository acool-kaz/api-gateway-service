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

func InitHandler() *Handler {
	log.Println("init http handler")
	return &Handler{}
}

func (h *Handler) InitRoutes() http.Handler {
	log.Println("init routes")

	router := chi.NewRouter()
	router.Use(middleware.Logger)

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "hello this is root of api gateway service")
	})

	return router
}
