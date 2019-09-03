package crm

import (
	"errors"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/crm"
	"github.com/daviddengcn/ljson"
)

type DeleteCustomerOpenInfoRequest struct {
	api.BaseRequest
}

type DeleteCustomerOpenInfoResponse struct {
	ErrorResp *api.ErrorResponnse         `json:"error_response,omitempty" codec:"error_response,omitempty"`
	Data      *DeleteCustomerOpenInfoData `json:"jingdong_crm_deleteCustomerOpenInfo_responce,omitempty" codec:"jingdong_crm_deleteCustomerOpenInfo_responce,omitempty"`
}

type DeleteCustomerOpenInfoData struct {
	Code   string                        `json:"code,omitempty" codec:"code,omitempty"`
	Result *DeleteCustomerOpenInfoResult `json:"deletecustomeropeninfo_result,omitempty" codec:"deletecustomeropeninfo_result,omitempty"`
}

type DeleteCustomerOpenInfoResult struct {
	Code string `json:"code,omitempty" codec:"code,omitempty"`
	Desc string `json:"desc,omitempty" codec:"desc,omitempty"`
	Data bool   `json:"data,omitempty" codec:"data,omitempty"`
}

// 获取单个SKU
func DeleteCustomerOpenInfo(req *DeleteCustomerOpenInfoRequest) (bool, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := crm.NewDeleteCustomerOpenInfoRequest()

	result, err := client.Execute(r.Request, req.Session)
	if err != nil {
		return false, err
	}
	if len(result) == 0 {
		return false, errors.New("no result info")
	}
	var response DeleteCustomerOpenInfoResponse
	err = ljson.Unmarshal(result, &response)
	if err != nil {
		return false, err
	}

	if response.ErrorResp != nil {
		return false, response.ErrorResp
	}

	return response.Data.Result.Data, nil
}
