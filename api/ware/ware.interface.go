package ware

type Ware struct {
	WareId               uint64      `json:"wareId,omitempty" codec:"wareId,omitempty"`                             // 商品id
	Title                string      `json:"title,omitempty" codec:"title,omitempty"`                               // 商品名称
	CategoryId           uint64      `json:"categoryId,omitempty" codec:"categoryId,omitempty"`                     // 类目ID
	BrandId              uint64      `json:"brandId,omitempty" codec:"brandId,omitempty"`                           // 品牌ID
	TemplateId           uint64      `json:"templateId,omitempty" codec:"templateId,omitempty"`                     // 关联板式id
	TransportId          uint64      `json:"transportId,omitempty" codec:"transportId,omitempty"`                   // 运费模板ID
	WareStatus           int         `json:"wareStatus,omitempty" codec:"wareStatus,omitempty"`                     // 商品状态 1:从未上架 2:自主下架 4:系统下架 8:上架 513:从未上架待审 514:自主下架待审 516:系统下架待审 520:上架待审核 1028:系统下架审核失败
	OuterId              string      `json:"outerId,omitempty" codec:"outerId,omitempty"`                           // 外部ID
	ItemNum              string      `json:"itemNum,omitempty" codec:"itemNum,omitempty"`                           // 商品货号
	BarCode              string      `json:"barCode,omitempty" codec:"barCode,omitempty"`                           // 条形码
	WareLocation         int         `json:"wareLocation,omitempty" codec:"wareLocation,omitempty"`                 // 商品产地
	Modified             int64       `json:"modified,omitempty" codec:"modified,omitempty"`                         // 商品最后一次修改时间
	Created              int64       `json:"created,omitempty" codec:"created,omitempty"`                           // 商品创建时间，只读属性
	OfflineTime          int64       `json:"offlineTime,omitempty" codec:"offlineTime,omitempty"`                   // 最后下架时间
	OnlineTime           int64       `json:"onlineTime,omitempty" codec:"onlineTime,omitempty"`                     // 最后上架时间
	ColType              int         `json:"colType,omitempty" codec:"colType,omitempty"`                           // 合作模式
	Delivery             int         `json:"delivery,omitempty" codec:"delivery,omitempty"`                         // 发货地
	AdWords              *AdWords    `json:"adWords,omitempty" codec:"adWords,omitempty"`                           // 商品广告词
	Wrap                 string      `json:"wrap,omitempty" codec:"wrap,omitempty"`                                 // 包装规格
	PackListing          string      `json:"packListing,omitempty" codec:"packListing,omitempty"`                   // 包装清单
	Weight               float64     `json:"weight,omitempty" codec:"weight,omitempty"`                             // 重
	Width                float64     `json:"width,omitempty" codec:"width,omitempty"`                               // 宽
	Height               float64     `json:"height,omitempty" codec:"height,omitempty"`                             // 高度
	Length               float64     `json:"length,omitempty" codec:"length,omitempty"`                             // 长
	Props                []*SaleAttr `json:"props,omitempty" codec:"props,omitempty"`                               // 商品属性
	Features             []*Feature  `json:"features,omitempty" codec:"features,omitempty"`                         // 特殊属性集合
	Images               []*Image    `json:"images,omitempty" codec:"images,omitempty"`                             // 商品图片
	ShopCategorys        []int       `json:"shopCategorys,omitempty" codec:"shopCategorys,omitempty"`               // 店内分类
	MobileDesc           string      `json:"mobileDesc,omitempty" codec:"mobileDesc,omitempty"`                     // 移动端详情
	Introduction         string      `json:"introduction,omitempty" codec:"introduction,omitempty"`                 // PC端详情
	ZhuangBaIntroduction string      `json:"zhuangBaIntroduction,omitempty" codec:"zhuangBaIntroduction,omitempty"` // 装吧详情
	ZhuangBaId           string      `json:"zhuangBaId,omitempty" codec:"zhuangBaId,omitempty"`                     // 商品描述装吧实例ID
	IntroductionUseFlag  string      `json:"introductionUseFlag,omitempty" codec:"introductionUseFlag,omitempty"`   // 商品描述使用标识 ,0：使用默认的商品描述,1：使用装吧商详
	AfterSales           string      `json:"afterSales,omitempty" codec:"afterSales,omitempty"`                     // 售后服务
	Logo                 string      `json:"logo,omitempty" codec:"logo,omitempty"`                                 // 商品主图
	MarketPrice          float64     `json:"marketPrice,omitempty" codec:"marketPrice,omitempty"`                   // 市场价
	CostPrice            float64     `json:"costPrice,omitempty" codec:"costPrice,omitempty"`                       // 成本价
	JdPrice              float64     `json:"jdPrice,omitempty" codec:"jdPrice,omitempty"`                           // 京东价
	BrandName            string      `json:"brandName,omitempty" codec:"brandName,omitempty"`                       // 品牌名称
	StockNum             int         `json:"stockNum,omitempty" codec:"stockNum,omitempty"`                         // 总库存数
	CategorySecId        int         `json:"categorySecId,omitempty" codec:"categorySecId,omitempty"`               // 二级叶子类目
	ShopId               int         `json:"shopId,omitempty" codec:"shopId,omitempty"`                             // 商家对应的shopID，只读属性
	PromiseId            int         `json:"promiseId,omitempty" codec:"promiseId,omitempty"`                       // 时效模板ID
	MultiCategoryId      int         `json:"multiCategoryId,omitempty" codec:"multiCategoryId,omitempty"`           // 四级类目ID
	MultiCateProps       []*SaleAttr `json:"multiCateProps,omitempty" codec:"multiCateProps,omitempty"`             // 四级类目属性列表
	SellPoint            string      `json:"sellPoint,omitempty" codec:"sellPoint,omitempty"`                       // 卖点
	WareTax              *WareTax    `json:"wareTax,omitempty" codec:"wareTax,omitempty"`                           // 税率
	AfterSaleDesc        string      `json:"afterSaleDesc,omitempty" codec:"afterSaleDesc,omitempty"`               // 售后图文详情
	ZhuangBaMobileDesc   string      `json:"zhuangBaMobileDesc,omitempty" codec:"zhuangBaMobileDesc,omitempty"`     // 移动版装吧商详
	MobileZhuangBaId     string      `json:"mobileZhuangBaId,omitempty" codec:"mobileZhuangBaId,omitempty"`         // 移动版装吧实例ID
	MobileDescUseFlag    string      `json:"mobileDescUseFlag,omitempty" codec:"mobileDescUseFlag,omitempty"`       // 移动版商品描述使用标识,0：使用默认的移动商详；1：使用装吧移动版商详
	FitCaseHtmlPc        string      `json:"fitCaseHtmlPc,omitempty" codec:"fitCaseHtmlPc,omitempty"`               // 装修案例PC版描述
	FitCaseHtmlApp       string      `json:"fitCaseHtmlApp,omitempty" codec:"fitCaseHtmlApp,omitempty"`             // 装修案例移动版描述
	SpecialServices      []string    `json:"specialServices,omitempty" codec:"specialServices,omitempty"`           // 特色服务,装修类目才可填写,装修类必填,最大为5,每个值最长为8个字符
	ParentId             uint64      `json:"parentId,omitempty" codec:"parentId,omitempty"`                         // 商品父ID
	WareGroupId          uint64      `json:"wareGroupId,omitempty" codec:"wareGroupId,omitempty"`                   // 商品分组ID
	BusinessType         string      `json:"businessType,omitempty" codec:"businessType,omitempty"`                 // 商品业务类型
	DesignConcept        string      `json:"designConcept,omitempty" codec:"designConcept,omitempty"`               // 商品设计理念,适用范围是toplife类目
	IsArchival           bool        `json:"isArchival,omitempty" codec:"isArchival,omitempty"`                     // 是否归档商品
	TemplateIds          string      `json:"templateIds,omitempty" codec:"templateIds,omitempty"`                   // 关联版式

	Skus           []*Sku `json:"skus,omitempty" codec:"skus,omitempty"`
	OnlineTimeStr  string `json:"OnlineTimeStr,omitempty" codec:"OnlineTimeStr,omitempty"`   // 最后下架时间
	OfflineTimeStr string `json:"offlineTimeStr,omitempty" codec:"offlineTimeStr,omitempty"` // 最后下架时间
}

