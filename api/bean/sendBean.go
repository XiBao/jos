package bean

import (
	"errors"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/api/util"
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/bean"
	"github.com/daviddengcn/ljson"
)

type SendBeanRequest struct {
	api.BaseRequest
	RequestId   string `json:"requestId,omitempty" codec:"requestId,omitempty"` // 防重入ID
	BeanNum     uint32 `json:"beanNum" codec:"beanNum"`                         // 发豆数量
	AccountId   uint64 `json:"accountId" codec:"accountId"`                     // 京豆账户ID
	AccountCode string `json:"accountCode" codec:"accountCode"`                 // 京豆账户,当accountType=1，则accountCode为商家id;当accountType=2，则accouontCode为供应商简码
	SendWay     uint8  `json:"sendWay" codec:"sendWay"`                         // 发豆方式,目前只支持按PIN发豆，固定值为1
	SendCode    string `json:"sendCode" codec:"sendCode"`                       // 发豆参数,当sendWay=1, 为用户pin,目前只支持按PIN发豆，固定值为1
	AccountType uint8  `json:"accountType" codec:"accountType"`                 // 京豆账户类型,1:商家账户,2:供应商账户,3:品牌商账户
	Context     string `json:"context,omitempty" codec:"context,omitempty"`     // 发豆说明
	PlanId      uint64 `json:"planId" codec:"planId"`                           // 计划ID
	Rfld        string `json:"rfld,omitempty" codec:"rfld,omitempty"`           // 外部活动id
}

type SendBeanResponse struct {
	ErrorResp *ErrorResponnse `json:"error_response,omitempty" codec:"error_response,omitempty"`
	Data      *SendBeanData   `json:"jingdong_pop_bean_sendBean_responce,omitempty" codec:"jingdong_pop_bean_sendBean_responce,omitempty"`
}

type SendBeanData struct {
	Code      string          `json:"code,omitempty" codec:"code,omitempty"`
	ErrorDesc string          `json:"error_description,omitempty" codec:"error_description,omitempty"`
	Result    *SendBeanResult `json:"beanSendResult,omitempty" codec:"beanSendResult,omitempty"`
}

type SendBeanResult struct {
	Result        uint32 `json:"result" codec:"result"`
	RemainBeanNum uint64 `json:"remainBeanNum" codec:"remainBeanNum"`
	Desc          string `json:"desc" codec:"desc"`
}

func SendBean(req *SendBeanRequest) (*SendBeanResult, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := bean.NewSendBeanRequest()
	r.SetBeanNum(req.BeanNum)
	r.SetAccountId(req.AccountId)
	r.SetAccountCode(req.AccountCode)
	r.SetSendWay(req.SendWay)
	r.SetSendCode(req.SendCode)
	r.SetAccountType(req.AccountType)
	r.SetPlanId(req.PlanId)
	if len(req.RequestId) > 0 {
		r.SetRequestId(req.RequestId)
	}
	if len(req.Context) > 0 {
		r.SetContext(req.Context)
	}
	if len(req.Rfld) > 0 {
		r.SetRfld(req.Rfld)
	}

	result, err := client.Execute(r.Request, req.Session)
	if err != nil {
		return nil, err
	}
	result = util.RemoveJsonSpace(result)

	var response SendBeanResponse
	err = ljson.Unmarshal(result, &response)
	if err != nil {
		return nil, err
	}
	if response.ErrorResp != nil {
		return nil, response.ErrorResp
	}
	if response.Data.Code != "0" {
		return nil, errors.New(response.Data.ErrorDesc)
	}
	if response.Data.Result == nil {
		return nil, errors.New("No send bean result.")
	}

	return response.Data.Result, nil
}
