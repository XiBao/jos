package keyword

import (
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/ads/dsp"
)

type KuaicheKeywordAddV2Request struct {
	Request *sdk.Request
}

type KuaicheKeywordAddV2RequestKeyword struct {
	KeywordName  string  `json:"keywordName"`  // 关键词名称
	KeywordPrice float64 `json:"keywordPrice"` // 关键词出价,最多保留一位小数
	Type         uint    `json:"type"`         // 关键词类型 8：切词 4：短语匹配 1：精确匹配
}

type KuaicheKeywordAddV2RequestData struct {
	KeywordList []KuaicheKeywordAddV2RequestKeyword `json:"keywordList"` // 关键词列表
	AdgroupId   uint64                              `json:"adGroupId"`   // 单元id
}

// create new request
func NewKuaicheKeywordAddV2Request() (req *KuaicheKeywordAddV2Request) {
	request := sdk.Request{MethodName: "jingdong.ads.dsp.rtb.keyword.add.v2", Params: make(map[string]interface{}, 2)}
	req = &KuaicheKeywordAddV2Request{
		Request: &request,
	}
	return
}

func (req *KuaicheKeywordAddV2Request) SetData(data *KuaicheKeywordAddV2RequestData) {
	req.Request.Params["data"] = data
}

func (req *KuaicheKeywordAddV2Request) GetData() *KuaicheKeywordAddV2RequestData {
	data, found := req.Request.Params["data"]
	if found {
		return data.(*KuaicheKeywordAddV2RequestData)
	}
	return nil
}

func (req *KuaicheKeywordAddV2Request) SetSystem(system *dsp.JdDspPlatformGatewayApiVoParamExt) {
	req.Request.Params["system"] = system
}

func (req *KuaicheKeywordAddV2Request) GetSystem() *dsp.JdDspPlatformGatewayApiVoParamExt {
	system, found := req.Request.Params["system"]
	if found {
		return system.(*dsp.JdDspPlatformGatewayApiVoParamExt)
	}
	return nil
}
