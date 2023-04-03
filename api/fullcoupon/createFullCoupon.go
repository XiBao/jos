package fullcoupon

import (
	"errors"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/fullcoupon"
	"github.com/daviddengcn/ljson"
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

type CreateFullCouponResponseData struct {
	ReturnType *CreateFullCouponResponseReturnType `json:"returnType,omitempty" codec:"returnType,omitempty"`
}

type CreateFullCouponResponseReturnType struct {
	Msg     string `json:"msg,omitempty" codec:"msg,omitempty"`         // 状态码描述
	Code    string `json:"code,omitempty" codec:"code,omitempty"`       // 状态码
	Success bool   `json:"success,omitempty" codec:"success,omitempty"` // 请求是否成功
	Data    uint64 `json:"data,omitempty" codec:"data,omitempty"`       // 促销ID
}

func CreateFullCoupon(req *CreateFullCouponRequest) (*CreateFullCouponResponseReturnType, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := fullcoupon.NewCreateFullCouponRequest()
	r.SetParam(req.Param)

	result, err := client.Execute(r.Request, req.Session)
	if err != nil {
		return nil, err
	}
	if len(result) == 0 {
		return nil, errors.New("no result.")
	}

	var response CreateFullCouponResponse
	err = ljson.Unmarshal(result, &response)
	if err != nil {
		return nil, err
	}
	if response.ErrorResp != nil {
		return nil, response.ErrorResp
	}

	if response.Data == nil || response.Data.ReturnType == nil {
		return nil, errors.New("no result.")
	}

	if !response.Data.ReturnType.Success {
		return nil, errors.New(response.Data.ReturnType.Msg)
	}

	return response.Data.ReturnType, nil
}
