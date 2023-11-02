package crm

import (
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

func (r WriteCustomerInfoResponse) IsError() bool {
	return r.ErrorResp != nil || r.Data == nil || r.Data.IsError()
}

func (r WriteCustomerInfoResponse) Error() string {
	if r.ErrorResp != nil {
		return r.ErrorResp.Error()
	}
	if r.Data != nil {
		return r.Data.Error()
	}
	return "no result data"
}

type WriteCustomerInfoData struct {
	Code   string                   `json:"code,omitempty" codec:"code,omitempty"`
	Result *WriteCustomerInfoResult `json:"writecustomerinfo_result,omitempty" codec:"writecustomerinfo_result,omitempty"`
}

func (r WriteCustomerInfoData) IsError() bool {
	return r.Result == nil || r.Result.IsError()
}

func (r WriteCustomerInfoData) Error() string {
	if r.Result != nil {
		return r.Result.Error()
	}
	return "no result data"
}

type WriteCustomerInfoResult struct {
	Code string `json:"code,omitempty" codec:"code,omitempty"`
	Desc string `json:"desc,omitempty" codec:"desc,omitempty"`
	Data bool   `json:"data,omitempty" codec:"data,omitempty"`
}

func (r WriteCustomerInfoResult) IsError() bool {
	return r.Code != "200"
}

func (r WriteCustomerInfoResult) Error() string {
	return sdk.ErrorString(r.Code, r.Desc)
}

// 获取单个SKU
func WriteCustomerInfo(req *WriteCustomerInfoRequest) (bool, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := crm.NewWriteCustomerInfoRequest()
	r.SetLinkUrl(req.LinkUrl)

	var response WriteCustomerInfoResponse
	if err := client.Execute(r.Request, req.Session, &response); err != nil {
		return false, err
	}
	return response.Data.Result.Data, nil
}
