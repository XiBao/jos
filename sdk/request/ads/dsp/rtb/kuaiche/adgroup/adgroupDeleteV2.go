package adgroup

import (
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/ads/dsp"
)

type KuaicheAdgroupDeleteV2Request struct {
	Request *sdk.Request
}

type KuaicheAdgroupDeleteV2RequestData struct {
	Ids []uint64 `json:"ids"` // 单元id集合
}

// create new request
func NewKuaicheAdgroupDeleteV2Request() (req *KuaicheAdgroupDeleteV2Request) {
	request := sdk.Request{MethodName: "jingdong.ads.dsp.rtb.kuaiche.group.delete.v2", Params: make(map[string]interface{}, 2)}
	req = &KuaicheAdgroupDeleteV2Request{
		Request: &request,
	}
	return
}

func (req *KuaicheAdgroupDeleteV2Request) SetData(data *KuaicheAdgroupDeleteV2RequestData) {
	req.Request.Params["data"] = data
}

func (req *KuaicheAdgroupDeleteV2Request) GetData() *KuaicheAdgroupDeleteV2RequestData {
	data, found := req.Request.Params["data"]
	if found {
		return data.(*KuaicheAdgroupDeleteV2RequestData)
	}
	return nil
}

func (req *KuaicheAdgroupDeleteV2Request) SetSystem(system *dsp.JdDspPlatformGatewayApiVoParamExt) {
	req.Request.Params["system"] = system
}

func (req *KuaicheAdgroupDeleteV2Request) GetSystem() *dsp.JdDspPlatformGatewayApiVoParamExt {
	system, found := req.Request.Params["system"]
	if found {
		return system.(*dsp.JdDspPlatformGatewayApiVoParamExt)
	}
	return nil
}
