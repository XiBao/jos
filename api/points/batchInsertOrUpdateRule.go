package points

import (
	"encoding/json"
	"errors"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/points"
)

type BatchInsertOrUpdateRuleRequest struct {
	api.BaseRequest
	Multiple   float64 `json:"multiple"`   //兑换倍数
	CreateTime string  `json:"createTime"` //创建记录时间  2006-01-02 15:04:05
	ModifyTime string  `json:"modifyTime"` //记录修改时间  2006-01-02 15:04:05
}

type BatchInsertOrUpdateRuleResponse struct {
	ErrorResp *api.ErrorResponnse          `json:"error_response,omitempty" codec:"error_response,omitempty"`
	Data      *BatchInsertOrUpdateRuleData `json:"jingdong_points_jos_batchInsertOrUpdateRule_responce,omitempty" codec:"jingdong_points_jos_batchInsertOrUpdateRule_responce,omitempty"`
}

type BatchInsertOrUpdateRuleData struct {
	JsfResult *BatchInsertOrUpdateRuleJsfResult `json:"jsfResult,omitempty" codec:"jsfResult,omitempty"`
}

type BatchInsertOrUpdateRuleJsfResult struct {
	Code   string `json:"code,omitempty" codec:"code,omitempty"`     //返回码
	Desc   string `json:"desc,omitempty" codec:"desc,omitempty"`     //返回描述
	Result bool   `json:"result,omitempty" codec:"result,omitempty"` //是否成功
}

// TODO 设置积分规则   按商家后台规则进行设置
func BatchInsertOrUpdateRule(req *BatchInsertOrUpdateRuleRequest) (bool, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := points.NewBatchInsertOrUpdateRuleRequest()

	if req.Multiple > 0 {
		r.SetMultiple(req.Multiple)
	}

	if len(req.CreateTime) > 0 {
		r.SetCreateTime(req.CreateTime)
	}

	if len(req.ModifyTime) > 0 {
		r.SetModifyTime(req.ModifyTime)
	}

	result, err := client.Execute(r.Request, req.Session)
	if err != nil {
		return false, err
	}
	if len(result) == 0 {
		return false, errors.New("no result info")
	}
	var response BatchInsertOrUpdateRuleResponse
	err = json.Unmarshal(result, &response)
	if err != nil {
		return false, err
	}

	if response.ErrorResp != nil {
		return false, response.ErrorResp
	}

	if response.Data.JsfResult.Code != "200" {
		return false, errors.New(response.Data.JsfResult.Desc)
	}

	return response.Data.JsfResult.Result, nil

}
