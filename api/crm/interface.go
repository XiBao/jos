package crm

type ReturnType struct {
	Code string `json:"code,omitempty" codec:"code,omitempty"` //状态码
	Desc string `json:"desc,omitempty" codec:"desc,omitempty"` //参数描述
	Data bool   `json:"data,omitempty" codec:"data,omitempty"` //是否成功
}
