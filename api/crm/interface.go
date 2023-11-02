package crm

import "github.com/XiBao/jos/sdk"

type ReturnType struct {
	Code string `json:"code,omitempty" codec:"code,omitempty"` //状态码
	Desc string `json:"desc,omitempty" codec:"desc,omitempty"` //参数描述
	Data bool   `json:"data,omitempty" codec:"data,omitempty"` //是否成功
}

func (r ReturnType) IsError() bool {
	return r.Code != "200"
}

func (r ReturnType) Error() string {
	return sdk.ErrorString(r.Code, r.Desc)
}
