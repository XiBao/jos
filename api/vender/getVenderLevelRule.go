package vender

import (
	"context"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/vender"
)

type GetVenderLevelRuleRequest struct {
	api.BaseRequest
}

type GetVenderLevelRuleResponse struct {
	ErrorResp *api.ErrorResponnse     `json:"error_response,omitempty" codec:"error_response,omitempty"`
	Data      *GetVenderLevelRuleData `json:"jingdong_pop_vender_getVenderLevelRule_responce,omitempty" codec:"jingdong_pop_vender_getVenderLevelRule_responce,omitempty"`
}

func (r GetVenderLevelRuleResponse) IsError() bool {
	return r.ErrorResp != nil || r.Data == nil || r.Data.IsError()
}

func (r GetVenderLevelRuleResponse) Error() string {
	if r.ErrorResp != nil {
		return r.ErrorResp.Error()
	}
	if r.Data != nil {
		return r.Data.Error()
	}
	return "no result data"
}

type GetVenderLevelRuleData struct {
	ReturnType *GetVenderLevelRuleReturnType `json:"returnType,omitempty" codec:"returnType,omitempty"`
}

func (r GetVenderLevelRuleData) IsError() bool {
	return r.ReturnType == nil || r.ReturnType.IsError()
}

func (r GetVenderLevelRuleData) Error() string {
	if r.ReturnType != nil {
		return r.ReturnType.Error()
	}
	return "no result data"
}

type GetVenderLevelRuleReturnType struct {
	Desc string             `json:"desc,omitempty" codec:"desc,omitempty"`
	Code string             `json:"code,omitempty" codec:"code,omitempty"`
	List []ShopLevelRuleDTO `json:"shopLevelRuleDTOList,omitempty" codec:"shopLevelRuleDTOList,omitempty"`
}

func (r GetVenderLevelRuleReturnType) IsError() bool {
	return r.Code != "200"
}

func (r GetVenderLevelRuleReturnType) Error() string {
	return sdk.ErrorString(r.Code, r.Desc)
}

type ShopLevelRuleDTO struct {
	VenderId          int64  `json:"venderId,omitempty" codec:"venderId,omitempty"`                   // 商家id
	CustomerLevel     int64  `json:"customerLevel,omitempty" codec:"customerLevel,omitempty"`         // 店铺会员等级
	CustomerLevelName string `json:"customerLevelName,omitempty" codec:"customerLevelName,omitempty"` // 店铺会员名称
	MinOrderPrice     int64  `json:"minOrderPrice,omitempty" codec:"minOrderPrice,omitempty"`         // 满足该级别的最低订单额
	MaxOrderPrice     int64  `json:"maxOrderPrice,omitempty" codec:"maxOrderPrice,omitempty"`         // 满足该级别的最高订单额
	MinOrderCount     int64  `json:"minOrderCount,omitempty" codec:"minOrderCount,omitempty"`         // 满足该级别的最低订单量
	MaxOrderCount     int64  `json:"maxOrderCount,omitempty" codec:"maxOrderCount,omitempty"`         // 满足该级别的最高订单量
}

// TODO 获取店铺等级体系规则
func GetVenderLevelRule(ctx context.Context, req *GetVenderLevelRuleRequest) ([]ShopLevelRuleDTO, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := vender.NewGetVenderLevelRuleRequest()

	var response GetVenderLevelRuleResponse
	if err := client.Execute(ctx, r.Request, req.Session, &response); err != nil {
		return nil, err
	}
	return response.Data.ReturnType.List, nil
}
