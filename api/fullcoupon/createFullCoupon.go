package fullcoupon

import (
	"fmt"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/fullcoupon"
)

// 创建满额返券
type CreateFullCouponRequest struct {
	api.BaseRequest
	Param *fullcoupon.CreateFullCouponParam `json:"param,omitempty" codec:"param,omitempty"` // 创建入参
}

type CreateFullCouponResponse struct {
	ErrorResp *api.ErrorResponnse           `json:"error_response,omitempty" codec:"error_response,omitempty"`
	Data      *CreateFullCouponResponseData `json:"jingdong_fullCoupon_createFullCoupon_responce,omitempty" codec:"jingdong_fullCoupon_createFullCoupon_responce,omitempty"`
}

func (r CreateFullCouponResponse) IsError() bool {
	return r.ErrorResp != nil || r.Data == nil || r.Data.IsError()
}

func (r CreateFullCouponResponse) Error() string {
	if r.ErrorResp != nil {
		return r.ErrorResp.Error()
	}
	if r.Data != nil {
		return r.Data.Error()
	}
	return "no result data"
}

type CreateFullCouponResponseData struct {
	ReturnType *CreateFullCouponResponseReturnType `json:"returnType,omitempty" codec:"returnType,omitempty"`
}

func (r CreateFullCouponResponseData) IsError() bool {
	return r.ReturnType == nil || r.ReturnType.IsError()
}

func (r CreateFullCouponResponseData) Error() string {
	if r.ReturnType != nil {
		return r.ReturnType.Error()
	}
	return "no result data"
}

type CreateFullCouponResponseReturnType struct {
	Msg     string `json:"msg,omitempty" codec:"msg,omitempty"`         // 状态码描述
	Code    string `json:"code,omitempty" codec:"code,omitempty"`       // 状态码
	Success bool   `json:"success,omitempty" codec:"success,omitempty"` // 请求是否成功
	Data    uint64 `json:"data,omitempty" codec:"data,omitempty"`       // 促销ID
}

func (r CreateFullCouponResponseReturnType) IsError() bool {
	return !r.Success
}

func (r CreateFullCouponResponseReturnType) Error() string {
	return fmt.Sprintf("code:%s, msg:%s", r.Code, r.Msg)
}

func CreateFullCoupon(req *CreateFullCouponRequest) (*CreateFullCouponResponseReturnType, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := fullcoupon.NewCreateFullCouponRequest()
	r.SetParam(req.Param)

	var response CreateFullCouponResponse
	if err := client.Execute(r.Request, req.Session, &response); err != nil {
		return nil, err
	}
	return response.Data.ReturnType, nil
}
