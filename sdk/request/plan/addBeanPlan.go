package plan

import (
	"github.com/XiBao/jos/sdk"
)

type AddBeanPlanRequest struct {
	Request *sdk.Request
}

// create new request
func NewAddBeanPlanRequest() (req *AddBeanPlanRequest) {
	request := sdk.Request{MethodName: "jingdong.pop.plan.addBeanPlan", Params: make(map[string]interface{}, 17)}
	req = &AddBeanPlanRequest{
		Request: &request,
	}
	return
}

func (req *AddBeanPlanRequest) SetRequestId(requestId string) {
	req.Request.Params["requestId"] = requestId
}

func (req *AddBeanPlanRequest) GetRequestId() string {
	requestId, found := req.Request.Params["requestId"]
	if found {
		return requestId.(string)
	}
	return ""
}

func (req *AddBeanPlanRequest) SetServiceMoneyBudget(serviceMoneyBudget uint32) {
	req.Request.Params["serviceMoneyBudget"] = serviceMoneyBudget
}

func (req *AddBeanPlanRequest) GetServiceMoneyBudget() uint32 {
	serviceMoneyBudget, found := req.Request.Params["serviceMoneyBudget"]
	if found {
		return serviceMoneyBudget.(uint32)
	}
	return 0
}

func (req *AddBeanPlanRequest) SetAccountType(accountType uint8) {
	req.Request.Params["accountType"] = accountType
}

func (req *AddBeanPlanRequest) GetAccountType() uint8 {
	accountType, found := req.Request.Params["accountType"]
	if found {
		return accountType.(uint8)
	}
	return 0
}

func (req *AddBeanPlanRequest) SetAccountCode(accountCode string) {
	req.Request.Params["accountCode"] = accountCode
}

func (req *AddBeanPlanRequest) GetAccountCode() string {
	accountCode, found := req.Request.Params["accountCode"]
	if found {
		return accountCode.(string)
	}
	return ""
}

func (req *AddBeanPlanRequest) SetSendTimes(sendTimes uint8) {
	req.Request.Params["sendTimes"] = sendTimes
}

func (req *AddBeanPlanRequest) GetSendTimes() uint8 {
	sendTimes, found := req.Request.Params["sendTimes"]
	if found {
		return sendTimes.(uint8)
	}
	return 0
}

func (req *AddBeanPlanRequest) SetType(aType uint8) {
	req.Request.Params["type"] = aType
}

func (req *AddBeanPlanRequest) GetType() uint8 {
	aType, found := req.Request.Params["type"]
	if found {
		return aType.(uint8)
	}
	return 0
}

func (req *AddBeanPlanRequest) SetModifyMode(modifyMode uint8) {
	req.Request.Params["modifyMode"] = modifyMode
}

func (req *AddBeanPlanRequest) GetModifyMode() uint8 {
	modifyMode, found := req.Request.Params["modifyMode"]
	if found {
		return modifyMode.(uint8)
	}
	return 0
}

func (req *AddBeanPlanRequest) SetContent(content string) {
	req.Request.Params["content"] = content
}

func (req *AddBeanPlanRequest) GetContent() string {
	content, found := req.Request.Params["content"]
	if found {
		return content.(string)
	}
	return ""
}

func (req *AddBeanPlanRequest) SetAccountId(accountId uint64) {
	req.Request.Params["accountId"] = accountId
}

func (req *AddBeanPlanRequest) GetAccountId() uint64 {
	accountId, found := req.Request.Params["accountId"]
	if found {
		return accountId.(uint64)
	}
	return 0
}

func (req *AddBeanPlanRequest) SetBudgetNum(budgetNum uint32) {
	req.Request.Params["budgetNum"] = budgetNum
}

func (req *AddBeanPlanRequest) GetBudgetNum() uint32 {
	budgetNum, found := req.Request.Params["budgetNum"]
	if found {
		return budgetNum.(uint32)
	}
	return 0
}

func (req *AddBeanPlanRequest) SetName(name string) {
	req.Request.Params["name"] = name
}

func (req *AddBeanPlanRequest) GetName() string {
	name, found := req.Request.Params["name"]
	if found {
		return name.(string)
	}
	return ""
}

func (req *AddBeanPlanRequest) SetRfld(rfld string) {
	req.Request.Params["rfld"] = rfld
}

func (req *AddBeanPlanRequest) GetRfld() string {
	rfld, found := req.Request.Params["rfld"]
	if found {
		return rfld.(string)
	}
	return ""
}

func (req *AddBeanPlanRequest) SetBeginTime(beginTime string) {
	req.Request.Params["beginTime"] = beginTime
}

func (req *AddBeanPlanRequest) GetBeginTime() string {
	beginTime, found := req.Request.Params["beginTime"]
	if found {
		return beginTime.(string)
	}
	return ""
}

func (req *AddBeanPlanRequest) SetEndTime(endTime string) {
	req.Request.Params["endTime"] = endTime
}

func (req *AddBeanPlanRequest) GetEndTime() string {
	endTime, found := req.Request.Params["endTime"]
	if found {
		return endTime.(string)
	}
	return ""
}

func (req *AddBeanPlanRequest) SetSendMode(sendMode uint8) {
	req.Request.Params["sendMode"] = sendMode
}

func (req *AddBeanPlanRequest) GetSendMode() uint8 {
	sendMode, found := req.Request.Params["sendMode"]
	if found {
		return sendMode.(uint8)
	}
	return 0
}

func (req *AddBeanPlanRequest) SetSendRule(sendRule uint8) {
	req.Request.Params["sendRule"] = sendRule
}

func (req *AddBeanPlanRequest) GetSendRule() uint8 {
	sendRule, found := req.Request.Params["sendRule"]
	if found {
		return sendRule.(uint8)
	}
	return 0
}

func (req *AddBeanPlanRequest) SetPinRiskLevel(pinRiskLevel uint8) {
	req.Request.Params["pinRiskLevel"] = pinRiskLevel
}

func (req *AddBeanPlanRequest) GetPinRiskLevel() uint8 {
	pinRiskLevel, found := req.Request.Params["pinRiskLevel"]
	if found {
		return pinRiskLevel.(uint8)
	}
	return 0
}
