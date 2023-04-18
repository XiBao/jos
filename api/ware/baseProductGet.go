package ware

import (
	"encoding/json"
	"errors"
	"strconv"
	"strings"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/ware"
)

type WareBaseProductGetRequest struct {
	api.BaseRequest
	Ids        []uint64 `json:"ids,omitempty" codec:"ids,omitempty"`                 //
	BaseFields string   `json:"base_fields,omitempty" codec:"base_fields,omitempty"` // 自定义返回字段
}

type WareBaseProductGetResponse struct {
	ErrorResp *api.ErrorResponnse     `json:"error_response,omitempty" codec:"error_response,omitempty"`
	Data      *WareBaseProductGetData `json:"jingdong_new_ware_baseproduct_get_responce,omitempty" codec:"jingdong_new_ware_baseproduct_get_responce,omitempty"`
}

type WareBaseProductGetData struct {
	Code   string          `json:"code,omitempty" codec:"code,omitempty"`
	Result []*ProductsBase `json:"listproductbase_result,omitempty" codec:"listproductbase_result,omitempty"`
}

// 获取单个SKU
func WareBaseProductGet(req *WareBaseProductGetRequest) ([]*ProductsBase, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := ware.NewWareBaseproductGetRequest()
	var ids []string
	for _, v := range req.Ids {
		ids = append(ids, strconv.FormatUint(v, 10))
	}
	r.SetIds(strings.Join(ids, ","))
	r.SetBaseFields(req.BaseFields)

	result, err := client.Execute(r.Request, req.Session)
	if err != nil {
		return nil, err
	}
	if len(result) == 0 {
		return nil, errors.New("No result.")
	}
	var response WareBaseProductGetResponse
	err = json.Unmarshal(result, &response)
	if err != nil {
		return nil, err
	}

	if response.ErrorResp != nil {
		return nil, response.ErrorResp
	}
	return response.Data.Result, nil
}
