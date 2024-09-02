package crm

import (
	"context"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/crm"
)

type GetMemeberDiscountRequest struct {
	api.BaseRequest
}

type GetMemeberDiscountResponse struct {
	ErrorResp *api.ErrorResponnse     `json:"error_response,omitempty" codec:"error_response,omitempty"`
	Data      *GetMemeberDiscountData `json:"jingdong_pop_crm_getMemeberDiscount_responce,omitempty" codec:"jingdong_pop_crm_getMemeberDiscount_responce,omitempty"`
}

func (r GetMemeberDiscountResponse) IsError() bool {
	return r.ErrorResp != nil || r.Data == nil || r.Data.IsError()
}

func (r GetMemeberDiscountResponse) Error() string {
	if r.ErrorResp != nil {
		return r.ErrorResp.Error()
	}
	if r.Data != nil {
		return r.Data.Error()
	}
	return "no result data"
}

type GetMemeberDiscountData struct {
	ReturnType *GetMemeberDiscountReturnType `json:"returnType,omitempty" codec:"returnType,omitempty"`
}

func (r GetMemeberDiscountData) IsError() bool {
	return r.ReturnType == nil || r.ReturnType.IsError()
}

func (r GetMemeberDiscountData) Error() string {
	if r.ReturnType != nil {
		return r.ReturnType.Error()
	}
	return "no result data"
}

type GetMemeberDiscountReturnType struct {
	Desc string                `json:"desc,omitempty" codec:"desc,omitempty"` // 返回值code码
	Code string                `json:"code,omitempty" codec:"code,omitempty"` // 返回值code码描述
	Data []ShopRuleDiscountDTO `json:"data,omitempty" codec:"data,omitempty"` // 折扣信息数组   返回值：code码为200时，可能为空；code码为400、500时，为空；
}

func (r GetMemeberDiscountReturnType) IsError() bool {
	return r.Code != "200"
}

func (r GetMemeberDiscountReturnType) Error() string {
	return sdk.ErrorString(r.Code, r.Desc)
}

type ShopRuleDiscountDTO struct {
	CurGradeName string `json:"curGradeName,omitempty" codec:"curGradeName,omitempty"` // 当前会员店铺等级名称
	CurGrade     string `json:"curGrade,omitempty" codec:"curGrade,omitempty"`         // 当前会员店铺等级(1、2、3、4、5)，最少1个等级，最多5个等级
	VenderId     uint64 `json:"venderId,omitempty" codec:"venderId,omitempty"`         // 商家Id
	Discount     string `json:"discount,omitempty" codec:"discount,omitempty"`         // 会员折扣(1-9.9),为空表示未设置折扣
}

// TODO 查询会员折扣信息
func GetMemeberDiscount(ctx context.Context, req *GetMemeberDiscountRequest) ([]ShopRuleDiscountDTO, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := crm.NewGetMemeberDiscountRequest()

	var response GetMemeberDiscountResponse
	if err := client.Execute(ctx, r.Request, req.Session, &response); err != nil {
		return nil, err
	}
	return response.Data.ReturnType.Data, nil
}
