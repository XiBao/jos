package ware

import (
	"context"
	"strconv"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/ware"
)

type PriceGetRequest struct {
	api.BaseRequest
	SkuId uint64 `json:"sku_id,omitempty" codec:"sku_id,omitempty"` //
}

type PriceGetResponse struct {
	ErrorResp *api.ErrorResponnse `json:"error_response,omitempty" codec:"error_response,omitempty"`
	Data      *PriceChangesData   `json:"jingdong_ware_price_get_responce,omitempty" codec:"jingdong_ware_price_get_responce,omitempty"`
}

func (r *PriceGetResponse) IsError() bool {
	return r.ErrorResp != nil || r.Data == nil || len(r.Data.PriceChanges) == 0
}

func (r *PriceGetResponse) Error() string {
	if r.ErrorResp != nil {
		return r.ErrorResp.Error()
	}
	return "no result data"
}

type PriceChangesData struct {
	PriceChanges []PriceChangeOrig `json:"price_changes,omitempty" codec:"price_changes,omitempty"`
}

type PriceChangeOrig struct {
	SkuId       string `json:"sku_id,omitempty" codec:"sku_id,omitempty"`
	Price       string `json:"price,omitempty" codec:"price,omitempty"`
	MarketPrice string `json:"market_price,omitempty" codec:"market_price,omitempty"`
}

type PriceChange struct {
	SkuId       uint64  `json:"sku_id,omitempty" codec:"sku_id,omitempty"`
	Price       float64 `json:"price,omitempty" codec:"price,omitempty"`
	MarketPrice float64 `json:"market_price,omitempty" codec:"market_price,omitempty"`
}

// 获取单个SKU
func PriceGet(ctx context.Context, req *PriceGetRequest) (*PriceChange, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := ware.NewWarePriceGetRequest()
	r.SetSkuId(sdk.StringsJoin("J_", strconv.FormatUint(req.SkuId, 10)))

	var response PriceGetResponse
	err := client.Execute(ctx, r.Request, req.Session, &response)
	if err != nil {
		return nil, err
	}
	priceChange := new(PriceChange)
	priceChange.SkuId = req.SkuId
	priceChange.Price, err = strconv.ParseFloat(response.Data.PriceChanges[0].Price, 64)
	if err != nil {
		return nil, err
	}
	priceChange.MarketPrice, err = strconv.ParseFloat(response.Data.PriceChanges[0].MarketPrice, 64)
	if err != nil {
		return nil, err
	}

	return priceChange, nil
}
