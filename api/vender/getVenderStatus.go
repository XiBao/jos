package vender

import (
	"errors"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/vender"
	"github.com/daviddengcn/ljson"
)

type GetVenderStatusRequest struct {
	api.BaseRequest
}

type GetVenderStatusResponse struct {
	ErrorResp *api.ErrorResponnse  `json:"error_response,omitempty" codec:"error_response,omitempty"`
	Data      *GetVenderStatusData `json:"jingdong_pop_vender_getVenderStatus_responce,omitempty" codec:"jingdong_pop_vender_getVenderStatus_responce,omitempty"`
}

type GetVenderStatusData struct {
	ReturnType *VenderStatusReturnType `json:"returnType,omitempty" codec:"returnType,omitempty"`
}

type VenderStatusReturnType struct {
	Status uint   `json:"status"` //会员体系状态：0:未开启状态；1:ISV计算；2：官方计算
	Code   string `json:"code"`   //200：成功，201：信息不存在，400：参数错误，500：系统错误
	Desc   string `json:"desc"`   //成功，信息不存在，参数错误，服务端异常
}

//TODO 查询会员体系状态
func GetVenderStatus(req *GetVenderStatusRequest) (uint, error) {

	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := vender.NewGetVenderStatusRequest()

	result, err := client.Execute(r.Request, req.Session)
	if err != nil {
		return 0, err
	}
	if len(result) == 0 {
		return 0, errors.New("no result info")
	}
	var response GetVenderStatusResponse
	err = ljson.Unmarshal(result, &response)
	if err != nil {
		return 0, err
	}

	if response.ErrorResp != nil {
		return 0, response.ErrorResp
	}

	if response.Data.ReturnType.Code != "200" {
		return 0, errors.New(response.Data.ReturnType.Desc)
	}
	return response.Data.ReturnType.Status, nil
}
