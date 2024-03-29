package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"majoo/config"
	"majoo/controllers"
	"majoo/helper"
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

	if os.Getenv("ENABLE_SEED") == "true" {
		helper.UserSeed(DbConn)
		helper.MerchantSeed(DbConn)
		helper.OutletSeed(DbConn)
		helper.TransactionSeed(DbConn)
	}

	if err != nil {
		panic(err)
	}
	// defer to close db connection
	defer config.CloseDbConnection(DbConn)

	UserRepository := repositories.NewUserRepository(DbConn)
	AuthService := services.NewAuthService(UserRepository)
	AuthController := controllers.NewAuthController(AuthService)

	MerchantRepository := repositories.NewMerchantRepository(DbConn)
	MerchantService := services.NewMerchantService(MerchantRepository)
	MerchantController := controllers.NewMerchantController(MerchantService)

	OutletRepository := repositories.NewOutletRepository(DbConn)
	OutletService := services.NewOutletService(OutletRepository)
	OutletController := controllers.NewOutletController(OutletService, MerchantService)

	r := routes.ProvideRoutes(AuthController, MerchantController, OutletController)
	r.Static("/public", "./public")

	err = r.Run(fmt.Sprintf(":%s", os.Getenv("APP_PORT")))

	if err != nil {
		panic("error to run app")
	}

}
