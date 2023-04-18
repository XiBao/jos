package asset

import (
	"encoding/json"
	"errors"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/api/util"
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/asset"
)

type BenefitSendRequest struct {
	api.BaseRequest

	TypeId      int    `json:"type_id" codec:"type_id"`             // 1:流量, 2:E卡,3:plus,4:爱奇艺,8:红包,10:京豆
	ItemId      int    `json:"item_id" codec:"item_id"`             // 资产项id
	Quantity    int    `json:"quantity" codec:"quantity"`           // 发放数量
	UserPin     string `json:"user_pin" codec:"user_pin"`           // 用户pin
	Token       string `json:"token" codec:"token"`                 // 创建活动计划返回的token
	RequestId   string `json:"request_id" codec:"request_id"`       // 请求唯一标识，防重，建议使用uuid 最长36位
	Remark      string `json:"remark" codec:"remark"`               // 发放备注
	Ip          string `json:"ip" codec:"ip"`                       // 发放用户ip
	OpenIdBuyer string `json:"open_id_buyer" codec:"open_id_buyer"` // 用户pin
}

type BenefitSendResponse struct {
	ErrorResp *api.ErrorResponnse `json:"error_response,omitempty" codec:"error_response,omitempty"`
	Response  *BenefitSendRes     `json:"jingdong_asset_benefit_send_responce,omitempty" codec:"jingdong_asset_benefit_send_responce,omitempty"`
}

type BenefitSendRes struct {
	Code string           `json:"code,omitempty" codec:"code,omitempty"`
	Res  *BenefitSendData `json:"response,omitempty" codec:"response,omitempty"`
}

type BenefitSendData struct {
	Code    string                        `json:"code,omitempty" codec:"code,omitempty"`
	Message string                        `json:"message,omitempty" codec:"message,omitempty"`
	Data    *BenefitSendDataConsumptionId `json:"data,omitempty" codec:"data,omitempty"`
}

type BenefitSendDataConsumptionId struct {
	ConsumptionId int `json:"consumption_id" codec:"consumption_id"`
}

func BenefitSend(req *BenefitSendRequest) (int, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := asset.NewBenefitSendRequest()
	r.SetTypeId(req.TypeId)
	r.SetItemId(req.ItemId)
	r.SetQuantity(req.Quantity)
	r.SetUserPin(req.UserPin)
	r.SetToken(req.Token)
	r.SetRequestId(req.RequestId)
	r.SetRemark(req.Remark)
	r.SetIp(req.Ip)
	r.SetOpenIdBuyer(req.OpenIdBuyer)

	result, err := client.Execute(r.Request, req.Session)
	if err != nil {
		return 0, err
	}
	result = util.RemoveJsonSpace(result)

	var response BenefitSendResponse
	err = json.Unmarshal(result, &response)
	if err != nil {
		return 0, err
	}
	if response.ErrorResp != nil {
		return 0, response.ErrorResp
	}
	if response.Response == nil || response.Response.Res == nil {
		return 0, errors.New("No response.")
	}
	if response.Response.Res.Code != "200" && response.Response.Res.Message != "" {
		return 0, errors.New(response.Response.Res.Message)
	}
	if response.Response.Res.Data == nil {
		return 0, errors.New("No result.")
	}

	return response.Response.Res.Data.ConsumptionId, nil
}
