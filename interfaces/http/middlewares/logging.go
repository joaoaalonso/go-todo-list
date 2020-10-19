package middlewares

import (
	"net/http"
	"os"

	"github.com/gorilla/handlers"
)

// LoggingHandler log all requestes
func LoggingHandler(h http.Handler) http.Handler {
	return handlers.LoggingHandler(os.Stdout, h)
}
