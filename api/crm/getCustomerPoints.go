package crm

import (
	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/crm"
)

type GetCustomerPointsRequest struct {
	api.BaseRequest

	CustomerPin string `json:"customer_pin,omitempty" codec:"customer_pin,omitempty"` //
}

type GetCustomerPointsResponse struct {
	ErrorResp *api.ErrorResponnse    `json:"error_response,omitempty" codec:"error_response,omitempty"`
	Data      *GetCustomerPointsData `json:"jingdong_pop_crm_getCustomerPoints_responce,omitempty" codec:"jingdong_pop_crm_getCustomerPoints_responce,omitempty"`
}

func (r GetCustomerPointsResponse) IsError() bool {
	return r.ErrorResp != nil || r.Data == nil
}

func (r GetCustomerPointsResponse) Error() string {
	if r.ErrorResp != nil {
		return r.ErrorResp.Error()
	}
	return "no result data"
}

type GetCustomerPointsData struct {
	Code   string `json:"code,omitempty" codec:"code,omitempty"`
	Result int64  `json:"getcustomerpoints_result,omitempty" codec:"getcustomerpoints_result,omitempty"`
}

func GetCustomerPoints(req *GetCustomerPointsRequest) (int64, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := crm.NewGetCustomerPointsRequest()
	r.SetCustomerPin(req.CustomerPin)

	var response GetCustomerPointsResponse
	if err := client.Execute(r.Request, req.Session, &response); err != nil {
		return 0, err
	}
	return response.Data.Result, nil
}
