package api

import (
	"fmt"

	"github.com/srivardhanreddy01/webapplication_go/api/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDB() *gorm.DB {
	dsn := "root:DBpassword@tcp(127.0.0.1:3306)/godatabase?parseTime=true"

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
