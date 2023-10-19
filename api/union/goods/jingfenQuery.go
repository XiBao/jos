package goods

import (
	"encoding/json"
	"errors"
	"strconv"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/union/goods"
)

type JingfenQueryRequest struct {
	api.BaseRequest
	EliteId   uint   `json:"eliteId"`             // 频道id：1-好券商品,2-超级大卖场,10-9.9专区,22-热销爆品,24-数码家电,25-超市,26-母婴玩具,27-家具日用,28-美妆穿搭,29-医药保健,30-图书文具,31-今日必推,32-王牌好货,33-秒杀商品,34-拼购商品
	PageIndex uint   `json:"pageIndex,omitempty"` // 页码，默认1
	PageSize  uint   `json:"pageSize,omitempty"`  // 每页数量，默认20，上限50
	SortName  string `json:"sortName,omitempty"`  // 排序字段(price：单价, commissionShare：佣金比例, commission：佣金， inOrderCount30DaysSku：sku维度30天引单量，comments：评论数，goodComments：好评数)
	Sort      string `json:"sort,omitempty"`      // asc,desc升降序,默认降序
}

type JingfenQueryResponse struct {
	ErrorResp *api.ErrorResponnse       `json:"error_response,omitempty"`
	Data      *JingfenQueryResponseData `json:"jd_union_open_goods_jingfen_query_response,omitempty"`
}

type JingfenQueryResponseData struct {
	Result string `json:"result,omitempty"`
}

type JingfenQueryResult struct {
	Code       int64         `json:"code,omitempty"`
	Message    string        `json:"message,omitempty"`
	TotalCount uint          `json:"totalCount,omitempty"`
	Data       []JFGoodsResp `json:"data,omitempty"`
}

type JFGoodsResp struct {
	Category              *JFCategory          `json:"categoryInfo,omitempty"`           // 类目信息
	Comments              uint                 `json:"comments,omitempty"`               // 评论数
	Commission            *JFCommission        `json:"commissionInfo,omitempty"`         // 佣金信息
	Coupon                *JFCouponList        `json:"couponInfo,omitempty"`             // 优惠券信息，返回内容为空说明该SKU无可用优惠券
	GoodCommentsShare     float64              `json:"goodCommentsShare,omitempty"`      // 商品好评率
	Images                *JFImageList         `json:"imageInfo,omitempty"`              // 图片信息
	InOrderCount30Days    uint                 `json:"inOrderCount30Days,omitempty"`     // 30天引单数量
	MaterialUrl           string               `json:"materialUrl,omitempty"`            // 商品落地页，当入参为联盟商品ID时，materialUrl拼接为：jingfen.jd.com/detail/{itemId}.html；当入参为原始商品ID时，materialUrl拼接为：item.jd.com/{skuId}.html
	Price                 *JFPrice             `json:"priceInfo,omitempty"`              // 价格信息
	Shop                  *JFShop              `json:"shopInfo,omitempty"`               // 店铺信息
	SkuId                 uint64               `json:"skuId,omitempty"`                  // 商品ID
	SkuName               string               `json:"skuName,omitempty"`                // 商品名称
	IsHot                 uint                 `json:"isHot,omitempty"`                  // 是否爆款，1：是，0：否 已废弃，请勿使用
	ProductId             uint64               `json:"spuid,omitempty"`                  // spuid，其值为同款商品的主skuid
	BrandCode             string               `json:"brandCode,omitempty"`              // 品牌code
	BrandName             string               `json:"brandName,omitempty"`              // 品牌名
	Owner                 string               `json:"owner,omitempty"`                  // g=自营，p=pop
	Pingou                *JFPingou            `json:"pinGouInfo,omitempty"`             // 拼购信息
	Resource              *JFResource          `json:"resourceInfo,omitempty"`           // 资源信息
	InOrderCount30DaysSku uint                 `json:"inOrderCount30DaysSku,omitempty"`  // 30天引单数量(sku维度)
	Seckill               *JFSeckill           `json:"seckillInfo,omitempty"`            // 秒杀信息
	JxFlags               []uint               `json:"jxFlags,omitempty"`                // 京喜商品类型，1京喜、2京喜工厂直供、3京喜优选（包含3时可在京东APP购买）
	Video                 *JFVideoList         `json:"videoInfo,omitempty"`              // 视频信息
	Document              *JFDocument          `json:"documentInfo,omitempty"`           // 商品段子信息，emoji表情等
	Book                  *JFBook              `json:"bookInfo,omitempty"`               // 图书信息
	ForbidTypes           []uint               `json:"forbidTypes,omitempty"`            // 0普通商品，10微信京东购物小程序禁售，11微信京喜小程序禁售
	DeliveryType          uint                 `json:"deliveryType,omitempty"`           // 京东配送 1：是，0：不是
	SkuLabel              *JFSkuLabel          `json:"skuLabelInfo,omitempty"`           // 商品标签
	PromotionLabel        []JFPromotionLabel   `json:"promotionLabelInfoList,omitempty"` // 商品促销标签集
	SecondPrice           []JFSecondPrice      `json:"secondPriceInfoList,omitempty"`    // 双价格
	PreSale               *JFPreSale           `json:"preSaleInfo,omitempty"`            // 预售信息
	Reserve               *JFReserve           `json:"reserveInfo,omitempty"`            // 预约信息
	SolitaireActivity     *JFSolitaireActivity `json:"solitaireActivity,omitempty"`      // 订单接龙活动信息
	IsOversea             uint                 `json:"isOversea,omitempty"`              // 是否全球购商品 1：是
	CompanyType           uint                 `json:"companyType,omitempty"`            // 2：POP自然人小店
	PurchasePrice         *JFPurchasePrice     `json:"purchasePriceInfo,omitempty"`      // 到手价明细
	Bonus                 []JFBonus            `json:"bonusInfoList,omitempty"`          // 联盟奖励活动集合
	ActivityCard          *JFActivityCard      `json:"activityCardInfo,omitempty"`       // 超市购物卡明细
	SmartDocument         []JFSmartDocument    `json:"smartDocumentInfoList,omitempty"`  // GPT算法智能生成的推广文案集合,内含多个风格
	TotalCount            uint                 `json:"totalCount,omitempty"`             // 有效商品总数量，上限1w
}

