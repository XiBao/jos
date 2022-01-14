package vender

type BasicVenderInfo struct {
	Id         uint64 `json:"id,omitempty" codec:"id,omitempty"`
	Status     int    `json:"status,omitempty" codec:"status,omitempty"`
	ShopName   string `json:"shopName,omitempty" codec:"shopName,omitempty"`
	ShopId     uint64 `json:"shopId,omitempty" codec:"shopId,omitempty"`
	VenderCode string `json:"venderCode,omitempty" codec:"venderCode,omitempty"`
	VenderType uint8  `json:"venderType,omitempty" codec:"venderType,omitempty"` // 1为自营，0为pop
}

type ShopInfo struct {
	VenderId         uint64 `json:"vender_id,omitempty" codec:"vender_id,omitempty"`
	ShopId           uint64 `json:"shop_id,omitempty" codec:"shop_id,omitempty"`
	ShopName         string `json:"shop_name,omitempty" codec:"shop_name,omitempty"`
	OpenTime         int64  `json:"open_time,omitempty" codec:"open_time,omitempty"`
	LogoUrl          string `json:"logo_url,omitempty" codec:"logo_url,omitempty"`
	Brief            string `json:"brief,omitempty" codec:"brief,omitempty"`
	CategoryMain     int    `json:"category_main,omitempty" codec:"category_main,omitempty"`
	CategoryMainName string `json:"category_main_name,omitempty" codec:"category_main_name,omitempty"`
	ColType          uint8  `json:"col_type" codec:"col_type"` // 商家类型: 0：SOP 1：FBP 2：LBP 5：SOPL
	OpenTimeStr      string `json:"open_time_str,omitempty" codec:"open_time_str,omitempty"`
}
