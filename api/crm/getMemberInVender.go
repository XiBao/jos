package crm

import (
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
	Data      *GetMemberInVenderData `json:"jingdong_pop_crm_getMemberInVender_responce,omitempty" codec:"jingdong_pop_crm_getMemberInVender_responce,omitempty"`
}

func (r GetMemberInVenderResponse) IsError() bool {
	return r.ErrorResp != nil || r.Data == nil
}

func (r GetMemberInVenderResponse) Error() string {
	if r.ErrorResp != nil {
		return r.ErrorResp.Error()
	}
	return "no result data"
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

	var response GetMemberInVenderResponse
	if err := client.Execute(r.Request, req.Session, &response); err != nil {
		return nil, err
	}
	return response.Data.Result.CustomerInfoEs, nil
}
