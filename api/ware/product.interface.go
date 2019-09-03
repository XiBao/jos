package ware

type ProductsBase struct {
	SkuId             uint64 `json:"skuId,omitempty" codec:"skuId,omitempty"`                         // sku编号
	Name              string `json:"name,omitempty" codec:"name,omitempty"`                           // 商品名称
	IsDelete          string `json:"isDelete,omitempty" codec:"isDelete,omitempty"`                   // 状态 0: 已删除，1: 正常
	State             string `json:"state,omitempty" codec:"state,omitempty"`                         // 上下柜状态 0: 未上架 1: 已上架 2: 非图书商品表示预上架，图书商品表示前台屏蔽 10: POP商品删除状态
	BarCode           string `json:"barCode,omitempty" codec:"barCode,omitempty"`                     // 产品编码
	ErpPid            string `json:"erpPid,omitempty" codec:"erpPid,omitempty"`                       // 商品编号 主skuId
	Color             string `json:"color,omitempty" codec:"color,omitempty"`                         // 	颜色
	ColorSequence     string `json:"colorSequence,omitempty" codec:"colorSequence,omitempty"`         // 	颜色顺序
	Size              string `json:"size,omitempty" codec:"size,omitempty"`                           // 	尺码
	SizeSequence      string `json:"sizeSequence,omitempty" codec:"sizeSequence,omitempty"`           // 	尺码顺序
	Upc               string `json:"upc,omitempty" codec:"upc,omitempty"`                             // 	Upc码
	SkuMark           string `json:"skuMark,omitempty" codec:"skuMark,omitempty"`                     // 	sku标示
	SaleDate          string `json:"saleDate,omitempty" codec:"saleDate,omitempty"`                   // 上下柜日期
	Cid2              string `json:"cid2,omitempty" codec:"cid2,omitempty"`                           // 商品的第二分类(3级)
	ValueWeight       string `json:"valueWeight,omitempty" codec:"valueWeight,omitempty"`             // 重量
	Weight            string `json:"weight,omitempty" codec:"weight,omitempty"`                       // 重量
	ProductArea       string `json:"productArea,omitempty" codec:"productArea,omitempty"`             // 	商品产地(默认:中国大陆)
	Wserve            string `json:"wserve,omitempty" codec:"wserve,omitempty"`                       // 	质保(无质保,一年质保,二年质保等)
	Allnum            string `json:"allnum,omitempty" codec:"allnum,omitempty"`                       // 	图片标签 默认0 没有标签
	MaxPurchQty       string `json:"maxPurchQty,omitempty" codec:"maxPurchQty,omitempty"`             // 	24小时限购数量
	BrandId           string `json:"brandId,omitempty" codec:"brandId,omitempty"`                     // 	品牌号
	ValuePayFirst     string `json:"valuePayFirst,omitempty" codec:"valuePayFirst,omitempty"`         // 	是否先款商品
	Length            string `json:"length,omitempty" codec:"length,omitempty"`                       // 长度
	Width             string `json:"width,omitempty" codec:"width,omitempty"`                         // 宽度
	Height            string `json:"height,omitempty" codec:"height,omitempty"`                       // 商家类型
	VenderType        string `json:"venderType,omitempty" codec:"venderType,omitempty"`               // 主商品名称
	Pname             string `json:"pname,omitempty" codec:"pname,omitempty"`                         // 	商品产地(默认:中国大陆)
	Issn              string `json:"issn,omitempty" codec:"issn,omitempty"`                           // 	是否序列化
	SafeDays          string `json:"safeDays,omitempty" codec:"safeDays,omitempty"`                   // 	保质期
	SaleUnit          string `json:"saleUnit,omitempty" codec:"saleUnit,omitempty"`                   // 	销售单位
	PackSpecification string `json:"packSpecification,omitempty" codec:"packSpecification,omitempty"` // 	包装规格
	Category          string `json:"category,omitempty" codec:"category,omitempty"`                   // 	分类信息
	ShopCategorys     string `json:"shopCategorys,omitempty" codec:"shopCategorys,omitempty"`         // 店内分类
	Phone             string `json:"phone,omitempty" codec:"phone,omitempty"`                         // 售后电话
	Site              string `json:"site,omitempty" codec:"site,omitempty"`                           // 官网
	Ebrand            string `json:"ebrand,omitempty" codec:"ebrand,omitempty"`                       // 	英文品牌
	Model             string `json:"model,omitempty" codec:"model,omitempty"`                         // 	型号
	ImagePath         string `json:"imagePath,omitempty" codec:"imagePath,omitempty"`                 // 	图片路径
	ShopName          string `json:"shopName,omitempty" codec:"shopName,omitempty"`                   // 	店铺名称
	Url               string `json:"url,omitempty" codec:"url,omitempty"`                             // 	未定义
	VenderId          string `json:"venderId,omitempty" codec:"venderId,omitempty"`                   // 	商家编号

	Price       float64 `json:"price,omitempty" codec:"price,omitempty"`
	MarketPrice float64 `json:"market_price,omitempty" codec:"market_price,omitempty"`
}
