package main

import (
	"context"
	"log"
	"time"

	"github.com/guimassoqueto/go-web-rankings/app"
	"github.com/guimassoqueto/go-web-rankings/website"
	"github.com/jackc/pgx/v4/pgxpool"
)

func main() {
	dbPool, err := pgxpool.Connect(context.Background(), "postgres://postgres:password@localhost:5432/website?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer dbPool.Close()

	websiteRepository := website.NewPostgresSQLPGXRepository(dbPool)
	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	defer cancel()

	app.RunRepositoryDemo(ctx, websiteRepository)
}