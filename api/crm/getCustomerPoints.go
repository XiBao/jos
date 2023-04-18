package crm

import (
	"encoding/json"
	"errors"

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
	Data      *GetCustomerPointsData `json:"jingdong_pop_crm_getCustomerPoints_responce",omitempty" codec:"jingdong_pop_crm_getCustomerPoints_responce",omitempty"`
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

	result, err := client.Execute(r.Request, req.Session)
	if err != nil {
		return 0, err
	}
	if len(result) == 0 {
		return 0, errors.New("No result info.")
	}
	var response GetCustomerPointsResponse
	err = json.Unmarshal(result, &response)
	if err != nil {
		return 0, err
	}

	if response.ErrorResp != nil {
		return 0, response.ErrorResp
	}

	return response.Data.Result, nil
}
