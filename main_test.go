package main

import (
	"io/ioutil"
	"log"
	"os"
	"strings"
	"testing"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"github.com/seeya/student-teacher-assessment/app/queries"
	"github.com/seeya/student-teacher-assessment/platform/database"
)

func TestMain(m *testing.M) {
	_ = godotenv.Load(".env")
	db, err := database.OpenTestDBConnection()

	if err != nil {
		log.Fatal(err)
	}

	ApiQuery := &queries.ApiQuery{DB: db}

	executeSQLFile(db, "./platform/migrations/000001_create_init_tables.down.sql")
	executeSQLFile(db, "./platform/migrations/000001_create_init_tables.up.sql")

	if err != nil {
		log.Println("Error creating test database: ", err)
	}

	ApiQuery.SeedTeachers()
	ApiQuery.SeedStudents()
	ApiQuery.SeedTeacherStudentRelationship()

	// Run the tests
	os.Exit(m.Run())
}

func executeSQLFile(db *sqlx.DB, filepath string) error {
	content, err := ioutil.ReadFile(filepath)
	if err != nil {
		return err
	}

	queries := strings.Split(string(content), ";")

	for _, query := range queries {
		query = strings.TrimSpace(query)
		if query == "" {
			continue
		}

		_, err = db.Exec(query)
		if err != nil {
			return err
		}
	}

	return nil
}
