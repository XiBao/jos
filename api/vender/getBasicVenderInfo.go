package vender

import (
	"context"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/vender"
)

type GetBasicVenderInfoRequest struct {
	api.BaseRequest
	ColNames []string `json:"colNames,omitempty" codec:"colNames,omitempty"`
	Source   int      `json:"source" codec:"source"`
}

type GetBasicVenderInfoResponse struct {
	ErrorResp *api.ErrorResponnse       `json:"error_response,omitempty" codec:"error_response,omitempty"`
	Data      *GetBasicVenderInfoResult `json:"jingdong_vender_vbinfo_getBasicVenderInfoByVenderId_responce,omitempty" codec:"jingdong_vender_vbinfo_getBasicVenderInfoByVenderId_responce,omitempty"`
}

func (r GetBasicVenderInfoResponse) IsError() bool {
	return r.ErrorResp != nil || r.Data == nil || r.Data.IsError()
}

func (r GetBasicVenderInfoResponse) Error() string {
	if r.ErrorResp != nil {
		return r.ErrorResp.Error()
	}
	if r.Data != nil {
		return r.Data.Error()
	}
	return "no result data"
}

type GetBasicVenderInfoResult struct {
	Code      string                 `json:"code,omitempty" codec:"code,omitempty"`
	ErrorDesc string                 `json:"error_description,omitempty" codec:"error_description,omitempty"`
	Result    *BasicVenderInfoResult `json:"getbasicvenderinfobyvenderid_result,omitempty" codec:"getbasicvenderinfobyvenderid_result,omitempty"`
}

func (r GetBasicVenderInfoResult) IsError() bool {
	return r.Code != "0" || r.Result == nil || r.Result.IsError()
}

func (r GetBasicVenderInfoResult) Error() string {
	if r.Result != nil && r.Result.IsError() {
		return r.Result.Error()
	}
	return sdk.ErrorString(r.Code, r.ErrorDesc)
}

type BasicVenderInfoResult struct {
	Success       bool             `json:"success,omitempty" codec:"success,omitempty"`
	ErrorCode     string           `json:"errorCode,omitempty" codec:"errorCode,omitempty"`
	ErrorMsg      string           `json:"errorMsg,omitempty" codec:"errorMsg,omitempty"`
	TotalNum      int              `json:"totalNum,omitempty" codec:"totalNum,omitempty"`
	VenderBasicVO *BasicVenderInfo `json:"venderBasicVO,omitempty" codec:"venderBasicVO,omitempty"`
}

func (r BasicVenderInfoResult) IsError() bool {
	return !r.Success
}

func (r BasicVenderInfoResult) Error() string {
	return sdk.ErrorString(r.ErrorCode, r.ErrorMsg)
}

// 店铺信息查询
func GetBasicVenderInfo(ctx context.Context, req *GetBasicVenderInfoRequest) (*BasicVenderInfo, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := vender.NewGetBasicVenderInfoRequest()
	r.SetSource(req.Source)
	if req.ColNames != nil {
		r.SetColNames(req.ColNames)
	}

	var response GetBasicVenderInfoResponse
	if err := client.Execute(ctx, r.Request, req.Session, &response); err != nil {
		return nil, err
	}
	return response.Data.Result.VenderBasicVO, nil
}
