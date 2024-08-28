package adkckeyword

import (
	"github.com/XiBao/jos/sdk"
)

type KeywordInsertRequest struct {
	Request *sdk.Request
}

func NewkeywordInsertRequest() (req *KeywordInsertRequest) {
	request := sdk.Request{MethodName: "jingdong.dsp.adkckeyword.newkeyword.insert", Params: make(map[string]interface{}, 4)}
	req = &KeywordInsertRequest{
		Request: &request,
	}
	return
}

func (req *KeywordInsertRequest) SetName(name string) {
	req.Request.Params["name"] = name
}

func (req *KeywordInsertRequest) GetName() string {
	name, found := req.Request.Params["name"]
	if found {
		return name.(string)
	}
	return ""
}

func (req *KeywordInsertRequest) SetAdGroupId(adGroupId uint64) {
	req.Request.Params["adGroupId"] = adGroupId
}

func (req *KeywordInsertRequest) GetAdGroupId() uint64 {
	adGroupId, found := req.Request.Params["adGroupId"]
	if found {
		return adGroupId.(uint64)
	}
	return 0
}

func (req *KeywordInsertRequest) SetPrice(price float64) {
	req.Request.Params["price"] = price
}

func (req *KeywordInsertRequest) GetPrice() float64 {
	price, found := req.Request.Params["price"]
	if found {
		return price.(float64)
	}
	return 0
}

func (req *KeywordInsertRequest) SetType(Type uint8) {
	req.Request.Params["type"] = Type
}

func (req *KeywordInsertRequest) GetType() uint8 {
	Type, found := req.Request.Params["type"]
	if found {
		return Type.(uint8)
	}
	return 0
}
