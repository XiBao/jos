package shorturl

import (
	"github.com/XiBao/jos/sdk"
)

type GenerateURLFastestRequest struct {
	Request *sdk.Request
}

// create new request
func NewGenerateURLFastestRequest() (req *GenerateURLFastestRequest) {
	request := sdk.Request{MethodName: "jingdong.shorturl.generateURLFastest", IsUnionGW: false, Params: make(map[string]interface{}, 4)}
	req = &GenerateURLFastestRequest{
		Request: &request,
	}
	return
}

func (req *GenerateURLFastestRequest) SetDomain(domain string) {
	req.Request.Params["domain"] = domain
}

func (req *GenerateURLFastestRequest) GetDomain() string {
	domain, found := req.Request.Params["domain"]
	if found {
		return domain.(string)
	}
	return ""
}

func (req *GenerateURLFastestRequest) SetLength(length uint) {
	req.Request.Params["length"] = length
}

func (req *GenerateURLFastestRequest) GetLength() uint {
	length, found := req.Request.Params["length"]
	if found {
		return length.(uint)
	}
	return 0
}

func (req *GenerateURLFastestRequest) SetRealUrl(realUrl string) {
	req.Request.Params["realUrl"] = realUrl
}

func (req *GenerateURLFastestRequest) GetRealUrl() string {
	realUrl, found := req.Request.Params["realUrl"]
	if found {
		return realUrl.(string)
	}
	return ""
}

func (req *GenerateURLFastestRequest) SetExpiredDays(expiredDays uint) {
	req.Request.Params["expiredDays"] = expiredDays
}

func (req *GenerateURLFastestRequest) GetExpiredDays() uint {
	expiredDays, found := req.Request.Params["expiredDays"]
	if found {
		return expiredDays.(uint)
	}
	return 0
}
