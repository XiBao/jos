package adgroup

import (
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/ads/dsp"
)

type KuaicheAdgroupUpdateFeeV2Request struct {
	Request *sdk.Request
}

type KuaicheAdgroupUpdateFeeV2RequestAutoBiddingModuleVO struct {
	DayCost            uint `json:"dayCost"`            // 统一日消耗（范围50-计划设置的日预算值）,出价控制方式为预算控制必须填写
	TcpaMaxClickBid    uint `json:"tcpaMaxClickBid"`    // CPC上限(范围 2-9999)，选择预算控制非系统托管需要填写
	BiddingControlType uint `json:"biddingControlType"` // 控制方式，1系统托管，3指定上限，若指定上限需和tcpaMaxClickBid配合使用
}

type KuaicheAdgroupUpdateFeeV2RequestData struct {
	MobilePriceCoef     uint64                                               `json:"mobilePriceCoef"`               // 快车单元无线出价系数
	InSearchFee         float64                                              `json:"inSearchFee"`                   // 快车pc搜索出价
	Id                  uint64                                               `json:"id"`                            // 快车单元id
	RecommendFee        float64                                              `json:"recommendFee"`                  // 快车pc推荐出价
	AutoBiddingModuleVO *KuaicheAdgroupUpdateFeeV2RequestAutoBiddingModuleVO `json:"autoBiddingModuleVO,omitempty"` // 智能出价相关信息
}

// create new request
func NewKuaicheAdgroupUpdateFeeV2Request() (req *KuaicheAdgroupUpdateFeeV2Request) {
	request := sdk.Request{MethodName: "jingdong.ads.dsp.rtb.kuaiche.group.updatefee.v2", Params: make(map[string]interface{}, 2)}
	req = &KuaicheAdgroupUpdateFeeV2Request{
		Request: &request,
	}
	return
}

func (req *KuaicheAdgroupUpdateFeeV2Request) SetData(data *KuaicheAdgroupUpdateFeeV2RequestData) {
	req.Request.Params["data"] = data
}

func (req *KuaicheAdgroupUpdateFeeV2Request) GetData() *KuaicheAdgroupUpdateFeeV2RequestData {
	data, found := req.Request.Params["data"]
	if found {
		return data.(*KuaicheAdgroupUpdateFeeV2RequestData)
	}
	return nil
}

func (req *KuaicheAdgroupUpdateFeeV2Request) SetSystem(system *dsp.JdDspPlatformGatewayApiVoParamExt) {
	req.Request.Params["system"] = system
}

func (req *KuaicheAdgroupUpdateFeeV2Request) GetSystem() *dsp.JdDspPlatformGatewayApiVoParamExt {
	system, found := req.Request.Params["system"]
	if found {
		return system.(*dsp.JdDspPlatformGatewayApiVoParamExt)
	}
	return nil
}
