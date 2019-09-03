package request

import (
	"github.com/XiBao/jos/sdk"
)

type EnqueryRegistrationDataCountRequest struct {
	Request *sdk.Request
}

// create new request
func NewEnqueryRegistrationDataCount() (req *EnqueryRegistrationDataCountRequest) {
	request := sdk.Request{MethodName: "jingdong.acty.enqueryRegistrationDataCount", Params: make(map[string]interface{}, 4)}
	req = &EnqueryRegistrationDataCountRequest{
		Request: &request,
	}
	return
}

func (req *EnqueryRegistrationDataCountRequest) SetSkuId(skuId uint64) {
	req.Request.Params["skuId"] = skuId
}

func (req *EnqueryRegistrationDataCountRequest) GetSkuId() uint64 {
	skuId, found := req.Request.Params["skuId"]
	if found {
		return skuId.(uint64)
	}
	return 0
}

func (req *EnqueryRegistrationDataCountRequest) SetOrderId(orderId uint64) {
	req.Request.Params["orderId"] = orderId
}

func (req *EnqueryRegistrationDataCountRequest) GetOrderId() uint64 {
	orderId, found := req.Request.Params["orderId"]
	if found {
		return orderId.(uint64)
	}
	return 0
}

func (req *EnqueryRegistrationDataCountRequest) SetBeginDate(beginDate string) {
	req.Request.Params["beginDate"] = beginDate
}

func (req *EnqueryRegistrationDataCountRequest) GetBeginDate() string {
	beginDate, found := req.Request.Params["beginDate"]
	if found {
		return beginDate.(string)
	}
	return ""
}

func (req *EnqueryRegistrationDataCountRequest) SetEndDate(endDate string) {
	req.Request.Params["endDate"] = endDate
}

func (req *EnqueryRegistrationDataCountRequest) GetEndDate() string {
	endDate, found := req.Request.Params["endDate"]
	if found {
		return endDate.(string)
	}
	return ""
}
