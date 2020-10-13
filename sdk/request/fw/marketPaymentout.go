package fw

import "github.com/XiBao/jos/sdk"

type MarketPaymentoutRequest struct {
	Request *sdk.Request
}

// create new request
func NewMarketPaymentoutRequest() (req *MarketPaymentoutRequest) {
	request := sdk.Request{MethodName: "jingdong.fw.market.paymentout", Params: make(map[string]interface{}, 4)}
	req = &MarketPaymentoutRequest{
		Request: &request,
	}
	return
}

func (req *MarketPaymentoutRequest) SetRequestNo(requestNo string) {
	req.Request.Params["requestNo"] = requestNo
}

func (req *MarketPaymentoutRequest) GetRequestNo() string {
	requestNo, found := req.Request.Params["requestNo"]
	if found {
		return requestNo.(string)
	}
	return ""
}

func (req *MarketPaymentoutRequest) SetActivityId(activityId uint) {
	req.Request.Params["activityId"] = activityId
}

func (req *MarketPaymentoutRequest) GetActivityId() uint {
	activityId, found := req.Request.Params["activityId"]
	if found {
		return activityId.(uint)
	}
	return 0
}

func (req *MarketPaymentoutRequest) SetAppId(appId string) {
	req.Request.Params["appId"] = appId
}

func (req *MarketPaymentoutRequest) GetAppId() string {
	appId, found := req.Request.Params["appId"]
	if found {
		return appId.(string)
	}
	return ""
}

func (req *MarketPaymentoutRequest) SetPrice(price uint) {
	req.Request.Params["price"] = price
}

func (req *MarketPaymentoutRequest) GetPrice() uint {
	price, found := req.Request.Params["price"]
	if found {
		return price.(uint)
	}
	return 0
}

func (req *MarketPaymentoutRequest) SetIsMainService(isMainService bool) {
	req.Request.Params["isMainService"] = isMainService
}

func (req *MarketPaymentoutRequest) GetIsMainService() bool {
	isMainService, found := req.Request.Params["isMainService"]
	if found {
		return isMainService.(bool)
	}
	return false
}

func (req *MarketPaymentoutRequest) SetServiceCycle(serviceCycle uint) {
	req.Request.Params["serviceCycle"] = serviceCycle
}

func (req *MarketPaymentoutRequest) GetServiceCycle() uint {
	serviceCycle, found := req.Request.Params["serviceCycle"]
	if found {
		return serviceCycle.(uint)
	}
	return 0
}

func (req *MarketPaymentoutRequest) SetSkuId(skuId uint64) {
	req.Request.Params["skuId"] = skuId
}

func (req *MarketPaymentoutRequest) GetSkuId() uint64 {
	skuId, found := req.Request.Params["skuId"]
	if found {
		return skuId.(uint64)
	}
	return 0
}

func (req *MarketPaymentoutRequest) SetServiceCode(serviceCode string) {
	req.Request.Params["serviceCode"] = serviceCode
}

func (req *MarketPaymentoutRequest) GetServiceCode() string {
	serviceCode, found := req.Request.Params["serviceCode"]
	if found {
		return serviceCode.(string)
	}
	return ""
}

func (req *MarketPaymentoutRequest) SetOrderNum(orderNum uint) {
	req.Request.Params["orderNum"] = orderNum
}

func (req *MarketPaymentoutRequest) GetOrderNum() uint {
	orderNum, found := req.Request.Params["orderNum"]
	if found {
		return orderNum.(uint)
	}
	return 0
}

func (req *MarketPaymentoutRequest) SetItemCode(itemCode string) {
	req.Request.Params["itemCode"] = itemCode
}

func (req *MarketPaymentoutRequest) GetItemCode() string {
	itemCode, found := req.Request.Params["itemCode"]
	if found {
		return itemCode.(string)
	}
	return ""
}

func (req *MarketPaymentoutRequest) SetOutOrderId(outOrderId uint) {
	req.Request.Params["outOrderId"] = outOrderId
}

func (req *MarketPaymentoutRequest) GetOutOrderId() uint {
	outOrderId, found := req.Request.Params["outOrderId"]
	if found {
		return outOrderId.(uint)
	}
	return 0
}

func (req *MarketPaymentoutRequest) SetValue1(value1 string) {
	req.Request.Params["value1"] = value1
}

func (req *MarketPaymentoutRequest) GetValue1() string {
	value1, found := req.Request.Params["value1"]
	if found {
		return value1.(string)
	}
	return ""
}

func (req *MarketPaymentoutRequest) SetResultPageType(resultPageType uint) {
	req.Request.Params["resultPageType"] = resultPageType
}

func (req *MarketPaymentoutRequest) GetResultPageType() uint {
	resultPageType, found := req.Request.Params["resultPageType"]
	if found {
		return resultPageType.(uint)
	}
	return 0
}

func (req *MarketPaymentoutRequest) SetSuccessUrl(successUrl string) {
	req.Request.Params["successUrl"] = successUrl
}

func (req *MarketPaymentoutRequest) GetSuccessUrl() string {
	successUrl, found := req.Request.Params["successUrl"]
	if found {
		return successUrl.(string)
	}
	return ""
}

func (req *MarketPaymentoutRequest) SetIp(ip string) {
	req.Request.Params["ip"] = ip
}

func (req *MarketPaymentoutRequest) GetIp() string {
	ip, found := req.Request.Params["ip"]
	if found {
		return ip.(string)
	}
	return ""
}
