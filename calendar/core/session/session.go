package session

import (
	"crypto/md5"
	"fmt"
	"calendar/core/cache"
	"calendar/core/opt"
	"calendar/core/uuid"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	DefaultKey = "nob_session"
)

type Values map[interface{}]interface{}

type session struct {
	id     string
	values Values
	ctx    *gin.Context
}

type Session interface {
	Get(key interface{}) interface{}
	Set(key interface{}, val interface{})
	Delete(key interface{})
	Clear()
	Save() error
	Create() string
}

var idFactory *uuid.IDFactory

func New(name string) gin.HandlerFunc {
	var err error
	idFactory, err = uuid.New(int64(opt.Config().Node))
	if err != nil {
		panic(fmt.Sprintf("calendar err: %s", err.Error()))
	}

	return func(c *gin.Context) {
		sess := &session{
			id:     c.Query(name),
			values: make(Values),
			ctx:    c,
		}
		err = cache.Get(sess.key(), &sess.values)
		if err != nil {
			sess.id = ""
		}

		c.Set(DefaultKey, sess)
		c.Next()
	}
}

func (s *session) key() string {
	if s.id == "" {
		return ""
	}
	return fmt.Sprintf("nobsess_%s", s.id)
}

func (s *session) Get(key interface{}) interface{} {
	return s.values[key]
}

func (s *session) Set(key interface{}, val interface{}) {
	s.values[key] = val
}

func (s *session) Delete(key interface{}) {
	delete(s.values, key)
}

func (s *session) Clear() {
	for k := range s.values {
		s.Delete(k)
	}
}

func (s *session) Save() error {
	if s.key() != "" {
		return cache.Set(s.key(), s.values, 2*time.Hour)
	}
	return nil
}

func (s *session) Create() string {
	s.id = fmt.Sprintf("%x", md5.Sum([]byte(idFactory.String())))
	return s.id
}

func Default(c *gin.Context) Session {
	return c.MustGet(DefaultKey).(Session)
}
