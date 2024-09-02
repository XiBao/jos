package adkckeyword

import (
	"github.com/XiBao/jos/sdk"
)

type CategorypricesuggestQueryRequest struct {
	Request *sdk.Request
}

func NewCategorypricesuggestQueryRequest() (req *CategorypricesuggestQueryRequest) {
	request := sdk.Request{MethodName: "jingdong.dsp.adkckeyword.categorypricesuggest.query", Params: make(map[string]interface{}, 2)}
	req = &CategorypricesuggestQueryRequest{
		Request: &request,
	}
	return
}

func (req *CategorypricesuggestQueryRequest) SetKey(key uint64) {
	req.Request.Params["key"] = key
}

func (req *CategorypricesuggestQueryRequest) GetKey() uint64 {
	key, found := req.Request.Params["key"]
	if found {
		return key.(uint64)
	}
	return 0
}

func (req *CategorypricesuggestQueryRequest) SetMobileType(mobileType uint8) {
	req.Request.Params["mobileType"] = mobileType
}

func (req *CategorypricesuggestQueryRequest) GetMobileType() uint8 {
	mobileType, found := req.Request.Params["mobileType"]
	if found {
		return mobileType.(uint8)
	}
	return 0
}
