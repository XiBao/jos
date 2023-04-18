package vender

import (
	"encoding/json"
	"errors"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/vender"
)

type SyncBindRequest struct {
	api.BaseRequest
	Birthday      string `json:"birthday,omitempty" codec:"birthday,omitempty"`           // 出生日期（yyyy-MM-dd）
	Gender        string `json:"gender,omitempty" codec:"gender,omitempty"`               // 性别（0-女; 1-男; 3-未知）
	City          string `json:"city,omitempty" codec:"city,omitempty"`                   // 城市
	Channel       uint16 `json:"channel,omitempty" codec:"channel,omitempty"`             // 渠道码（101-卡包；102-店铺首页；103-app支付完成页；601-ISV服务；999-默认渠道；888-CRM-SHOP）
	CardNo        string `json:"cardNo,omitempty" codec:"cardNo,omitempty"`               // 会员卡号
	PhoneNo       string `json:"phoneNo,omitempty" codec:"phoneNo,omitempty"`             // 手机号
	CustomerLevel uint8  `json:"customerLevel,omitempty" codec:"customerLevel,omitempty"` // 会员等级（1-一等级；2-二等级；3-三等级；4-四等级；5-五等级）
	Extend        string `json:"extend,omitempty" codec:"extend,omitempty"`               // 扩展字段，Json数据；存储姓名、称谓、邮箱、密码等非必须字段
	CustomerType  uint8  `json:"customerType,omitempty" codec:"customerType,omitempty"`   // 会员类型（1-注册已完成；2-绑定；3-注册未完成；4-待激活；5-审核中；6-待购卡)
	Province      string `json:"province,omitempty" codec:"province,omitempty"`           // 省份
	Street        string `json:"street,omitempty" codec:"street,omitempty"`               // 街道
	CustomerPin   string `json:"customerPin,omitempty" codec:"customerPin,omitempty"`     // 京东用户PIN
	Status        uint8  `json:"status,omitempty" codec:"status,omitempty"`               // 状态（0-无效；1-有效；2-解绑中；3-已解绑）
}

type SyncBindResponse struct {
	ErrorResp *api.ErrorResponnse `json:"error_response,omitempty" codec:"error_response,omitempty"`
	Data      *SyncBindData       `json:"jingdong_pop_vender_syncBind_responce,omitempty" codec:"jingdong_pop_vender_syncBind_responce,omitempty"`
}

type SyncBindData struct {
	Code      string      `json:"code,omitempty" codec:"code,omitempty"`
	ErrorDesc string      `json:"error_description,omitempty" codec:"error_description,omitempty"`
	Result    *SyncResult `json:"returnResult,omitempty" codec:"returnResult,omitempty"`
}

type SyncResult struct {
	Desc string `json:"desc,omitempty" codec:"desc,omitempty"`
	Data bool   `json:"data,omitempty" codec:"data,omitempty"`
	Code string `json:"code,omitempty" codec:"code,omitempty"`
}

func SyncBind(req *SyncBindRequest) (bool, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := vender.NewVenderSyncBindRequest()
	r.SetPhoneNo(req.PhoneNo)
	r.SetCustomerPin(req.CustomerPin)
	r.SetStatus(req.Status)
	r.SetCustomerLevel(req.CustomerLevel)
	r.SetCustomerType(req.CustomerType)
	if req.Channel > 0 {
		r.SetChannel(req.Channel)
	} else {
		r.SetChannel(601)
	}
	if req.Birthday != "" {
		r.SetBirthday(req.Birthday)
	}
	if req.Gender != "" {
		r.SetGender(req.Gender)
	} else {
		r.SetGender("3")
	}
	if req.City != "" {
		r.SetCity(req.City)
	}
	if req.CardNo != "" {
		r.SetCardNo(req.CardNo)
	}
	if req.Extend != "" {
		r.SetExtend(req.Extend)
	}
	if req.Province != "" {
		r.SetProvince(req.Province)
	}
	if req.Street != "" {
		r.SetStreet(req.Street)
	}
	result, err := client.Execute(r.Request, req.Session)
	if err != nil {
		return false, err
	}
	if len(result) == 0 {
		return false, errors.New("no result.")
	}
	var response SyncBindResponse
	err = json.Unmarshal(result, &response)
	if err != nil {
		return false, err
	}
	if response.ErrorResp != nil {
		return false, response.ErrorResp
	}
	if response.Data.Code != "0" {
		return false, errors.New(response.Data.ErrorDesc)
	}

	if response.Data.Result.Code != "200" && response.Data.Result.Code != "0" {
		return false, errors.New(response.Data.Result.Desc)
	}

	return response.Data.Result.Data, nil
}
