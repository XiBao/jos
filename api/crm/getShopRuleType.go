package crm

import (
	"errors"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/crm"
	"github.com/daviddengcn/ljson"
)

type GetShopRuleTypeRequest struct {
	api.BaseRequest
}

type GetShopRuleTypeResponse struct {
	ErrorResp *api.ErrorResponnse  `json:"error_response,omitempty" codec:"error_response,omitempty"`
	Data      *GetShopRuleTypeData `json:"jingdong_pop_crm_getShopRuleType_responce,omitempty" codec:"jingdong_pop_crm_getShopRuleType_responce,omitempty"`
}

type GetShopRuleTypeData struct {
	ReturnResult *GetShopRuleTypeReturnResult `json:"returnResult,omitempty" codec:"returnResult,omitempty"`
}

type GetShopRuleTypeReturnResult struct {
	Code string `json:"code,omitempty" codec:"code,omitempty"` //状态码
	Desc string `json:"desc,omitempty" codec:"desc,omitempty"` //参数描述
	Data uint8  `json:"data,omitempty" codec:"data,omitempty"` //会员类型 0-未开启会员规则 1-店铺已购即会员规则 2-店铺开卡规则 3- 品牌开卡规则
}

//TODO 查询商家是否开通会 员开卡功能/开卡类 型
func GetShopRuleType(req *GetShopRuleTypeRequest) (uint8, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := crm.NewGetShopRuleTypeRequest()

	result, err := client.Execute(r.Request, req.Session)
	if err != nil {
		return 0, err
	}
	if len(result) == 0 {
		return 0, errors.New("no result info")
	}
	var response GetShopRuleTypeResponse
	err = ljson.Unmarshal(result, &response)
	if err != nil {
		return 0, err
	}

	if response.ErrorResp != nil {
		return 0, response.ErrorResp
	}

	if response.Data.ReturnResult.Code != `200` {
		return 0, errors.New(response.Data.ReturnResult.Desc)
	}

	return response.Data.ReturnResult.Data, nil
}
