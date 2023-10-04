package main

import (
	"context"
	"database/sql"
	"log"
	"time"

	"github.com/guimassoqueto/go-web-rankings/app"
	"github.com/guimassoqueto/go-web-rankings/website"
)

func main() {
	db, err := sql.Open("pgx", "postgres://postgres:password@localhost:5432/website")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	websiteRepository := website.NewPostgresSQLClassicRepository(db)
	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	defer cancel()

	app.RunRepositoryDemo(ctx, websiteRepository)
}