package ware

import (
	"encoding/json"
	"errors"
	"fmt"
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

type PriceChangesData struct {
	PriceChanges []*PriceChangeOrig `json:"price_changes,omitempty" codec:"price_changes,omitempty"`
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
func PriceGet(req *PriceGetRequest) (*PriceChange, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := ware.NewWarePriceGetRequest()
	r.SetSkuId(fmt.Sprintf("J_%d", req.SkuId))

	result, err := client.Execute(r.Request, req.Session)
	if err != nil {
		return nil, err
	}
	if len(result) == 0 {
		return nil, errors.New("No result.")
	}
	var response PriceGetResponse
	err = json.Unmarshal(result, &response)
	if err != nil {
		return nil, err
	}

	if response.ErrorResp != nil || len(response.Data.PriceChanges) == 0 {
		return nil, response.ErrorResp
	}

	priceChange := &PriceChange{}
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
