package goods

import (
	"context"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/union/goods"
)

type QueryRequest struct {
	api.BaseRequest
	Cid1                 uint64   `json:"cid1,omitempty"`                 // 一级类目id
	Cid2                 uint64   `json:"cid2,omitempty"`                 // 二级类目id
	Cid3                 uint64   `json:"cid3,omitempty"`                 // 三级类目id
	PageIndex            uint     `json:"pageIndex,omitempty"`            // 页码，默认1
	PageSize             uint     `json:"pageSize,omitempty"`             // 每页数量，单页数最大30，默认20
	SkuIds               []uint64 `json:"skuIds,omitempty"`               // skuid集合(一次最多支持查询20个sku)，数组类型开发时记得加[]
	Keyword              string   `json:"keyword,omitempty"`              // 关键词，字数同京东商品名称一致，目前未限制字数个数；支持cps长链接查询（https://union-click.jd.com）；支持cps短链接查询（https://u.jd.com）;
	PriceFrom            float64  `json:"pricefrom,omitempty"`            // 商品券后价格下限
	PriceTo              float64  `json:"priceto,omitempty"`              // 商品券后价格上限
	CommissionShareStart uint     `json:"commissionShareStart,omitempty"` // 佣金比例区间开始
	CommissionShareEnd   uint     `json:"commissionShareEnd,omitempty"`   // 佣金比例区间结束
	Owner                string   `json:"owner,omitempty"`                // 商品类型：自营[g]，POP[p]
	SortName             string   `json:"sortName,omitempty"`             // 排序字段(price：单价, commissionShare：佣金比例, commission：佣金， inOrderCount30Days：30天引单量， inOrderComm30Days：30天支出佣金)
	Sort                 string   `json:"sort,omitempty"`                 // asc,desc升降序,默认降序
	IsCoupon             uint     `json:"isCoupon,omitempty"`             // 是否是优惠券商品，1：有优惠券
	IsPG                 uint     `json:"isPG,omitempty"`                 // 是否是拼购商品，1：拼购商品
	PingouPriceStart     float64  `json:"pingouPriceStart,omitempty"`     // 拼购价格区间开始
	PingouPriceEnd       float64  `json:"pingouPriceEnd,omitempty"`       // 拼购价格区间结束
	IsHot                uint     `json:"isHot,omitempty"`                // 已废弃，请勿使用
	BrandCode            string   `json:"brandCode,omitempty"`            // 品牌code
	ShopId               uint64   `json:"shopId,omitempty"`               // 店铺Id
	HasContent           uint     `json:"hasContent,omitempty"`           // 1：查询内容商品；其他值过滤掉此入参条件。
	HasBestCoupon        uint     `json:"hasBestCoupon,omitempty"`        // 1：查询有最优惠券商品；其他值过滤掉此入参条件。（查询最优券需与isCoupon同时使用）
	Pid                  string   `json:"pid,omitempty"`                  // 联盟id_应用iD_推广位id
	Fields               string   `json:"fields,omitempty"`               // 支持出参数据筛选，逗号','分隔，目前可用：videoInfo(视频信息),hotWords(热词),similar(相似推荐商品),documentInfo(段子信息，智能文案),skuLabelInfo（商品标签）,promotionLabelInfo（商品促销标签）,stockState（商品库存）,companyType（小店标识）,purchasePriceInfo（到手价）,purchaseBPriceInfo（普惠到手价）
	ForbidTypes          string   `json:"forbidTypes,omitempty"`          // 10微信京东购物小程序禁售，11微信京喜小程序禁售
	JxFlags              []uint   `json:"jxFlags,omitempty"`              // 京喜商品类型，1京喜、2京喜工厂直供、3京喜优选，入参多个值表示或条件查询
	ShopLevelFrom        float64  `json:"shopLevelFrom,omitempty"`        // 支持传入0.0、2.5、3.0、3.5、4.0、4.5、4.9，默认为空表示不筛选评分
	Isbn                 string   `json:"isbn,omitempty"`                 // 图书编号
	SpuId                uint64   `json:"spuId,omitempty"`                // 主商品spuId
	CouponUrl            string   `json:"couponUrl,omitempty"`            // 优惠券链接
	DeliveryType         uint     `json:"deliveryType,omitempty"`         // 京东配送 1：是，0：不是
	EliteType            []uint   `json:"eliteType,omitempty"`            // 资源位17：极速版商品
	IsSeckill            uint     `json:"isSeckill,omitempty"`            // 是否秒杀商品。1：是
	IsPresale            uint     `json:"isPresale,omitempty"`            // 是否预售商品。1：是
	IsReserve            uint     `json:"isReserve,omitempty"`            // 是否预约商品。1:是
	BonusId              uint64   `json:"bonusId,omitempty"`              // 奖励活动ID
	Area                 string   `json:"area,omitempty"`                 // 区域地址（查区域价格）
	IsOversea            uint     `json:"isOversea,omitempty"`            // 是否全球购商品 1：是
	UserIdType           uint     `json:"userIdType,omitempty"`           // 用户ID类型，传入此参数可获得个性化推荐结果。当前userIdType支持的枚举值包括：8、16、32、64、128、32768。userIdType和userId需同时传入，且一一对应。userIdType各枚举值对应的userId含义如下：8(安卓移动设备Imei); 16(苹果移动设备Openudid)；32(苹果移动设备idfa); 64(安卓移动设备imei的md5编码，32位，大写，匹配率略低);128(苹果移动设备idfa的md5编码，32位，大写，匹配率略低); 32768(安卓移动设备oaid); 131072(安卓移动设备oaid的md5编码，32位，大写)
	UserId               string   `json:"userId,omitempty"`               // userIdType对应的用户设备ID，传入此参数可获得个性化推荐结果，userIdType和userId需同时传入
	ChannelId            uint64   `json:"channelId,omitempty"`            // 渠道关系ID
	Ip                   string   `json:"ip,omitempty"`                   // 客户端ip
	ProvinceId           uint     `json:"provinceId,omitempty"`           // 省Id
	CityId               uint     `json:"cityId,omitempty"`               // 市Id
	CountryId            uint     `json:"countryId,omitempty"`            // 县Id
	TownId               uint     `json:"townId,omitempty"`               // 镇Id
	ItemIds              []string `json:"itemIds,omitempty"`              // 联盟商品ID集合(一次最多支持查询20个itemId)，为字符串数组类型，开发时记得加[]
}

