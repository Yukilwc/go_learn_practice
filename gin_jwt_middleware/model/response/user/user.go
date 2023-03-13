package user

import (
	dbModel "gjm/model/db"
)

type LoginRes struct {
	ID            int    `json:"id"`
	UserName      string `json:"userName"`
	Email         string `json:"email"`
	NickName      string `json:"nickName"`
	Authorization string `json:"authorization"`
}

// 返回res生成
func LoginResFromDBLogin(dbLogin *dbModel.User, auth string) LoginRes {
	return LoginRes{
		Authorization: auth,
		ID:            int(dbLogin.ID),
		UserName:      dbLogin.UserName,
		Email:         dbLogin.Email,
		NickName:      dbLogin.NickName,
	}
}
