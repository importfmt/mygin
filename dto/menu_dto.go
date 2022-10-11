package dto

// MenuDto 主菜单
type MenuDto struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Children []MenuDto `json:"children"`
	Path string `json:"path"`
} 