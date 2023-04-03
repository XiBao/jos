package fullcoupon

type PromoListInfo struct {
	PromoId       uint64 `json:"promoId,omitempty"`       // 促销ID
	PromoName     string `json:"promoName,omitempty"`     // 促销名称
	PromoType     int    `json:"promoType,omitempty"`     // 促销类型 固定20
	MemberType    int    `json:"memberType,omitempty"`    // 会员类型 1-京东会员，2-企业会员
	MemberLevel   int    `json:"memberLevel,omitempty"`   // 会员级别：1-普通会员 2-PLUS 3-PLUS+普通 4-新人会员 5-学生用户
	MemberName    string `json:"memberName,omitempty"`    // 会员级别展示文案
	Status        int    `json:"status,omitempty"`        // 状态 全部：-1 ；系统未审核：1；人工未审核：5；驳回：11；未开始：2；进行中：3；已暂停：4；已结束：6；即将结束：20
	StatusName    string `json:"statusName,omitempty"`    // 状态展示文案
	WareGrade     int    `json:"wareGrade,omitempty"`     // 参与商品级 1：SKU级；2：SPU级
	WareGradeName string `json:"wareGradeName,omitempty"` // 参与商品级别名称 1：SKU级；2：SPU级
	StartTime     string `json:"startTime,omitempty"`     // 开始时间
	EndTime       string `json:"endTime,omitempty"`       // 结束时间
}

type PromoDetailsPlatform struct {
	PlatName   string `json:"platName"`   // 平台名称 默认全平台
	SelectType uint   `json:"selectType"` // 一级推广平台的选择方式：0全部1限选(非全部渠道)
}

type PromoDetailsCouponInfo struct {
	CouponQuota     float64 `json:"couponQuota"`     // 满减额度
	Discount        float64 `json:"discount"`        // 优惠额度
	ValidateDays    uint    `json:"validateDays"`    // 有效天数
	CouponType      uint    `json:"couponType"`      // 优惠券类型 ：0-商品券 1-店铺券
	OrderDoneTime   string  `json:"orderDoneTime"`   // 完成订单时间
	LimitFirstOrder uint    `json:"limitFirstOrder"` // 是否首单返券：1-是 2-否
	OrderDayLimit   uint    `json:"orderDayLimit"`   // 订单完成多少天内有效
	ValidityType    uint    `json:"validityType"`    // 有效期类型：5为绝对时间类型，1为相对时间类型
	ValidBeginTime  string  `json:"validBeginTime"`  // 券开始时间
	ValidEndTime    string  `json:"validEndTime"`    // 券结束时间
	StoreNum        uint    `json:"storeNum"`        // 库存数量
}

type PromoDetailsOrderMode struct {
	Quota      float64                `json:"quota"`      // 满金额
	CouponInfo PromoDetailsCouponInfo `json:"couponInfo"` // 满额返券优惠券信息
}

type PromoDetailsInfo struct {
	WareGrade     uint                    `json:"wareGrade"`                // 参与商品级别 1：SKU级；2：SPU级
	PromoName     string                  `json:"promoName"`                // 促销名称
	StartTime     string                  `json:"startTime"`                // 开始时间
	PromoId       uint64                  `json:"promoId"`                  // 促销ID
	EndTime       string                  `json:"endTime"`                  // 结束时间
	WareGradeName string                  `json:"wareGradeName"`            // 参与商品级别名称 1：SKU级；2：SPU级
	Platform      PromoDetailsPlatform    `json:"fullCouponPlatForm"`       // 推广平台
	OrderModes    []PromoDetailsOrderMode `json:"fullCouponOrderModesList"` // 满额返券促销内容
}

type PromoWare struct {
	WareId     uint64 `json:"wareId"`     // 商品ID
	WareName   string `json:"wareName"`   // 商品名称
	WareImg    string `json:"wareImg"`    // 商品主图
	SyncStatus uint   `json:"syncStatus"` // 状态 1-上架 2-下架
	JdPrice    string `json:"jdPrice"`    // 商品价格
}

type PromoSku struct {
	SkuName    string `json:"skuName"`    // sku名称
	JdPrice    string `json:"jdPrice"`    // 商品价格
	SkuId      uint64 `json:"skuId"`      // sku编码
	SyncStatus uint   `json:"syncStatus"` // 状态 1-上架 2-下架
	ImgRui     string `json:"imgRui"`     // sku主图
}

type PromoTrendData struct {
	LastCustPriceNum int     `json:"lastCustPriceNum"` // 较前一天客单件数 前端直接添加% 不需要乘100
	EndDate          string  `json:"endDate"`          // 促销结束时间
	OrdProNum        uint    `json:"ordProNum"`        // 成交商品件数
	AccOrdAmt        float64 `json:"accOrdAmt"`        // 累计成交金额
	ToOrdRate        float64 `json:"toOrdRate"`        // 成交转化率
	AccOrdNum        uint    `json:"accOrdNum"`        // 累计成交单量
	TimeStr          string  `json:"timeStr"`          // 数据对应日期
	Platform         string  `json:"platform"`         // 推广平台
	LastOrdProNum    uint    `json:"lastOrdProNum"`    // 较前一天成交商品数 前端直接添加% 不需要乘100
	CustPriceAvg     float64 `json:"custPriceAvg"`     // 客单价
	LastCustPriceAvg uint    `json:"lastCustPriceAvg"` // 较前一天客单价 前端直接添加% 不需要乘100
	Member           string  `json:"member"`           // 促销人群
	PromoId          uint64  `json:"promoId"`          // 促销ID
	CustPriceNum     float64 `json:"custPriceNum"`     // 客单件数
	LastToOrdRate    uint    `json:"lastToOrdRate"`    // 较前一天成交转化率 前端直接添加% 不需要乘100
	AccOrdProNum     uint    `json:"accOrdProNum"`     // 累计成交商品件数
	AccOrdCustNum    uint    `json:"accOrdCustNum"`    // 累计成交客户数
	PromoType        int     `json:"promoType"`        // 促销类型
	OrdCustNum       uint    `json:"ordCustNum"`       // 成交客户数
	Favor            string  `json:"favor"`            // 优惠条件
	DataName         string  `json:"dataName"`         // 名称
	LastOrdCustNum   uint    `json:"lastOrdCustNum"`   // 较前一天成交客户数 前端直接添加% 不需要乘100
	OrdAmt           float64 `json:"ordAmt"`           // 成交金额
	OrdNum           uint    `json:"ordNum"`           // 成交单量
	LastOrdAmt       uint    `json:"lastOrdAmt"`       // 较前一天成交金额 前端直接添加% 不需要乘100
	LastOrdNum       uint    `json:"lastOrdNum"`       // 较前一天成交单量 前端直接添加% 不需要乘100
	StartDate        string  `json:"startDate"`        // 促销开始时间
}

type PromoFlow struct {
	SkuId  uint64 `json:"skuId"`  // skuId
	Remark string `json:"remark"` // 审核原因
}
