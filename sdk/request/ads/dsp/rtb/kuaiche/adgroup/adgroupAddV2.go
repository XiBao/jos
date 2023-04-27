package adgroup

import (
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/ads/dsp"
)

type KuaicheAdgroupAddV2Request struct {
	Request *sdk.Request
}

type KuaicheAdgroupAddV2RequestData struct {
	AdList               []KuaicheAdgroupAddV2RequestAd     `json:"adList,omitempty"`             // 创意集合,商品类型计划必须先填写创意，腰带店铺可不填写
	NewAreaIds           []string                           `json:"newAreaIds,omitempty"`         // 地域id列表,可通过jingdong.ads.dsp.rtb.area.query.v2接口获取对应的地域id
	CrowdVO              *KuaicheAdgroupAddV2RequestCrowdVO `json:"crowdVO"`                      // 推荐人群，默认开启一个智能人群（推荐人群已下线）
	FeeDecimal           uint                               `json:"feeDecimal"`                   // 单元pc出价
	InSearchFee          uint                               `json:"inSearchFee"`                  // 搜索广告位 PC端出价
	CampaignId           uint64                             `json:"campaignId"`                   // 商品类型的计划id
	MobilePriceCoef      uint                               `json:"mobilePriceCoef"`              // 无线出价系数
	Name                 string                             `json:"name"`                         // 单元名称
	AutomatedBiddingType uint                               `json:"automatedBiddingType"`         // 出价控制方式，2：系统智能调价（投放目标为自定义的系统智能托管） 256：预算控制(投方目标可选择点击，下单，成交) 若使用智能出价必填，腰带店铺填写0：手动出价
	DeliveryTarget       uint                               `json:"deliveryTarget,omitempty"`     // 投放目标： 1, 自定义, 2, 点击, 3,下单, 4, 成交，选择智能出价控制必须填写
	DayCost              uint                               `json:"dayCost"`                      // 统一日消耗（范围50-计划设置的日预算值）,出价控制方式为预算控制必须填写
	TcpaMaxClickBid      uint                               `json:"tcpaMaxClickBid,omitempty"`    // CPC上限(范围 2-9999)，选择预算控制非系统托管需要填写
	PremiumCoef          uint                               `json:"premiumCoef,omitempty"`        // 溢价系数，投放目标为自定义，出价控制方式为系统智能调价选择自定义上限必填 30-300
	BiddingControlType   uint                               `json:"biddingControlType,omitempty"` // 控制方式，1系统托管，3指定上限，若指定上限需和tcpaMaxClickBid配合使用
	OrientationRange     uint                               `json:"orientationRange,omitempty"`   // 智能出价定向范围，投放目标为自定义，出价控制方式为智能调价必填，1：关键词定向，2：商品定向，3：关键词+商品定向
	ShopId               uint64                             `json:"shopId,omitempty"`             // 店铺id，腰带店铺单元必填
}

type KuaicheAdgroupAddV2RequestAd struct {
	SkuId       string `json:"skuId"`                 // skuId
	AdName      string `json:"adName"`                // 创意名称
	ImgUrl      string `json:"imgUrl,omitempty"`      // 图片地址,不填写则为sku默认主图，当sku为特殊类目时，需要自定义标题与图片
	CustomTitle string `json:"customTitle,omitempty"` // 创意标题，长度为10-30字符,不填写则为sku默认标题，当sku为特殊类目时，需要自定义标题与图片
}

type KuaicheAdgroupAddV2RequestCrowdVO struct {
	IsUsed uint `json:"isUsed"` // 是否启用智能定向，0：不启用，1：启用
}

// create new request
func NewKuaicheAdgroupAddV2Request() (req *KuaicheAdgroupAddV2Request) {
	request := sdk.Request{MethodName: "jingdong.ads.dsp.rtb.kuaiche.productgroup.add.v2", Params: make(map[string]interface{}, 2)}
	req = &KuaicheAdgroupAddV2Request{
		Request: &request,
	}
	return
}

func (req *KuaicheAdgroupAddV2Request) SetData(data *KuaicheAdgroupAddV2RequestData) {
	req.Request.Params["data"] = data
}

func (req *KuaicheAdgroupAddV2Request) GetData() *KuaicheAdgroupAddV2RequestData {
	data, found := req.Request.Params["data"]
	if found {
		return data.(*KuaicheAdgroupAddV2RequestData)
	}
	return nil
}

func (req *KuaicheAdgroupAddV2Request) SetSystem(system *dsp.JdDspPlatformGatewayApiVoParamExt) {
	req.Request.Params["system"] = system
}

func (req *KuaicheAdgroupAddV2Request) GetSystem() *dsp.JdDspPlatformGatewayApiVoParamExt {
	system, found := req.Request.Params["system"]
	if found {
		return system.(*dsp.JdDspPlatformGatewayApiVoParamExt)
	}
	return nil
}
