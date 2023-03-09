package user

type Login struct {
	UserName string `json:"userName"`
	Password string `json:"password"`
}

type Register struct {
	UserName string `json:"userName"`
	NickName string `json:"nickName"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
