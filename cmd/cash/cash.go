package cash

import (
	"fmt"
	"github.com/go-redis/redis/v7"
	"github.com/infinitss13/innotaxiuser/configs"
)

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

func (cash *RedisCash) SetValue(key string, value string) {
	status := cash.Client.Set(key, value, cash.Connection.RedisExpires)
	fmt.Println(status)
}

func (cash *RedisCash) GetValue(key string) (bool, error) {
	err := cash.Client.Get(key)
	if err.Err() == redis.Nil {
		return false, err.Err()
	} else {
		return true, nil
	}
}