type JFBonus struct {
	Id            uint64  `json:"id,omitempty"`            // 活动ID
	Name          string  `json:"name,omitempty"`          // 活动名称
	State         uint    `json:"state,omitempty"`         // 活动状态， 1 ：已提报， 2 ：预热中， 3： 进行中， 4 ：已结束
	BeginTime     uint64  `json:"beginTime,omitempty"`     // 活动开始时间
	EndTime       uint64  `json:"endTime,omitempty"`       // 活动结束时间
	ActivityType  uint    `json:"activityType,omitempty"`  // 活动类型，0：其他， 1：超级补贴:，2：佣金比例活动
	EstimateBonus float64 `json:"estimateBonus,omitempty"` // 预估奖励金额
}

type JFSolitaireActivity struct {
	ActivityId    uint64  `json:"activityId,omitempty"`    // 接龙活动id，订单接龙商品链接（推广订单接龙商品时用该链接转链）： https://item.jd.com/?activityId=xxxx%26skuId=xxxx%26page=chain
	GroupPrice    float64 `json:"groupPrice,omitempty"`    // 成团价
	GroupProgress uint    `json:"groupProgress,omitempty"` // 成团进度（0-100）
	Reason        string  `json:"reason,omitempty"`        // 推荐理由
}

type JFSmartDocument struct {
	DocumentType uint   `json:"documentType,omitempty"` // 智能文案类型，1：最优推荐，2：知乎风格， 3：小红书风格， 4：社群风格
	DocumentName string `json:"documentName,omitempty"` // 文案名称
	Document     string `json:"document,omitempty"`     // 智能文案内容
}

type JFActivityCard struct {
	Amount       uint `json:"amount,omitempty"`       // 超市卡金额
	ActivityType uint `json:"activityType,omitempty"` // 活动类型，购物返卡：3
	ExpireDay    uint `json:"expireDay,omitempty"`    // 超市卡有效天数
}

type JFPurchasePrice struct {
	Code                   uint               `json:"code,omitempty"`                   // 返回码
	Message                string             `json:"message,omitempty"`                // 返回消息
	PurchasePrice          float64            `json:"purchasePrice,omitempty"`          // 到手价
	ThresholdPrice         float64            `json:"thresholdPrice,omitempty"`         // 门槛价金额，计算到手价的基准价
	BasisPriceType         uint               `json:"basisPriceType,omitempty"`         // 依据的价格类型，1、京东价 ,2 Plus价，7 粉丝价，8 新人价，9学生价，10 陪伴计划价（双价格新增）
	PromotionLabelInfoList []JFPromotionLabel `json:"promotionLabelInfoList,omitempty"` // 商品促销标签集
	CouponList             []JFCoupon         `json:"couponList,omitempty"`             // 优惠券集合
}

