package main

import (
	"context"
	"log"
	"time"

	"github.com/guimassoqueto/go-web-rankings/app"
	"github.com/guimassoqueto/go-web-rankings/website"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	gormDb, err := gorm.Open(postgres.Open("postgres://postgres:password@localhost:5432/website?sslmode=disable"))
	if err != nil {
		log.Fatal(err)
	}
	websiteRepository := website.NewPostgresSQLGORMRepository(gormDb)

	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	defer cancel()

	app.RunRepositoryDemo(ctx, websiteRepository)
}