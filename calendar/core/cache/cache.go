package cache

import (
	"errors"
	"time"
)

type Driver interface {
	Get(key string, v interface{}) error
	Set(key string, v interface{}, expire time.Duration) error
}

var ErrDriver = errors.New("empty driver")

var cacheDriver Driver

func RegisterDriver(driver Driver) {
	cacheDriver = driver
}

func Get(key string, v interface{}) error {
	if cacheDriver == nil {
		return ErrDriver
	}

	return cacheDriver.Get(key, v)
}

func Set(key string, v interface{}, expire time.Duration) error {
	if cacheDriver == nil {
		return ErrDriver
	}

	return cacheDriver.Set(key, v, expire)
}
