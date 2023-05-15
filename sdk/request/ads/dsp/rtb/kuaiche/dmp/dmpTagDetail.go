package dmp

import (
	"github.com/XiBao/jos/sdk"
)

type KuaicheDmpNewTagDetailRequest struct {
	Request *sdk.Request
}

// create new request
func NewKuaicheDmpNewTagDetailRequest() (req *KuaicheDmpNewTagDetailRequest) {
	request := sdk.Request{MethodName: "jingdong.dmp.new.tag.detail", Params: make(map[string]interface{}, 6)}
	req = &KuaicheDmpNewTagDetailRequest{
		Request: &request,
	}
	return
}

func (req *KuaicheDmpNewTagDetailRequest) SetAccessPin(accessPin string) {
	req.Request.Params["accessPin"] = accessPin
}

func (req *KuaicheDmpNewTagDetailRequest) GetAccessPin() string {
	accessPin, found := req.Request.Params["accessPin"]
	if found {
		return accessPin.(string)
	}
	return ""
}

func (req *KuaicheDmpNewTagDetailRequest) SetAuthType(authType string) {
	req.Request.Params["authType"] = authType
}

func (req *KuaicheDmpNewTagDetailRequest) GetAuthType() string {
	authType, found := req.Request.Params["authType"]
	if found {
		return authType.(string)
	}
	return ""
}

func (req *KuaicheDmpNewTagDetailRequest) SetTagId(tagId uint64) {
	req.Request.Params["tagId"] = tagId
}

func (req *KuaicheDmpNewTagDetailRequest) GetTagId() uint64 {
	tagId, found := req.Request.Params["tagId"]
	if found {
		return tagId.(uint64)
	}
	return 0
}

func (req *KuaicheDmpNewTagDetailRequest) SetCrowdId(crowdId int64) {
	req.Request.Params["crowdId"] = crowdId
}

func (req *KuaicheDmpNewTagDetailRequest) GetCrowdId() int64 {
	crowdId, found := req.Request.Params["crowdId"]
	if found {
		return crowdId.(int64)
	}
	return 0
}

func (req *KuaicheDmpNewTagDetailRequest) SetIndustryHot(industryHot string) {
	req.Request.Params["industryHot"] = industryHot
}

func (req *KuaicheDmpNewTagDetailRequest) GetIndustryHot() string {
	industryHot, found := req.Request.Params["industryHot"]
	if found {
		return industryHot.(string)
	}
	return ""
}

func (req *KuaicheDmpNewTagDetailRequest) SetCoverageRate(coverageRate string) {
	req.Request.Params["coverageRate"] = coverageRate
}

func (req *KuaicheDmpNewTagDetailRequest) GetCoverageRate() string {
	coverageRate, found := req.Request.Params["coverageRate"]
	if found {
		return coverageRate.(string)
	}
	return ""
}
