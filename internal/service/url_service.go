package service

import (
	"UrlShortener/internal/cache"
	error2 "UrlShortener/internal/error"
	"UrlShortener/internal/model"
	"UrlShortener/internal/persist"
	"context"
)

func NewUrlService(urlRepository persist.UrlRepository, urlCache cache.UrlCache, kgsClient KgsClient) *UrlService {
	return &UrlService{
		urlRepository: urlRepository,
		urlCache:      urlCache,
		kgsClient:     kgsClient,
	}
}

type UrlService struct {
	urlRepository persist.UrlRepository
	urlCache      cache.UrlCache
	kgsClient     KgsClient
}

func (s *UrlService) CreateNewShortenedUrl(longUrl string) (model.ShortenedUrl, error) {
	if s.isLongUrlExists(longUrl) {
		return model.ShortenedUrl{}, error2.LongUrlAlreadyExist{LongUrl: longUrl}
	}

	r, err := s.kgsClient.GenKey(context.Background(), &KeyRequest{Url: longUrl})

	if err != nil {
		return model.ShortenedUrl{}, err
	}

	shortenedUrl := model.ShortenedUrl{
		LongUrl:  longUrl,
		ShortUrl: r.Key,
	}

	err = s.urlRepository.Store(&shortenedUrl)
	if err != nil {
		return model.ShortenedUrl{}, err
	}

	return shortenedUrl, nil
}

func (s *UrlService) GetLongUrl(shortUrl string) (string, error) {
	if longUrl, err := s.urlCache.GetLongUrl(shortUrl); err == nil {
		return longUrl, nil
	}

	longUrl, err := s.urlRepository.FindLongFromShort(shortUrl)
	if err != nil {
		return "", err
	}

	err = s.urlCache.Store(shortUrl, longUrl)
	if err != nil {
		return "", err
	}

	return longUrl, nil
}

func (s *UrlService) isLongUrlExists(longUrl string) bool {
	return s.urlRepository.ExistsLong(longUrl)
}

func (s *UrlService) isShortUrlExists(shortUrl string) bool {
	return s.urlRepository.ExistsShort(shortUrl)
}
