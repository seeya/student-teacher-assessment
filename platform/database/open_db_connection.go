package database

import (
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func OpenDBConnection() (*sqlx.DB, error) {
	// connect to mysql database using sqlx
	// parseTime=true is required to parse mysql DATETIME to go time.Time
	db, err := sqlx.Connect("mysql",
		fmt.Sprintf("root:%s@tcp(%s:%s)/%s?parseTime=true",
			os.Getenv("MYSQL_ROOT_PASSWORD"),
			os.Getenv("MYSQL_HOST"),
			os.Getenv("MYSQL_PORT"),
			os.Getenv("MYSQL_DATABASE")))

	if err != nil {
		return nil, err
	}

	fmt.Printf("Connected to database: %v", db)

	return db, nil

}
