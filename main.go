package main

import (
	"encoding/csv"
	"log"
	"net/http"
	"os"

	"github.com/srivardhanreddy01/webapplication_go/api"
	"github.com/srivardhanreddy01/webapplication_go/api/models"
	endpoint "github.com/srivardhanreddy01/webapplication_go/cmd/endpoint"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func main() {
	db := api.InitDB()
	csvFile, err := os.Open("opt/user.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer csvFile.Close()

	csvReader := csv.NewReader(csvFile)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	users := []models.User{}

	for _, record := range records {
		firstname := record[0]
		lastname := record[1]
		email := record[2]
		password := record[3]

		var existingUser models.User
		result := db.Where("Email = ?", email).First(&existingUser)

		if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
			log.Fatal(result.Error)
		} else if result.Error == gorm.ErrRecordNotFound {

			hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
			if err != nil {
				log.Fatal(err)
			}

			// currentTime := time.Now()

			user := models.User{
				FirstName:    firstname,
				LastName:     lastname,
				Email:        email,
				HashPassword: string(hashedPassword),
				// AccountCreated: currentTime,
				// AccountUpdated: currentTime,
			}

			users = append(users, user)
		}

	}

	for _, user := range users {
		db.Create(&user)
	}

	log.Println("Starting the API server...")
	err = http.ListenAndServe(":8081", endpoint.NewRouter(db))
	if err != nil {
		log.Fatal("Server error:", err)
	}
}
