package user

import (
	"fmt"
	userModel "gjm/model/request/user"
	systemModel "gjm/model/system"
	"gjm/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func LoginController(ctx *gin.Context) {
	var req userModel.LoginRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println("登录信息:", req.Name)
	if req.Name == "admin" && req.Password == "123456" {
		fmt.Println("账号密码验证正确")
		const TokenExpireDuration = time.Hour * 24 * 7
		auth, err := utils.GetToken(systemModel.TokenInfo{Name: req.Name, Id: 1}, TokenExpireDuration)
		if err == nil {
			data := map[string]any{
				"code": 0,
				"msg":  "success",
				"data": map[string]any{
					"name":          req.Name,
					"authorization": auth,
				},
			}
			ctx.JSON(http.StatusOK, data)

		} else {
			ctx.JSON(http.StatusOK, gin.H{"code": 1, "msg": err.Error(), "data": nil})
		}
	} else {
		ctx.JSON(http.StatusOK, gin.H{"code": 1, "msg": "密码错误", "data": nil})
	}
}
func RegisterController(ctx *gin.Context) {

}
func ListController(ctx *gin.Context) {
	info := ctx.MustGet("authInfo").(systemModel.TokenInfo)
	data := map[string]any{
		"list": []any{},
		"info": info,
	}

	ctx.JSON(http.StatusOK, gin.H{"code": 1, "msg": "", "data": data})
}
