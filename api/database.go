package api

import (
	"fmt"
	"os"

	"github.com/srivardhanreddy01/webapplication_go/api/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDB() *gorm.DB {
	// Read environment variables
	dbHost := os.Getenv("DB_HOST")
	dbUserName := os.Getenv("DB_HOST_NAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	// Construct DSN using environment variables
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", dbUserName, dbPassword, dbHost, dbPort, dbName)


	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("failed to connect to the database")
	}

	initialMigration()
	return db
}

func initialMigration() {
	modelsToMigrate := []interface{}{
		&models.User{},
		&models.Assignment{},
	}

	for _, model := range modelsToMigrate {
		if err := db.AutoMigrate(model); err != nil {
			fmt.Println(err.Error())
			panic("failed to migrate the schema")
		}
	}
}
