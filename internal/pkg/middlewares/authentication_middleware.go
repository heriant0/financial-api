package middlewares

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type AccessTokenVerifier interface {
	VerifyAccessToken(tokenString string) (string, error)
}

func AuthenticationMiddleware(tokenMaker AccessTokenVerifier) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		accessToken := tokenFromHeader(ctx)
		if accessToken == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized access"})
			ctx.Abort()
			return
		}

		// verify
		sub, err := tokenMaker.VerifyAccessToken(accessToken)
		if err != nil {
			log.Error(fmt.Errorf("error verify access token %w", err))
			ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized access"})
			ctx.Abort()
			return
		}

		ctx.Set("user_id", sub)

		ctx.Next()
	}
}

func tokenFromHeader(ctx *gin.Context) string {
	var accessToken string

	bearerToken := ctx.Request.Header.Get("Authorization")
	fields := strings.Fields(bearerToken)

	if len(fields) != 0 && fields[0] == "Bearer" {
		accessToken = fields[1]
	}

	return accessToken
}
