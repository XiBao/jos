package dsp

import (
	"github.com/XiBao/jos/api"
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

func (r BalanceGetResponse) IsError() bool {
	return r.ErrorResp != nil || r.Data == nil || r.Data.IsError()
}

func (r BalanceGetResponse) Error() string {
	if r.ErrorResp != nil {
		return r.ErrorResp.Error()
	}
	if r.Data != nil {
		return r.Data.Error()
	}
	return "no result data"
}

type BalanceGetData struct {
	Code      string            `json:"code,omitempty" codec:"code,omitempty"`
	ErrorDesc string            `json:"error_description,omitempty" codec:"error_description,omitempty"`
	Result    *BalanceGetResult `json:"getaccountbalance_result,omitempty" codec:"getaccountbalance_result,omitempty"`
}

func (r BalanceGetData) IsError() bool {
	return r.Code != "0" || r.Result == nil || r.Result.IsError()
}

func (r BalanceGetData) Error() string {
	if r.Result != nil && r.Result.IsError() {
		return r.Result.Error()
	}
	if r.Code != "0" {
		return sdk.ErrorString(r.Code, r.ErrorDesc)
	}
	return "no result data"
}

type BalanceGetResult struct {
	Data       *BalanceGetResultData `json:"value,omitempty" codec:"value,omitempty"`
	ResultCode string                `json:"resultCode,omitempty" codec:"resultCode,omitempty"`
	ErrorMsg   string                `json:"errorMsg,omitempty" codec:"errorMsg,omitempty"`
	Success    bool                  `json:"success" codec:"success"`
}

func (r BalanceGetResult) IsError() bool {
	return !r.Success || r.Data == nil
}

func (r BalanceGetResult) Error() string {
	if !r.Success {
		return sdk.ErrorString(r.ResultCode, r.ErrorMsg)
	}
	return "no result data"
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

	var response BalanceGetResponse
	if err := client.Execute(r.Request, req.Session, &response); err != nil {
		return nil, err
	}
	return response.Data.Result.Data, nil
}
