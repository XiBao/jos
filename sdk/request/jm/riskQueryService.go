package jm

import (
	"github.com/XiBao/jos/sdk"
)

type RiskQueryServiceRequest struct {
	Request *sdk.Request
}

type RiskQueryServiceRequestParam1 struct {
	AppName     string                 `json:"appName"`          // 系统字段，应用名
	Time        string                 `json:"time"`             // 事件时间，一般为当前时间
	ExtendMap   map[string]interface{} `json:"extendMap"`        // 业务扩展字段集合，一次最多传入20个key
	UseType     string                 `json:"useType"`          // 风控授权通过后，提供的接入点名称
	Pin         string                 `json:"pin"`              // jd账号
	SubSys      string                 `json:"subSys,emitempty"` // 场景名
	OpenIdBuyer string                 `json:"open_id_buyer"`    // jd账号
	XidBuyer    string                 `json:"xid_buyer"`        // jd账号
}

// create new request
func NewRiskQueryServiceRequest() (req *RiskQueryServiceRequest) {
	request := sdk.Request{MethodName: "jingdong.risk.RiskQueryService", Params: make(map[string]interface{}, 1)}
	req = &RiskQueryServiceRequest{
		Request: &request,
	}
	return
}

func (req *RiskQueryServiceRequest) SetParam1(param1 *RiskQueryServiceRequestParam1) {
	req.Request.Params["param1"] = param1
}

func (req *RiskQueryServiceRequest) GetParam1() *RiskQueryServiceRequestParam1 {
	param1, found := req.Request.Params["param1"]
	if found {
		return param1.(*RiskQueryServiceRequestParam1)
	}
	return nil
}
