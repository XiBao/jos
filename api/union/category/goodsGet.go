package category

import (
	"errors"
	"strconv"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/union/category"
	"github.com/daviddengcn/ljson"
)

type GoodsGetRequest struct {
	api.BaseRequest
	ParentId uint64 `json:"parentId"` // 父类目id(一级父类目为0)
	Grade    uint   `json:"grade"`    // 类目级别(类目级别 0，1，2 代表一、二、三级类目)
}

type GoodsGetResponse struct {
	ErrorResp *api.ErrorResponnse   `json:"error_response,omitempty"`
	Data      *GoodsGetResponseData `json:"jd_union_open_category_goods_get_responce,omitempty"`
}

type GoodsGetResponseData struct {
	Result string `json:"getResult,omitempty"`
}

type GoodsGetResult struct {
	Code    int64          `json:"code,omitempty"`
	Message string         `json:"message,omitempty"`
	Data    []CategoryResp `json:"data,omitempty"`
}

type CategoryResp struct {
	Id       uint64 `json:"id,omitempty"`       // 类目Id
	Name     string `json:"name,omitempty"`     // 类目名称
	Grade    uint   `json:"grade,omitempty"`    // 类目级别(类目级别 0，1，2 代表一、二、三级类目)
	ParentId uint64 `json:"parentId,omitempty"` // 父类目Id
}

//  商品类目查询
func GoodsGet(req *GoodsGetRequest) ([]CategoryResp, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := category.NewGoodsGetRequest()
	goodsReq := &category.GoodsGetReq{
		ParentId: req.ParentId,
		Grade:    req.Grade,
	}
	r.SetReq(goodsReq)

	result, err := client.Execute(r.Request, req.Session)
	if err != nil {
		return nil, err
	}
	var response GoodsGetResponse
	err = ljson.Unmarshal(result, &response)
	if err != nil {
		return nil, err
	}
	if response.Data == nil || response.Data.Result == "" {
		return nil, errors.New("no result")
	}
	var resp GoodsGetResult
	err = ljson.Unmarshal([]byte(response.Data.Result), &resp)
	if err != nil {
		return nil, err
	}
	if resp.Code != 200 {
		return nil, &api.ErrorResponnse{Code: strconv.FormatInt(resp.Code, 10), ZhDesc: resp.Message}
	}

	return resp.Data, nil
}
