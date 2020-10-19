package middlewares

import (
	"net/http"

	"github.com/gorilla/handlers"
)

// CompressHandler handle response compression
func CompressHandler(h http.Handler) http.Handler {
	return handlers.CompressHandler(h)
}
