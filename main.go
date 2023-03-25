package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/seeya/student-teacher-assessment/app/controllers"
	"github.com/seeya/student-teacher-assessment/app/queries"
	"github.com/seeya/student-teacher-assessment/platform/database"
)

func main() {
	_ = godotenv.Load(".env")

	db, err := database.OpenDBConnection()

	if err != nil {
		log.Fatal(err)
	}

	ApiQuery := &queries.ApiQuery{DB: db}

	ApiQuery.SeedTeachers()
	ApiQuery.SeedStudents()

	app := fiber.New()
	api := app.Group("/api")

	// Dependency Injection
	controllers.InitApiRoutes(api, ApiQuery)

	err = app.Listen(fmt.Sprintf(":%s", os.Getenv("API_PORT")))

	if err != nil {
		log.Printf("Failed to listen on port %s", os.Getenv("API_PORT"))
	}

}