type QueryResponse struct {
	ErrorResp *api.ErrorResponnse `json:"error_response,omitempty"`
	Data      *QueryResponseData  `json:"jd_union_open_goods_query_responce,omitempty"`
}

func (r QueryResponse) IsError() bool {
	return r.ErrorResp != nil || r.Data == nil || r.Data.Result == ""
}

func (r QueryResponse) Error() string {
	if r.ErrorResp != nil {
		return r.ErrorResp.Error()
	}
	return "no result data"
}

type QueryResponseData struct {
	Result string `json:"queryResult,omitempty"`
}

type QueryResult struct {
	Code       int64       `json:"code,omitempty"`
	Message    string      `json:"message,omitempty"`
	TotalCount uint        `json:"totalCount,omitempty"`
	Data       []GoodsResp `json:"data,omitempty"`
}

func (r QueryResult) IsError() bool {
	return r.Code != 200
}

func (r QueryResult) Error() string {
	return sdk.ErrorString(r.Code, r.Message)
}

type GoodsResp struct {
	Category           *JFCategory        `json:"categoryInfo,omitempty"`           // 类目信息
	Comments           uint               `json:"comments,omitempty"`               // 评论数
	Commission         *JFCommission      `json:"commissionInfo,omitempty"`         // 佣金信息
	Coupon             *JFCouponList      `json:"couponInfo,omitempty"`             // 优惠券信息，返回内容为空说明该SKU无可用优惠券
	GoodCommentsShare  float64            `json:"goodCommentsShare,omitempty"`      // 商品好评率
	Images             *JFImageList       `json:"imageInfo,omitempty"`              // 图片信息
	InOrderCount30Days uint               `json:"inOrderCount30Days,omitempty"`     // 30天引单数量
	MaterialUrl        string             `json:"materialUrl,omitempty"`            // 商品落地页，当入参为联盟商品ID时，materialUrl拼接为：jingfen.jd.com/detail/{itemId}.html；当入参为原始商品ID时，materialUrl拼接为：item.jd.com/{skuId}.html
	Price              *JFPrice           `json:"priceInfo,omitempty"`              // 价格信息
	Shop               *JFShop            `json:"shopInfo,omitempty"`               // 店铺信息
	SkuId              uint64             `json:"skuId,omitempty"`                  // 商品ID
	SkuName            string             `json:"skuName,omitempty"`                // 商品名称
	IsHot              uint               `json:"isHot,omitempty"`                  // 是否爆款，1：是，0：否 已废弃，请勿使用
	ProductId          uint64             `json:"spuid,omitempty"`                  // spuid，其值为同款商品的主skuid
	BrandCode          string             `json:"brandCode,omitempty"`              // 品牌code
	BrandName          string             `json:"brandName,omitempty"`              // 品牌名
	Owner              string             `json:"owner,omitempty"`                  // g=自营，p=pop
	Pingou             *JFPingou          `json:"pinGouInfo,omitempty"`             // 拼购信息
	Video              *JFVideoList       `json:"videoInfo,omitempty"`              // 视频信息
	Comment            *JFCommentList     `json:"commentInfo,omitempty"`            // 评价信息
	JxFlags            []uint             `json:"jxFlags,omitempty"`                // 京喜商品类型，1京喜、2京喜工厂直供、3京喜优选（包含3时可在京东APP购买）
	Document           *JFDocument        `json:"documentInfo,omitempty"`           // 商品段子信息，emoji表情等
	Book               *JFBook            `json:"bookInfo,omitempty"`               // 图书信息
	Spec               *JFSpec            `json:"specInfo,omitempty"`               // 扩展信息
	StockState         uint               `json:"stockState,omitempty"`             // 库存状态：1有货、0无货（供tob选品场景参考，toc场景不适用）
	EliteType          []int              `json:"eliteType,omitempty"`              // 资源位17：极速版商品
	ForbidTypes        []int              `json:"forbidTypes,omitempty"`            // 0普通商品，10微信京东购物小程序禁售，11微信京喜小程序禁售
	DeliveryType       uint               `json:"deliveryType,omitempty"`           // 京东配送 1：是，0：不是
	SkuLabel           *JFSkuLabel        `json:"skuLabelInfo,omitempty"`           // 商品标签
	PromotionLabel     []JFPromotionLabel `json:"promotionLabelInfoList,omitempty"` // 商品促销标签集
	SecondPrice        []JFSecondPrice    `json:"secondPriceInfoList,omitempty"`    // 双价格
	Seckill            *JFSeckill         `json:"seckillInfo,omitempty"`            // 秒杀信息
	PreSale            *JFPreSale         `json:"preSaleInfo,omitempty"`            // 预售信息
	Reserve            *JFReserve         `json:"reserveInfo,omitempty"`            // 预约信息
	IsOversea          uint               `json:"isOversea,omitempty"`              // 是否全球购商品 1：是
	CompanyType        uint               `json:"companyType,omitempty"`            // 2：POP自然人小店
	PurchasePrice      *JFPurchasePrice   `json:"purchasePriceInfo,omitempty"`      // 到手价明细
	ItemId             string             `json:"itemId,omitempty"`                 // 联盟商品ID
	ActivityCard       *JFActivityCard    `json:"activityCardInfo,omitempty"`       // 超市购物卡明细
	SmartDocument      []JFSmartDocument  `json:"smartDocumentInfoList,omitempty"`  // GPT算法智能生成的推广文案集合,内含多个风格
	TotalCount         uint               `json:"totalCount,omitempty"`             // 有效商品总数量，上限1w
	HotWords           string             `json:"hotWords,omitempty"`               // 日常top10的热搜词，按小时更新
	SimilarSkuList     []uint64           `json:"similarSkuList,omitempty"`         // 相似推荐商品skuId集合
}

