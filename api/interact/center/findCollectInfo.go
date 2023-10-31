package center

import (
	"fmt"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/interact/center"
)

type FindCollectInfoRequest struct {
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

type FindCollectInfoResponse struct {
	ErrorResp *api.ErrorResponnse  `json:"error_response,omitempty" codec:"error_response,omitempty"`
	Data      *FindCollectInfoData `json:"jingdong_interact_center_api_service_read_findCollectInfo_responce,omitempty" codec:"jingdong_interact_center_api_service_read_findCollectInfo_responce,omitempty"`
}

func (r FindCollectInfoResponse) IsError() bool {
	return r.ErrorResp != nil || r.Data == nil || r.Data.IsError()
}

func (r FindCollectInfoResponse) Error() string {
	if r.ErrorResp != nil {
		return r.ErrorResp.Error()
	}
	if r.Data != nil {
		return r.Data.Error()
	}
	return "no result data"
}

type FindCollectInfoData struct {
	Code      string                 `json:"code,omitempty" codec:"code,omitempty"`
	ErrorDesc string                 `json:"error_description,omitempty" codec:"error_description,omitempty"`
	Result    *FindCollectInfoResult `json:"giftActivityResults,omitempty" codec:"giftActivityResults,omitempty"`
}

func (r FindCollectInfoData) IsError() bool {
	return r.Code != "0" || r.Result == nil || r.Result.IsError()
}

func (r FindCollectInfoData) Error() string {
	if r.Result != nil && r.Result.IsError() {
		return r.Result.Error()
	}
	if r.Code != "0" {
		return fmt.Sprintf("code:%s, msg:%s", r.Code, r.ErrorDesc)
	}
	return "no result data"
}

type FindCollectInfoDetails struct {
	ValidateDay    uint16 `json:"validateDay,omitempty" codec:"validateDay,omitempty"`
	PrizeStartTime uint64 `json:"prizeStartTime,omitempty" codec:"prizeStartTime,omitempty"`
	VenderId       uint64 `json:"venderId,omitempty" codec:"venderId,omitempty"`
	PrizeType      uint8  `json:"prizeType,omitempty" codec:"prizeType,omitempty"`
	CouponId       uint64 `json:"couponId,omitempty" codec:"couponId,omitempty"`
	Discount       uint16 `json:"discount,omitempty" codec:"discount,omitempty"`
	Id             uint64 `json:"id,omitempty" codec:"id,omitempty"`
	ActivityId     uint64 `json:"activityId,omitempty" codec:"activityId,omitempty"`
	Quota          uint32 `json:"quota,omitempty" codec:"quota,omitempty"`
	SendCount      uint32 `json:"sendCount,omitempty" codec:"sendCount,omitempty"`
	DrawCount      uint32 `json:"drawCount,omitempty" codec:"drawCount,omitempty"`
	BatchKey       string `json:"batchKey,omitempty" codec:"batchKey,omitempty"`
	CollectTimes   uint16 `json:"collectTimes,omitempty" codec:"collectTimes,omitempty"`
	PrizeEndTime   uint64 `json:"prizeEndTime,omitempty" codec:"prizeEndTime,omitempty"`
	PrizeId        uint64 `json:"prizeId,omitempty" codec:"prizeId,omitempty"`
}

type FindCollectInfoResult struct {
	Data []*FindCollectInfoDetails `json:"data,omitempty" codec:"data,omitempty"`
	Code uint32                    `json:"code" codec:"code"`
	Msg  string                    `json:"msg" codec:"msg"`
}

func (r FindCollectInfoResult) IsError() bool {
	return r.Code != 200
}

func (r FindCollectInfoResult) Error() string {
	return fmt.Sprintf("code:%d, msg:%s", r.Code, r.Msg)
}

func FindCollectInfo(req *FindCollectInfoRequest) ([]*FindCollectInfoDetails, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := center.NewFindCollectInfoRequest()
	r.SetAppName(req.AppName)
	r.SetAppId(req.AppId)
	r.SetChannel(req.Channel)
	r.SetPin(req.Pin)
	r.SetActivityId(req.ActivityId)
	r.SetType(req.Type)

	var response FindCollectInfoResponse
	if err := client.Execute(r.Request, req.Session, &response); err != nil {
		return nil, err
	}

	return response.Data.Result.Data, nil
}
