package crm

import (
	"errors"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/crm"
	"github.com/daviddengcn/ljson"
)

type SetShopLevelRuleRequest struct {
	api.BaseRequest
	CustomerLevelName []string `json:"customerLevelName,omitempty"` //按顺序填写店铺会员等级名称
}

type SetShopLevelRuleResponse struct {
	ErrorResp *api.ErrorResponnse   `json:"error_response,omitempty" codec:"error_response,omitempty"`
	Data      *SetShopLevelRuleData `json:"jingdong_pop_crm_setShopLevelRule_responce,omitempty" codec:"jingdong_pop_crm_setShopLevelRule_responce,omitempty"`
}

type SetShopLevelRuleData struct {
	ReturnType *ReturnType `json:"returnType,omitempty" codec:"returnType,omitempty"`
}

//TODO 修改会员体系规则
func SetShopLevelRule(req *SetShopLevelRuleRequest) (bool, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := crm.NewSetShopLevelRuleRequest()

	if len(req.CustomerLevelName) > 0 {
		r.SetCustomerLevelName(req.CustomerLevelName)
	}
	result, err := client.Execute(r.Request, req.Session)
	if err != nil {
		return false, err
	}
	if len(result) == 0 {
		return false, errors.New("no result info")
	}
	var response SetShopLevelRuleResponse
	err = ljson.Unmarshal(result, &response)
	if err != nil {
		return false, err
	}

	if response.ErrorResp != nil {
		return false, response.ErrorResp
	}

	if response.Data.ReturnType.Code != "200" {
		return false, errors.New(response.Data.ReturnType.Desc)
	}

	return response.Data.ReturnType.Data, nil

}
