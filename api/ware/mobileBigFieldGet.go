package ware

import (
	"context"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/ware"
)

type MobileBigFieldRequest struct {
	api.BaseRequest
	SkuId uint64 `json:"sku_id" codec:"sku_id"`
}

type MobileBigFieldResponse struct {
	ErrorResp *api.ErrorResponnse `json:"error_response,omitempty" codec:"error_response,omitempty"`
	Data      *MobileBigFieldData `json:"jingdong_new_ware_mobilebigfield_get_responce,omitempty" codec:"jingdong_new_ware_mobilebigfield_get_responce,omitempty"`
}

func (r MobileBigFieldResponse) IsError() bool {
	return r.ErrorResp != nil || r.Data == nil || r.Data.IsError()
}

func (r MobileBigFieldResponse) Error() string {
	if r.ErrorResp != nil {
		return r.ErrorResp.Error()
	}
	if r.Data.IsError() {
		return r.Data.Error()
	}
	return "no result data"
}

type MobileBigFieldData struct {
	Code      string `json:"code,omitempty" codec:"code,omitempty"`
	ErrorDesc string `json:"error_description,omitempty" codec:"error_description,omitempty"`
	Result    string `json:"result,omitempty" codec:"result,omitempty"`
}

func (r MobileBigFieldData) IsError() bool {
	return r.Code != "0" || r.Result == ""
}

func (r MobileBigFieldData) Error() string {
	return sdk.ErrorString(r.Code, r.ErrorDesc)
}

func MobileBigField(ctx context.Context, req *MobileBigFieldRequest) (string, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := ware.NewMobileBigFieldRequest()
	r.SetSkuId(req.SkuId)

	var response MobileBigFieldResponse
	if err := client.Execute(ctx, r.Request, req.Session, &response); err != nil {
		return "", err
	}
	return response.Data.Result, nil
}
