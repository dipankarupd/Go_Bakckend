package config

// this deals with connection with the sql database

import (
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

func Connect() {

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")

	// Create the connection string
	connectionString := dbUser + ":" + dbPassword + "@(" + dbHost + ")/" + dbName + "?charset=utf8&parseTime=True&loc=Local"

	// Open a database connection
	d, err := gorm.Open("mysql", connectionString)
	if err != nil {
		panic(err)
	}

	// Assign the connection to the global variable
	db = d
}
