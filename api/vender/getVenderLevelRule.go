package vender

import (
	"errors"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/vender"
	"github.com/daviddengcn/ljson"
)

type GetVenderLevelRuleRequest struct {
	api.BaseRequest
}

type GetVenderLevelRuleResponse struct {
	ErrorResp *api.ErrorResponnse     `json:"error_response,omitempty" codec:"error_response,omitempty"`
	Data      *GetVenderLevelRuleData `json:"jingdong_pop_vender_getVenderLevelRule_responce,omitempty" codec:"jingdong_pop_vender_getVenderLevelRule_responce,omitempty"`
}

type GetVenderLevelRuleData struct {
	ReturnType *GetVenderLevelRuleReturnType `json:"returnType,omitempty" codec:"returnType,omitempty"`
}

type GetVenderLevelRuleReturnType struct {
	Desc string              `json:"desc,omitempty" codec:"desc,omitempty"`
	Code string              `json:"code,omitempty" codec:"code,omitempty"`
	List []*ShopLevelRuleDTO `json:"shopLevelRuleDTOList,omitempty" codec:"shopLevelRuleDTOList,omitempty"`
}

type ShopLevelRuleDTO struct {
	VenderId          int64  `json:"venderId,omitempty" codec:"venderId,omitempty"`                   //商家id
	CustomerLevel     int64  `json:"customerLevel,omitempty" codec:"customerLevel,omitempty"`         //店铺会员等级
	CustomerLevelName string `json:"customerLevelName,omitempty" codec:"customerLevelName,omitempty"` //店铺会员名称
	MinOrderPrice     int64  `json:"minOrderPrice,omitempty" codec:"minOrderPrice,omitempty"`         //满足该级别的最低订单额
	MaxOrderPrice     int64  `json:"maxOrderPrice,omitempty" codec:"maxOrderPrice,omitempty"`         //满足该级别的最高订单额
	MinOrderCount     int64  `json:"minOrderCount,omitempty" codec:"minOrderCount,omitempty"`         //满足该级别的最低订单量
	MaxOrderCount     int64  `json:"maxOrderCount,omitempty" codec:"maxOrderCount,omitempty"`         //满足该级别的最高订单量
}

// TODO 获取店铺等级体系规则
func GetVenderLevelRule(req *GetVenderLevelRuleRequest) ([]*ShopLevelRuleDTO, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := vender.NewGetVenderLevelRuleRequest()

	result, err := client.Execute(r.Request, req.Session)
	if err != nil {
		return nil, err
	}
	if len(result) == 0 {
		return nil, errors.New("no result info")
	}
	var response GetVenderLevelRuleResponse
	err = ljson.Unmarshal(result, &response)
	if err != nil {
		return nil, err
	}

	if response.ErrorResp != nil {
		return nil, response.ErrorResp
	}

	if response.Data.ReturnType.Code != "200" {
		return nil, errors.New(response.Data.ReturnType.Desc)
	}
	return response.Data.ReturnType.List, nil
}
