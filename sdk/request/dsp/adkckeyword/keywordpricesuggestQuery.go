package adkckeyword

import (
	"github.com/XiBao/jos/sdk"
)

type KeywordpricesuggestQueryRequest struct {
	Request *sdk.Request
}

func NewKeywordpricesuggestQueryRequest() (req *KeywordpricesuggestQueryRequest) {
	request := sdk.Request{MethodName: "jingdong.dsp.adkckeyword.keywordpricesuggest.query", Params: make(map[string]interface{}, 2)}
	req = &KeywordpricesuggestQueryRequest{
		Request: &request,
	}
	return
}

func (req *KeywordpricesuggestQueryRequest) SetKey(key string) {
	req.Request.Params["key"] = key
}

func (req *KeywordpricesuggestQueryRequest) GetKey() string {
	key, found := req.Request.Params["key"]
	if found {
		return key.(string)
	}
	return ""
}

func (req *KeywordpricesuggestQueryRequest) SetMobileType(mobileType uint8) {
	req.Request.Params["mobileType"] = mobileType
}

func (req *KeywordpricesuggestQueryRequest) GetMobileType() uint8 {
	mobileType, found := req.Request.Params["mobileType"]
	if found {
		return mobileType.(uint8)
	}
	return 0
}