type JFReserve struct {
	Price                float64 `json:"price,omitempty"`                // 预约价格
	Type                 uint    `json:"type,omitempty"`                 // 预约类型： 1：预约购买资格（仅预约的用户才可以进行购买）； 5：预约抽签（仅中签用户可购买）
	Status               uint    `json:"status,omitempty"`               // 1：等待预约 2：预约中 3：等待抢购/抽签中 4：抢购中 5：抢购结束
	StartTime            uint64  `json:"startTime,omitempty"`            // 预定开始时间
	EndTime              uint64  `json:"endTime,omitempty"`              // 预定结束时间
	PanicBuyingStartTime uint64  `json:"panicBuyingStartTime,omitempty"` // 抢购开始时间
	PanicBuyingEndTime   uint64  `json:"panicBuyingEndTime,omitempty"`   // 抢购结束时间
}

type JFPreSale struct {
	CurrentPrice     float64 `json:"currentPrice,omitempty"`     // 预售价格
	Earnest          float64 `json:"earnest,omitempty"`          // 订金金额（定金不能超过预售总价的20%）
	PreSalePayType   uint    `json:"preSalePayType,omitempty"`   // 预售支付类型：1.仅全款 2.定金、全款均可 5.一阶梯仅定金
	DiscountType     uint    `json:"discountType,omitempty"`     // 1: 定金膨胀 2: 定金立减
	DepositWorth     float64 `json:"depositWorth,omitempty"`     // 定金膨胀金额（定金可抵XXX）【废弃】
	PreAmountDeposit float64 `json:"preAmountDeposit,omitempty"` // 立减金额
	PreSaleStartTime uint64  `json:"preSaleStartTime,omitempty"` // 定金开始时间
	PreSaleEndTime   uint64  `json:"preSaleEndTime,omitempty"`   // 定金结束时间
	BalanceStartTime uint64  `json:"balanceStartTime,omitempty"` // 尾款开始时间
	BalanceEndTime   uint64  `json:"balanceEndTime,omitempty"`   // 尾款结束时间
	ShipTime         uint64  `json:"shipTime,omitempty"`         // 预计发货时间
	PreSaleStatus    uint    `json:"preSaleStatus,omitempty"`    // 预售状态（0 未开始；1 预售中；2 预售结束；3 尾款进行中；4 尾款结束）
	AmountDeposit    uint    `json:"amountDeposit,omitempty"`    // 定金膨胀金额（定金可抵XXX）
}

type JFSecondPrice struct {
	SecondPriceType uint    `json:"secondPriceType,omitempty"` // 双价格类型：18新人价，2plus会员价
	SecondPrice     float64 `json:"secondPrice,omitempty"`     // 价格（资源位238新人价请使用此价格）
}

type JFPromotionLabel struct {
	PromotionLabel   string `json:"promotionLabel,omitempty"`   // 商品促销文案
	LabelName        string `json:"labelName,omitempty"`        // 促销标签名称
	StartTime        uint64 `json:"startTime,omitempty"`        // 促销开始时间
	EndTime          uint64 `json:"endTime,omitempty"`          // 促销结束时间
	PromotionLableId uint64 `json:"promotionLableId,omitempty"` // 促销ID
}

type JFSkuLabel struct {
	Is7ToReturn    uint           `json:"is7ToReturn,omitempty"`    // 0：不支持； 1或null：支持7天无理由退货； 2：支持90天无理由退货； 4：支持15天无理由退货； 6：支持30天无理由退货；
	Fxg            uint           `json:"fxg,omitempty"`            // 1：放心购商品
	FxgServiceList []JFFxgService `json:"fxgServiceList,omitempty"` // 放心购商品子标签集合
}

type JFFxgService struct {
	ServiceName string `json:"serviceName,omitempty"` // 服务名称
}

type JFSpec struct {
	Size     string `json:"size,omitempty"`     // 尺寸
	Color    string `json:"color,omitempty"`    // 颜色
	Spec     string `json:"spec,omitempty"`     // 自定义属性
	SpecName string `json:"specName,omitempty"` // 自定义属性名称
}

