package service

import (
	"sync"

	"github.com/cerami-craft-shop/merchant-backend/repository/db/dao"
)

type ItemService struct{
	itemDao dao.ItemDao
}

var once sync.Once
var itemService *ItemService

func GetItemService(itd dao.ItemDao) *ItemService {
	once.Do(func() {
		itemService = &ItemService{
			itemDao: itd,
		}
	})
	return itemService
}

func (is *ItemService) GetItemList() (resp interface{}, err error) {
	itemList, err := is.itemDao.GetItemList()
	if err != nil {
		return nil, err
	}
	return itemList, nil
}
