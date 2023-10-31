package center

import (
	"fmt"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/sdk"
	center "github.com/XiBao/jos/sdk/request/interact/center/evaluate"
)

type GetEvaluateActivityByIdRequest struct {
	api.BaseRequest
	AppName    string `json:"appName" codec:"appName"`
	Channel    uint8  `json:"channel" codec:"channel"`
	ActivityId uint64 `json:"activityId" codec:"activityId"`
}

type GetEvaluateActivityByIdResponse struct {
	ErrorResp *api.ErrorResponnse          `json:"error_response,omitempty" codec:"error_response,omitempty"`
	Data      *GetEvaluateActivityByIdData `json:"jingdong_interact_center_vender_read_evaluate_getActivityById_responce,omitempty" codec:"jingdong_interact_center_vender_read_evaluate_getActivityById_responce,omitempty"`
}

func (r GetEvaluateActivityByIdResponse) IsError() bool {
	return r.ErrorResp != nil || r.Data == nil || r.Data.IsError()
}

func (r GetEvaluateActivityByIdResponse) Error() string {
	if r.ErrorResp != nil {
		return r.ErrorResp.Error()
	}
	if r.Data != nil {
		return r.Data.Error()
	}
	return "no result data"
}

type GetEvaluateActivityByIdData struct {
	Code      string            `json:"code,omitempty" codec:"code,omitempty"`
	ErrorDesc string            `json:"error_description,omitempty" codec:"error_description,omitempty"`
	Result    *EvaluateActivity `json:"EvaluateActivity,omitempty" codec:"GiftActivityResults,omitempty"`
}

func (r GetEvaluateActivityByIdData) IsError() bool {
	return r.Code != "0" || r.Result == nil
}

func (r GetEvaluateActivityByIdData) Error() string {
	if r.Code != "0" {
		return fmt.Sprintf("code:%s, msg:%s", r.Code, r.ErrorDesc)
	}
	return "no result data"
}

type EvaluateActivity struct {
	VenderId           uint64 `json:"vender_id" codec:"vender_id"`
	EndTime            uint64 `json:"end_time" codec:"end_time"`
	Type               string `json:"type" codec:"type"`
	Creator            string `json:"creator" codec:"creator"`
	StartTime          uint64 `json:"start_time" codec:"start_time"`
	Created            uint64 `json:"created" codec:"created"`
	RewardTime         uint64 `json:"reward_time" codec:"reward_time"`
	Name               string `json:"name" codec:"name"`
	RfId               uint64 `json:"rfId" codec:"rfId"`
	Id                 uint64 `json:"id" codec:"id"`
	Platform           uint   `json:"platform" codec:"platform"`
	Status             uint8  `json:"status" codec:"status"`
	Modified           uint64 `json:"modified" codec:"modified"`
	WordRequirement    uint   `json:"word_requirement" codec:"word_requirement"`
	PictureRequirement uint   `json:"picture_requirement" codec:"picture_requirement"`
	VedioRequirement   uint   `json:"vedio_requirement" codec:"vedio_requirement"`
}

func GetEvaluateActivityById(req *GetEvaluateActivityByIdRequest) (*EvaluateActivity, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := center.NewGetEvaluateActivityByIdRequest()
	r.SetAppName(req.AppName)
	r.SetChannel(req.Channel)
	r.SetActivityId(req.ActivityId)

	var response GetEvaluateActivityByIdResponse
	if err := client.Execute(r.Request, req.Session, &response); err != nil {
		return nil, err
	}
	return response.Data.Result, nil
}
