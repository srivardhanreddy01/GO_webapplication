package handlers

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func isDatabaseHealthy() (bool, error) {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	dbHost := os.Getenv("DB_HOST")
	dbUserName := os.Getenv("DB_HOST_NAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	// Construct DSN using environment variables

	// Display the original and escaped passwords
	fmt.Println("Original Password:", dbPassword)

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", dbUserName, dbPassword, dbHost, dbPort, dbName)

	//dsn := "root:Sripragna$1@tcp(127.0.0.1:3306)/godatabase?parseTime=true"

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return false, err
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		fmt.Println("Error pinging the database:", err)
		return false, nil
	}

	return true, nil
}

func HealthzHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	w.Header().Set("Pragma", "no-cache")

	databaseIsHealthy, err := isDatabaseHealthy()

	if err != nil {
		fmt.Println("Error checking database health:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if databaseIsHealthy {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusServiceUnavailable)
	}
}
