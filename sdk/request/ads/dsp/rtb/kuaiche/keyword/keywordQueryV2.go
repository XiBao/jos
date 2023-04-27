package keyword

import (
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/ads/dsp"
)

type KuaicheKeywordQueryV2Request struct {
	Request *sdk.Request
}

type KuaicheKeywordQueryV2RequestData struct {
	CampaignId uint64 `json:"campaignId"` // 计划id
	AdgroupId  uint64 `json:"adGroupId"`  // 单元id
}

// create new request
func NewKuaicheKeywordQueryV2Request() (req *KuaicheKeywordQueryV2Request) {
	request := sdk.Request{MethodName: "jingdong.ads.dsp.rtb.keyword.query.v2", Params: make(map[string]interface{}, 2)}
	req = &KuaicheKeywordQueryV2Request{
		Request: &request,
	}
	return
}

func (req *KuaicheKeywordQueryV2Request) SetData(data *KuaicheKeywordQueryV2RequestData) {
	req.Request.Params["data"] = data
}

func (req *KuaicheKeywordQueryV2Request) GetData() *KuaicheKeywordQueryV2RequestData {
	data, found := req.Request.Params["data"]
	if found {
		return data.(*KuaicheKeywordQueryV2RequestData)
	}
	return nil
}

func (req *KuaicheKeywordQueryV2Request) SetSystem(system *dsp.JdDspPlatformGatewayApiVoParamExt) {
	req.Request.Params["system"] = system
}

func (req *KuaicheKeywordQueryV2Request) GetSystem() *dsp.JdDspPlatformGatewayApiVoParamExt {
	system, found := req.Request.Params["system"]
	if found {
		return system.(*dsp.JdDspPlatformGatewayApiVoParamExt)
	}
	return nil
}
