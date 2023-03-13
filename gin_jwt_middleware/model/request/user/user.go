package user

type Login struct {
	UserName string `json:"userName" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type Register struct {
	UserName string `json:"userName" validate:"required"`
	NickName string `json:"nickName" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required,myPassword"`
}
