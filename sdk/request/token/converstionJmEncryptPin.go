package token

import (
	"github.com/XiBao/jos/sdk"
)

type ConverstionJmEncryptPinRequest struct {
	Request *sdk.Request
}

// create new request
func NewConverstionJmEncryptPinRequest() (req *ConverstionJmEncryptPinRequest) {
	request := sdk.Request{MethodName: "jingdong.TokenToPinCenter.converstionJmEncryptPin", Params: make(map[string]interface{}, 2)}
	req = &ConverstionJmEncryptPinRequest{
		Request: &request,
	}
	return
}

func (req *ConverstionJmEncryptPinRequest) SetEncryptPin(encryptPin string) {
	req.Request.Params["encryptPin"] = encryptPin
}

func (req *ConverstionJmEncryptPinRequest) GetEncryptPin() string {
	encryptPin, found := req.Request.Params["encryptPin"]
	if found {
		return encryptPin.(string)
	}
	return ""
}

func (req *ConverstionJmEncryptPinRequest) SetAppKey(appKey string) {
	req.Request.Params["appKey"] = appKey
}

func (req *ConverstionJmEncryptPinRequest) GetAppKey() string {
	appKey, found := req.Request.Params["appKey"]
	if found {
		return appKey.(string)
	}
	return ""
}
