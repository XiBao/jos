package vender

import (
	"context"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/vender"
)

type GetVenderStatusRequest struct {
	api.BaseRequest
}

type GetVenderStatusResponse struct {
	ErrorResp *api.ErrorResponnse  `json:"error_response,omitempty" codec:"error_response,omitempty"`
	Data      *GetVenderStatusData `json:"jingdong_pop_vender_getVenderStatus_responce,omitempty" codec:"jingdong_pop_vender_getVenderStatus_responce,omitempty"`
}

func (r GetVenderStatusResponse) IsError() bool {
	return r.ErrorResp != nil || r.Data == nil || r.Data.IsError()
}

func (r GetVenderStatusResponse) Error() string {
	if r.ErrorResp != nil {
		return r.ErrorResp.Error()
	}
	if r.Data != nil {
		return r.Data.Error()
	}
	return "no result data"
}

type GetVenderStatusData struct {
	ReturnType *VenderStatusReturnType `json:"returnType,omitempty" codec:"returnType,omitempty"`
}

func (r GetVenderStatusData) IsError() bool {
	return r.ReturnType == nil || r.ReturnType.IsError()
}

func (r GetVenderStatusData) Error() string {
	if r.ReturnType != nil {
		return r.ReturnType.Error()
	}
	return "no result data"
}

type VenderStatusReturnType struct {
	Status uint   `json:"status"` // 会员体系状态：0:未开启状态；1:ISV计算；2：官方计算
	Code   string `json:"code"`   // 200：成功，201：信息不存在，400：参数错误，500：系统错误
	Desc   string `json:"desc"`   // 成功，信息不存在，参数错误，服务端异常
}

func (r VenderStatusReturnType) IsError() bool {
	return r.Code != "200"
}

func (r VenderStatusReturnType) Error() string {
	return sdk.ErrorString(r.Code, r.Desc)
}

// TODO 查询会员体系状态
func GetVenderStatus(ctx context.Context, req *GetVenderStatusRequest) (uint, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := vender.NewGetVenderStatusRequest()

	var response GetVenderStatusResponse
	if err := client.Execute(ctx, r.Request, req.Session, &response); err != nil {
		return 0, err
	}
	return response.Data.ReturnType.Status, nil
}
