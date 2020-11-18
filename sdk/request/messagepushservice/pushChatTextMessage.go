package messagepushservice

import (
	"github.com/XiBao/jos/sdk"
)

type PushChatTextMessageRequest struct {
	Request *sdk.Request
}

// create new request
func NewPushChatTextMessageRequest() (req *PushChatTextMessageRequest) {
	request := sdk.Request{MethodName: "jingdong.MessagePushService.pushChatTextMessage", Params: make(map[string]interface{}, 12)}
	req = &PushChatTextMessageRequest{
		Request: &request,
	}
	return
}

func (req *PushChatTextMessageRequest) SetAccessToken(accessToken string) {
	req.Request.Params["accessToken"] = accessToken
}

func (req *PushChatTextMessageRequest) GetAccessToken() string {
	accessToken, found := req.Request.Params["accessToken"]
	if found {
		return accessToken.(string)
	}
	return ""
}

func (req *PushChatTextMessageRequest) SetAspId(aspId string) {
	req.Request.Params["aspid"] = aspId
}

func (req *PushChatTextMessageRequest) GetAspId() string {
	aspId, found := req.Request.Params["aspid"]
	if found {
		return aspId.(string)
	}
	return ""
}

func (req *PushChatTextMessageRequest) SetAccessId(accessId string) {
	req.Request.Params["accessid"] = accessId
}

func (req *PushChatTextMessageRequest) GetAccessId() string {
	accessId, found := req.Request.Params["accessid"]
	if found {
		return accessId.(string)
	}
	return ""
}

func (req *PushChatTextMessageRequest) SetFromPin(fromPin string) {
	req.Request.Params["fromPin"] = fromPin
}

func (req *PushChatTextMessageRequest) GetFromPin() string {
	fromPin, found := req.Request.Params["fromPin"]
	if found {
		return fromPin.(string)
	}
	return ""
}

func (req *PushChatTextMessageRequest) SetFromApp(fromApp string) {
	req.Request.Params["fromApp"] = fromApp
}

func (req *PushChatTextMessageRequest) GetFromApp() string {
	fromApp, found := req.Request.Params["fromApp"]
	if found {
		return fromApp.(string)
	}
	return ""
}

func (req *PushChatTextMessageRequest) SetFromClientType(fromClientType string) {
	req.Request.Params["fromClientType"] = fromClientType
}

func (req *PushChatTextMessageRequest) GetFromClientType() string {
	fromClientType, found := req.Request.Params["fromClientType"]
	if found {
		return fromClientType.(string)
	}
	return ""
}

func (req *PushChatTextMessageRequest) SetOpenIdSeller(openIdSeller string) {
	req.Request.Params["open_id_seller"] = openIdSeller
}

func (req *PushChatTextMessageRequest) GetOpenIdSeller() string {
	openIdSeller, found := req.Request.Params["open_id_seller"]
	if found {
		return openIdSeller.(string)
	}
	return ""
}

func (req *PushChatTextMessageRequest) SetToPin(toPin string) {
	req.Request.Params["toPin"] = toPin
}

func (req *PushChatTextMessageRequest) GetToPin() string {
	toPin, found := req.Request.Params["toPin"]
	if found {
		return toPin.(string)
	}
	return ""
}

func (req *PushChatTextMessageRequest) SetToApp(toApp string) {
	req.Request.Params["toApp"] = toApp
}

func (req *PushChatTextMessageRequest) GetToApp() string {
	toApp, found := req.Request.Params["toApp"]
	if found {
		return toApp.(string)
	}
	return ""
}

func (req *PushChatTextMessageRequest) SetToClientType(toClientType string) {
	req.Request.Params["toClientType"] = toClientType
}

func (req *PushChatTextMessageRequest) GetToClientType() string {
	toClientType, found := req.Request.Params["toClientType"]
	if found {
		return toClientType.(string)
	}
	return ""
}

func (req *PushChatTextMessageRequest) SetOpenIdBuyer(openIdBuyer string) {
	req.Request.Params["open_id_buyer"] = openIdBuyer
}

func (req *PushChatTextMessageRequest) GetOpenIdBuyer() string {
	openIdBuyer, found := req.Request.Params["open_id_buyer"]
	if found {
		return openIdBuyer.(string)
	}
	return ""
}

func (req *PushChatTextMessageRequest) SetContent(content string) {
	req.Request.Params["content"] = content
}

func (req *PushChatTextMessageRequest) GetContent() string {
	content, found := req.Request.Params["content"]
	if found {
		return content.(string)
	}
	return ""
}
