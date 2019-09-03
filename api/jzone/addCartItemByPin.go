package jzone

import (
	"errors"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/jzone"
	"github.com/daviddengcn/ljson"
)

type AddCartItemByPinRequest struct {
	api.BaseRequest
	Pin    string `json:"pin,omitempty" codec:"pin,omitempty"`       //加密pin
	ItemId uint64 `json:"itemId,omitempty" codec:"itemId,omitempty"` //skuid
	Num    uint64 `json:"num,omitempty" codec:"num,omitempty"`       //数量
}

type AddCartItemByPinResponse struct {
	ErrorResp *api.ErrorResponnse   `json:"error_response,omitempty" codec:"error_response,omitempty"`
	Data      *AddCartItemByPinData `json:"jingdong_jzone_addCartItemByPin_responce,omitempty" codec:"jingdong_jzone_addCartItemByPin_responce,omitempty"`
}

type AddCartItemByPinData struct {
	ReturnType *AddCartItemByPinReturnType `json:"returnType,omitempty" codec:"returnType,omitempty"`
	Code       string                      `json:"code"`
}

type AddCartItemByPinReturnType struct {
	Message string `json:"message,omitempty" codec:"message,omitempty"`
	Code    string `json:"code,omitempty" codec:"code,omitempty"`
}

//TODO  通过Pin将商品加入用户购物车
func AddCartItemByPin(req *AddCartItemByPinRequest) (bool, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := jzone.NewAddCartItemByPinRequest()
	r.SetPin(req.Pin)
	r.SetItemId(req.ItemId)
	if req.Num > 0 {
		r.SetNum(req.Num)
	}
	r.SetNum(1)

	result, err := client.Execute(r.Request, req.Session)
	if err != nil {
		return false, err
	}
	if len(result) == 0 {
		return false, errors.New("No result info.")
	}
	var response AddCartItemByPinResponse
	err = ljson.Unmarshal(result, &response)
	if err != nil {
		return false, err
	}

	if response.ErrorResp != nil {
		return false, response.ErrorResp
	}

	if response.Data.ReturnType.Code != "0" {
		return false, errors.New(response.Data.ReturnType.Message)
	}

	return true, nil
}
