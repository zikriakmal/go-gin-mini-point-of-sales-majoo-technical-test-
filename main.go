package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"majoo/config"
	"majoo/controllers"
	"majoo/models"
	"majoo/repositories"
	"majoo/routes"
	"majoo/services"
	"os"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		panic("cannot load env")
	}

	// init db
	DbConn := config.InitDatabase()

	err = DbConn.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(
		&models.User{},
		&models.Merchant{},
		&models.Transaction{},
		&models.Outlet{},
	)

	if err != nil {
		panic(err)
	}
	// defer to close db connection
	defer config.CloseDbConnection(DbConn)

	UserRepository := repositories.NewUserRepository(DbConn)
	AuthService := services.NewAuthService(UserRepository)
	AuthController := controllers.NewAuthController(AuthService)

	OutletRepository := repositories.NewOutletRepository(DbConn)
	OutletService := services.NewOutletService(OutletRepository)
	OutletController := controllers.NewOutletController(OutletService)

	r := routes.ProvideRoutes(AuthController, OutletController)

	err = r.Run(fmt.Sprintf(":%s", os.Getenv("APP_PORT")))

	if err != nil {
		panic("error to run app")
	}

}
