package dmp

import (
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/ads/dsp"
)

type KuaicheDmpListV2Request struct {
	Request *sdk.Request
}

type KuaicheDmpListV2RequestData struct {
	PageSize            int    `json:"pageSize,omitempty"`            // 每页条数
	ClickOrOrderDay     uint   `json:"clickOrOrderDay"`               // 转化周期 当天:0 1天:1 3天:3 7天:7 15天: 15
	CampaignId          uint64 `json:"campaignId,omitempty"`          // 计划id
	OrderStatusCategory uint   `json:"orderStatusCategory,omitempty"` // 订单类型 空：全部订单 1：成交订单
	EndDay              string `json:"endDay"`                        // 结束时间
	GroupId             uint64 `json:"groupId,omitempty"`             // 单元id
	StartDay            string `json:"startDay"`                      // 开始时间
	Page                int    `json:"page"`                          // 页码
	ClickOrOrderCaliber uint   `json:"clickOrOrderCaliber"`           // 口径，0：点击，1：下单
	GiftFlag            uint   `json:"giftFlag"`                      // 包含赠品 0不含赠品 1含赠品
	CampaignTypes       []uint `json:"campaignTypes,omitempty"`       // 计划类型，默认返回商品计划和腰带店铺计划类型人群
}

// create new request
func NewKuaicheDmpListV2Request() (req *KuaicheDmpListV2Request) {
	request := sdk.Request{MethodName: "jingdong.ads.dsp.rtb.kuaiche.dmplist.v2", Params: make(map[string]interface{}, 2)}
	req = &KuaicheDmpListV2Request{
		Request: &request,
	}
	return
}

func (req *KuaicheDmpListV2Request) SetData(data *KuaicheDmpListV2RequestData) {
	req.Request.Params["data"] = data
}

func (req *KuaicheDmpListV2Request) GetData() *KuaicheDmpListV2RequestData {
	data, found := req.Request.Params["data"]
	if found {
		return data.(*KuaicheDmpListV2RequestData)
	}
	return nil
}

func (req *KuaicheDmpListV2Request) SetSystem(system *dsp.JdDspPlatformGatewayApiVoParamExt) {
	req.Request.Params["system"] = system
}

func (req *KuaicheDmpListV2Request) GetSystem() *dsp.JdDspPlatformGatewayApiVoParamExt {
	system, found := req.Request.Params["system"]
	if found {
		return system.(*dsp.JdDspPlatformGatewayApiVoParamExt)
	}
	return nil
}
