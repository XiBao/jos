package asset

import (
	"context"

	"github.com/XiBao/jos/api"
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

func (r BenefitSendResponse) IsError() bool {
	return r.ErrorResp != nil || r.Response == nil || r.Response.IsError()
}

func (r BenefitSendResponse) Error() string {
	if r.ErrorResp != nil {
		return r.ErrorResp.Error()
	}
	if r.Response != nil {
		return r.Response.Error()
	}
	return "no result data"
}

type BenefitSendRes struct {
	Code string           `json:"code,omitempty" codec:"code,omitempty"`
	Res  *BenefitSendData `json:"response,omitempty" codec:"response,omitempty"`
}

func (r BenefitSendRes) IsError() bool {
	return r.Res != nil || r.Res.IsError()
}

func (r BenefitSendRes) Error() string {
	if r.Res != nil {
		return r.Res.Error()
	}
	return "no result data"
}

type BenefitSendData struct {
	Code    string                        `json:"code,omitempty" codec:"code,omitempty"`
	Message string                        `json:"message,omitempty" codec:"message,omitempty"`
	Data    *BenefitSendDataConsumptionId `json:"data,omitempty" codec:"data,omitempty"`
}

func (r BenefitSendData) IsError() bool {
	return r.Code != "200"
}

func (r BenefitSendData) Error() string {
	return sdk.ErrorString(r.Code, r.Message)
}

type BenefitSendDataConsumptionId struct {
	ConsumptionId int `json:"consumption_id" codec:"consumption_id"`
}

func BenefitSend(ctx context.Context, req *BenefitSendRequest) (int, error) {
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

	var response BenefitSendResponse
	if err := client.Execute(ctx, r.Request, req.Session, &response); err != nil {
		return 0, err
	}
	return response.Response.Res.Data.ConsumptionId, nil
}
