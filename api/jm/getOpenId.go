package jm

import (
    "errors"
    . "github.com/XiBao/jos/api"
    "github.com/XiBao/jos/sdk"
    "github.com/XiBao/jos/sdk/request/jm"
    "github.com/daviddengcn/ljson"
)

type GetOpenIdRequest struct {
    BaseRequest
    Source string `json:"source" codec:"source"` //  01:京东App，02：微信
    Token  string `json:"token" codec:"token"`
}

type GetOpenIdResponse struct {
    ErrorResp *ErrorResponnse                      `json:"error_response,omitempty" codec:"error_response,omitempty"`
    Data      GetOpenIdResponseData `json:"jingdong_pop_jm_center_user_getOpenId_responce"`
}

type GetOpenIdReturnType struct {
    Message   string `json:"message"`
    OpenID    string `json:"open_id"`
    Pin       string `json:"pin"`
    RequestID string `json:"requestId"`
    Code      int    `json:"code"`
}

type GetOpenIdResponseData struct {
    ReturnType GetOpenIdReturnType `json:"returnType"`
}

func GetOpenId(req GetOpenIdRequest) (GetOpenIdReturnType, error) {
    client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
    client.Debug = req.Debug
    r := jm.NewGetOpenIdRequest()
    r.SetSource(req.Source)
    r.SetToken(req.Token)
    result, err := client.Execute(r.Request, req.Session)
    if err != nil {
        return GetOpenIdReturnType{}, err
    }
    if len(result) == 0 {
        return GetOpenIdReturnType{}, errors.New("no result.")
    }
    var response GetOpenIdResponse
    err = ljson.Unmarshal(result, &response)
    if err != nil {
        return GetOpenIdReturnType{}, err
    }
    if response.ErrorResp != nil {
        return GetOpenIdReturnType{}, response.ErrorResp
    }

    if response.Data.ReturnType.Code != 0 && response.Data.ReturnType.Code != 200 {
        return GetOpenIdReturnType{}, errors.New(response.Data.ReturnType.Message)
    }

    return response.Data.ReturnType, nil

}
