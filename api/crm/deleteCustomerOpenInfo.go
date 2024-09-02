package crm

import (
	"context"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/crm"
)

type DeleteCustomerOpenInfoRequest struct {
	api.BaseRequest
}

type DeleteCustomerOpenInfoResponse struct {
	ErrorResp *api.ErrorResponnse         `json:"error_response,omitempty" codec:"error_response,omitempty"`
	Data      *DeleteCustomerOpenInfoData `json:"jingdong_crm_deleteCustomerOpenInfo_responce,omitempty" codec:"jingdong_crm_deleteCustomerOpenInfo_responce,omitempty"`
}

func (r DeleteCustomerOpenInfoResponse) IsError() bool {
	return r.ErrorResp != nil || r.Data == nil || r.Data.IsError()
}

func (r DeleteCustomerOpenInfoResponse) Error() string {
	if r.ErrorResp != nil {
		return r.ErrorResp.Error()
	}
	if r.Data != nil {
		return r.Data.Error()
	}
	return "no result data"
}

type DeleteCustomerOpenInfoData struct {
	Code   string                        `json:"code,omitempty" codec:"code,omitempty"`
	Result *DeleteCustomerOpenInfoResult `json:"deletecustomeropeninfo_result,omitempty" codec:"deletecustomeropeninfo_result,omitempty"`
}

func (r DeleteCustomerOpenInfoData) IsError() bool {
	return r.Result == nil || r.Result.IsError()
}

func (r DeleteCustomerOpenInfoData) Error() string {
	if r.Result != nil && r.Result.IsError() {
		return r.Result.Error()
	}
	return "no result data"
}

type DeleteCustomerOpenInfoResult struct {
	Code string `json:"code,omitempty" codec:"code,omitempty"`
	Desc string `json:"desc,omitempty" codec:"desc,omitempty"`
	Data bool   `json:"data,omitempty" codec:"data,omitempty"`
}

func (r DeleteCustomerOpenInfoResult) IsError() bool {
	return r.Code != "200"
}

func (r DeleteCustomerOpenInfoResult) Error() string {
	return sdk.ErrorString(r.Code, r.Desc)
}

// 获取单个SKU
func DeleteCustomerOpenInfo(ctx context.Context, req *DeleteCustomerOpenInfoRequest) (bool, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := crm.NewDeleteCustomerOpenInfoRequest()

	var response DeleteCustomerOpenInfoResponse
	if err := client.Execute(ctx, r.Request, req.Session, &response); err != nil {
		return false, err
	}
	return response.Data.Result.Data, nil
}