type JFBook struct {
	Isbn string `json:"isbn,omitempty"` // 图书编号
}

type JFDocument struct {
	Document string `json:"document,omitempty"` // 描述文案
	Discount string `json:"discount,omitempty"` // 优惠力度文案
}

type JFCommentList struct {
	List []JFComment `json:"commentList,omitempty"` // 评价集合
}

type JFComment struct {
	Content string `json:"content,omitempty"` // 评价内容
}

type JFSeckill struct {
	OriPrice  float64 `json:"seckillOriPrice,omitempty"`  // 秒杀价原价
	Price     float64 `json:"seckillPrice,omitempty"`     // 秒杀价
	StartTime int64   `json:"seckillStartTime,omitempty"` // 秒杀开始时间(时间戳，毫秒)
	EndTime   int64   `json:"seckillEndTime,omitempty"`   // 秒杀结束时间(时间戳，毫秒)
}

type JFResource struct {
	EliteId   uint   `json:"eliteId,omitempty"`   // 频道id
	EliteName string `json:"eliteName,omitempty"` // 频道名称
}

type JFVideoList struct {
	List []JFVideo `json:"videoList,omitempty"`
}

type JFVideo struct {
	Width     uint   `json:"width,omitempty"`     // 宽
	Height    uint   `json:"high,omitempty"`      // 高
	ImageUrl  string `json:"imageUrl,omitempty"`  // 视频图片地址
	VideoType uint   `json:"videoType,omitempty"` // 1:主图，2：商详
	PlayUrl   string `json:"playUrl,omitempty"`   // 播放地址
	PlayType  string `json:"playType,omitempty"`  // low：标清，high：高清
	Duration  uint   `json:"duration,omitempty"`  // 时长(单位:s)
}

type JFPingou struct {
	Price     float64 `json:"pingouPrice,omitempty"`     // 拼购价格
	TmCount   uint    `json:"pingouTmCount,omitempty"`   // 拼购成团所需人数
	PingouUrl string  `json:"pingouUrl,omitempty"`       // 拼购落地页url
	StartTime int64   `json:"pingouStartTime,omitempty"` // 拼购开始时间(时间戳，毫秒)
	EndTime   int64   `json:"pingouEndTime,omitempty"`   // 拼购结束时间(时间戳，毫秒)
}

type JFShop struct {
	Name                          string  `json:"shopName,omitempty"`                      // 店铺名称（或供应商名称）
	Id                            uint64  `json:"shopId,omitempty"`                        // 店铺Id
	ShopLevel                     float64 `json:"shopLevel,omitempty"`                     // 店铺评分
	ShopLabel                     string  `json:"shopLabel,omitempty"`                     // 1：京东好店 https://img12.360buyimg.com/schoolbt/jfs/t1/80828/19/2993/908/5d14277aEbb134d76/889d5265315e11ed.png
	UserEvaluateScore             string  `json:"userEvaluateScore,omitempty"`             // 用户评价评分（仅pop店铺有值）
	CommentFactorScoreRankGrade   string  `json:"commentFactorScoreRankGrade,omitempty"`   // 用户评价评级（仅pop店铺有值）
	LogisticsLvyueScore           string  `json:"logisticsLvyueScore,omitempty"`           // 物流履约评分（仅pop店铺有值）
	LogisticsFactorScoreRankGrade string  `json:"logisticsFactorScoreRankGrade,omitempty"` // 物流履约评级（仅pop店铺有值）
	AfterServiceScore             string  `json:"afterServiceScore,omitempty"`             // 售后服务评分（仅pop店铺有值）
	AfsFactorScoreRankGrade       string  `json:"afsFactorScoreRankGrade,omitempty"`       // 售后服务评级（仅pop店铺有值）
	ScoreRankRate                 string  `json:"scoreRankRate,omitempty"`                 // 店铺风向标（仅pop店铺有值）
}

type JFPrice struct {
	Price             float64 `json:"price"`
	LowestPrice       float64 `json:"lowestPrice,omitempty"`       // 促销价
	LowestPriceType   float64 `json:"lowestPriceType,omitempty"`   // 促销价类型，1：无线价格；2：拼购价格； 3：秒杀价格；4：预售价格
	LowestCouponPrice float64 `json:"lowestCouponPrice,omitempty"` // 券后价（有无券都返回此字段）
	HistoryPriceDay   float64 `json:"historyPriceDay,omitempty"`   // 历史最低价天数（例：当前券后价最近180天最低）
}

