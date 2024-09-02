package market

import (
	"context"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/market"
)

type GetUserCartSkuRequest struct {
	api.BaseRequest
	Pin         string `json:"pin" codec:"pin"`                     // 用户PIN
	OpenIdBuyer string `json:"open_id_buyer" codec:"open_id_buyer"` // 用户PIN
	XidBuyer    string `json:"xid_buyer" codec:"xid_buyer"`         // 用户PIN
}

type GetUserCartSkuResponse struct {
	ErrorResp *api.ErrorResponnse `json:"error_response,omitempty" codec:"error_response,omitempty"`
	Data      *GetUserCartSkuData `json:"jingdong_market_bdp_userBehavior_getUserCartSku_responce,omitempty" codec:"jingdong_market_bdp_userBehavior_getUserCartSku_responce,omitempty"`
}

func (r GetUserCartSkuResponse) IsError() bool {
	return r.ErrorResp != nil || r.Data == nil
}

func (r GetUserCartSkuResponse) Error() string {
	if r.ErrorResp != nil {
		return r.ErrorResp.Error()
	}
	return "no result data"
}

type GetUserCartSkuData struct {
	Code     string    `json:"code,omitempty" codec:"code,omitempty"`
	CartSkus []CartSku `json:"returnType,omitempty" codec:"returnType,omitempty"`
}

// 获取用户加购的sku数据
func GetUserCartSku(ctx context.Context, req *GetUserCartSkuRequest) ([]CartSku, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := market.NewGetUserCartSkuRequest()
	r.SetPin(req.Pin)
	r.SetOpenIdBuyer(req.OpenIdBuyer)
	r.SetXidBuyer(req.XidBuyer)

	var response GetUserCartSkuResponse
	if err := client.Execute(ctx, r.Request, req.Session, &response); err != nil {
		return nil, err
	}

	return response.Data.CartSkus, nil
}
