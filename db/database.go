package db

import (
	"errors"
	"os"

	"github.com/jinzhu/gorm"
	//Required by GORM
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var connection *gorm.DB

//Connection Opens a database connection based on DATABASE_URL
func Connection() (*gorm.DB, error) {
	var err error

	if connection == nil {
		dbURL := os.Getenv("DATABASE_URL")
		connection, err = gorm.Open("postgres", dbURL)
	}

	return connection, err
}

//TestConnection returns a database connection to the test URL
func TestConnection() (*gorm.DB, error) {
	os.Setenv("DATABASE_URL", os.Getenv("TEST_DATABASE_URL"))
	return Connection()
}

//CloseConnection closes opened connection
func CloseConnection() error {
	if connection == nil {
		return errors.New("Connection is not open")
	}

	err := connection.Close()
	if err == nil {
		connection = nil
	}

	return err
}
