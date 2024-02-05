package pkg1

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"proj_struct/configs" // Adjust the import path

)

func LoadEnv() error {
	if err := godotenv.Load(); err != nil {
		return fmt.Errorf("Error loading .env file")
	}
	return nil
}

func SetupDB() (*sql.DB, error) {
	if err := LoadEnv(); err != nil {
		return nil, err
	}

	// Load database configuration
	dbConfig := configs.LoadDBConfig()

	// Build DSN (Data Source Name)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.Name,
	)

	// Open a database connection
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println("Error connecting to the database:", err)
		os.Exit(1)
	}

	// Test the connection
	err = db.Ping()
	if err != nil {
		fmt.Println("Error pinging the database:", err)
		os.Exit(1)
	}

	fmt.Println("Connected to the database")

	return db, nil
}
