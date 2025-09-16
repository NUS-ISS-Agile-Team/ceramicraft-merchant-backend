package service

import (
	"github.com/cerami-craft-shop/merchant-backend/repository/db/dao"
)

type ItemService struct {
	itemDao dao.ItemDao
}

func GetItemService(itd dao.ItemDao) *ItemService {
	return &ItemService{
		itemDao: itd,
	}
}
func (is *ItemService) GetItemList() (resp interface{}, err error) {
	itemList, err := is.itemDao.GetItemList()
	if err != nil {
		return nil, err
	}
	return itemList, nil
}
