package app

import (
	"net/http"

	"github.com/gorilla/handlers"
)

func NewRouter(c *Controller) *http.ServeMux {
    router := http.NewServeMux()

    // Health Check
    router.HandleFunc("GET /ping", Fn(c.Ping))

    router.HandleFunc("GET /status/{id}", Fn(c.GetStatus))
    router.HandleFunc("GET /name/{id}", Fn(c.GetName))
    router.HandleFunc("POST /rsvp", Fn(c.RSVP))
    router.HandleFunc("GET /responses", Fn(c.GetAllResponses))

    return router
}

func NewCorsRouter(router *http.ServeMux, allowedOrigin string) http.Handler {
    corsHandler := handlers.CORS(
        handlers.AllowedOrigins([]string{"http://localhost:3000", "http://localhost:5173", allowedOrigin}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"}),
		handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
		handlers.AllowCredentials(),
	)

    return corsHandler(router)
}

type apiFunc func(http.ResponseWriter, *http.Request) error

func Fn(fn apiFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        err := fn(w, r)
        if err != nil {
            WriteJSON(w, http.StatusInternalServerError, ErrMsg(err.Error()))
        }
    }
}

