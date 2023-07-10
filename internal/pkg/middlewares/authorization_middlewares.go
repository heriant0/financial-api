package middlewares

import (
	"net/http"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
)

func AuthorizationMiddleware(
	enforcer *casbin.Enforcer,
	sub string,
	obj string,
	act string,
) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ok, err := enforcer.Enforce(sub, obj, act)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"message": "you cannot perform"})
			ctx.Abort()
		}

		if ok {
			ctx.Next()
		} else {
			ctx.JSON(http.StatusUnauthorized, gin.H{"message": "you cannot perform"})
			ctx.Abort()
		}
	}
}
