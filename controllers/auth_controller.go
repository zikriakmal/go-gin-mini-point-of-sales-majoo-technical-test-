package controllers

import (
	"github.com/gin-gonic/gin"
	"majoo/dto"
	"majoo/helper"
	"majoo/services"
	"net/http"
)

type AuthController interface {
	Login(ctx *gin.Context)
}

type authController struct {
	authService services.AuthService
}

func NewAuthController(authService services.AuthService) AuthController {
	return &authController{
		authService: authService,
	}
}

func (c *authController) Login(ctx *gin.Context) {
	var loginReqDto dto.LoginReqDto

	errDto := ctx.ShouldBind(&loginReqDto)
	if errDto != nil {
		response := helper.BuildFailResponse("Validation fail", helper.EmptyObject{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	authResult, err := c.authService.Login(loginReqDto.Username, loginReqDto.Password)
	if err != nil {
		response := helper.BuildFailResponse(err.Error(), helper.EmptyObject{})
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	response := helper.BuildSuccessResponse(authResult)
	ctx.JSON(http.StatusOK, response)
	return
}
