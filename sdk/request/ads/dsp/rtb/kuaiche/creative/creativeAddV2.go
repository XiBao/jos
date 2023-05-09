package creative

import (
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/ads/dsp"
)

type KuaicheCreativeAddV2Request struct {
	Request *sdk.Request
}

type KuaicheCreativeAddV2RequestData struct {
	AdList        []KuaicheCreativeAddV2RequestAd `json:"adList"`                  // 创意集合,商品类型计划必须先填写创意，腰带店铺可不填写
	ShowSalesWord string                          `json:"showSalesWord,omitempty"` // 活动推广，推广文案，不能修改，和原单元活动推广文案保持一致传入
	GroupId       uint64                          `json:"groupId"`                 // 所属单元ID
	Url           string                          `json:"url,omitempty"`           // 活动推广，活动地址，不能修改，和原单元活动地址保持一致传入
}

type KuaicheCreativeAddV2RequestAd struct {
	Id          uint64 `json:"id,omitempty"`          // 创意id，无创意id代表为新增创意
	SkuId       string `json:"skuId"`                 // skuId
	Name        string `json:"name,omitempty"`        // 创意名称
	ImgUrl      string `json:"imgUrl,omitempty"`      // 图片地址,不填写则为sku默认主图，当sku为特殊类目时，需要自定义标题与图片
	CustomTitle string `json:"customTitle,omitempty"` // 创意标题，长度为10-30字符,不填写则为sku默认标题，当sku为特殊类目时，需要自定义标题与图片
}

type KuaicheCreativeAddV2RequestCrowdVO struct {
	IsUsed uint `json:"isUsed"` // 是否启用智能定向，0：不启用，1：启用
}

// create new request
func NewKuaicheCreativeAddV2Request() (req *KuaicheCreativeAddV2Request) {
	request := sdk.Request{MethodName: "jingdong.ads.dsp.rtb.kc.ad.edit.v2", Params: make(map[string]interface{}, 2)}
	req = &KuaicheCreativeAddV2Request{
		Request: &request,
	}
	return
}

func (req *KuaicheCreativeAddV2Request) SetData(data *KuaicheCreativeAddV2RequestData) {
	req.Request.Params["data"] = data
}

func (req *KuaicheCreativeAddV2Request) GetData() *KuaicheCreativeAddV2RequestData {
	data, found := req.Request.Params["data"]
	if found {
		return data.(*KuaicheCreativeAddV2RequestData)
	}
	return nil
}

func (req *KuaicheCreativeAddV2Request) SetSystem(system *dsp.JdDspPlatformGatewayApiVoParamExt) {
	req.Request.Params["system"] = system
}

func (req *KuaicheCreativeAddV2Request) GetSystem() *dsp.JdDspPlatformGatewayApiVoParamExt {
	system, found := req.Request.Params["system"]
	if found {
		return system.(*dsp.JdDspPlatformGatewayApiVoParamExt)
	}
	return nil
}
