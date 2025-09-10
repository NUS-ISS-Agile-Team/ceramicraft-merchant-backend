package dao

import (
	"sync"

	"github.com/cerami-craft-shop/merchant-backend/repository/db"
	"github.com/cerami-craft-shop/merchant-backend/repository/db/model"
)

type ItemDao struct{}

var once sync.Once
var itemDao *ItemDao

func GetItemDao() *ItemDao {
	once.Do(func() {
		itemDao = &ItemDao{}
	})
	return itemDao
}

func (i *ItemDao) GetItemList() ([]*model.Item, error) {
	var itemList []*model.Item
	err := db.DB.Raw("SELECT * FROM items").Scan(&itemList).Error
	if err != nil {
		return nil, err
	}
	return itemList, nil
}
