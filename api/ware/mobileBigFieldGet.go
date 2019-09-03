package ware

import (
	"errors"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/api/util"
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/ware"
	"github.com/daviddengcn/ljson"
)

type MobileBigFieldRequest struct {
	api.BaseRequest
	SkuId uint64 `json:"sku_id" codec:"sku_id"`
}

type MobileBigFieldResponse struct {
	ErrorResp *api.ErrorResponnse `json:"error_response,omitempty" codec:"error_response,omitempty"`
	Data      *MobileBigFieldData `json:"jingdong_new_ware_mobilebigfield_get_responce,omitempty" codec:"jingdong_new_ware_mobilebigfield_get_responce,omitempty"`
}

type MobileBigFieldData struct {
	Code      string `json:"code,omitempty" codec:"code,omitempty"`
	ErrorDesc string `json:"error_description,omitempty" codec:"error_description,omitempty"`
	Result    string `json:"result,omitempty" codec:"result,omitempty"`
}

func MobileBigField(req *MobileBigFieldRequest) (string, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := ware.NewMobileBigFieldRequest()
	r.SetSkuId(req.SkuId)

	result, err := client.Execute(r.Request, req.Session)
	if err != nil {
		return "", err
	}
	result = util.RemoveJsonSpace(result)

	var response MobileBigFieldResponse
	err = ljson.Unmarshal(result, &response)
	if err != nil {
		return "", err
	}
	if response.ErrorResp != nil {
		return "", response.ErrorResp
	}
	if response.Data.Code != "0" {
		return "", errors.New(response.Data.ErrorDesc)
	}
	if response.Data.Result == "" {
		return "", errors.New("No mobile desc info.")
	}

	return response.Data.Result, nil
}
