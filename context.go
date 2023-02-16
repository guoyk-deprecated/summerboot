package summerboot

import (
	"github.com/go-resty/resty/v2"
	"github.com/guoyk93/summer"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"net/http"
)

type Context interface {
	summer.Context

	R() *resty.Request
	RDB() *redis.Client
	DB() *gorm.DB
}

type bootContext struct {
	summer.Context

	client *resty.Client
	db     *gorm.DB
	redis  *redis.Client
}

func (c *bootContext) RDB() *redis.Client {
	return c.redis
}

func (c *bootContext) DB() *gorm.DB {
	return c.db.WithContext(c)
}

func (c *bootContext) R() *resty.Request {
	return c.client.R().SetContext(c)
}

func (a *app[T]) CreateContext(rw http.ResponseWriter, req *http.Request) Context {
	c := &bootContext{
		Context: summer.BasicContext(rw, req),

		client: a.client,
		db:     a.db,
		redis:  a.redis,
	}
	return c
}
