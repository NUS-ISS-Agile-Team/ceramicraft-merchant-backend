package service

import (
	"errors"
	"testing"

	"github.com/cerami-craft-shop/merchant-backend/repository/db/dao"
	"github.com/golang/mock/gomock"
)

func TestItemService_GetItemList(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// 通过m来mock ItemDao接口
	m := dao.NewMockItemDao(ctrl)
	// 期望调用GetItemList方法，返回nil和错误
	m.EXPECT().GetItemList().Return(nil, errors.New("not exist"))

	// 将m作为参数传入service层的懒获取方法
	_, err := GetItemService(m).GetItemList()
	if err == nil {
		t.Fatal("expect err not nil")
	}
}
