package cache

import (
	"errors"

	"github.com/go-redis/redis/v7"
	"github.com/infinitss13/innotaxiuser/configs"
)

var UserSignedOut = errors.New("user have signed-out")

type RedisCache struct {
	Client     *redis.Client
	Connection configs.ConnectionRedis
}

func NewRedisCache() (RedisCache, error) {
	cl, err := NewClientRedis()
	if err != nil {
		return RedisCache{}, err
	}
	connect, err := configs.NewConnectionRedis()
	if err != nil {
		return RedisCache{}, err
	}
	return RedisCache{
		Client:     cl,
		Connection: connect,
	}, nil
}

type Cache interface {
	SetValue(key string, value string) error
	GetValue(key string) (bool, error)
}

func NewClientRedis() (*redis.Client, error) {
	connectionRedis, err := configs.NewConnectionRedis()
	if err != nil {
		return nil, err
	}
	return redis.NewClient(&redis.Options{
		Addr:        connectionRedis.RedisHost,
		DB:          connectionRedis.RedisDB,
		DialTimeout: connectionRedis.RedisExpires,
	}), nil
}

func (cash RedisCache) SetValue(key string, value string) error {
	status := cash.Client.Set(key, value, cash.Connection.RedisExpires)
	if status.Err() != nil {
		return status.Err()
	}
	return nil
}

func (cash RedisCache) GetValue(key string) (bool, error) {
	err := cash.Client.Get(key)
	if err.Err() == redis.Nil {
		return true, UserSignedOut
	} else {
		return false, nil
	}
}
