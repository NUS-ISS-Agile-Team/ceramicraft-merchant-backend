package service

import (
	"sync"

	"github.com/cerami-craft-shop/merchant-backend/repository/db/dao"
)

type ItemService struct{}

var once sync.Once
var itemService *ItemService

func GetItemService() *ItemService {
	once.Do(func() {
		itemService = &ItemService{}
	})
	return itemService
}

func (is *ItemService) GetItemList() (resp interface{}, err error) {
	itemList, err := dao.GetItemDao().GetItemList()
	if err != nil {
		return nil, err
	}
	return itemList, nil
}
