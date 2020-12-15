package center

import (
	"errors"

	"github.com/XiBao/jos/api/util"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/sdk"
	center "github.com/XiBao/jos/sdk/request/interact/center/evaluate"
	"github.com/daviddengcn/ljson"
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

type GetEvaluateActivityByIdData struct {
	Code      string            `json:"code,omitempty" codec:"code,omitempty"`
	ErrorDesc string            `json:"error_description,omitempty" codec:"error_description,omitempty"`
	Result    *EvaluateActivity `json:"EvaluateActivity,omitempty" codec:"GiftActivityResults,omitempty"`
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

	result, err := client.Execute(r.Request, req.Session)
	if err != nil {
		return nil, err
	}
	result = util.RemoveJsonSpace(result)

	var response GetEvaluateActivityByIdResponse
	err = ljson.Unmarshal(result, &response)
	if err != nil {
		return nil, err
	}
	if response.ErrorResp != nil {
		return nil, response.ErrorResp
	}
	if response.Data.Code != "0" {
		return nil, errors.New(response.Data.ErrorDesc)
	}
	if response.Data.Result == nil {
		return nil, errors.New("No result.")
	}

	return response.Data.Result, nil
}
