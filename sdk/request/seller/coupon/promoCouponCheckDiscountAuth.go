package coupon

import (
	"github.com/XiBao/jos/sdk"
)

type PromoCouponCheckDiscountAuthRequest struct {
	Request *sdk.Request
}

// create new request
func NewPromoCouponCheckDiscountAuthRequest() (req *PromoCouponCheckDiscountAuthRequest) {
	request := sdk.Request{MethodName: "jingdong.pop.promo.coupon.checkDiscountAuth", Params: make(map[string]interface{}, 4)}
	req = &PromoCouponCheckDiscountAuthRequest{
		Request: &request,
	}
	return
}

func (req *PromoCouponCheckDiscountAuthRequest) SetAppName(appName string) {
	req.Request.Params["appName"] = appName
}

func (req *PromoCouponCheckDiscountAuthRequest) GetAppName() string {
	appName, found := req.Request.Params["appName"]
	if found {
		return appName.(string)
	}
	return ""
}

func (req *PromoCouponCheckDiscountAuthRequest) SetAppId(appId string) {
	req.Request.Params["appId"] = appId
}

func (req *PromoCouponCheckDiscountAuthRequest) GetAppId() string {
	appId, found := req.Request.Params["appId"]
	if found {
		return appId.(string)
	}
	return ""
}

func (req *PromoCouponCheckDiscountAuthRequest) SetUuid(uuid string) {
	req.Request.Params["uuid"] = uuid
}

func (req *PromoCouponCheckDiscountAuthRequest) GetUuid() string {
	uuid, found := req.Request.Params["uuid"]
	if found {
		return uuid.(string)
	}
	return ""
}

func (req *PromoCouponCheckDiscountAuthRequest) SetIp(ip string) {
	req.Request.Params["ip"] = ip
}

func (req *PromoCouponCheckDiscountAuthRequest) GetIp() string {
	ip, found := req.Request.Params["ip"]
	if found {
		return ip.(string)
	}
	return ""
}
