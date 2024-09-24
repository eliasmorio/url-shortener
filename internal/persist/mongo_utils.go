package persist

import (
	"UrlShortener/internal/config"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoConfig struct {
	Url      string `env:"MONGO_URL" envDefault:"mongodb://localhost:27017"`
	Database string `env:"MONGO_DB" envDefault:"urlShortener"`
	Username string `env:"MONGO_USER"`
	Password string `env:"MONGO_PASS"`
}

func GetMongoConfig() MongoConfig {
	datasourceConfig := MongoConfig{}
	err := config.LoadConfig(&datasourceConfig)
	if err != nil {
		panic(err)
	}
	return datasourceConfig
}

func GetMongoClient(datasourceConfig MongoConfig) *mongo.Client {
	clientOptions := buildOptions(&datasourceConfig)
	client, err := mongo.Connect(nil, clientOptions)
	if err != nil {
		panic(err)
	}
	err = client.Ping(context.Background(), nil)
	if err != nil {
		panic(err)
	}
	return client
}

func buildOptions(config *MongoConfig) *options.ClientOptions {
	opts := options.Client().ApplyURI(config.Url)
	if config.Username != "" && config.Password != "" {
		opts.SetAuth(options.Credential{
			AuthMechanism: "SCRAM-SHA-256",
			Username:      config.Username,
			Password:      config.Password,
		})
	}
	return opts
}
