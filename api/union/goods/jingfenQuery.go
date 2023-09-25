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
	Category              *JFCategory   `json:"categoryInfo,omitempty"`          // 类目信息
	Comments              uint          `json:"comments,omitempty"`              // 评论数
	GoodCommentsShare     float64       `json:"goodCommentsShare,omitempty"`     // 商品好评率
	Commission            *JFCommission `json:"commissionInfo,omitempty"`        // 佣金信息
	Coupon                *JFCouponList `json:"couponInfo,omitempty"`            // 优惠券信息，返回内容为空说明该SKU无可用优惠券
	Images                *JFImageList  `json:"imageInfo,omitempty"`             //图片信息
	InOrderCount30Days    uint          `json:"inOrderCount30Days,omitempty"`    // 30天引单数量
	InOrderCount30DaysSku uint          `json:"inOrderCount30DaysSku,omitempty"` // 30天引单数量(sku维度)
	MaterialUrl           string        `json:"materialUrl,omitempty"`           // 商品落地页
	Price                 JFPrice       `json:"priceInfo"`                       // 价格信息
	Shop                  *JFShop       `json:"shopInfo,omitempty"`              // 店铺信息
	SkuId                 uint64        `json:"skuId,omitempty"`
	SkuName               string        `json:"skuName,omitempty"`      // 商品名称
	IsHot                 uint          `json:"isHot,omitempty"`        // 是否爆款，1：是，0：否
	ProductId             uint64        `json:"spuid,omitempty"`        // 其值为同款商品的主skuid
	BrandCode             string        `json:"brandCode,omitempty"`    // 品牌code
	BrandName             string        `json:"brandName,omitempty"`    // 品牌名
	Owner                 string        `json:"owner,omitempty"`        // g=自营，p=pop
	Pingou                *JFPingou     `json:"pinGouInfo,omitempty"`   // 拼购信息
	Video                 *JFVideoList  `json:"videoInfo,omitempty"`    // 视频信息
	Resource              *JFResource   `json:"resourceInfo,omitempty"` // 资源信息
	Seckill               *JFSeckill    `json:"seckillInfo,omitempty"`  // 秒杀信息
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
	Height    uint   `json:"height,omitempty"`    // 高
	ImageUrl  string `json:"imageUrl,omitempty"`  // 视频图片地址
	VideoType uint   `json:"videoType,omitempty"` // 1:主图，2：商详
	PlayUrl   string `json:"playUrl,omitempty"`   // 播放地址
	PlayType  string `json:"playType,omitempty"`  // low：标清，high：高清
}

type JFPingou struct {
	Price     float64 `json:"pingouPrice,omitempty"`     // 拼购价格
	PingouUrl string  `json:"pingouUrl,omitempty"`       // 拼购落地页url
	TmCount   uint    `json:"pingouTmCount,omitempty"`   // 拼购成团所需人数
	StartTime int64   `json:"pingouStartTime,omitempty"` // 拼购开始时间(时间戳，毫秒)
	EndTime   int64   `json:"pingouEndTime,omitempty"`   // 拼购结束时间(时间戳，毫秒)
}

type JFShop struct {
	Name string `json:"shopName,omitempty"` // 店铺名称（或供应商名称）
	Id   uint64 `json:"shopId,omitempty"`   // 店铺Id
}
type JFPrice struct {
	Price             float64 `json:"price"`
	LowestPrice       float64 `json:"lowestPrice,omitempty"`       // 最低价格
	LowestPriceType   float64 `json:"lowestPriceType,omitempty"`   // 最低价格类型，1：无线价格；2：拼购价格； 3：秒杀价格
	LowestCouponPrice float64 `json:"lowestCouponPrice,omitempty"` // 最低价后的优惠券价
}

type JFImageList struct {
	List []JFImage `json:"imageList,omitempty"`
}

type JFImage struct {
	Url string `json:"url,omitempty"`
}

type JFCouponList struct {
	List []JFCoupon `json:"couponList,omitempty"`
}

type JFCoupon struct {
	BindType     uint    `json:"bindType,omitempty"`     // 券种类 (优惠券种类：0 - 全品类，1 - 限品类（自营商品），2 - 限店铺，3 - 店铺限商品券)
	Discount     float64 `json:"discount,omitempty"`     // 券面额
	Link         string  `json:"link,omitempty"`         // 券链接
	PlatformType uint    `json:"platformType,omitempty"` // 券使用平台 (平台类型：0 - 全平台券，1 - 限平台券)
	Quota        float64 `json:"quota,omitempty"`        // 券消费限额
	GetStartTime int64   `json:"getStartTime,omitempty"` // 领取开始时间(时间戳，毫秒)
	GetEndTime   int64   `json:"getEndTime,omitempty"`   // 券领取结束时间(时间戳，毫秒)
	UseStartTime int64   `json:"useStartTime,omitempty"` // 券有效使用开始时间(时间戳，毫秒)
	UseEndTime   int64   `json:"useEndTime,omitempty"`   // 券有效使用结束时间(时间戳，毫秒)
	IsBest       uint    `json:"isBest,omitempty"`       // 最优优惠券，1：是；0：否
}

type JFCommission struct {
	Commission       float64 `json:"commission,omitempty"`       //佣金
	Rate             float64 `json:"commissionShare,omitempty"`  // 佣金比例
	CouponCommission float64 `json:"couponCommission,omitempty"` // 券后佣金
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
