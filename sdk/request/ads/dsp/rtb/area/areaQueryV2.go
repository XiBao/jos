package area

import (
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/ads/dsp"
)

type AreaQueryV2Request struct {
	Request *sdk.Request
}

// create new request
func NewAreaQueryV2Request() (req *AreaQueryV2Request) {
	request := sdk.Request{MethodName: "jingdong.ads.dsp.rtb.area.query.v2", Params: make(map[string]interface{}, 1)}
	req = &AreaQueryV2Request{
		Request: &request,
	}
	return
}

func (req *AreaQueryV2Request) SetSystem(system *dsp.JdDspPlatformGatewayApiVoParamExt) {
	req.Request.Params["system"] = system
}

func (req *AreaQueryV2Request) GetSystem() *dsp.JdDspPlatformGatewayApiVoParamExt {
	system, found := req.Request.Params["system"]
	if found {
		return system.(*dsp.JdDspPlatformGatewayApiVoParamExt)
	}
	return nil
}