// 查询商品及优惠券信息，返回的结果可调用转链接口生成单品或二合一推广链接。支持按SKUID、关键词、优惠券基本属性、是否拼购、是否爆款等条件查询，建议不要同时传入SKUID和其他字段，以获得较多的结果。支持按价格、佣金比例、佣金、引单量等维度排序。用优惠券链接调用转链接口时，需传入搜索接口link字段返回的原始优惠券链接，切勿对链接进行任何encode、decode操作，否则将导致转链二合一推广链接时校验失败。
func Query(ctx context.Context, req *QueryRequest) (*QueryResult, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := goods.NewQueryRequest()
	goodsReq := &goods.QueryReq{
		Cid1:                 req.Cid1,
		Cid2:                 req.Cid2,
		Cid3:                 req.Cid3,
		PageIndex:            req.PageIndex,
		PageSize:             req.PageSize,
		SkuIds:               req.SkuIds,
		Keyword:              req.Keyword,
		PriceFrom:            req.PriceFrom,
		PriceTo:              req.PriceTo,
		CommissionShareStart: req.CommissionShareStart,
		CommissionShareEnd:   req.CommissionShareEnd,
		Owner:                req.Owner,
		SortName:             req.SortName,
		Sort:                 req.Sort,
		IsCoupon:             req.IsCoupon,
		IsPG:                 req.IsPG,
		PingouPriceStart:     req.PingouPriceStart,
		PingouPriceEnd:       req.PingouPriceEnd,
		IsHot:                req.IsHot,
		BrandCode:            req.BrandCode,
		ShopId:               req.ShopId,
		HasContent:           req.HasContent,
		HasBestCoupon:        req.HasBestCoupon,
		Pid:                  req.Pid,
		Fields:               req.Fields,
		ForbidTypes:          req.ForbidTypes,
		JxFlags:              req.JxFlags,
		ShopLevelFrom:        req.ShopLevelFrom,
		Isbn:                 req.Isbn,
		SpuId:                req.SpuId,
		CouponUrl:            req.CouponUrl,
		DeliveryType:         req.DeliveryType,
		EliteType:            req.EliteType,
		IsSeckill:            req.IsSeckill,
		IsPresale:            req.IsPresale,
		IsReserve:            req.IsReserve,
		BonusId:              req.BonusId,
		Area:                 req.Area,
		IsOversea:            req.IsOversea,
		UserIdType:           req.UserIdType,
		UserId:               req.UserId,
		ChannelId:            req.ChannelId,
		Ip:                   req.Ip,
		ProvinceId:           req.ProvinceId,
		CityId:               req.CityId,
		CountryId:            req.CountryId,
		TownId:               req.TownId,
		ItemIds:              req.ItemIds,
	}
	r.SetGoodsReqDTO(goodsReq)

	var response QueryResponse
	if err := client.Execute(ctx, r.Request, req.Session, &response); err != nil {
		return nil, err
	}
	var resp QueryResult
	if err := client.Logger().DecodeJSON([]byte(response.Data.Result), &resp); err != nil {
		return nil, err
	}
	if resp.IsError() {
		return nil, resp
	}

	return &resp, nil
}
