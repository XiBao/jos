package keyword

import (
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/ads/dsp"
)

type KuaicheKeywordBootRecommendListV2Request struct {
	Request *sdk.Request
}

type KuaicheKeywordBootRecommendListV2RequestData struct {
	AdgroupId uint64 `json:"adGroupId"` // 单元id
	Keyword   string `json:"keyWord"`   // 核心词
}

// create new request
func NewKuaicheKeywordBootRecommendListV2Request() (req *KuaicheKeywordBootRecommendListV2Request) {
	request := sdk.Request{MethodName: "jingdong.ads.dsp.rtb.keyword.boot.recommend.list.v2", Params: make(map[string]interface{}, 2)}
	req = &KuaicheKeywordBootRecommendListV2Request{
		Request: &request,
	}
	return
}

func (req *KuaicheKeywordBootRecommendListV2Request) SetData(data *KuaicheKeywordBootRecommendListV2RequestData) {
	req.Request.Params["data"] = data
}

func (req *KuaicheKeywordBootRecommendListV2Request) GetData() *KuaicheKeywordBootRecommendListV2RequestData {
	data, found := req.Request.Params["data"]
	if found {
		return data.(*KuaicheKeywordBootRecommendListV2RequestData)
	}
	return nil
}

func (req *KuaicheKeywordBootRecommendListV2Request) SetSystem(system *dsp.JdDspPlatformGatewayApiVoParamExt) {
	req.Request.Params["system"] = system
}

func (req *KuaicheKeywordBootRecommendListV2Request) GetSystem() *dsp.JdDspPlatformGatewayApiVoParamExt {
	system, found := req.Request.Params["system"]
	if found {
		return system.(*dsp.JdDspPlatformGatewayApiVoParamExt)
	}
	return nil
}
