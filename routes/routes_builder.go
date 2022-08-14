package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"majoo/controllers"
	"majoo/helper"
	"majoo/middlewares"
)

type Routes struct {
	router *gin.Engine
}

func ProvideRoutes(
	authController controllers.AuthController,
	merchantController controllers.MerchantController,
	outletController controllers.OutletController) *gin.Engine {

	basePath := "/api/v1"

	r := Routes{router: gin.Default()}
	r.router.Use(middlewares.CORSMiddleware())

	r.router.GET(basePath, func(ctx *gin.Context) {
		ctx.JSON(200, "Majoo API")
	})
	//default if url not initialized
	r.router.NoRoute(func(c *gin.Context) {
		c.JSON(404, helper.BuildFailResponse("NOT FOUND", []string{}))
	})

	authRoutes := r.router.Group(fmt.Sprintf("%s/auth", basePath))
	authRoutes.POST("/login", authController.Login)

	merchantRoutes := r.router.Group(fmt.Sprintf("%s/merchants", basePath))
	merchantRoutes.Use(middlewares.AuthorizeJwt())
	{
		merchantRoutes.GET("", merchantController.GetAll)
		merchantRoutes.GET("/:merchantId", merchantController.Get)
		merchantRoutes.GET("/:merchantId/outlets", outletController.GetAll)
		merchantRoutes.GET("/:merchantId/outlets/:outletId", outletController.Get)
	}

	return r.router
}

func (r Routes) Run(addr string) error {
	return r.router.Run(addr)
}
