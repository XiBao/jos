package promotion

import (
	"errors"
	"fmt"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/sdk"
	promo "github.com/XiBao/jos/sdk/request/seller/promotion"
	"github.com/daviddengcn/ljson"
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

type UnitCreateResponseData struct {
	Code      string                    `json:"code,omitempty" codec:"code,omitempty"`
	ErrorDesc string                    `json:"error_description,omitempty" codec:"error_description,omitempty"`
	Result    *UnitCreateResponseResult `json:"result,omitempty" codec:"result,omitempty"`
}

type UnitCreateResponseResult struct {
	Code    string `json:"code,omitempty" codec:"code,omiempty"`
	Success bool   `json:"success,omitempty" codec:"success,omitempty"`
	Data    uint64 `json:"data,omitempty" codec:"data,omitempty"`
	Msg     string `json:"msg,omitempty" codec:"msg.omitempty"`
}

func UnitCreate(req *UnitCreateRequest) (uint64, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := promo.NewSellerPromotionUnitCreateRequest()

	r.SetClient(req.Client)
	r.SetPromoParam(req.PromoParam)

	result, err := client.Execute(r.Request, req.Session)
	fmt.Println(string(result))
	if err != nil {
		return 0, err
	}
	if len(result) == 0 {
		return 0, errors.New("no result.")
	}

	var response UnitCreateResponse
	err = ljson.Unmarshal(result, &response)
	if err != nil {
		return 0, err
	}
	if response.ErrorResp != nil {
		return 0, response.ErrorResp
	}
	if response.Data.Code != "0" {
		return 0, errors.New(response.Data.ErrorDesc)
	}
	if response.Data.Result == nil {
		return 0, errors.New("No create promotion result.")
	}
	if response.Data.Result.Code != "200" || !response.Data.Result.Success {
		if response.Data.Result.Msg == "" {
			return 0, errors.New("未知错误")
		} else {
			return 0, errors.New(response.Data.Result.Msg)
		}
	}

	return response.Data.Result.Data, nil
}
