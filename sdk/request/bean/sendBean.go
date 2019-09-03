package bean

import (
	"github.com/XiBao/jos/sdk"
)

type SendBeanRequest struct {
	Request *sdk.Request
}

// create new request
func NewSendBeanRequest() (req *SendBeanRequest) {
	request := sdk.Request{MethodName: "jingdong.pop.bean.sendBean", Params: make(map[string]interface{}, 10)}
	req = &SendBeanRequest{
		Request: &request,
	}
	return
}

func (req *SendBeanRequest) SetRequestId(requestId string) {
	req.Request.Params["requestId"] = requestId
}

func (req *SendBeanRequest) GetRequestId() string {
	requestId, found := req.Request.Params["requestId"]
	if found {
		return requestId.(string)
	}
	return ""
}

func (req *SendBeanRequest) SetBeanNum(beanNum uint32) {
	req.Request.Params["beanNum"] = beanNum
}

func (req *SendBeanRequest) GetBeanNum() uint32 {
	beanNum, found := req.Request.Params["beanNum"]
	if found {
		return beanNum.(uint32)
	}
	return 0
}

func (req *SendBeanRequest) SetAccountId(accountId uint64) {
	req.Request.Params["accountId"] = accountId
}

func (req *SendBeanRequest) GetAccountId() uint64 {
	accountId, found := req.Request.Params["accountId"]
	if found {
		return accountId.(uint64)
	}
	return 0
}

func (req *SendBeanRequest) SetAccountCode(accountCode string) {
	req.Request.Params["accountCode"] = accountCode
}

func (req *SendBeanRequest) GetAccountCode() string {
	accountCode, found := req.Request.Params["accountCode"]
	if found {
		return accountCode.(string)
	}
	return ""
}

func (req *SendBeanRequest) SetSendWay(sendWay uint8) {
	req.Request.Params["sendWay"] = sendWay
}

func (req *SendBeanRequest) GetSendWay() uint8 {
	sendWay, found := req.Request.Params["sendWay"]
	if found {
		return sendWay.(uint8)
	}
	return 0
}

func (req *SendBeanRequest) SetSendCode(sendCode string) {
	req.Request.Params["sendCode"] = sendCode
}

func (req *SendBeanRequest) GetSendCode() string {
	sendCode, found := req.Request.Params["sendCode"]
	if found {
		return sendCode.(string)
	}
	return ""
}

func (req *SendBeanRequest) SetAccountType(accountType uint8) {
	req.Request.Params["accountType"] = accountType
}

func (req *SendBeanRequest) GetAccountType() uint8 {
	accountType, found := req.Request.Params["accountType"]
	if found {
		return accountType.(uint8)
	}
	return 0
}

func (req *SendBeanRequest) SetContext(context string) {
	req.Request.Params["context"] = context
}

func (req *SendBeanRequest) GetContext() string {
	context, found := req.Request.Params["context"]
	if found {
		return context.(string)
	}
	return ""
}

func (req *SendBeanRequest) SetPlanId(planId uint64) {
	req.Request.Params["planId"] = planId
}

func (req *SendBeanRequest) GetPlanId() uint64 {
	planId, found := req.Request.Params["planId"]
	if found {
		return planId.(uint64)
	}
	return 0
}

func (req *SendBeanRequest) SetRfld(rfld string) {
	req.Request.Params["rfld"] = rfld
}

func (req *SendBeanRequest) GetRfld() string {
	rfld, found := req.Request.Params["rfld"]
	if found {
		return rfld.(string)
	}
	return ""
}
