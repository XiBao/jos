package crm

import (
	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/crm"
)

type IsPointsEnabledRequest struct {
	api.BaseRequest
}

type IsPointsEnabledResponse struct {
	ErrorResp *api.ErrorResponnse  `json:"error_response,omitempty" codec:"error_response,omitempty"`
	Data      *IsPointsEnabledData `json:"jingdong_pop_crm_isPointsEnabled_responce,omitempty" codec:"jingdong_pop_crm_isPointsEnabled_responce,omitempty"`
}

func (r IsPointsEnabledResponse) IsError() bool {
	return r.ErrorResp != nil || r.Data == nil || r.Data.IsError()
}

func (r IsPointsEnabledResponse) Error() string {
	if r.ErrorResp != nil {
		return r.ErrorResp.Error()
	}
	if r.Data != nil {
		return r.Data.Error()
	}
	return "no result data"
}

type IsPointsEnabledData struct {
	Result bool   `json:"ispointsenabled_result,omitempty" codec:"ispointsenabled_result,omitempty"`
	Code   string `json:"code,omitempty" codec:"code,omitempty"`
}

func (r IsPointsEnabledData) IsError() bool {
	return r.Code != "0"
}

func (r IsPointsEnabledData) Error() string {
	return r.Code
}

// 是否开启店铺积分功能
func IsPointsEnabled(req *IsPointsEnabledRequest) (bool, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := crm.NewIsPointsEnabledRequest()

	var response IsPointsEnabledResponse
	if err := client.Execute(r.Request, req.Session, &response); err != nil {
		return false, err
	}
	return response.Data.Result, nil

}
