package vender

import (
	"fmt"

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

func (r InfoGetResponse) IsError() bool {
	return r.ErrorResp != nil || r.Data == nil || r.Data.IsError()
}

func (r InfoGetResponse) Error() string {
	if r.ErrorResp != nil {
		return r.ErrorResp.Error()
	}
	if r.Data != nil {
		return r.Data.Error()
	}
	return "no result data"
}

type InfoGetSubResponse struct {
	Code       string      `json:"code,omitempty" codec:"code,omitempty"`
	ErrorDesc  string      `json:"error_description,omitempty" codec:"error_description,omitempty"`
	VenderInfo *VenderInfo `json:"vender_info_result,omitempty" codec:"vender_info_result,omitempty"`
}

func (r InfoGetSubResponse) IsError() bool {
	return r.Code != "0"
}

func (r InfoGetSubResponse) Error() string {
	return fmt.Sprintf("code:%s, msg:%s", r.Code, r.ErrorDesc)
}

// 店铺信息查询
func InfoGet(req *InfoGetRequest) (*VenderInfo, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := vender.NewSellerVenderInfoGetRequest()
	var response InfoGetResponse
	if err := client.Execute(r.Request, req.Session, &response); err != nil {
		return nil, err
	}
	return response.Data.VenderInfo, nil
}
