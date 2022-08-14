package controllers

import (
	"github.com/gin-gonic/gin"
	"majoo/dto"
	"majoo/services"
	"net/http"
)

type OutletController interface {
	GetAll(ctx *gin.Context)
	Get(ctx *gin.Context)
}

type outletController struct {
	OutletService services.OutletService
}

func NewOutletController(OutletService services.OutletService) OutletController {
	return &outletController{
		OutletService: OutletService,
	}
}

func (c *outletController) GetAll(ctx *gin.Context) {
	var request dto.OutletsReqDto
	err := ctx.BindQuery(&request)
	err = ctx.ShouldBindUri(&request)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, nil)
		return
	}
	outlets, err := c.OutletService.GetAll(request)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, nil)
		return
	}
	ctx.JSON(http.StatusOK, outlets)
}

func (c *outletController) Get(ctx *gin.Context) {
	var request dto.OutletReqDto
	err := ctx.ShouldBindUri(&request)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, nil)
		return
	}
	outlets, err := c.OutletService.Get(request)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, nil)
		return
	}
	ctx.JSON(http.StatusOK, outlets)
}
