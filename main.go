package main

import (
	"context"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/jubolanos/go-inventory-project/database"
	"github.com/jubolanos/go-inventory-project/settings"
	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		fx.Provide(
			context.Background,
			settings.New,
			database.New,
		),
		fx.Invoke(
			func(db *sqlx.DB) {
				_, err := db.Query("select * from USERS")
				if err != nil {
					panic(err)
				}
			},
		),
	)
	app.Run()
}
