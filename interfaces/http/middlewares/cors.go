package middlewares

import (
	"net/http"

	"github.com/gorilla/handlers"
)

// CorsHandler apple cors rules
func CorsHandler(h http.Handler) http.Handler {
	return handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(h)
}
