package acty

import (
	"context"

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

func (r EnqueryRegistrationDataCountResponse) IsError() bool {
	return r.Responce == nil || r.Responce.Result == nil
}

func (r EnqueryRegistrationDataCountResponse) Error() string {
	if r.Responce == nil || r.Responce.Result == nil {
		return "no result"
	}
	return ""
}

type EnqueryRegistrationDataCountResult struct {
	Message           string             `json:"message,omitempty" codec:"message,omitempty"`
	ResultCode        uint               `json:"resultCode" codec:"resultCode,omitempty"`
	Count             uint               `json:"count" codec:"count"`
	RegistrationItems []RegistrationItem `json:"registrationItems,omitempty" codec:"registrationItems,omitempty"`
}

func EnqueryRegistrationDataCount(ctx context.Context, req *EnqueryRegistrationDataCountRequest) (*EnqueryRegistrationDataCountResult, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := acty.NewEnqueryRegistrationDataCount()
	r.SetSkuId(req.SkuId)
	r.SetOrderId(req.OrderId)
	r.SetBeginDate(req.BeginDate)
	r.SetEndDate(req.EndDate)

	var response EnqueryRegistrationDataCountResponse
	if err := client.Execute(ctx, r.Request, req.Session, &response); err != nil {
		return nil, err
	}
	return response.Responce.Result, nil
}
