package middlewares

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/heriant0/financial-api/internal/pkg/handler"
	log "github.com/sirupsen/logrus"
)

type AccessTokenVerifier interface {
	VerifyAccessToken(tokenString string) (string, error)
}

func AuthenticationMiddleware(tokenMaker AccessTokenVerifier) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		accessToken := tokenFromHeader(ctx)
		if accessToken == "" {
			handler.ResponError(ctx, http.StatusUnauthorized, "Unauthorized access")
			ctx.Abort()
			return
		}

		// verify
		sub, err := tokenMaker.VerifyAccessToken(accessToken)
		if err != nil {
			log.Error(fmt.Errorf("error verify access token %w", err))
			handler.ResponError(ctx, http.StatusUnauthorized, "Unauthorized access")
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
