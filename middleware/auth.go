package middleware

import (
	"evently/security"
	"evently/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Authenticate(ctx *gin.Context) {
	token := ctx.Request.Header.Get("Authorization")
	if token == "" {
		ctx.AbortWithStatusJSON(http.StatusForbidden, utils.GetResponse("No access token provided", http.StatusForbidden))
		return
	}

	authResponse, err := security.ValidateToken(token)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, utils.GetResponse("Unauthorized", http.StatusUnauthorized))
		return
	}

	ctx.Set("auth", authResponse)
	ctx.Next()
}
