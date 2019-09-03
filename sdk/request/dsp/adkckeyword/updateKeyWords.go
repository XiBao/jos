package adkckeyword

import "github.com/XiBao/jos/sdk"

type UpdateKeyWordsRequest struct {
	Request *sdk.Request
}

func NewUpdateKeyWordsRequest() (req *UpdateKeyWordsRequest) {
	request := sdk.Request{MethodName: "jingdong.dsp.adkckeyword.updateKeyWords", Params: make(map[string]interface{}, 5)}
	req = &UpdateKeyWordsRequest{
		Request: &request,
	}
	return
}

func (req *UpdateKeyWordsRequest) SetName(name string) {
	req.Request.Params["name"] = name
}

func (req *UpdateKeyWordsRequest) GetName() string {
	name, found := req.Request.Params["name"]
	if found {
		return name.(string)
	}
	return ""
}

func (req *UpdateKeyWordsRequest) SetPrice(price string) {
	req.Request.Params["price"] = price
}

func (req *UpdateKeyWordsRequest) GetPrice() string {
	price, found := req.Request.Params["price"]
	if found {
		return price.(string)
	}
	return ""
}

func (req *UpdateKeyWordsRequest) SetType(Type string) {
	req.Request.Params["type"] = Type
}

func (req *UpdateKeyWordsRequest) GetType() string {
	Type, found := req.Request.Params["type"]
	if found {
		return Type.(string)
	}
	return ""
}

func (req *UpdateKeyWordsRequest) SetMobilePrice(mobilePrice string) {
	req.Request.Params["mobilePrice"] = mobilePrice
}

func (req *UpdateKeyWordsRequest) GetMobilePrice() string {
	mobilePrice, found := req.Request.Params["mobilePrice"]
	if found {
		return mobilePrice.(string)
	}
	return ""
}

func (req *UpdateKeyWordsRequest) SetAdGroupId(adGroupId uint64) {
	req.Request.Params["adGroupId"] = adGroupId
}

func (req *UpdateKeyWordsRequest) GetAdGroupId() uint64 {
	adGroupId, found := req.Request.Params["adGroupId"]
	if found {
		return adGroupId.(uint64)
	}
	return 0
}
