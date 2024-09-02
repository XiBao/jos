package jzone

import (
	"context"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/jzone"
)

type AddCartItemByPinRequest struct {
	api.BaseRequest
	Pin    string `json:"pin,omitempty" codec:"pin,omitempty"`       // 加密pin
	ItemId uint64 `json:"itemId,omitempty" codec:"itemId,omitempty"` // skuid
	Num    uint64 `json:"num,omitempty" codec:"num,omitempty"`       // 数量
}

type AddCartItemByPinResponse struct {
	ErrorResp *api.ErrorResponnse   `json:"error_response,omitempty" codec:"error_response,omitempty"`
	Data      *AddCartItemByPinData `json:"jingdong_jzone_addCartItemByPin_responce,omitempty" codec:"jingdong_jzone_addCartItemByPin_responce,omitempty"`
}

func (r AddCartItemByPinResponse) IsError() bool {
	return r.ErrorResp != nil || r.Data == nil || r.Data.IsError()
}

func (r AddCartItemByPinResponse) Error() string {
	if r.ErrorResp != nil {
		return r.ErrorResp.Error()
	}
	if r.Data != nil {
		return r.Data.Error()
	}
	return "no result data"
}

type AddCartItemByPinData struct {
	ReturnType *AddCartItemByPinReturnType `json:"returnType,omitempty" codec:"returnType,omitempty"`
	Code       string                      `json:"code"`
}

func (r AddCartItemByPinData) IsError() bool {
	return r.ReturnType == nil || r.ReturnType.IsError()
}

func (r AddCartItemByPinData) Error() string {
	if r.ReturnType != nil {
		return r.ReturnType.Error()
	}
	return "no result data"
}

type AddCartItemByPinReturnType struct {
	Message string `json:"message,omitempty" codec:"message,omitempty"`
	Code    string `json:"code,omitempty" codec:"code,omitempty"`
}

func (r AddCartItemByPinReturnType) IsError() bool {
	return r.Code != "0"
}

func (r AddCartItemByPinReturnType) Error() string {
	return sdk.ErrorString(r.Code, r.Message)
}

// TODO  通过Pin将商品加入用户购物车
func AddCartItemByPin(ctx context.Context, req *AddCartItemByPinRequest) (bool, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := jzone.NewAddCartItemByPinRequest()
	r.SetPin(req.Pin)
	r.SetItemId(req.ItemId)
	if req.Num > 0 {
		r.SetNum(req.Num)
	}
	r.SetNum(1)

	var response AddCartItemByPinResponse
	if err := client.Execute(ctx, r.Request, req.Session, &response); err != nil {
		return false, err
	}
	return true, nil
}
