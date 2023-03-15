package user

import (
	"errors"
	"fmt"
	"gjm/global"
	dbModel "gjm/model/db"
	userModel "gjm/model/request/user"
	userResModel "gjm/model/response/user"
	systemModel "gjm/model/system"
	"gjm/utils"
	"log"
	"regexp"

	resModel "gjm/model/response"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/mitchellh/mapstructure"
	"gorm.io/gorm"
)

func myPasswordValidator(fl validator.FieldLevel) bool {
	ps := fl.Field().String()
	re := regexp.MustCompile(`^.{6,}$`)
	return re.MatchString(ps)
}
func LoginController(ctx *gin.Context) {
	var req userModel.Login
	jsonMap := utils.PostJson2Map(ctx)
	// 进行一些特殊处理
	if err := mapstructure.Decode(jsonMap, &req); err != nil {
		resModel.FailWithMessage(err.Error(), ctx)
	}
	fmt.Println("登录参数:", jsonMap, req)
	//  validator
	validate := validator.New()
	validate.RegisterValidation("myPassword", myPasswordValidator)
	err := validate.Struct(req)
	if err != nil {
		fmt.Println("validate error:", err.Error())
		if _, ok := err.(*validator.InvalidValidationError); ok {
			resModel.FailWithMessage(err.Error(), ctx)
			return
		}
		err := err.(validator.ValidationErrors)[0]
		switch err.Tag() {
		case "required":
			resModel.FailWithMessage(err.Field()+"字段不能为空", ctx)
		default:
			resModel.FailWithMessage(err.Error(), ctx)
		}
		return
	}
	findUser := dbModel.User{}
	result := global.DB.Where(&dbModel.User{UserName: req.UserName}).Take(&findUser)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		resModel.FailWithMessage("用户名不存在", ctx)
		return
	}
	fmt.Println("用户信息:", findUser)
	if req.Password == findUser.Password {
		fmt.Println("账号密码验证正确")
		auth, err := utils.GetToken(systemModel.TokenInfo{UserName: findUser.UserName, Id: int(findUser.ID)}, global.CONFIG.JWT.ExpiresTime)
		if err == nil {
			resModel.OkWithDetailed(userResModel.LoginResFromDBLogin(&findUser, auth), "登录成功", ctx)
		} else {
			resModel.FailWithMessage(err.Error(), ctx)
		}
	} else {
		resModel.FailWithMessage("密码错误", ctx)
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
	//  validator
	validate := validator.New()
	validate.RegisterValidation("myPassword", myPasswordValidator)
	err = validate.Struct(register)
	if err != nil {
		fmt.Println("validate error:", err.Error())
		if _, ok := err.(*validator.InvalidValidationError); ok {
			resModel.FailWithMessage(err.Error(), ctx)
			return
		}
		err := err.(validator.ValidationErrors)[0]
		switch err.Tag() {
		case "required":
			resModel.FailWithMessage(err.Field()+"字段不能为空", ctx)
		case "myPassword":
			resModel.FailWithMessage("密码至少为6位字符组成", ctx)
		default:
			resModel.FailWithMessage(err.Error(), ctx)
		}
	}

	// 校验重复
	findUser := dbModel.User{}
	result := global.DB.Where(&dbModel.User{UserName: register.UserName}).Take(&findUser)
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
	resModel.OkWithDetailed(data, "success", ctx)
}
