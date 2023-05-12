package dmp

import (
	"github.com/XiBao/jos/sdk"
)

type KuaicheDmpCommonCrowdQueryRequest struct {
	Request *sdk.Request
}

// create new request
func NewKuaicheDmpCommonCrowdQueryRequest() (req *KuaicheDmpCommonCrowdQueryRequest) {
	request := sdk.Request{MethodName: "jingdong.dmp.common.crowd.query", Params: make(map[string]interface{}, 14)}
	req = &KuaicheDmpCommonCrowdQueryRequest{
		Request: &request,
	}
	return
}

func (req *KuaicheDmpCommonCrowdQueryRequest) SetAdGroupAdType(adGroupAdType uint) {
	req.Request.Params["adGroupAdType"] = adGroupAdType
}

func (req *KuaicheDmpCommonCrowdQueryRequest) GetAdGroupAdType() uint {
	adGroupAdType, found := req.Request.Params["adGroupAdType"]
	if found {
		return adGroupAdType.(uint)
	}
	return 0
}

func (req *KuaicheDmpCommonCrowdQueryRequest) SetAdGroupBidPrice(adGroupBidPrice uint) {
	req.Request.Params["adGroupBidPrice"] = adGroupBidPrice
}

func (req *KuaicheDmpCommonCrowdQueryRequest) GetAdGroupBidPrice() uint {
	adGroupBidPrice, found := req.Request.Params["adGroupBidPrice"]
	if found {
		return adGroupBidPrice.(uint)
	}
	return 0
}

func (req *KuaicheDmpCommonCrowdQueryRequest) SetAdGroupBillingType(adGroupBillingType uint) {
	req.Request.Params["adGroupBillingType"] = adGroupBillingType
}

func (req *KuaicheDmpCommonCrowdQueryRequest) GetAdGroupBillingType() uint {
	adGroupBillingType, found := req.Request.Params["adGroupBillingType"]
	if found {
		return adGroupBillingType.(uint)
	}
	return 0
}

func (req *KuaicheDmpCommonCrowdQueryRequest) SetAdGroupId(adGroupId uint64) {
	req.Request.Params["adGroupId"] = adGroupId
}

func (req *KuaicheDmpCommonCrowdQueryRequest) GetAdGroupId() uint64 {
	adGroupId, found := req.Request.Params["adGroupId"]
	if found {
		return adGroupId.(uint64)
	}
	return 0
}

func (req *KuaicheDmpCommonCrowdQueryRequest) SetBusinessType(businessType uint) {
	req.Request.Params["businessType"] = businessType
}

func (req *KuaicheDmpCommonCrowdQueryRequest) GetBusinessType() uint {
	businessType, found := req.Request.Params["businessType"]
	if found {
		return businessType.(uint)
	}
	return 0
}

func (req *KuaicheDmpCommonCrowdQueryRequest) SetCrowdName(crowdName string) {
	req.Request.Params["crowdName"] = crowdName
}

func (req *KuaicheDmpCommonCrowdQueryRequest) GetCrowdName() string {
	crowdName, found := req.Request.Params["crowdName"]
	if found {
		return crowdName.(string)
	}
	return ""
}

func (req *KuaicheDmpCommonCrowdQueryRequest) SetCrowdTabType(crowdTabType uint) {
	req.Request.Params["crowdTabType"] = crowdTabType
}

func (req *KuaicheDmpCommonCrowdQueryRequest) GetCrowdTabType() uint {
	crowdTabType, found := req.Request.Params["crowdTabType"]
	if found {
		return crowdTabType.(uint)
	}
	return 0
}

func (req *KuaicheDmpCommonCrowdQueryRequest) SetFirstSenceCategory(firstSenceCategory string) {
	req.Request.Params["firstSenceCategory"] = firstSenceCategory
}

func (req *KuaicheDmpCommonCrowdQueryRequest) GetFirstSenceCategory() string {
	firstSenceCategory, found := req.Request.Params["firstSenceCategory"]
	if found {
		return firstSenceCategory.(string)
	}
	return ""
}

func (req *KuaicheDmpCommonCrowdQueryRequest) SetPageIndex(pageIndex int) {
	req.Request.Params["pageIndex"] = pageIndex
}

func (req *KuaicheDmpCommonCrowdQueryRequest) GetPageIndex() int {
	pageIndex, found := req.Request.Params["pageIndex"]
	if found {
		return pageIndex.(int)
	}
	return 0
}

func (req *KuaicheDmpCommonCrowdQueryRequest) SetPageSize(pageSize int) {
	req.Request.Params["pageSize"] = pageSize
}

func (req *KuaicheDmpCommonCrowdQueryRequest) GetPageSize() int {
	pageSize, found := req.Request.Params["pageSize"]
	if found {
		return pageSize.(int)
	}
	return 0
}

func (req *KuaicheDmpCommonCrowdQueryRequest) SetSecondSenceCategory(secondSenceCategory string) {
	req.Request.Params["secondSenceCategory"] = secondSenceCategory
}

func (req *KuaicheDmpCommonCrowdQueryRequest) GetSecondSenceCategory() string {
	secondSenceCategory, found := req.Request.Params["secondSenceCategory"]
	if found {
		return secondSenceCategory.(string)
	}
	return ""
}

func (req *KuaicheDmpCommonCrowdQueryRequest) SetResourcesList(resourcesList string) {
	req.Request.Params["resourcesList"] = resourcesList
}

func (req *KuaicheDmpCommonCrowdQueryRequest) GetResourcesList() string {
	resourcesList, found := req.Request.Params["resourcesList"]
	if found {
		return resourcesList.(string)
	}
	return ""
}

func (req *KuaicheDmpCommonCrowdQueryRequest) SetAccessPin(accessPin string) {
	req.Request.Params["accessPin"] = accessPin
}

func (req *KuaicheDmpCommonCrowdQueryRequest) GetAccessPin() string {
	accessPin, found := req.Request.Params["accessPin"]
	if found {
		return accessPin.(string)
	}
	return ""
}

func (req *KuaicheDmpCommonCrowdQueryRequest) SetAuthType(authType string) {
	req.Request.Params["authType"] = authType
}

func (req *KuaicheDmpCommonCrowdQueryRequest) GetAuthType() string {
	authType, found := req.Request.Params["authType"]
	if found {
		return authType.(string)
	}
	return ""
}
