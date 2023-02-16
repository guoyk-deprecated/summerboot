package summerboot

import (
	"context"
	"github.com/go-resty/resty/v2"
)

func (a *app[T]) setupResty() {
	a.Component("resty").
		Startup(func(ctx context.Context) (err error) {
			a.client = resty.New()
			return
		})
}
