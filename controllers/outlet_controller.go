package controllers

import (
	"github.com/gin-gonic/gin"
	"majoo/dto"
	"majoo/helper"
	"majoo/services"
	"net/http"
)

type OutletController interface {
	GetAll(ctx *gin.Context)
	Get(ctx *gin.Context)
}

type outletController struct {
	outletService   services.OutletService
	merchantService services.MerchantService
}

func NewOutletController(OutletService services.OutletService, merchantService services.MerchantService) OutletController {
	return &outletController{
		outletService:   OutletService,
		merchantService: merchantService,
	}
}

func (c *outletController) GetAll(ctx *gin.Context) {
	var request dto.OutletsReqDto
	err := ctx.BindQuery(&request)
	err = ctx.ShouldBindUri(&request)
	userId := ctx.MustGet("user_id").(int64)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, nil)
		return
	}
	_, err = c.merchantService.Get(userId, dto.MerchantReqDto{MerchantId: request.MerchantId})

	if err != nil {
		if err.Error() == "record not found" {
			ctx.AbortWithStatusJSON(http.StatusNotFound, helper.BuildFailResponse("NOT FOUND", []string{}))
			return
		}
	}

	outlets, err := c.outletService.GetAll(request)
	if err != nil {
		if err.Error() == "record not found" {
			ctx.AbortWithStatusJSON(http.StatusNotFound, helper.BuildFailResponse("NOT FOUND", []string{}))
			return
		}
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, nil)
		return
	}
	ctx.JSON(http.StatusOK, helper.BuildSuccessResponse(outlets))
}

func (c *outletController) Get(ctx *gin.Context) {
	var request dto.OutletReqDto
	userId := ctx.MustGet("user_id").(int64)
	err := ctx.ShouldBindUri(&request)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, helper.BuildFailResponse("NOT FOUND", []string{}))
		return
	}

	_, err = c.merchantService.Get(userId, dto.MerchantReqDto{MerchantId: request.MerchantId})
	if err != nil {
		if err.Error() == "record not found" {
			ctx.AbortWithStatusJSON(http.StatusNotFound, helper.BuildFailResponse("NOT FOUND", []string{}))
			return
		}
	}
	outlet, err := c.outletService.Get(request)
	if err != nil {
		if err.Error() == "record not found" {
			ctx.AbortWithStatusJSON(http.StatusNotFound, helper.BuildFailResponse("NOT FOUND", []string{}))
			return
		}
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, []string{})
		return
	}
	ctx.JSON(http.StatusOK, helper.BuildSuccessResponse(outlet))
}
