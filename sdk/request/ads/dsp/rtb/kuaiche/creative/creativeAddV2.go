package creative

import (
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/ads/dsp"
)

type KuaicheCreativeAddV2Request struct {
	Request *sdk.Request
}

type KuaicheCreativeAddV2RequestData struct {
	AdList  []KuaicheCreativeAddV2RequestAd `json:"adList"`  // 创意集合,商品类型计划必须先填写创意，腰带店铺可不填写
	GroupId uint64                          `json:"groupId"` // 所属单元ID
}

type KuaicheCreativeAddV2RequestAd struct {
	CustomTitle   string `json:"customTitle,omitempty"`   // 创意标题，长度为10-30字符,不填写则为sku默认标题，当sku为特殊类目时，需要自定义标题与图片
	SkuId         string `json:"skuId"`                   // skuId
	Name          string `json:"name,omitempty"`          // 创意名称
	Url           string `json:"url,omitempty"`           // 落地页
	ImgUrl        string `json:"imgUrl,omitempty"`        // 图片地址,不填写则为sku默认主图，当sku为特殊类目时，需要自定义标题与图片
	ShowSalesWord string `json:"showSalesWord,omitempty"` // 推广文案
}

// create new request
func NewKuaicheCreativeAddV2Request() (req *KuaicheCreativeAddV2Request) {
	request := sdk.Request{MethodName: "jingdong.ads.dsp.rtb.kc.ad.add.v2", Params: make(map[string]interface{}, 2)}
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
