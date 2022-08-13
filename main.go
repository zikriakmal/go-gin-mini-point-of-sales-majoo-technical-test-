package main

import (
	"github.com/joho/godotenv"
	"log"
	"majoo/config"
	"majoo/models"
)

func main() {

	// load env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	// init db
	DbConn := config.InitDatabase()

	err = DbConn.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&models.User{}, &models.Merchant{}, &models.Transaction{}, &models.Outlet{})

	// Add table suffix when creating tables

	if err != nil {
		panic(err)
	}

	// close db when all finish
	defer config.CloseDbConnection(DbConn)

	// create jwt instance

	// defer to close db connection

	// define port

	// execute gin run

}
