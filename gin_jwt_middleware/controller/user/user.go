package user

import (
	"errors"
	"fmt"
	"gjm/global"
	dbModel "gjm/model/db"
	userModel "gjm/model/request/user"
	systemModel "gjm/model/system"
	"gjm/utils"
	"log"
	"net/http"
	"strings"
	"time"

	resModel "gjm/model/response"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func LoginController(ctx *gin.Context) {
	var req userModel.Login
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println("登录信息:", req.UserName)
	if req.UserName == "admin" && req.Password == "123456" {
		fmt.Println("账号密码验证正确")
		const TokenExpireDuration = time.Hour * 24 * 7
		auth, err := utils.GetToken(systemModel.TokenInfo{Name: req.UserName, Id: 1}, TokenExpireDuration)
		if err == nil {
			data := map[string]any{
				"code": 0,
				"msg":  "success",
				"data": map[string]any{
					"name":          req.UserName,
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
	// TODO:很多逻辑应该迁移到services中
	fmt.Println("in  RegisterController")
	var register userModel.Register
	err := ctx.ShouldBindJSON(&register)
	if err != nil {
		fmt.Println("绑定错误:", register)
		resModel.FailWithMessage("参数绑定错误"+err.Error(), ctx)
		return
	}
	fmt.Println("注册:", register)
	// TODO:改成基于tag的校验
	// 校验非空
	if strings.TrimSpace(register.UserName) == "" {
		resModel.FailWithMessage("参数不能为空", ctx)
		return
	}
	if strings.TrimSpace(register.NickName) == "" {
		resModel.FailWithMessage("参数不能为空", ctx)
		return
	}
	if strings.TrimSpace(register.Email) == "" {
		resModel.FailWithMessage("参数不能为空", ctx)
		return
	}
	if strings.TrimSpace(register.Password) == "" {
		resModel.FailWithMessage("参数不能为空", ctx)
		return
	}
	// 校验重复
	findUser := dbModel.User{}
	result := global.DB.Where(&dbModel.User{NickName: register.NickName}).Take(&findUser)
	if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		log.Println("存在同名的:", findUser)
		resModel.FailWithMessage("用户名已存在", ctx)
		return
	}
	result = global.DB.Where(&dbModel.User{Email: register.Email}).Take(&findUser)
	if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		log.Println("存在同邮箱的:", findUser)
		resModel.FailWithMessage("邮箱已存在", ctx)
		return
	}
	newUser := dbModel.User{
		NickName: register.NickName,
		Password: register.Password,
		UserName: register.UserName,
		Email:    register.Email,
	}
	err = global.DB.Create(&newUser).Error
	if err != nil {
		resModel.FailWithMessage(err.Error(), ctx)
	} else {
		resModel.OkWithDetailed(map[string]any{
			"id": newUser.ID,
		}, "注册成功", ctx)
	}

}
func ListController(ctx *gin.Context) {
	info := ctx.MustGet("authInfo").(systemModel.TokenInfo)
	data := map[string]any{
		"list": []any{},
		"info": info,
	}

	ctx.JSON(http.StatusOK, gin.H{"code": 1, "msg": "", "data": data})
}
