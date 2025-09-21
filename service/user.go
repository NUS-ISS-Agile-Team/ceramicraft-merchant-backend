package service

import (
	"crypto/md5"
	"encoding/hex"

	"github.com/cerami-craft-shop/merchant-backend/pkg/utils/jwt"
	"github.com/cerami-craft-shop/merchant-backend/types"
)

type UserService struct{}

const (
	adminEmail string = "admin@qq.com"
	adminPwd   string = "e10adc3949ba59abbe56e057f20f883e" // 123456
)

func GetUserService() *UserService {
	return &UserService{}
}

func (u *UserService) Login(userId string, pwd string) (resp interface{}, err error) {
	loginResp := &types.UserLoginResp{}

	if userId != adminEmail {
		loginResp.Success = false
		return loginResp, nil
	}

	md5Pwd := md5.Sum([]byte(pwd))
	md5PwdStr := hex.EncodeToString(md5Pwd[:]) // 转成16进制字符串
	if md5PwdStr != adminPwd {
		loginResp.Success = false
		return loginResp, nil
	}

	token, err := jwt.GenerateToken(userId)
	if err != nil {
		return nil, err
	}

	loginResp.Success = true
	loginResp.Token = token
	return loginResp, nil
}
