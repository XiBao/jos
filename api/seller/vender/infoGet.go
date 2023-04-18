package vender

import (
	"encoding/json"
	"errors"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/seller/vender"
)

type InfoGetRequest struct {
	api.BaseRequest
}
type InfoGetResponse struct {
	ErrorResp *api.ErrorResponnse `json:"error_response,omitempty" codec:"error_response,omitempty"`
	Data      *InfoGetSubResponse `json:"jingdong_seller_vender_info_get_responce,omitempty" codec:"jingdong_seller_vender_info_get_responce,omitempty"`
}

type InfoGetSubResponse struct {
	Code       string      `json:"code,omitempty" codec:"code,omitempty"`
	ErrorDesc  string      `json:"error_description,omitempty" codec:"error_description,omitempty"`
	VenderInfo *VenderInfo `json:"vender_info_result,omitempty" codec:"vender_info_result,omitempty"`
}

// 店铺信息查询
func InfoGet(req *InfoGetRequest) (*VenderInfo, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := vender.NewSellerVenderInfoGetRequest()
	result, err := client.Execute(r.Request, req.Session)
	if err != nil {
		return nil, err
	}
	if len(result) == 0 {
		return nil, errors.New("no result.")
	}

	var response InfoGetResponse
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

	return response.Data.VenderInfo, nil
}
