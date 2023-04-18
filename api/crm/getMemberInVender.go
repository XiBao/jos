package crm

import (
	"encoding/json"
	"errors"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/crm"
)

type GetMemberInVenderRequest struct {
	api.BaseRequest
	CustomerPin string `json:"customer_pin,omitempty" codec:"customer_pin,omitempty"` //
}

type GetMemberInVenderResponse struct {
	ErrorResp *api.ErrorResponnse    `json:"error_response,omitempty" codec:"error_response,omitempty"`
	Data      *GetMemberInVenderData `json:"jingdong_pop_crm_getMemberInVender_responce",omitempty" codec:"jingdong_pop_crm_getMemberInVender_responce",omitempty"`
}

type GetMemberInVenderData struct {
	Code   string                    `json:"code,omitempty" codec:"code,omitempty"`
	Result *GetMemberInVenderSubData `json:"getmemberinvender_result,omitempty" codec:"getmemberinvender_result,omitempty"`
}

type GetMemberInVenderSubData struct {
	CustomerInfoEs *CustomerInfoEs `json:"customerInfoEs,omitempty" codec:"customerInfoEs,omitempty"`
}

func GetMemberInVender(req *GetMemberInVenderRequest) (*CustomerInfoEs, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := crm.NewGetMemberInVenderRequest()
	r.SetCustomerPin(req.CustomerPin)

	result, err := client.Execute(r.Request, req.Session)
	if err != nil {
		return nil, err
	}
	if len(result) == 0 {
		return nil, errors.New("No result info.")
	}
	var response GetMemberInVenderResponse
	err = json.Unmarshal(result, &response)
	if err != nil {
		return nil, err
	}

	if response.ErrorResp != nil {
		return nil, response.ErrorResp
	}

	if response.Data.Result == nil {
		return nil, nil
	}

	return response.Data.Result.CustomerInfoEs, nil
}
