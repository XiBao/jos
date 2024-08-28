package crm

import (
	"context"

	. "github.com/XiBao/jos/api"
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/crm"
)

type GetCustomerRequest struct {
	BaseRequest
	CustomerPin string `json:"customerPin,omitempty" codec:"customerPin,omitempty"`
}

type GetCustomerResponse struct {
	ErrorResp *ErrorResponnse      `json:"error_response,omitempty" codec:"error_response,omitempty"`
	Response  GetCustomerResponse1 `json:"jingdong_pop_crm_customer_getCustomer_responce,omitempty" codec:"jingdong_pop_crm_customer_getCustomer_responce,omitempty"`
}

func (r GetCustomerResponse) IsError() bool {
	return r.ErrorResp != nil || r.Response.IsError()
}

func (r GetCustomerResponse) Error() string {
	if r.ErrorResp != nil {
		return r.ErrorResp.Error()
	}
	return r.Response.Error()
}

type GetCustomerResponse1 struct {
	Result GetCustomerResult `json:"returnResult,omitempty" codec:"returnResult,omitempty"`
}

func (r GetCustomerResponse1) IsError() bool {
	return r.Result.IsError()
}

func (r GetCustomerResponse1) Error() string {
	return r.Result.Error()
}

type GetCustomerResult struct {
	Desc string     `json:"desc,omitempty" codec:"desc,omitempty"`
	Code string     `json:"code,omitempty" codec:"code,omitempty"`
	Data CardMember `json:"data,omitempty" codec:"data,omitempty"`
}

func (r GetCustomerResult) IsError() bool {
	return r.Code != "200"
}

func (r GetCustomerResult) Error() string {
	return sdk.ErrorString(r.Code, r.Desc)
}

func GetCustomer(ctx context.Context, req GetCustomerRequest) (CardMember, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := crm.GetCustomerRequest()
	r.SetCustomerPin(req.CustomerPin)

	var response GetCustomerResponse
	if err := client.Execute(ctx, r.Request, req.Session, &response); err != nil {
		return CardMember{}, err
	}
	return response.Response.Result.Data, nil
}
