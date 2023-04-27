package keyword

import (
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/ads/dsp"
)

type KuaicheKeywordSkuRecommendListV2Request struct {
	Request *sdk.Request
}

type KuaicheKeywordSkuRecommendListV2RequestData struct {
	AdgroupId      uint64 `json:"adGroupId"`      // 单元id
	SkuId          uint64 `json:"skuId"`          // skuId
	DevType        uint   `json:"devType"`        // 设备类型,1=全部，2=无线，3=pc
	AdKeywordTypes []uint `json:"adKeywordTypes"` // 关键词标签，热词：12或者20，潜力词：72或者80，热+潜：76或者84，无标签：8或者16
}

// create new request
func NewKuaicheKeywordSkuRecommendListV2Request() (req *KuaicheKeywordSkuRecommendListV2Request) {
	request := sdk.Request{MethodName: "jingdong.ads.dsp.rtb.keyword.sku.recommend.list.v2", Params: make(map[string]interface{}, 2)}
	req = &KuaicheKeywordSkuRecommendListV2Request{
		Request: &request,
	}
	return
}

func (req *KuaicheKeywordSkuRecommendListV2Request) SetData(data *KuaicheKeywordSkuRecommendListV2RequestData) {
	req.Request.Params["data"] = data
}

func (req *KuaicheKeywordSkuRecommendListV2Request) GetData() *KuaicheKeywordSkuRecommendListV2RequestData {
	data, found := req.Request.Params["data"]
	if found {
		return data.(*KuaicheKeywordSkuRecommendListV2RequestData)
	}
	return nil
}

func (req *KuaicheKeywordSkuRecommendListV2Request) SetSystem(system *dsp.JdDspPlatformGatewayApiVoParamExt) {
	req.Request.Params["system"] = system
}

func (req *KuaicheKeywordSkuRecommendListV2Request) GetSystem() *dsp.JdDspPlatformGatewayApiVoParamExt {
	system, found := req.Request.Params["system"]
	if found {
		return system.(*dsp.JdDspPlatformGatewayApiVoParamExt)
	}
	return nil
}
