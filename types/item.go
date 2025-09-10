package types

// BasePage 基础分页请求
// 注意：这里假设BasePage已在其他地方定义
// 如果未定义，需要添加BasePage结构体

// ItemListReq 获取商品列表请求

type ItemListReq struct {
	BasePage
}

// ItemListResp 获取商品列表响应
type ItemListResp struct {
	Items []ItemResp `json:"items"`
	Total int64      `json:"total"`
}

// ItemResp 商品详情响应
type ItemResp struct {
	ID          uint      `json:"id"`
	CategoryID  uint      `json:"category_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	ImgPath     string    `json:"img_path"`
	Price       string    `json:"price"`
	OnSale      bool      `json:"on_sale"`
	Num         int       `json:"num"`
}
