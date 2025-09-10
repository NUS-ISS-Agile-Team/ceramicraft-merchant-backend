package types

// 基础分页参数
type BasePage struct {
	PageNum  int `form:"page_num"`
	PageSize int `form:"page_size"`
}