type Sku struct {
	WareId               uint64      `json:"wareId,omitempty" codec:"wareId,omitempty"`       // 商品ID
	SkuId                uint64      `json:"skuId,omitempty" codec:"skuId,omitempty"`         // skuId
	Status               uint8       `json:"status,omitempty" codec:"status,omitempty"`       // 状态,只读属性. 1:上架 2:下架 4:删除
	SaleAttrs            []*SaleAttr `json:"saleAttrs,omitempty" codec:"saleAttrs,omitempty"` // 销售属性集合
	Features             []*Feature  `json:"features,omitempty" codec:"features,omitempty"`   // 特殊属性集合
	JdPrice              float64     `json:"jdPrice,omitempty" codec:"jdPrice,omitempty"`     // 京东价
	PromoPrice           float64     `json:"promoPrice,omitempty", codec:"promoPrice,omitempty"`
	OuterId              string      `json:"outerId,omitempty" codec:"outerId,omitempty"`                           // 外部ID
	BarCode              string      `json:"barCode,omitempty" codec:"barCode,omitempty"`                           // sku条形码
	CategoryId           uint64      `json:"categoryId,omitempty" codec:"categoryId,omitempty"`                     // 类目id
	ImgTag               int         `json:"imgTag,omitempty" codec:"imgTag,omitempty"`                             // 图片标签
	Logo                 string      `json:"logo,omitempty" codec:"logo,omitempty"`                                 // sku颜色的主图
	SkuName              string      `json:"skuName,omitempty" codec:"skuName,omitempty"`                           // sku名称
	StockNum             int         `json:"stockNum,omitempty" codec:"stockNum,omitempty"`                         // 总库存数
	WareTitle            string      `json:"wareTitle,omitempty" codec:"wareTitle,omitempty"`                       // 商品名称
	FixedDeliveryTime    string      `json:"fixedDeliveryTime,omitempty" codec:"fixedDeliveryTime,omitempty"`       // 大件商品固定发货时效 格式：订单开始日期,订单结束日期,发货日期,完成发货天数
	RelativeDeliveryTime string      `json:"relativeDeliveryTime,omitempty" codec:"relativeDeliveryTime,omitempty"` // 大件商品相对发货时效（完成发货天数）
	ParentId             uint64      `json:"parentId,omitempty" codec:"parentId,omitempty"`                         // 父id
	Modified             int64       `json:"modified,omitempty" codec:"modified,omitempty"`                         // 修改时间
	Created              int64       `json:"created,omitempty" codec:"created,omitempty"`                           // 创建时间
	MultiCateProps       []*SaleAttr `json:"multiCateProps,omitempty" codec:"multiCateProps,omitempty"`             // 多级SKU属性
	Props                []*SaleAttr `json:"props,omitempty" codec:"props,omitempty"`                               // SKU属性
	Capacity             string      `json:"capacity,omitempty"`
	MobileDesc           string      `json:"mobileDesc" coded:"mobileDesc"` // 手机详情
	XbOnlineSts          uint8       `json:"xb_online_sts,omitempty" codec:"xb_online_sts,omitempty"`
	XbAuditSts           uint8       `json:"xb_audit_sts,omitempty" codec:"xb_audit_sts,omitempty"`
	XbUpdatedAt          string      `json:"xb_updated_at,omitempty" codec:"xb_updated_at,omitempty"`
	TrackUrl             string      `json:"track_url,omitempty" codec:"trace_url,omitempty"`   // 京东追踪链接
	PromoLimit           float64     `json:"promoLimit,omitempty" codec:"promoLimit,omitempty"` // 促销最低折扣
}

