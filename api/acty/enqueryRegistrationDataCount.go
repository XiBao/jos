package acty

import (
	"encoding/json"
	"errors"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/acty"
)

type EnqueryRegistrationDataCountRequest struct {
	api.BaseRequest
	SkuId     uint64 `json:"skuId,omitempty" codec:"skuId,omitempty"`
	OrderId   uint64 `json:"orderId,omitempty" codec:"orderId,omitempty"`
	BeginDate string `json:"beginDate,omitempty" codec:"beginDate,omitempty"`
	EndDate   string `json:"endDate,omitempty" codec:"endDate,omitempty"`
}

type EnqueryRegistrationDataCountResponse struct {
	Responce *EnqueryRegistrationDataCountResponce `json:"jingdong_acty_enqueryRegistrationDataCount_responce,omitempty" codec:"jingdong_acty_enqueryRegistrationDataCount_responce,omitempty"`
}

type EnqueryRegistrationDataCountResponce struct {
	Code   string                              `json:"code,omitempty" codec:"code,omitempty"`
	Result *EnqueryRegistrationDataCountResult `json:"queryregistrationdatacount_result,omitempty" codec:"queryregistrationdatacount_result,omitempty"`
}

type EnqueryRegistrationDataCountResult struct {
	Message           string              `json:"message,omitempty" codec:"message,omitempty"`
	ResultCode        uint                `json:"resultCode" codec:"resultCode,omitempty"`
	Count             uint                `json:"count" codec:"count"`
	RegistrationItems []*RegistrationItem `json:"registrationItems,omitempty" codec:"registrationItems,omitempty"`
}

func EnqueryRegistrationDataCount(req *EnqueryRegistrationDataCountRequest) (*EnqueryRegistrationDataCountResult, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := acty.NewEnqueryRegistrationDataCount()
	r.SetSkuId(req.SkuId)
	r.SetOrderId(req.OrderId)
	r.SetBeginDate(req.BeginDate)
	r.SetEndDate(req.EndDate)

	result, err := client.Execute(r.Request, req.Session)
	if err != nil {
		return nil, err
	}
	if len(result) == 0 {
		return nil, errors.New("No result.")
	}

	var response EnqueryRegistrationDataCountResponse
	err = json.Unmarshal(result, &response)
	if err != nil {
		return nil, err
	}
	if response.Responce == nil {
		return nil, errors.New("No result.")
	}
	if response.Responce.Result == nil {
		return nil, errors.New("No result.")
	}

	return response.Responce.Result, nil
}
