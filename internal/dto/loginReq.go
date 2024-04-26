package dto

type LoginReq struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}
