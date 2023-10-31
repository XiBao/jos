package center

import (
	"fmt"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/interact/center"
)

type WriteCollectCouponRequest struct {
	api.BaseRequest
	AppName    string `json:"appName,omitempty" codec:"appName,omitempty"` // 调用方应用名称，新接口接入必须联系产品，出现问题概不负责，且有权利追求责任及接口降级
	AppId      uint64 `json:"appId,omitempty" codec:"appId,omitempty"`
	Channel    uint8  `json:"channel,omitempty" codec:"channel,omitempty"`
	Pin        string `json:"pin,omitempty" codec:"pin,omitempty"`
	ActivityId uint64 `json:"activityId,omitempty" codec:"activityId,omitempty"`
	Ip         string `json:"ip,omitempty" codec:"ip,omitempty"`
	RuleId     uint64 `json:"ruleId,omitempty" codec:"ruleId,omitempty"`
	RfId       string `json:"rfId,omitempty" codec:"rfId,omitempty"`
	Source     uint8  `json:"source,omitempty" codec:"source,omitempty"`
	Type       uint8  `json:"type,omitempty" codec:"type,omitempty"`
}

type WriteCollectCouponResponse struct {
	ErrorResp *api.ErrorResponnse     `json:"error_response,omitempty" codec:"error_response,omitempty"`
	Data      *WriteCollectCouponData `json:"jingdong_interact_center_api_service_write_collectCoupon_responce,omitempty" codec:"jingdong_interact_center_api_service_write_collectCoupon_responce,omitempty"`
}

func (r WriteCollectCouponResponse) IsError() bool {
	return r.ErrorResp != nil || r.Data == nil || r.Data.IsError()
}

func (r WriteCollectCouponResponse) Error() string {
	if r.ErrorResp != nil {
		return r.ErrorResp.Error()
	}
	if r.Data != nil {
		return r.Data.Error()
	}
	return "no result data"
}

type WriteCollectCouponData struct {
	Code      string                    `json:"code,omitempty" codec:"code,omitempty"`
	ErrorDesc string                    `json:"error_description,omitempty" codec:"error_description,omitempty"`
	Result    *WriteCollectCouponResult `json:"GiftActivityResults,omitempty" codec:"GiftActivityResults,omitempty"`
}

func (r WriteCollectCouponData) IsError() bool {
	return r.Code != "0" || r.Result == nil
}

func (r WriteCollectCouponData) Error() string {
	if r.Code != "0" {
		return fmt.Sprintf("code:%s, msg:%s", r.Code, r.ErrorDesc)
	}
	return "no result data"
}

type WriteCollectCouponResult struct {
	Data bool   `json:"data,omitempty" codec:"data,omitempty"`
	Code uint32 `json:"code" codec:"code"`
	Msg  string `json:"msg" codec:"msg"`
}

func WriteCollectCoupon(req *WriteCollectCouponRequest) (bool, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := center.NewWriteCollectCouponRequest()
	r.SetAppName(req.AppName)
	r.SetAppId(req.AppId)
	r.SetChannel(req.Channel)
	r.SetPin(req.Pin)
	r.SetActivityId(req.ActivityId)
	r.SetIp(req.Ip)
	r.SetRuleId(req.RuleId)
	r.SetRfId(req.RfId)
	r.SetSource(req.Source)
	r.SetType(req.Type)

	var response WriteCollectCouponResponse
	if err := client.Execute(r.Request, req.Session, &response); err != nil {
		return false, err
	}
	return response.Data.Result.Data, nil
}
