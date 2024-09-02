package goods

import (
	"github.com/XiBao/jos/sdk"
)

type QueryReq struct {
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

type QueryRequest struct {
	Request *sdk.Request
}

// create new request
func NewQueryRequest() (req *QueryRequest) {
	request := sdk.Request{MethodName: "jd.union.open.goods.query", IsUnionGW: true, Params: make(map[string]interface{}, 1)}
	req = &QueryRequest{
		Request: &request,
	}
	return
}

func (req *QueryRequest) SetGoodsReqDTO(goodsReq *QueryReq) {
	req.Request.Params["goodsReqDTO"] = goodsReq
}

func (req *QueryRequest) GetGoodsReqDTO() *QueryReq {
	goodsReq, found := req.Request.Params["goodsReqDTO"]
	if found {
		return goodsReq.(*QueryReq)
	}
	return nil
}
