package promotion

import (
	"github.com/XiBao/jos/sdk"
)

type SellerPromotionUnitCreateRequest struct {
	Request *sdk.Request
}

type PromotionUnitClientExtentsion struct {
	ApiVersion string `json:"apiVersion"`
}

type PromotionUnitClient struct {
	AppName          string                         `json:"appName"`
	BusinessIdentity string                         `json:"businessIdentity"`
	AppId            string                         `json:"appId,omitempty"`
	UserAgent        string                         `json:"userAgent,omitempty"`
	Uuid             string                         `json:"uuid,omitempty"`
	Ip               string                         `json:"ip,omitempty"`
	Site             string                         `json:"site,omitempty"`
	Language         string                         `json:"language,omitempty"`
	Extension        *PromotionUnitClientExtentsion `json:"extension,omitempty"`
}

type PromotionUnitParamSku struct {
	PromoPrice float64 `json:"promoPrice"`
	SkuId      uint64  `json:"skuId"`
	LimitNum   uint    `json:"limitNum,omitempty"`
}

type PromotionUnitParamMember struct {
	UserClass uint `json:"userClass,omitempty"`
	UserLevel uint `json:"userLevel,omitempty"`
}

type PromotionUnitParamPlatform struct {
	SelectType        uint8  `json:"selectType,omitempty"`        // 0：全平台 1：限平台。如果不传，默认为全平台；如果传0，则channel，channelSelectType，platform参数可以不传
	Channel           int8   `json:"channel,omitempty"`           // 推广渠道 -1:全部渠道 1:京东全平台
	ChannelSelectType uint8  `json:"channelSelectType,omitempty"` // 推广平台下的渠道选择方式 0:全部 1:限选
	Platform          string `json:"platform,omitempty"`          // 推广平台，多个之间用逗号隔开
}

type PromotionUnitParam struct {
	Name             string                    `json:"name"`
	BeginTime        string                    `json:"beginTime"`
	EndTime          string                    `json:"endTime"`
	PromoLabel       uint8                     `json:"promoLabel"` // 促销标签 1-表示限量 2-表示限时 0-表示普通
	BusiCode         string                    `json:"busiCode"`
	PreviewTime      string                    `json:"previewTime,omitempty"`
	PerMaxNum        uint                      `json:"perMaxNum,omitempty"`
	PerMinNum        uint                      `json:"perMinNum,omitempty"`
	LimitWay         uint8                     `json:"limitWay,omitempty"` // 1、限ip、账号 2、限ip 3、限账号
	BuyLimit         bool                      `json:"buyLimit"`           // 是否设置售空数量 true:表示设置售空数量 false:表示不设置售空数量
	Slogan           string                    `json:"slogan,omitempty"`
	UnitPromoSkuList []PromotionUnitParamSku   `json:"unitPromoSkuList"`
	Member           *PromotionUnitParamMember `json:"member,omitempty"`
}

// create new request
func NewSellerPromotionUnitCreateRequest() (req *SellerPromotionUnitCreateRequest) {
	request := sdk.Request{MethodName: "jingdong.pop.promo.unit.createUnitPromo", Params: make(map[string]interface{}, 2)}
	req = &SellerPromotionUnitCreateRequest{
		Request: &request,
	}
	return
}

func (req *SellerPromotionUnitCreateRequest) SetClient(client *PromotionUnitClient) {
	req.Request.Params["client"] = client
}

func (req *SellerPromotionUnitCreateRequest) GetClient() *PromotionUnitClient {
	client, found := req.Request.Params["client"]
	if found {
		return client.(*PromotionUnitClient)
	}
	return nil
}

func (req *SellerPromotionUnitCreateRequest) SetPromoParam(param *PromotionUnitParam) {
	req.Request.Params["promoParam"] = param
}

func (req *SellerPromotionUnitCreateRequest) GetPromoParam() *PromotionUnitParam {
	param, found := req.Request.Params["promoParam"]
	if found {
		return param.(*PromotionUnitParam)
	}
	return nil
}
