package vender

type VenderInfo struct {
	VenderId uint64 `json:"vender_id,omitempty" codec:"vender_id,omitempty"`         // 商家编号
	ShopId   uint64 `json:"shop_id,omitempty" codec:"shop_id,omitempty"`             // 商家类型: 0：SOP 1：FBP 2：LBP 5：SOPL
	ShopName string `json:"shop_name,omitempty" codec:"shop_name,omitempty"`         // 店铺编号
	ColType  uint8  `json:"col_type" codec:"col_type"`                               // 商家类型: 0：SOP 1：FBP 2：LBP 5：SOPL
	CateMain int    `json:"category_main,omitempty" codec:"category_main,omitempty"` // 主营类目编号
}
