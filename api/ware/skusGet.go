package ware

import (
	"encoding/json"
	"errors"
	"strings"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/ware"
)

type SkusGetRequest struct {
	api.BaseRequest
	WareIds []string `json:"ware_ids,omitempty" codec:"ware_ids,omitempty"`
	Fields  []string `json:"fields,omitempty" codec:"fields,omitempty"`
}

type SkusGetResponse struct {
	ErrorResp *api.ErrorResponnse `json:"error_response,omitempty" codec:"error_response,omitempty"`
	Data      *SkusGetSubResponse `json:"ware_skus_get_response,omitempty" codec:"ware_skus_get_response,omitempty"`
}

type SkusGetSubResponse struct {
	Code      string  `json:"code,omitempty" codec:"code,omitempty"`
	ErrorDesc string  `json:"error_description,omitempty" codec:"error_description,omitempty"`
	Skus      []*Sku2 `json:"skus,omitempty" codec:"skus,omitempty"`
}

// 根据条件检索订单信息 （仅适用于SOP、LBP，SOPL类型，FBP类型请调取FBP订单检索 ）
func SkusGet(req *SkusGetRequest) ([]*Sku2, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := ware.NewWareSkusGetRequest()
	r.SetWareIds(strings.Join(req.WareIds, ","))
	r.SetFields(strings.Join(req.Fields, ","))
	result, err := client.Execute(r.Request, req.Session)
	if err != nil {
		return nil, err
	}
	if len(result) == 0 {
		return nil, errors.New("No sku info.")
	}

	var response SkusGetResponse
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

	return response.Data.Skus, nil
}
