package crm

import (
	"context"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/crm"
)

type GetShopRuleTypeRequest struct {
	api.BaseRequest
}

type GetShopRuleTypeResponse struct {
	ErrorResp *api.ErrorResponnse  `json:"error_response,omitempty" codec:"error_response,omitempty"`
	Data      *GetShopRuleTypeData `json:"jingdong_pop_crm_getShopRuleType_responce,omitempty" codec:"jingdong_pop_crm_getShopRuleType_responce,omitempty"`
}

func (r GetShopRuleTypeResponse) IsError() bool {
	return r.ErrorResp != nil || r.Data == nil || r.Data.IsError()
}

func (r GetShopRuleTypeResponse) Error() string {
	if r.ErrorResp != nil {
		return r.ErrorResp.Error()
	}
	if r.Data != nil {
		return r.Data.Error()
	}
	return "no result data"
}

type GetShopRuleTypeData struct {
	ReturnResult *GetShopRuleTypeReturnResult `json:"returnResult,omitempty" codec:"returnResult,omitempty"`
}

func (r GetShopRuleTypeData) IsError() bool {
	return r.ReturnResult == nil || r.ReturnResult.IsError()
}

func (r GetShopRuleTypeData) Error() string {
	if r.ReturnResult != nil {
		return r.ReturnResult.Error()
	}
	return "no result data"
}

type GetShopRuleTypeReturnResult struct {
	Code string `json:"code,omitempty" codec:"code,omitempty"` // 状态码
	Desc string `json:"desc,omitempty" codec:"desc,omitempty"` // 参数描述
	Data uint8  `json:"data,omitempty" codec:"data,omitempty"` // 会员类型 0-未开启会员规则 1-店铺已购即会员规则 2-店铺开卡规则 3- 品牌开卡规则
}

func (r GetShopRuleTypeReturnResult) IsError() bool {
	return r.Code != "200"
}

func (r GetShopRuleTypeReturnResult) Error() string {
	return sdk.ErrorString(r.Code, r.Desc)
}

// TODO 查询商家是否开通会 员开卡功能/开卡类 型
func GetShopRuleType(ctx context.Context, req *GetShopRuleTypeRequest) (uint8, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := crm.NewGetShopRuleTypeRequest()

	var response GetShopRuleTypeResponse
	if err := client.Execute(ctx, r.Request, req.Session, &response); err != nil {
		return 0, err
	}
	return response.Data.ReturnResult.Data, nil
}
