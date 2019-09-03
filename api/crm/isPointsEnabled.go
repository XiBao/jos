package crm

import (
	"errors"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/crm"
	"github.com/daviddengcn/ljson"
)

type IsPointsEnabledRequest struct {
	api.BaseRequest
}

type IsPointsEnabledResponse struct {
	ErrorResp *api.ErrorResponnse  `json:"error_response,omitempty" codec:"error_response,omitempty"`
	Data      *IsPointsEnabledData `json:"jingdong_pop_crm_isPointsEnabled_responce,omitempty" codec:"jingdong_pop_crm_isPointsEnabled_responce,omitempty"`
}

type IsPointsEnabledData struct {
	Result bool   `json:"ispointsenabled_result,omitempty" codec:"ispointsenabled_result,omitempty"`
	Code   string `json:"code,omitempty" codec:"code,omitempty"`
}

//是否开启店铺积分功能
func IsPointsEnabled(req *IsPointsEnabledRequest) (bool, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := crm.NewIsPointsEnabledRequest()

	result, err := client.Execute(r.Request, req.Session)
	if err != nil {
		return false, err
	}
	if len(result) == 0 {
		return false, errors.New("no result info")
	}
	var response IsPointsEnabledResponse
	err = ljson.Unmarshal(result, &response)
	if err != nil {
		return false, err
	}
	if response.ErrorResp != nil {
		return false, response.ErrorResp
	}

	if response.Data.Code != "0" {
		return false, errors.New(response.Data.Code)
	}

	return response.Data.Result, nil

}
