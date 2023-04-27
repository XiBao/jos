package keyword

import (
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/ads/dsp"
)

type KuaicheKeywordUpdatePerPriceV2Request struct {
	Request *sdk.Request
}

type KuaicheKeywordUpdatePerPriceV2RequestData struct {
	KeywordMobilePrice float64 `json:"keywordMobilePrice"` // 关键词无线出价
	KeywordPrice       float64 `json:"keywordPrice"`       // 关键词pc出价
	AdGroupId          uint64  `json:"adGroupId"`          // 单元id
	Id                 uint64  `json:"id"`                 // 关键词id
	CampaignType       uint    `json:"campaignType"`       // 计划类型
}

// create new request
func NewKuaicheKeywordUpdatePerPriceV2Request() (req *KuaicheKeywordUpdatePerPriceV2Request) {
	request := sdk.Request{MethodName: "jingdong.ads.dsp.rtb.keyword.updatePerKeyword.v2", Params: make(map[string]interface{}, 2)}
	req = &KuaicheKeywordUpdatePerPriceV2Request{
		Request: &request,
	}
	return
}

func (req *KuaicheKeywordUpdatePerPriceV2Request) SetData(data []KuaicheKeywordUpdatePerPriceV2RequestData) {
	req.Request.Params["data"] = data
}

func (req *KuaicheKeywordUpdatePerPriceV2Request) GetData() []KuaicheKeywordUpdatePerPriceV2RequestData {
	data, found := req.Request.Params["data"]
	if found {
		return data.([]KuaicheKeywordUpdatePerPriceV2RequestData)
	}
	return nil
}

func (req *KuaicheKeywordUpdatePerPriceV2Request) SetSystem(system *dsp.JdDspPlatformGatewayApiVoParamExt) {
	req.Request.Params["system"] = system
}

func (req *KuaicheKeywordUpdatePerPriceV2Request) GetSystem() *dsp.JdDspPlatformGatewayApiVoParamExt {
	system, found := req.Request.Params["system"]
	if found {
		return system.(*dsp.JdDspPlatformGatewayApiVoParamExt)
	}
	return nil
}
