package web

import (
	"UrlShortener/internal/service"
	"log/slog"
	"net/http"
)

type UrlController struct {
	service *service.UrlService
}

func NewUrlController(service *service.UrlService) *UrlController {
	return &UrlController{service: service}
}

func (c *UrlController) RedirectToLongUrl(writer http.ResponseWriter, request *http.Request) {
	shortUrl := getPathParam(request, "short_url")
	longUrl, err := c.service.GetLongUrl(shortUrl)
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}

	writer.Header().Set("Location", longUrl)
	writer.WriteHeader(http.StatusMovedPermanently)
	writer.Header().Set("cache-control", "public")

	return
}

type CreateNewShortenedUrlRequest struct {
	LongUrl string `json:"longUrl"`
}

func (c *UrlController) CreateNewShortenedUrl(writer http.ResponseWriter, request *http.Request) {
	var requestBody CreateNewShortenedUrlRequest
	if err := getRequestBody(request, &requestBody); err != nil {
		slog.Error("Error parsing request body: ", err)
		setResponseStatusCode(writer, http.StatusBadRequest)
		return
	}

	shortenedUrl, err := c.service.CreateNewShortenedUrl(requestBody.LongUrl)
	if err != nil {
		slog.Error("Error creating new shortened url: ", err)
		setResponseStatusCode(writer, http.StatusBadRequest)
		return
	}
	setResponseJson(writer, shortenedUrl)
	return
}
