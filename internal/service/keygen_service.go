package service

import (
	"UrlShortener/internal/config"
	"context"
	"encoding/base64"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type KeygenService struct {
}

type KeygenConfig struct {
	Url string `env:"KGS_URL" envDefault:"localhost:8081"`
}

func GetKeygenConfig() KeygenConfig {
	keygenConfig := KeygenConfig{}
	err := config.LoadConfig(&keygenConfig)
	if err != nil {
		panic(err)
	}
	return keygenConfig
}

func NewGrpcConnection(keygenConfig KeygenConfig) *grpc.ClientConn {
	conn, err := grpc.Dial(keygenConfig.Url, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	return conn
}

func NewKeygenClientServiceFromConfig(keygenConfig KeygenConfig) KgsClient {
	conn := NewGrpcConnection(keygenConfig)
	return NewKgsClient(conn)
}

func NewKeygenServer() *KeygenService {
	return &KeygenService{}
}

func (s *KeygenService) mustEmbedUnimplementedKgsServer() {
	return
}

func (s *KeygenService) GenKey(_ context.Context, request *KeyRequest) (*KeyResponse, error) {
	//TODO implement Bloom filter, better key generation
	longUrl := request.Url
	key := encodeUrl(longUrl)
	shortKey := key[:7]
	return &KeyResponse{Key: shortKey}, nil
}

func encodeUrl(url string) string {
	return base64.StdEncoding.EncodeToString([]byte(url))
}
