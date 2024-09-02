package promotion

import (
	"context"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/sdk"
	promo "github.com/XiBao/jos/sdk/request/seller/promotion"
)

type UnitCreateRequest struct {
	api.BaseRequest
	Client     *promo.PromotionUnitClient `json:"client" codec:"client"`
	PromoParam *promo.PromotionUnitParam  `json:"promoParam" codec:"promoParam"`
}

type UnitCreateResponse struct {
	ErrorResp *api.ErrorResponnse     `json:"error_response,omitempty" codec:"error_response,omitempty"`
	Data      *UnitCreateResponseData `json:"jingdong_pop_promo_unit_createUnitPromo_responce,omitempty" codec:"jingdong_pop_promo_unit_createUnitPromo_responce,omitempty"`
}

func (r UnitCreateResponse) IsError() bool {
	return r.ErrorResp != nil || r.Data == nil || r.Data.IsError()
}

func (r UnitCreateResponse) Error() string {
	if r.ErrorResp != nil {
		return r.ErrorResp.Error()
	}
	if r.Data != nil {
		return r.Data.Error()
	}
	return "no result data"
}

type UnitCreateResponseData struct {
	Code      string                    `json:"code,omitempty" codec:"code,omitempty"`
	ErrorDesc string                    `json:"error_description,omitempty" codec:"error_description,omitempty"`
	Result    *UnitCreateResponseResult `json:"result,omitempty" codec:"result,omitempty"`
}

func (r UnitCreateResponseData) IsError() bool {
	return r.Code != "0" || r.Result == nil || r.Result.IsError()
}

func (r UnitCreateResponseData) Error() string {
	if r.Result != nil && r.Result.IsError() {
		return r.Result.Error()
	}
	return sdk.ErrorString(r.Code, r.ErrorDesc)
}

type UnitCreateResponseResult struct {
	Code    string `json:"code,omitempty" codec:"code,omitempty"`
	Success bool   `json:"success,omitempty" codec:"success,omitempty"`
	Data    uint64 `json:"data,omitempty" codec:"data,omitempty"`
	Msg     string `json:"msg,omitempty" codec:"msg,omitempty"`
}

func (r UnitCreateResponseResult) IsError() bool {
	return r.Code != "200" || !r.Success
}

func (r UnitCreateResponseResult) Error() string {
	return sdk.ErrorString(r.Code, r.Msg)
}

func UnitCreate(ctx context.Context, req *UnitCreateRequest) (uint64, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := promo.NewSellerPromotionUnitCreateRequest()

	r.SetClient(req.Client)
	r.SetPromoParam(req.PromoParam)

	var response UnitCreateResponse
	if err := client.PostExecute(ctx, r.Request, req.Session, &response); err != nil {
		return 0, err
	}
	return response.Data.Result.Data, nil
}
