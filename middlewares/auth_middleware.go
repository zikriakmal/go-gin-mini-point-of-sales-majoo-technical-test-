package middlewares

import (
	"github.com/gin-gonic/gin"
	"majoo/helper"
	"net/http"
	"strings"
)

func AuthorizeJwt() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			response := helper.BuildFailResponse("No token. Please put your token", helper.EmptyObject{})
			c.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}

		strToken := strings.Split(authHeader, " ")
		validateRes, err := helper.ValidateJwt(strToken[1])

		if err != nil {
			response := helper.BuildFailResponse("Unauthorized", helper.EmptyObject{})
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		c.Set("user_id", int64(validateRes["user_id"].(float64)))
		c.Next()
	}
}
