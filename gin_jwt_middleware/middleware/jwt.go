package middleware

import (
	"errors"
	"fmt"
	"gjm/global"
	"gjm/model/response"
	"gjm/utils"
	"net/http"
	"strings"
	"time"

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
		if err == nil {
			fmt.Println("parse success:", claims)
			ctx.Set("authInfo", claims.Info)
			ctx.Next()
			return
		}
		// 判断下过期
		if errors.Is(err, jwt.ErrTokenExpired) {
			// 在buffer time期间内，进行刷新，实现高频率操作用户能无感知刷新，需要前端逻辑配合
			bufferTime := time.Duration(global.CONFIG.JWT.BufferTime) * time.Second
			nowTime := time.Now()
			// expiredTime := .Unix()
			// fmt.Println("三个unix时间/时间段:", nowTime, expiredTime, bufferTime)
			finalTime := claims.ExpiresAt.Add(bufferTime)
			fmt.Println("buffer time", bufferTime.Seconds())
			fmt.Println("token过期时间,", claims.ExpiresAt.String())
			fmt.Println("token最后刷新截止时间,", finalTime.String())
			fmt.Println("当前时间,", nowTime.String())
			if finalTime.Before(nowTime) {
				fmt.Println("已经超出刷新时间")
				ctx.Redirect(http.StatusMovedPermanently, "/home")
				ctx.Abort()
				return

			} else {
				fmt.Println("在刷新时间内")
				fmt.Println("parse success:", claims)
				// 生成一个新的token
				auth, err := utils.GetToken(claims.Info, global.CONFIG.JWT.ExpiresTime)
				if err != nil {
					response.FailWithMessage(err.Error(), ctx)
					ctx.Abort()
					return
				}
				ctx.Header("new-token", auth)
				ctx.Set("authInfo", claims.Info)
				ctx.Next()
				return
			}
		} else {
			// 过期以外的错误
			ctx.Redirect(http.StatusMovedPermanently, "/home")
			ctx.Abort()
			return
		}
	}
}
