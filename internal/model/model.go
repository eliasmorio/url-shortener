package model

type ShortenedUrl struct {
	LongUrl  string
	ShortUrl string
	ExpireAt int64
}