type JFImageList struct {
	List       []JFImage `json:"imageList,omitempty"`
	WhiteImage string    `json:"whiteImage,omitempty"` // 白底图
}

type JFImage struct {
	Url string `json:"url,omitempty"`
}

type JFCouponList struct {
	List []JFCoupon `json:"couponList,omitempty"`
}

type JFCoupon struct {
	BindType      uint    `json:"bindType,omitempty"`      // 券种类 (优惠券种类：0 - 全品类，1 - 限品类（自营商品），2 - 限店铺，3 - 店铺限商品券)
	Discount      float64 `json:"discount,omitempty"`      // 券面额
	Link          string  `json:"link,omitempty"`          // 券链接
	PlatformType  uint    `json:"platformType,omitempty"`  // 券使用平台 (平台类型：0 - 全平台券，1 - 限平台券)
	Quota         float64 `json:"quota,omitempty"`         // 券消费限额
	GetStartTime  int64   `json:"getStartTime,omitempty"`  // 领取开始时间(时间戳，毫秒)
	GetEndTime    int64   `json:"getEndTime,omitempty"`    // 券领取结束时间(时间戳，毫秒)
	UseStartTime  int64   `json:"useStartTime,omitempty"`  // 券有效使用开始时间(时间戳，毫秒)
	UseEndTime    int64   `json:"useEndTime,omitempty"`    // 券有效使用结束时间(时间戳，毫秒)
	IsBest        uint    `json:"isBest,omitempty"`        // 最优优惠券，1：是；0：否
	HotValue      uint    `json:"hotValue,omitempty"`      // 券热度，值越大热度越高，区间:[0,10]
	IsInputCoupon uint    `json:"isInputCoupon,omitempty"` // 入参couponUrl优惠券链接搜索对应的券，1 是 ，0 否
	CouponStyle   uint    `json:"couponStyle,omitempty"`   // 优惠券分类 0：满减券，3：满折券
}

type JFCommission struct {
	Commission          float64 `json:"commission,omitempty"`          //佣金
	CommissionShare     float64 `json:"commissionShare,omitempty"`     // 佣金比例
	CouponCommission    float64 `json:"couponCommission,omitempty"`    // 券后佣金，（促销价-优惠券面额）*佣金比例
	PlusCommissionShare float64 `json:"plusCommissionShare,omitempty"` // plus佣金比例，plus用户购买推广者能获取到的佣金比例
	IsLock              uint    `json:"isLock,omitempty"`              // 是否锁定佣金比例：1是，0否
	StartTime           uint64  `json:"startTime,omitempty"`           // 计划开始时间（时间戳，毫秒）
	EndTime             uint64  `json:"endTime,omitempty"`             // 计划结束时间（时间戳，毫秒）
}

type JFCategory struct {
	Cid1     uint64 `json:"cid1,omitempty"`
	Cid1Name string `json:"cid1Name,omitempty"`
	Cid2     uint64 `json:"cid2,omitempty"`
	Cid2Name string `json:"cid2Name,omitempty"`
	Cid3     uint64 `json:"cid3,omitempty"`
	Cid3Name string `json:"cid3Name,omitempty"`
}

// 京东联盟精选优质商品，每日更新，可通过频道ID查询各个频道下的精选商品。
func JingfenQuery(req *JingfenQueryRequest) (*JingfenQueryResult, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := goods.NewJingfenQueryRequest()
	goodsReq := &goods.GoodsReq{
		EliteId:   req.EliteId,
		PageIndex: req.PageIndex,
		PageSize:  req.PageSize,
		SortName:  req.SortName,
		Sort:      req.Sort,
	}
	r.SetGoodsReq(goodsReq)

	result, err := client.Execute(r.Request, req.Session)
	if err != nil {
		return nil, err
	}
	var response JingfenQueryResponse
	err = json.Unmarshal(result, &response)
	if err != nil {
		return nil, err
	}
	if response.Data == nil || response.Data.Result == "" {
		return nil, errors.New("no result")
	}
	var resp JingfenQueryResult
	err = json.Unmarshal([]byte(response.Data.Result), &resp)
	if err != nil {
		return nil, err
	}
	if resp.Code != 200 {
		return nil, &api.ErrorResponnse{Code: strconv.FormatInt(resp.Code, 10), ZhDesc: resp.Message}
	}

	return &resp, nil
}
