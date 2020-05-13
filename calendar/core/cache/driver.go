package cache

import (
	"bytes"
	"encoding/gob"
	"time"

	"github.com/go-redis/redis"
)

type RedisDriver struct {
	client *redis.Client
}

func NewRedisDriver(client *redis.Client) *RedisDriver {
	return &RedisDriver{
		client: client,
	}
}

func (r *RedisDriver) Get(key string, v interface{}) error {
	buf, err := r.client.Get(key).Bytes()
	if err != nil {
		return err
	}

	reader := bytes.NewBuffer(buf)
	dec := gob.NewDecoder(reader)
	err = dec.Decode(v)
	return err
}

func (r *RedisDriver) Set(key string, v interface{}, exp time.Duration) error {
	var writer bytes.Buffer

	enc := gob.NewEncoder(&writer)
	err := enc.Encode(v)
	if err != nil {
		return err
	}

	return r.client.Set(key, writer.Bytes(), exp).Err()
}
