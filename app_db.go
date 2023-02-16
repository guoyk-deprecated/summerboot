package summerboot

import (
	"context"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
	"strings"
)

func (a *app[T]) setupDB() {
	var dsn string
	if dsn = strings.TrimSpace(os.Getenv("MYSQL_DSN")); dsn == "" {
		return
	}

	a.Component("db").
		Startup(func(ctx context.Context) (err error) {
			a.db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
			return
		}).
		Check(func(ctx context.Context) error {
			return a.db.WithContext(ctx).Select("SELECT 1").Error
		})
}
