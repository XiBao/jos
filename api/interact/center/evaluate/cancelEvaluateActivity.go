package center

import (
	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/sdk"
	center "github.com/XiBao/jos/sdk/request/interact/center/evaluate"
)

type CancelEvaluateActivityRequest struct {
	api.BaseRequest
	AppName string `json:"appName" codec:"appName"`
	Channel uint8  `json:"channel" codec:"channel"`
	Id      uint64 `json:"id" codec:"id"`
}

type CancelEvaluateActivityResponse struct {
	ErrorResp *api.ErrorResponnse         `json:"error_response,omitempty" codec:"error_response,omitempty"`
	Data      *CancelEvaluateActivityData `json:"jingdong_com_jd_interact_center_api_service_write_EvaluateActivityWriteService_cancelActivity_responce,omitempty" codec:"jingdong_com_jd_interact_center_api_service_write_EvaluateActivityWriteService_cancelActivity_responce,omitempty"`
}

func (r CancelEvaluateActivityResponse) IsError() bool {
	return r.ErrorResp != nil || r.Data == nil || r.Data.IsError()
}

func (r CancelEvaluateActivityResponse) Error() string {
	if r.ErrorResp != nil {
		return r.ErrorResp.Error()
	}
	if r.Data != nil {
		return r.Data.Error()
	}
	return "no result data"
}

type CancelEvaluateActivityData struct {
	Code      string `json:"code,omitempty" codec:"code,omitempty"`
	ErrorDesc string `json:"error_description,omitempty" codec:"error_description,omitempty"`
	Result    bool   `json:"result,omitempty" codec:"result,omitempty"`
}

func (r CancelEvaluateActivityData) IsError() bool {
	return r.Code != "0"
}

func (r CancelEvaluateActivityData) Error() string {
	return sdk.ErrorString(r.Code, r.ErrorDesc)
}

func CancelEvaluateActivity(req *CancelEvaluateActivityRequest) (bool, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := center.NewCancelEvaluateActivityRequest()
	r.SetAppName(req.AppName)
	r.SetChannel(req.Channel)
	r.SetId(req.Id)

	var response CancelEvaluateActivityResponse
	if err := client.Execute(r.Request, req.Session, &response); err != nil {
		return false, err
	}

	return response.Data.Result, nil
}
