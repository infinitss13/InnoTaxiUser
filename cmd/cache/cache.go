package cache

import (
	"errors"

	"github.com/go-redis/redis/v7"
	"github.com/infinitss13/innotaxiuser/configs"
)

var UserSignedOut error = errors.New("user have signed-out")

type RedisCash struct {
	Client     *redis.Client
	Connection configs.ConnectionRedis
}

func NewRedisCash() (RedisCash, error) {
	cl, err := NewClientRedis()
	if err != nil {
		return RedisCash{}, err
	}
	connect, err := configs.NewConnectionRedis()
	if err != nil {
		return RedisCash{}, err
	}
	return RedisCash{
		Client:     cl,
		Connection: connect,
	}, nil
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

func (cash *RedisCash) SetValue(key string, value string) error {
	status := cash.Client.Set(key, value, cash.Connection.RedisExpires)
	if status.Err() != nil {
		return status.Err()
	}
	return nil
}

func (cash *RedisCash) GetValue(key string) (bool, error) {
	err := cash.Client.Get(key)
	if err.Err() == redis.Nil {
		return false, UserSignedOut
	} else {
		return true, nil
	}
}
