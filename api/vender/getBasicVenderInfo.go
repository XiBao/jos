package vender

import (
	"encoding/json"
	"errors"

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

type GetBasicVenderInfoResult struct {
	Code      string                 `json:"code,omitempty" codec:"code,omitempty"`
	ErrorDesc string                 `json:"error_description,omitempty" codec:"error_description,omitempty"`
	Result    *BasicVenderInfoResult `json:"getbasicvenderinfobyvenderid_result,omitempty" codec:"getbasicvenderinfobyvenderid_result,omitempty"`
}

type BasicVenderInfoResult struct {
	Success       bool             `json:"success,omitempty" codec:"success,omitempty"`
	ErrorCode     string           `json:"errorCode,omitempty" codec:"errorCode,omitempty"`
	ErrorMsg      string           `json:"errorMsg,omitempty" codec:"errorMsg,omitempty"`
	TotalNum      int              `json:"totalNum,omitempty" codec:"totalNum,omitempty"`
	VenderBasicVO *BasicVenderInfo `json:"venderBasicVO,omitempty" codec:"venderBasicVO,omitempty"`
}

// 店铺信息查询
func GetBasicVenderInfo(req *GetBasicVenderInfoRequest) (*BasicVenderInfo, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := vender.NewGetBasicVenderInfoRequest()
	r.SetSource(req.Source)
	if req.ColNames != nil {
		r.SetColNames(req.ColNames)
	}
	result, err := client.Execute(r.Request, req.Session)
	if err != nil {
		return nil, err
	}
	if len(result) == 0 {
		return nil, errors.New("no result.")
	}
	var response GetBasicVenderInfoResponse
	err = json.Unmarshal(result, &response)
	if err != nil {
		return nil, err
	}
	if response.ErrorResp != nil {
		return nil, response.ErrorResp
	}
	if response.Data.Code != "0" {
		return nil, errors.New(response.Data.ErrorDesc)
	}

	if response.Data.Result.ErrorMsg != "" {
		return nil, errors.New(response.Data.Result.ErrorMsg)
	}

	return response.Data.Result.VenderBasicVO, nil
}
