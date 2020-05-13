package core

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
)

func NewSessionStore(cfg *SessionConfig) (sessions.Store,error) {
	switch cfg.Driver {
	case "redis":
		store,err := redis.NewStore(10, "tcp", fmt.Sprintf("%s:%d",cfg.Host,cfg.Port), cfg.Password, []byte(cfg.KeyPairs))
		store.Options(sessions.Options{
			Path:     "/",
			Domain:   "",
			MaxAge:   30 * 60,
			Secure:   false,
			HttpOnly: true,
		})

		return store,err
	default:

	}

	return nil,fmt.Errorf("invalid session config")
}
