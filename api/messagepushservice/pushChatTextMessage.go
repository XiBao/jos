package messagepushservice

import (
	"errors"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/messagepushservice"
	"github.com/daviddengcn/ljson"
)

type PushChatTextMessageRequest struct {
	api.BaseRequest
	AccessToken    string `json:"accessToken,omitempty" codec:"accessToken,omitempty"`       // 咚咚服务器动态分配的访问Token
	AspId          string `json:"aspid,omitempty" codec:"aspid,omitempty"`                   // 咚咚注册的应用服务提供商ID
	AccessId       string `json:"accessid,omitempty" codec:"accessid,omitempty"`             // 访问ID，可用于请求去重，在3分钟之内同一accessid的请求咚咚视为相同请求，只会处理第一个成功接收到的,通过uuid.toString()生成
	FromPin        string `json:"fromPin,omitempty" codec:"fromPin,omitempty"`               // 发送方标识
	FromApp        string `json:"fromApp,omitempty" codec:"fromApp,omitempty"`               // app标识
	FromClientType string `json:"fromClientType,omitempty" codec:"fromClientType,omitempty"` // 终端类型可填写成gw
	OpenIdSeller   string `json:"open_id_seller,omitempty" codec:"open_id_seller,omitempty"` // 发送方标识
	ToPin          string `json:"toPin,omitempty" codec:"toPin,omitempty"`                   // 接收方标识
	ToApp          string `json:"toApp,omitempty" codec:"toApp,omitempty"`                   // app标识
	ToClientType   string `json:"toClientType,omitempty" codec:"toClientType,omitempty"`     // 终端类型可不填
	OpenIdBuyer    string `json:"open_id_buyer,omitempty" codec:"open_id_buyer,omitempty"`   // 接收方标识
	Content        string `json:"content,omitempty" codec:"content,omitempty"`               // 文本内容
}

type PushChatTextMessageResponse struct {
	ErrorResp  *api.ErrorResponnse            `json:"error_response,omitempty" codec:"error_response,omitempty"`
	ReturnType *PushChatTextMessageReturnType `json:"jingdong_MessagePushService_pushChatTextMessage_responce,omitempty" codec:"jingdong_MessagePushService_pushChatTextMessage_responce,omitempty"`
}

type PushChatTextMessageReturnType struct {
	Code          string `json:"code,omitempty" codec:"code,omitempty"`
	ErrMsg        string `json:"errmsg,omitempty" codec:"errmsg,omitempty"`
	MsgId         string `json:"msgid,omitempty" codec:"msgid,omitempty"`
	AccId         string `json:"accid,omitempty" codec:"accid,omitempty"`
	SendTimestamp string `json:"sendTimestamp,omitempty" codec:"sendTimestamp,omitempty"`
}

// 新提供发送咚咚消息接口，方便打标pin
func PushChatTextMessage(req *PushChatTextMessageRequest) (*PushChatTextMessageReturnType, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := messagepushservice.NewPushChatTextMessageRequest()
	r.SetAccessToken(req.AccessToken)
	r.SetAspId(req.AspId)
	r.SetAccessId(req.AccessId)
	r.SetFromPin(req.FromPin)
	r.SetFromApp(req.FromApp)
	r.SetOpenIdSeller(req.OpenIdSeller)
	r.SetToPin(req.ToPin)
	r.SetToApp(req.ToApp)
	r.SetOpenIdBuyer(req.OpenIdBuyer)
	r.SetContent(req.Content)
	if req.FromClientType != "" {
		r.SetFromClientType(req.FromClientType)
	}
	if req.ToClientType != "" {
		r.SetToClientType(req.ToClientType)
	}

	result, err := client.Execute(r.Request, req.Session)
	if err != nil {
		return nil, err
	}
	if len(result) == 0 {
		return nil, errors.New("No return type.")
	}
	var response PushChatTextMessageResponse
	err = ljson.Unmarshal(result, &response)
	if err != nil {
		return nil, err
	}

	if response.ErrorResp != nil {
		return nil, response.ErrorResp
	}

	return response.ReturnType, nil
}
