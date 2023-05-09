package dmp

import (
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/ads/dsp"
)

type KuaicheDmpOperateRequest struct {
	Request *sdk.Request
}

type KuaicheDmpOperateRequestCrowdVOS struct {
	IsUsed       uint   `json:"isUsed"`                 // 是否启用 1启用 0 不启用
	AdGroupPrice uint   `json:"adGroupPrice,omitempty"` // 人群溢价
	CrowdId      uint64 `json:"crowdId"`                // 人群id
}

type KuaicheDmpOperateRequestDmpVO struct {
	JosOperate uint                               `json:"josOperate"` // jos人群操作类 0:解绑 1:绑定
	AdgroupId  uint64                             `json:"adGroupId"`  // 单元id
	CrowdVOS   []KuaicheDmpOperateRequestCrowdVOS `json:"crowdVOS"`   // 人群实体类集合
}

// create new request
func NewKuaicheDmpOperateRequest() (req *KuaicheDmpOperateRequest) {
	request := sdk.Request{MethodName: "jingdong.ads.dsp.rtb.kuaiche.dmp.operate", Params: make(map[string]interface{}, 2)}
	req = &KuaicheDmpOperateRequest{
		Request: &request,
	}
	return
}

func (req *KuaicheDmpOperateRequest) SetDmpVO(dmpVO *KuaicheDmpOperateRequestDmpVO) {
	req.Request.Params["dmpVO"] = dmpVO
}

func (req *KuaicheDmpOperateRequest) GetDmpVO() *KuaicheDmpOperateRequestDmpVO {
	dmpVO, found := req.Request.Params["dmpVO"]
	if found {
		return dmpVO.(*KuaicheDmpOperateRequestDmpVO)
	}
	return nil
}

func (req *KuaicheDmpOperateRequest) SetParamExt(paramExt *dsp.JdDspPlatformGatewayApiVoParamExt) {
	req.Request.Params["paramExt"] = paramExt
}

func (req *KuaicheDmpOperateRequest) GetParamExt() *dsp.JdDspPlatformGatewayApiVoParamExt {
	paramExt, found := req.Request.Params["paramExt"]
	if found {
		return paramExt.(*dsp.JdDspPlatformGatewayApiVoParamExt)
	}
	return nil
}
