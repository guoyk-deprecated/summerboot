package summerboot

import (
	"github.com/go-resty/resty/v2"
	"github.com/guoyk93/summer"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type App[T Context] interface {
	summer.App[T]
}

type app[T Context] struct {
	summer.App[T]

	redis  *redis.Client
	db     *gorm.DB
	client *resty.Client
}

func New[T Context](cf summer.ContextFactory[T], opts ...summer.Option) App[T] {
	a := &app[T]{}
	a.App = summer.New(cf, opts...)

	a.setupRedis()
	a.setupDB()
	a.setupResty()

	return a
}
