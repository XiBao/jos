package center

import (
	"errors"

	"github.com/XiBao/jos/api/util"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/sdk"
	center "github.com/XiBao/jos/sdk/request/interact/center/evaluate"
	"github.com/daviddengcn/ljson"
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

type CancelEvaluateActivityData struct {
	Code      string `json:"code,omitempty" codec:"code,omitempty"`
	ErrorDesc string `json:"error_description,omitempty" codec:"error_description,omitempty"`
	Result    bool   `json:"result,omitempty" codec:"result,omitempty"`
}

func CancelEvaluateActivity(req *CancelEvaluateActivityRequest) (bool, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := center.NewCancelEvaluateActivityRequest()
	r.SetAppName(req.AppName)
	r.SetChannel(req.Channel)
	r.SetId(req.Id)

	result, err := client.Execute(r.Request, req.Session)
	if err != nil {
		return false, err
	}
	result = util.RemoveJsonSpace(result)

	var response CancelEvaluateActivityResponse
	err = ljson.Unmarshal(result, &response)
	if err != nil {
		return false, err
	}
	if response.ErrorResp != nil {
		return false, response.ErrorResp
	}
	if response.Data.Code != "0" {
		return false, errors.New(response.Data.ErrorDesc)
	}

	return response.Data.Result, nil
}
