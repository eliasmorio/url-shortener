package main

import (
	"UrlShortener/internal/cache"
	"UrlShortener/internal/logging"
	"UrlShortener/internal/persist"
	"UrlShortener/internal/service"
	"UrlShortener/internal/web"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log/slog"
	"net/http"
)

func main() {
	logging.Init()

	slog.Info("Starting Web Server API")

	slog.Info("Init Redis Cache")
	redisConfig := cache.GetRedisConfig()
	urlCache := cache.NewUrlCacheFromConfig(redisConfig)

	slog.Info("Init MongoDB Database")
	mongoConfig := persist.GetMongoConfig()
	urlRepository := persist.NewUrlRepositoryFromConfig(mongoConfig)

	slog.Info("Init Grpc Client")
	keygenConfig := service.GetKeygenConfig()
	keygenService := service.NewKeygenClientServiceFromConfig(keygenConfig)

	urlService := service.NewUrlService(urlRepository, urlCache, keygenService)

	slog.Info("Start Web Server")
	urlController := web.NewUrlController(urlService)

	webServer := mux.NewRouter()

	webServer.Use(web.LoggingMiddleware)

	webServer.HandleFunc("/", urlController.CreateNewShortenedUrl).Methods("PUT")
	webServer.HandleFunc("/{short_url}", urlController.RedirectToLongUrl).Methods("GET")

	webServer.Handle("/metrics", promhttp.Handler())

	serverConfig := web.GetWebServerConfig()

	slog.Info("Web server started at " + serverConfig.Port)
	err := http.ListenAndServe(":"+serverConfig.Port, webServer)
	if err != nil {
		panic(err)
	}
}
