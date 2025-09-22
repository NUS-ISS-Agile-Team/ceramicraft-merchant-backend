package service

import (
	"testing"

	"github.com/cerami-craft-shop/merchant-backend/pkg/utils/jwt"
	"github.com/cerami-craft-shop/merchant-backend/types"
)

func TestUserService_Login(t *testing.T) {
	u := &UserService{}
	resp, err := u.Login("admin@qq.com", "123456")
	if err != nil {
		t.Fatalf("TestUserService_Login, got err: %v", err)
	}
	if !resp.(*types.UserLoginResp).Success {
		t.Fatalf("expect login success, got false")
	}
	expectedToken, _ := jwt.GenerateToken("admin@qq.com")
	if resp.(*types.UserLoginResp).Token != expectedToken {
		t.Fatalf("expect token: %s, got: %s", expectedToken, resp.(*types.UserLoginResp).Token)
	}

	resp, err = u.Login("admin@qq.com", "1234")
	if err != nil {
		t.Fatalf("TestUserService_Login, got err: %v", err)
	}
	if resp.(*types.UserLoginResp).Success {
		t.Fatalf("expect login fail, got true")
	}
}
