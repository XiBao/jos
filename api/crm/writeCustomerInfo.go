package crm

import (
	"encoding/json"
	"errors"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/crm"
)

type WriteCustomerInfoRequest struct {
	api.BaseRequest

	LinkUrl string `json:"linkUrl,omitempty" codec:"linkUrl,omitempty"`
}

type WriteCustomerInfoResponse struct {
	ErrorResp *api.ErrorResponnse    `json:"error_response,omitempty" codec:"error_response,omitempty"`
	Data      *WriteCustomerInfoData `json:"jingdong_crm_writeCustomerInfo_responce,omitempty" codec:"jingdong_crm_writeCustomerInfo_responce,omitempty"`
}

type WriteCustomerInfoData struct {
	Code   string                   `json:"code,omitempty" codec:"code,omitempty"`
	Result *WriteCustomerInfoResult `json:"writecustomerinfo_result,omitempty" codec:"writecustomerinfo_result,omitempty"`
}

type WriteCustomerInfoResult struct {
	Code string `json:"code,omitempty" codec:"code,omitempty"`
	Desc string `json:"desc,omitempty" codec:"desc,omitempty"`
	Data bool   `json:"data,omitempty" codec:"data,omitempty"`
}

// 获取单个SKU
func WriteCustomerInfo(req *WriteCustomerInfoRequest) (bool, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := crm.NewWriteCustomerInfoRequest()
	r.SetLinkUrl(req.LinkUrl)

	result, err := client.Execute(r.Request, req.Session)
	if err != nil {
		return false, err
	}
	if len(result) == 0 {
		return false, errors.New("no result info")
	}
	var response WriteCustomerInfoResponse
	err = json.Unmarshal(result, &response)
	if err != nil {
		return false, err
	}

	if response.ErrorResp != nil {
		return false, response.ErrorResp
	}

	return response.Data.Result.Data, nil
}
