package types

type UserLoginReq struct {
	UserName string `form:"user_name" json:"user_name"`
	Password string `form:"password" json:"password"`
}

type UserLoginResp struct {
	Token   string
	Success bool
}
