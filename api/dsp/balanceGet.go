package dsp

import (
	"encoding/json"
	"errors"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/api/util"
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/dsp"
)

type BalanceGetRequest struct {
	api.BaseRequest
}

type BalanceGetResponse struct {
	ErrorResp *api.ErrorResponnse `json:"error_response,omitempty" codec:"error_response,omitempty"`
	Data      *BalanceGetData     `json:"jingdong_dsp_balance_get_responce,omitempty" codec:"jingdong_dsp_balance_get_responce,omitempty"`
}

type BalanceGetData struct {
	Code      string            `json:"code,omitempty" codec:"code,omitempty"`
	ErrorDesc string            `json:"error_description,omitempty" codec:"error_description,omitempty"`
	Result    *BalanceGetResult `json:"getaccountbalance_result,omitempty" codec:"getaccountbalance_result,omitempty"`
}

type BalanceGetResult struct {
	Data       *BalanceGetResultData `json:"value,omitempty" codec:"value,omitempty"`
	ResultCode string                `json:"resultCode,omitempty" codec:"resultCode,omitempty"`
	ErrorMsg   string                `json:"errorMsg,omitempty" codec:"errorMsg,omitempty"`
	Success    bool                  `json:"success" codec:"success"`
}

type BalanceGetResultData struct {
	TotalAmount     float64 `json:"totalAmount" codec:"totalAmount"`
	AvailableAmount float64 `json:"availableAmount" codec:"availableAmount"`
	FreezeAmount    float64 `json:"freezeAmount" codec:"freezeAmount"`
}

func BalanceGet(req *BalanceGetRequest) (*BalanceGetResultData, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := dsp.NewBalanceGetRequest()

	result, err := client.Execute(r.Request, req.Session)
	if err != nil {
		return nil, err
	}
	result = util.RemoveJsonSpace(result)

	var response BalanceGetResponse
	err = json.Unmarshal(result, &response)
	if err != nil {
		return nil, err
	}
	if response.ErrorResp != nil {
		return nil, response.ErrorResp
	}
	if response.Data.Code != "0" && response.Data.ErrorDesc != "" {
		return nil, errors.New(response.Data.ErrorDesc)
	}
	if response.Data.Result == nil {
		return nil, errors.New("No dsp balance result.")
	} else if !response.Data.Result.Success {
		if response.Data.Result.ErrorMsg != "" {
			return nil, errors.New(response.Data.Result.ErrorMsg)
		} else {
			return nil, errors.New("未知错误")
		}
	}

	return response.Data.Result.Data, nil
}