type Sku2 struct {
	SkuId       uint64 `json:"sku_id,omitempty" codec:"sku_id,omitempty"`             // sku的id
	WareId      uint64 `json:"ware_id,omitempty" codec:"ware_id,omitempty"`           // sku所属商品id
	ShopId      uint64 `json:"shop_id,omitempty" codec:"shop_id,omitempty"`           // 店铺id
	Status      string `json:"status,omitempty" codec:"status,omitempty"`             // sku状态 有效：Valid 无效：Invalid
	Attributes  string `json:"attributes,omitempty" codec:"attributes,omitempty"`     // sku的销售属性组合字符串（颜色，大小，等等，可通过类目API获取某类目下的销售属性）,格式是aid1:vid1;aid2:vid2
	JdPrice     string `json:"jd_price,omitempty" codec:"jd_price,omitempty"`         // 京东价,精确到2位小数，单位元
	CostPrice   string `json:"cost_price,omitempty" codec:"cost_price,omitempty"`     // 进货价, 精确到2位小数，单位元
	MarketPrice string `json:"market_price,omitempty" codec:"market_price,omitempty"` // 市场价, 精确到2位小数，单位元
	StockNum    int    `json:"stock_num,omitempty" codec:"stock_num,omitempty"`       // 库存
	OuterId     string `json:"outer_id,omitempty" codec:"outer_id,omitempty"`         // 外部id,商家设置的外部id
	Created     string `json:"created,omitempty" codec:"created,omitempty"`           // sku创建时间 时间格式：yyyy-MM-ddHH:mm:ss
	Modified    string `json:"modified,omitempty" codec:"modified,omitempty"`         // sku修改时间 时间格式：yyyy-MM-ddHH:mm:ss
	ColorValue  string `json:"color_value,omitempty" codec:"color_value,omitempty"`   // 颜色对应的值
	SizeValue   string `json:"size_value,omitempty" codec:"size_value,omitempty"`     // 尺码对应的值

	Logo string `json:"logo,omitempty" codec:"logo,omitempty"` // sku颜色的主图
}

