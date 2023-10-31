package vender

import (
	"fmt"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/vender"
)

type GetCustomerBasicInfoRequest struct {
	api.BaseRequest
	Pin string `json:"pin"`
}

type GetCustomerBasicInfoResponse struct {
	ErrorResp *api.ErrorResponnse       `json:"error_response,omitempty" codec:"error_response,omitempty"`
	Data      *GetCustomerBasicInfoData `json:"jingdong_pop_vender_getCustomerBasicInfo_responce,omitempty" codec:"jingdong_pop_vender_getCustomerBasicInfo_responce,omitempty"`
}

func (r GetCustomerBasicInfoResponse) IsError() bool {
	return r.ErrorResp != nil || r.Data == nil || r.Data.IsError()
}

func (r GetCustomerBasicInfoResponse) Error() string {
	if r.ErrorResp != nil {
		return r.ErrorResp.Error()
	}
	if r.Data != nil {
		return r.Data.Error()
	}
	return "no result data"
}

type GetCustomerBasicInfoData struct {
	Result *GetCustomerResult `json:"returnResult,omitempty" codec:"returnResult,omitempty"`
}

func (r GetCustomerBasicInfoData) IsError() bool {
	return r.Result == nil || r.Result.IsError()
}

func (r GetCustomerBasicInfoData) Error() string {
	if r.Result != nil {
		return r.Result.Error()
	}
	return "no result data"
}

type GetCustomerResult struct {
	Desc string          `json:"desc,omitempty" codec:"desc,omitempty"`
	Code string          `json:"code,omitempty" codec:"code,omitempty"`
	Info *CardMemberInfo `json:"data,omitempty" codec:"data,omitempty"`
}

func (r GetCustomerResult) IsError() bool {
	return r.Code != "200"
}

func (r GetCustomerResult) Error() string {
	return fmt.Sprintf("code:%s, msg:%s", r.Code, r.Desc)
}

type CardMemberInfo struct {
	Birthday      string `json:"birthday,omitempty" codec:"birthday,omitempty"`           // 出生日期（yyyy-MM-dd）
	CustomerType  uint8  `json:"customerType,omitempty" codec:"customerType,omitempty"`   // 1	会员类型（1-注册已完成；2-绑定；3-注册未完成；4-待激活；5-审核中；6-待购卡)
	Status        uint8  `json:"status,omitempty" codec:"status,omitempty"`               // 1	状态（0-无效；1-有效；2-解绑中；3-已解绑）
	Extend        string `json:"extend,omitempty" codec:"extend,omitempty"`               // 1	状态（0-无效；1-有效；2-解绑中；3-已解绑）
	CustomerLevel uint8  `json:"customerLevel,omitempty" codec:"customerLevel,omitempty"` // 1	会员等级（1-一等级；2-二等级；3-三等级；4-四等级；5-五等级）
	VenderId      uint64 `json:"venderId,omitempty" codec:"venderId,omitempty"`           // 商家ID
	City          string `json:"city,omitempty" codec:"city,omitempty"`                   // 城市
	CardNo        string `json:"cardNo,omitempty" codec:"cardNo,omitempty"`               // 会员卡号
	PhoneNo       string `json:"phoneNo,omitempty" codec:"phoneNo,omitempty"`             // 手机号
	Province      string `json:"province,omitempty" codec:"province,omitempty"`           // 省份
	Street        string `json:"street,omitempty" codec:"street,omitempty"`               // 街道
	CustomerPin   string `json:"customerPin,omitempty" codec:"customerPin,omitempty"`     // 京东用户PIN
	Gender        string `json:"gender,omitempty" codec:"gender,omitempty"`               // 性别（0-女; 1-男; 3-未知）
	Channel       uint   `json:"channel,omitempty" codec:"channel,omitempty"`             // 601	渠道码（101-卡包；102-店铺首页；103-app支付完成页；。。。601-ISV服务；999-默认渠道；888-CRM-SHOP）
}

func GetCustomerBasicInfo(req *GetCustomerBasicInfoRequest) (*CardMemberInfo, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := vender.NewGetCustomerBasicInfoRequest()

	if len(req.Pin) > 0 {
		r.SetCustomerPin(req.Pin)
	}

	var response GetCustomerBasicInfoResponse
	if err := client.Execute(r.Request, req.Session, &response); err != nil {
		return nil, err
	}
	return response.Data.Result.Info, nil

}
