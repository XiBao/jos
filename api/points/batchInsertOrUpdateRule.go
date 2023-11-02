package points

import (
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

func (r BatchInsertOrUpdateRuleResponse) IsError() bool {
	return r.ErrorResp != nil || r.Data == nil || r.Data.IsError()
}

func (r BatchInsertOrUpdateRuleResponse) Error() string {
	if r.ErrorResp != nil {
		return r.ErrorResp.Error()
	}
	if r.Data != nil {
		return r.Data.Error()
	}
	return "no result data"
}

type BatchInsertOrUpdateRuleData struct {
	JsfResult *BatchInsertOrUpdateRuleJsfResult `json:"jsfResult,omitempty" codec:"jsfResult,omitempty"`
}

func (r BatchInsertOrUpdateRuleData) IsError() bool {
	return r.JsfResult == nil || r.JsfResult.IsError()
}

func (r BatchInsertOrUpdateRuleData) Error() string {
	if r.JsfResult != nil {
		return r.JsfResult.Error()
	}
	return "no result data"
}

type BatchInsertOrUpdateRuleJsfResult struct {
	Code   string `json:"code,omitempty" codec:"code,omitempty"`     //返回码
	Desc   string `json:"desc,omitempty" codec:"desc,omitempty"`     //返回描述
	Result bool   `json:"result,omitempty" codec:"result,omitempty"` //是否成功
}

func (r BatchInsertOrUpdateRuleJsfResult) IsError() bool {
	return r.Code != "200"
}

func (r BatchInsertOrUpdateRuleJsfResult) Error() string {
	return sdk.ErrorString(r.Code, r.Desc)
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

	var response BatchInsertOrUpdateRuleResponse
	if err := client.Execute(r.Request, req.Session, &response); err != nil {
		return false, err
	}

	return response.Data.JsfResult.Result, nil

}
