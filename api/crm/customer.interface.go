package crm

type CustomerInfoEs struct {
	VenderId           uint64 `json:"venderId,omitempty" codec:"venderId,omitempty"`                     // 商家Id
	ShopId             uint64 `json:"shopId,omitempty" codec:"shopId,omitempty"`                         // 店铺Id
	CompanyId          uint64 `json:"companyId,omitempty" codec:"companyId,omitempty"`                   // 公司Id
	UserId             int64  `json:"userId,omitempty" codec:"userId,omitempty"`                         // 用户Id
	NickName           string `json:"nickName,omitempty" codec:"nickName,omitempty"`                     // 用户昵称
	LevelAtJd          string `json:"levelAtJd,omitempty" codec:"levelAtJd,omitempty"`                   // jd等级
	TotalOrderPrice    int64  `json:"totalOrderPrice,omitempty" codec:"totalOrderPrice,omitempty"`       // 订单总金额 保存的是分
	TotalOrderCount    int    `json:"totalOrderCount,omitempty" codec:"totalOrderCount,omitempty"`       // 订单总数量
	TotalGoodsCount    int    `json:"totalGoodsCount,omitempty" codec:"totalGoodsCount,omitempty"`       // 商品数量
	CanceledOrderCount int    `json:"canceledOrderCount,omitempty" codec:"canceledOrderCount,omitempty"` // 退货次数
	AvgOrderPrice      int64  `json:"avgOrderPrice,omitempty" codec:"avgOrderPrice,omitempty"`           // 平均客单价
	LastOrderDate      string `json:"lastOrderDate,omitempty" codec:"lastOrderDate,omitempty"`           // 上次订单时间
	OrderFrom          int    `json:"orderFrom,omitempty" codec:"orderFrom,omitempty"`                   // 订单来源
	FirstOrderDate     string `json:"firstOrderDate,omitempty" codec:"firstOrderDate,omitempty"`         // 首次订单时间
	CustomerStatus     uint8  `json:"customerStatus,omitempty" codec:"customerStatus,omitempty"`         // 会员状态 1：正式会员 2：预会员 3：流失会员
	LevelAtShop        int    `json:"levelAtShop,omitempty" codec:"levelAtShop,omitempty"`               // 店铺等级
	PcFlag             int    `json:"pcFlag,omitempty" codec:"pcFlag,omitempty"`                         // 是否pc下单
	PhoneFlag          int    `json:"phoneFlag,omitempty" codec:"phoneFlag,omitempty"`                   // 是否手机下单
	WxFlag             int    `json:"wxFlag,omitempty" codec:"wxFlag,omitempty"`                         // 是否微信下单
	IncreaseDate       string `json:"increaseDate,omitempty" codec:"increaseDate,omitempty"`             // 增量日期
	Created            string `json:"created,omitempty" codec:"created,omitempty"`                       // 创建时间
	Modified           string `json:"modified,omitempty" codec:"modified,omitempty"`                     // 修改时间
	HuanHuo            int    `json:"huanHuo,omitempty" codec:"huanHuo,omitempty"`                       // 换货次数
	Tuidan             int    `json:"tuidan,omitempty" codec:"tuidan,omitempty"`                         // 退单次数
	TuihuanMoney       int64  `json:"tuihuanMoney,omitempty" codec:"tuihuanMoney,omitempty"`             // 退换货金额
	Points             int    `json:"points,omitempty" codec:"points,omitempty"`                         // 积分
}

type Member struct {
	CustomerPin      string  `json:"customer_pin,omitempty" codec:"customer_pin,omitempty"`             // 用户在京东的会员账号
	Grade            string  `json:"grade,omitempty" codec:"grade,omitempty"`                           // 会员等级， 1:一星会员，2:二星会员 3:三星会员 4:四星会员 5:五星会员
	TradeCount       int     `json:"trade_count,omitempty" codec:"trade_count,omitempty"`               // 交易成功笔数
	TradeAmount      float64 `json:"trade_amount,omitempty" codec:"trade_amount,omitempty"`             // 交易成功的金额，单位为元
	CloseTradeCount  int     `json:"close_trade_count,omitempty" codec:"close_trade_count,omitempty"`   // 交易关闭的的笔数 退货订单量
	CloseTradeAmount float64 `json:"close_trade_amount,omitempty" codec:"close_trade_amount,omitempty"` // 交易关闭的金额 退货订单金额，单位为元
	ItemNum          int     `json:"item_num,omitempty" codec:"item_num,omitempty"`                     // 购买的商品件数
	AvgPrice         float64 `json:"avg_price,omitempty" codec:"avg_price,omitempty"`                   // 平均客单价，单位为元
	LastTradeTime    int64   `json:"last_trade_time,omitempty" codec:"last_trade_time,omitempty"`       // 最后交易时间。非sdk调用 请自行转换long到date

	AccountId        uint64 `json:"account_id,omitempty" codec:"account_id,omitempty"`
	JdUserId         int64  `json:"jd_user_id,omitempty" codec:"jd_user_id,omitempty"`
	CustomerEncryPin string `json:"customer_encrypin,omitempty" codec:"customer_encrypin,omitempty"`
}
