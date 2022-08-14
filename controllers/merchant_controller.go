package controllers

import (
	"github.com/gin-gonic/gin"
	"majoo/dto"
	"majoo/helper"
	"majoo/services"
	"net/http"
)

type MerchantController interface {
	GetAll(ctx *gin.Context)
	Get(ctx *gin.Context)
}

type merchantController struct {
	merchantService services.MerchantService
}

func NewMerchantController(merchantService services.MerchantService) MerchantController {
	return &merchantController{
		merchantService: merchantService,
	}
}

func (c *merchantController) GetAll(ctx *gin.Context) {
	var request dto.MerchantsReqDto
	err := ctx.BindQuery(&request)
	err = ctx.ShouldBindUri(&request)

	userId := ctx.MustGet("user_id").(int64)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, helper.BuildFailResponse("NOT FOUND", []string{}))
		return
	}
	merchants, err := c.merchantService.GetAll(userId, request)

	if err != nil {
		if err.Error() == "record not found" {
			ctx.AbortWithStatusJSON(http.StatusNotFound, helper.BuildFailResponse("NOT FOUND", []string{}))
			return
		}
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, nil)
		return
	}
	ctx.JSON(http.StatusOK, helper.BuildSuccessResponse(merchants))
}

func (c *merchantController) Get(ctx *gin.Context) {
	var request dto.MerchantReqDto
	err := ctx.ShouldBindUri(&request)
	userId := ctx.MustGet("user_id").(int64)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, nil)
		return
	}

	merchant, err := c.merchantService.Get(userId, request)

	if err != nil {
		if err.Error() == "record not found" {
			ctx.AbortWithStatusJSON(http.StatusNotFound, helper.BuildFailResponse("NOT FOUND", []string{}))
			return
		}
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, nil)
		return
	}
	ctx.JSON(http.StatusOK, helper.BuildSuccessResponse(merchant))
}
