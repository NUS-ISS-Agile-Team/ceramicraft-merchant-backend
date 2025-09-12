package model

import "gorm.io/gorm"

// 商品模型
type Item struct {
	gorm.Model
	CategoryID  uint   `gorm:"column:category_id;type:uint;not null"`
	Title       string `gorm:"column:title;type:varchar(255);not null"`
	Description string `gorm:"column:description;type:text;size:1000"`
	ImgPath     string `gorm:"column:img_path;type:varchar(255);not null"`
	Price       string `gorm:"column:price;type:varchar(50);not null"`
	OnSale      bool   `gorm:"column:on_sale;type:boolean;default:false;not null"`
	Num         int    `gorm:"column:num;type:int;default:0;not null"`
}

func (i *Item) TableName() string {
	return "items"
}