// 销售属性
type SaleAttr struct {
	AttrId         string   `json:"attrId,omitempty" codec:"attrId,omitempty"`                 // 属性ID
	AttrValueAlias []string `json:"attrValueAlias,omitempty" codec:"attrValueAlias,omitempty"` // 属性值的描述或者别名,与attrValues一一对应
	AttrValues     []string `json:"attrValues,omitempty" codec:"attrValues,omitempty"`         // 属性值ID
	Index          int      `json:"index,omitempty" codec:"index,omitempty"`                   // 销售属性维度 1：颜色，2：尺码
}

type Feature struct {
	FeatureCn    string `json:"featureCn,omitempty" codec:"featureCn,omitempty"`       // 特殊属性中文含义
	FeatureKey   string `json:"featureKey,omitempty" codec:"featureKey,omitempty"`     // 特殊属性key
	FeatureValue string `json:"featureValue,omitempty" codec:"featureValue,omitempty"` // 特殊属性value
}

type AdWords struct {
	Words    string `json:"words,omitempty" codec:"words,omitempty"`       // 广告词仅文字内容
	Url      string `json:"url,omitempty" codec:"url,omitempty"`           // 广告词链接地址
	UrlWords string `json:"urlWords,omitempty" codec:"urlWords,omitempty"` // 带链接的广告词
}

type Image struct {
	ColorId   string `json:"colorId,omitempty" codec:"colorId,omitempty"`     // 颜色
	ImgId     uint64 `json:"imgId,omitempty" codec:"imgId,omitempty"`         // 图片ID
	ImgIndex  int    `json:"imgIndex,omitempty" codec:"imgIndex,omitempty"`   // 图片顺序
	ImgUrl    string `json:"imgUrl,omitempty" codec:"imgUrl,omitempty"`       // 图片路径
	ImgZoneId string `json:"imgZoneId,omitempty" codec:"imgZoneId,omitempty"` // 图片空间中图片ID
}

type WareTax struct {
	TaxCode         string  `json:"taxCode,omitempty" codec:"taxCode,omitempty"`                 // 税收编码（固定长度19位）
	TaxRate         float64 `json:"taxRate,omitempty" codec:"taxRate,omitempty"`                 // 税率（范围0.00~1.00）
	IsTaxCheap      uint8   `json:"isTaxCheap,omitempty" codec:"isTaxCheap,omitempty"`           // 是否享受税收优惠政策（0：不享受；1：享受；）
	TaxCheapContent string  `json:"taxCheapContent,omitempty" codec:"taxCheapContent,omitempty"` // 享受税收优惠政策内容（最大长度32位）
	ZeroTaxRate     uint8   `json:"zeroTaxRate,omitempty" codec:"zeroTaxRate,omitempty"`         // 零税率（0：出口退税；1：免税；2：不征收；3：普通零税率；）
}
