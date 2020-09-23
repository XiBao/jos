package jm

import "github.com/XiBao/jos/sdk"

type GetOpenIdRequest struct {
    Request *sdk.Request
}

func NewGetOpenIdRequest() (req *GetOpenIdRequest) {
    request := sdk.Request{MethodName: "jingdong.pop.jm.center.user.getOpenId", Params: make(map[string]interface{})}
    req = &GetOpenIdRequest{
        Request: &request,
    }
    return
}


func (req *GetOpenIdRequest) SetSource(source string) {
    req.Request.Params["source"] = source
}

func (req *GetOpenIdRequest) GetSource() string {
    source, found := req.Request.Params["source"]
    if found {
        return source.(string)
    }
    return ""
}

func (req *GetOpenIdRequest) SetToken(token string) {
    req.Request.Params["token"] = token
}

func (req *GetOpenIdRequest) GetToken() string {
    token, found := req.Request.Params["token"]
    if found {
        return token.(string)
    }
    return ""
}
