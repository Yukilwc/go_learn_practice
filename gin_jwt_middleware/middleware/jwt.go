package middleware

import (
	"errors"
	"fmt"
	"gjm/model/response"
	"gjm/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func JWTAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		fmt.Println("jwt auth handler")
		auth := ctx.Request.Header.Get("Authorization")
		fmt.Println("auth", auth)
		if auth == "" {
			response.FailForbidden(nil, "", ctx)
			ctx.Abort()
			return
		}
		parts := strings.SplitN(auth, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			response.FailForbidden(nil, "请求头auth格式错误", ctx)
			ctx.Abort()
			return
		}
		claims, err := utils.ParseToken(parts[1])
		if err != nil {
			if errors.Is(err, jwt.ErrTokenExpired) {
				fmt.Println("token过期")
			}
			// 判断下过期
			// response.FailForbidden(nil, err.Error(), ctx)
			ctx.Redirect(http.StatusMovedPermanently, "/home")
			ctx.Abort()
			return
		}
		fmt.Println("parse success:", claims)
		ctx.Set("authInfo", claims.Info)
		ctx.Next()
	}
}
