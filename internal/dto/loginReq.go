package dto

type LoginReq struct {
	UserName string `json:"user_name" validate:"required"`
	Password string `json:"password"`
}
