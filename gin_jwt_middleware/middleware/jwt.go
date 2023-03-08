package middleware

import (
	"fmt"
	"gjm/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func JWTAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		fmt.Println("jwt auth handler")
		auth := ctx.Request.Header.Get("Authorization")
		fmt.Println("auth", auth)
		if auth == "" {
			ctx.Status(http.StatusForbidden)
			ctx.Abort()
			return
		}
		parts := strings.SplitN(auth, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			ctx.JSON(http.StatusForbidden, gin.H{"code": "1", "message": "请求头auth格式错误"})
			ctx.Abort()
			return
		}
		claims, err := utils.ParseToken(parts[1])
		if err != nil {
			ctx.JSON(http.StatusForbidden, gin.H{"code": "1", "message": err.Error()})
			ctx.Abort()
			return
		}
		fmt.Println("parse success:", claims)
		ctx.Set("authInfo", claims.Info)
		ctx.Next()
	}
}
