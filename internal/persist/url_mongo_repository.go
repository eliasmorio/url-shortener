package persist

import (
	"UrlShortener/internal/model"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	ShortenedUrlsCollection = "urls"
)

type UrlRepository interface {
	Store(shortenedUrl *model.ShortenedUrl) error
	FindLongFromShort(shortUrl string) (string, error)
	ExistsLong(longUrl string) bool
	ExistsShort(shortUrl string) bool
}

type urlMongoRepository struct {
	collection *mongo.Collection
}

func NewUrlRepositoryWCollection(collection *mongo.Collection) UrlRepository {
	return &urlMongoRepository{
		collection: collection,
	}
}

func NewUrlRepositoryFromConfig(config MongoConfig) UrlRepository {
	client := GetMongoClient(config)
	collection := client.Database(config.Database).Collection(ShortenedUrlsCollection)
	return NewUrlRepositoryWCollection(collection)
}

type shortenedUrlDocument struct {
	LongUrl  string `bson:"longUrl"`
	ShortUrl string `bson:"shortUrl"`
}

func mapToDocument(shortenedUrl *model.ShortenedUrl) *shortenedUrlDocument {
	return &shortenedUrlDocument{
		LongUrl:  shortenedUrl.LongUrl,
		ShortUrl: shortenedUrl.ShortUrl,
	}
}

func (r *urlMongoRepository) Store(shortenedUrl *model.ShortenedUrl) error {
	row := mapToDocument(shortenedUrl)
	_, err := r.collection.InsertOne(nil, row)
	return err
}

func (r *urlMongoRepository) FindLongFromShort(shortUrl string) (string, error) {
	var result shortenedUrlDocument
	err := r.collection.FindOne(nil, map[string]string{"shortUrl": shortUrl}).Decode(&result)
	if err != nil {
		return "", err
	}
	return result.LongUrl, nil
}

func (r *urlMongoRepository) ExistsLong(longUrl string) bool {
	return r.collection.FindOne(nil, map[string]string{"longUrl": longUrl}).Err() == nil
}

func (r *urlMongoRepository) ExistsShort(shortUrl string) bool {
	return r.collection.FindOne(nil, map[string]string{"shortUrl": shortUrl}).Err() == nil
}
