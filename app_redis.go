package summerboot

import (
	"context"
	"github.com/guoyk93/rg"
	"github.com/redis/go-redis/v9"
)

func (a *app[T]) setupRedis() {
	var redisURL string

	if redisURL = envOr("REDIS_URL", ""); redisURL == "" {
		return
	}

	a.Component("redis").
		Startup(func(ctx context.Context) (err error) {
			defer rg.Guard(&err)
			opts := rg.Must(redis.ParseURL(redisURL))
			a.redis = redis.NewClient(opts)
			return
		}).
		Check(func(ctx context.Context) error {
			return a.redis.Ping(ctx).Err()
		}).
		Shutdown(func(ctx context.Context) error {
			return a.redis.Close()
		})
}
