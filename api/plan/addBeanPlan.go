package plan

import (
	"fmt"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/plan"
)

type AddBeanPlanRequest struct {
	api.BaseRequest
	RequestId          string `json:"requestId,omitempty" codec:"requestId,omitempty"` // 可用来防重。时间有效性为60分钟，过了有效期，同样的requestId会认为是不同的调用
	ServiceMoneyBudget uint32 `json:"serviceMoneyBudget" codec:"serviceMoneyBudget"`   // 服务费预算,当为BeanPlanType.FEE时，必须大于0
	AccountType        uint8  `json:"accountType" codec:"accountType"`                 // 京豆账户类型,1:商家账户,2:供应商账户,3:品牌商账户
	AccountCode        string `json:"accountCode" codec:"accountCode"`                 // 京豆账户,当accountType=1，则accountCode为商家id;当accountType=2，则accouontCode为供应商简码
	SendTimes          uint8  `json:"sendTimes" codec:"sendTimes"`                     // 京豆每人发放次数,1-20之间
	AType              uint8  `json:"type" codec:"type"`                               // 计划类型,1:普通,2:含服务费,3:直发,8:回收,9:中转
	ModifyMode         uint8  `json:"modifyMode" codec:"modifyMode"`                   // 修改模式
	Content            string `json:"content,omitempty" codec:"content,omitempty"`     // 计划描述
	AccountId          uint64 `json:"accountId" codec:"accountId"`                     // 京豆账户ID
	BudgetNum          uint32 `json:"budgetNum" codec:"budgetNum"`                     // 预算京豆数量
	Name               string `json:"name" codec:"name"`                               // 计划名称
	Rfld               string `json:"rfld,omitempty" codec:"rfld,omitempty"`           // 外部活动id
	BeginTime          string `json:"beginTime" codec:"beginTime"`                     // 计划开始时间
	EndTime            string `json:"endTime" codec:"endTime"`                         // 计划结束时间
	SendMode           uint8  `json:"sendMode" codec:"sendMode"`                       // 发豆模式,0:不限制来源,1:限制来源
	SendRule           uint8  `json:"sendRule" codec:"sendRule"`                       // 京豆发放规则，1:京豆计划期间每人每天发放次数,2:京豆计划期间每人发放次数
	PinRiskLevel       uint8  `json:"pinRiskLevel" codec:"pinRiskLevel"`               // 风险等级,新风控参考枚举类：11-16,原风控0 5 10 ,数值越大风险越高
}

type AddBeanPlanResponse struct {
	ErrorResp *api.ErrorResponnse `json:"error_response,omitempty" codec:"error_response,omitempty"`
	Data      *AddBeanPlanData    `json:"jingdong_pop_plan_addBeanPlan_responce,omitempty" codec:"jingdong_pop_plan_addBeanPlan_responce,omitempty"`
}

func (r AddBeanPlanResponse) IsError() bool {
	return r.ErrorResp != nil || r.Data == nil || r.Data.IsError()
}

func (r AddBeanPlanResponse) Error() string {
	if r.ErrorResp != nil {
		return r.ErrorResp.Error()
	}
	if r.Data != nil {
		return r.Data.Error()
	}
	return "no result data"
}

type AddBeanPlanData struct {
	Code      string `json:"code,omitempty" codec:"code,omitempty"`
	ErrorDesc string `json:"error_description,omitempty" codec:"error_description,omitempty"`
	PlanId    uint64 `json:"planId,omitempty" codec:"planId,omitempty"`
}

func (r AddBeanPlanData) IsError() bool {
	return r.Code != "0"
}

func (r AddBeanPlanData) Error() string {
	return fmt.Sprintf("code:%s, msg:%s", r.Code, r.ErrorDesc)
}

func AddBeanPlan(req *AddBeanPlanRequest) (uint64, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := plan.NewAddBeanPlanRequest()
	r.SetServiceMoneyBudget(req.ServiceMoneyBudget)
	r.SetAccountCode(req.AccountCode)
	r.SetAccountType(req.AccountType)
	r.SetSendTimes(req.SendTimes)
	r.SetType(req.AType)
	r.SetModifyMode(req.ModifyMode)
	r.SetAccountId(req.AccountId)
	r.SetBudgetNum(req.BudgetNum)
	r.SetName(req.Name)
	r.SetBeginTime(req.BeginTime)
	r.SetEndTime(req.EndTime)
	r.SetSendMode(req.SendMode)
	r.SetSendRule(req.SendRule)
	r.SetPinRiskLevel(req.PinRiskLevel)
	if len(req.RequestId) > 0 {
		r.SetRequestId(req.RequestId)
	}
	if len(req.Content) > 0 {
		r.SetContent(req.Content)
	}
	if len(req.Rfld) > 0 {
		r.SetRfld(req.Rfld)
	}

	var response AddBeanPlanResponse
	if err := client.Execute(r.Request, req.Session, &response); err != nil {
		return 0, err
	}

	return response.Data.PlanId, nil
}
