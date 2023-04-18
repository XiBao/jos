package asset

import (
	"encoding/json"
	"errors"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/api/util"
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

type AccountBalanceQueryRes struct {
	Code string                   `json:"code,omitempty" codec:"code,omitempty"`
	Res  *AccountBalanceQueryData `json:"response,omitempty" codec:"response,omitempty"`
}

type AccountBalanceQueryData struct {
	Code      string            `json:"code,omitempty" codec:"code,omitempty"`
	ErrorDesc string            `json:"error_description,omitempty" codec:"error_description,omitempty"`
	Data      []*AccountBalance `json:"data,omitempty" codec:"data,omitempty"`
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

func AccountBalanceQuery(req *AccountBalanceQueryRequest) ([]*AccountBalance, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := asset.NewAccountBalanceQueryRequest()
	r.SetTypeId(req.TypeId)

	result, err := client.Execute(r.Request, req.Session)
	if err != nil {
		return nil, err
	}
	result = util.RemoveJsonSpace(result)

	var response AccountBalanceQueryResponse
	err = json.Unmarshal(result, &response)
	if err != nil {
		return nil, err
	}
	if response.ErrorResp != nil {
		return nil, response.ErrorResp
	}
	if response.Response == nil || response.Response.Res == nil {
		return nil, errors.New("No response.")
	}
	if response.Response.Res.Data == nil {
		return nil, errors.New("No result.")
	}

	return response.Response.Res.Data, nil
}
