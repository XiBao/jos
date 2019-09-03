package vender

import (
	"github.com/XiBao/jos/sdk"
)

type VenderCommonQueryRequest struct {
	Request *sdk.Request
}

// create new request
func NewVenderCommonQueryRequest() (req *VenderCommonQueryRequest) {
	request := sdk.Request{MethodName: "jingdong.data.vender.common.query", Params: make(map[string]interface{}, 2)}
	req = &VenderCommonQueryRequest{
		Request: &request,
	}
	return
}

func (req *VenderCommonQueryRequest) SetMethod(method string) {
	req.Request.Params["method"] = method
}

func (req *VenderCommonQueryRequest) GetMethod() string {
	method, found := req.Request.Params["method"]
	if found {
		return method.(string)
	}
	return ""
}

func (req *VenderCommonQueryRequest) SetInputPara(inputPara string) {
	req.Request.Params["input_para"] = inputPara
}

func (req *VenderCommonQueryRequest) GetInputPara() string {
	inputPara, found := req.Request.Params["inputPara"]
	if found {
		return inputPara.(string)
	}
	return ""
}
