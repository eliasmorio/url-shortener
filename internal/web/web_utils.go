package web

import (
	"UrlShortener/internal/config"
	"encoding/json"
	"github.com/gorilla/mux"
	"log/slog"
	"net/http"
)

type WebServerConfig struct {
	Port string `env:"PORT" envDefault:"8080"`
}

func GetWebServerConfig() WebServerConfig {
	webServerConfig := WebServerConfig{}
	err := config.LoadConfig(&webServerConfig)
	if err != nil {
		panic(err)
	}
	return webServerConfig
}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		slog.Info("Request received: " + r.RequestURI)
		next.ServeHTTP(w, r)
	})
}

func getPathParam(r *http.Request, paramName string) string {
	vars := mux.Vars(r)
	return vars[paramName]
}

func setResponseStatusCode(writer http.ResponseWriter, statusCode int) {
	writer.WriteHeader(statusCode)
}

func setResponseJsonBody(writer http.ResponseWriter, body interface{}) {
	err := json.NewEncoder(writer).Encode(body)
	if err != nil {
		setResponseStatusCode(writer, http.StatusInternalServerError)
		return
	}
}

func setResponseJson(writer http.ResponseWriter, body interface{}) {
	writer.Header().Set("Content-Type", "application/json")
	setResponseJsonBody(writer, body)
}

func getRequestBody(r *http.Request, body interface{}) error {
	return json.NewDecoder(r.Body).Decode(body)
}
