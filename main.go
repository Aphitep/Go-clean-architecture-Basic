package main

import (
	"fmt"
	"go-clean-arc/controller"
	"go-clean-arc/entities"
	"go-clean-arc/usecases"
	"log"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host     = "localhost"  // or the Docker service name if running in another container
	port     = 5432         // default PostgreSQL port
	user     = "myuser"     // as defined in docker-compose.yml
	password = "mypassword" // as defined in docker-compose.yml
	dbname   = "mydatabase" // as defined in docker-compose.yml
)

func main() {
	app := fiber.New()
	dsn := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Print("Error Something wrong")
	}

	if err := db.AutoMigrate(&entities.Book{}); err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	bookRepo := controller.NewBookDb(db)
	bookUsecase := usecases.NewBookService(bookRepo)
	bookHttp := controller.NewBookHttpHandler(bookUsecase)

	app.Post("/book", bookHttp.CreateBook)
	app.Get("/book", bookHttp.FirstBook)
	app.Listen(":8080")
}
