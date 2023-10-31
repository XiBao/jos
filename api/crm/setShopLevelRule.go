package crm

import (
	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/crm"
)

type SetShopLevelRuleRequest struct {
	api.BaseRequest
	CustomerLevelName []string `json:"customerLevelName,omitempty"` //按顺序填写店铺会员等级名称
}

type SetShopLevelRuleResponse struct {
	ErrorResp *api.ErrorResponnse   `json:"error_response,omitempty" codec:"error_response,omitempty"`
	Data      *SetShopLevelRuleData `json:"jingdong_pop_crm_setShopLevelRule_responce,omitempty" codec:"jingdong_pop_crm_setShopLevelRule_responce,omitempty"`
}

func (r SetShopLevelRuleResponse) IsError() bool {
	return r.ErrorResp != nil || r.Data == nil || r.Data.IsError()
}

func (r SetShopLevelRuleResponse) Error() string {
	if r.ErrorResp != nil {
		return r.ErrorResp.Error()
	}
	if r.Data != nil {
		return r.Data.Error()
	}
	return "no result data"
}

type SetShopLevelRuleData struct {
	ReturnType *ReturnType `json:"returnType,omitempty" codec:"returnType,omitempty"`
}

func (r SetShopLevelRuleData) IsError() bool {
	return r.ReturnType == nil || r.ReturnType.IsError()
}

func (r SetShopLevelRuleData) Error() string {
	if r.ReturnType != nil {
		return r.ReturnType.Error()
	}
	return "no result data"
}

// TODO 修改会员体系规则
func SetShopLevelRule(req *SetShopLevelRuleRequest) (bool, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := crm.NewSetShopLevelRuleRequest()

	if len(req.CustomerLevelName) > 0 {
		r.SetCustomerLevelName(req.CustomerLevelName)
	}

	var response SetShopLevelRuleResponse
	if err := client.Execute(r.Request, req.Session, &response); err != nil {
		return false, err
	}
	return response.Data.ReturnType.Data, nil

}
