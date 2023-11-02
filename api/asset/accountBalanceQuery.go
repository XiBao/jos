package asset

import (
	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/asset"
)

type AccountBalanceQueryRequest struct {
	api.BaseRequest

	TypeId uint8 `json:"type_id" codec:"type_id"` // 1:流量包, 2:E卡, 3:PLUS会员, 4:爱奇艺会员, 8:红包, 9:短信
}

type AccountBalanceQueryResponse struct {
	ErrorResp *api.ErrorResponnse     `json:"error_response,omitempty" codec:"error_response,omitempty"`
	Response  *AccountBalanceQueryRes `json:"jingdong_asset_account_balance_query_responce,omitempty" codec:"jingdong_asset_account_balance_query_responce,omitempty"`
}

func (r AccountBalanceQueryResponse) IsError() bool {
	return r.ErrorResp != nil || r.Response == nil || r.Response.IsError()
}

func (r AccountBalanceQueryResponse) Error() string {
	if r.ErrorResp != nil {
		return r.ErrorResp.Error()
	}
	if r.Response != nil {
		return r.Response.Error()
	}
	return "no result data"
}

type AccountBalanceQueryRes struct {
	Code string                   `json:"code,omitempty" codec:"code,omitempty"`
	Res  *AccountBalanceQueryData `json:"response,omitempty" codec:"response,omitempty"`
}

func (r AccountBalanceQueryRes) IsError() bool {
	return r.Res == nil || r.Res.IsError()
}

func (r AccountBalanceQueryRes) Error() string {
	if r.Res != nil {
		return r.Res.Error()
	}
	return "no result data"
}

type AccountBalanceQueryData struct {
	Code      string           `json:"code,omitempty" codec:"code,omitempty"`
	ErrorDesc string           `json:"error_description,omitempty" codec:"error_description,omitempty"`
	Data      []AccountBalance `json:"data,omitempty" codec:"data,omitempty"`
}

func (r AccountBalanceQueryData) IsError() bool {
	return r.Code != "200"
}

func (r AccountBalanceQueryData) Error() string {
	return sdk.ErrorString(r.Code, r.ErrorDesc)
}

type AccountBalance struct {
	TypeId          uint8  `json:"type_id" codec:"type_id"`
	TypeName        string `json:"type_name" codec:"type_name"`
	ItemId          uint8  `json:"item_id" codec:"item_id"`
	ItemName        string `json:"item_name" codec:"item_name"`
	Unit            string `json:"unit" codec:"unit"`
	QuantityTotal   uint   `json:"quantity_total" codec:"quantity_total"`
	QuantityFrozen  uint   `json:"quantity_frozen" codec:"quantity_frozen"`
	QuantityBalance uint   `json:"quantity_balance" codec:"quantity_balance"`
	Signed          bool   `json:"signed" codec:"signed"`
}

func AccountBalanceQuery(req *AccountBalanceQueryRequest) ([]AccountBalance, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := asset.NewAccountBalanceQueryRequest()
	r.SetTypeId(req.TypeId)

	var response AccountBalanceQueryResponse
	if err := client.Execute(r.Request, req.Session, &response); err != nil {
		return nil, err
	}
	return response.Response.Res.Data, nil
}
