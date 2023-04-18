package coupon

import (
	"encoding/json"
	"errors"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/sdk"
	coupon "github.com/XiBao/jos/sdk/request/seller/coupon"
)

type PromoCouponCreateRequest struct {
	api.BaseRequest
	ClientInfo       *coupon.SellerPromoCouponCreateClientInfo `json:"clientInfo" codec:"clientInfo"`
	CouponOuterParam *coupon.SellerPromoCouponCreateParam      `json:"couponOuterParam" codec:"couponOuterParam"`
}

type PromoCouponCreateResponse struct {
	ErrorResp *api.ErrorResponnse    `json:"error_response,omitempty" codec:"error_response,omitempty"`
	Data      *PromoCouponCreateData `json:"jingdong_pop_promo_coupon_createCoupon_responce,omitempty" codec:"jingdong_pop_promo_coupon_createCoupon_responce,omitempty"`
}

type PromoCouponCreateData struct {
	Code      string                   `json:"code,omitempty" codec:"code,omitempty"`
	ErrorDesc string                   `json:"error_description,omitempty" codec:"error_description,omitempty"`
	Result    *PromoCouponCreateResult `json:"returnType,omitempty" codec:"returnType,omitempty"`
}

type PromoCouponCreateResult struct {
	Code    string `json:"code,omitempty" codec:"code,omitempty"`
	Success bool   `json:"success,omitempty" codec:"success,omitempty"`
	Data    uint64 `json:"data,omitempty" codec:"data,omitempty"`
	Msg     string `json:"msg,omitempty" codec:"msg,omitempty"`
}

func PromoCouponCreate(req *PromoCouponCreateRequest) (uint64, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := coupon.NewSellerPromoCouponCreateRequest()

	r.SetClientInfo(req.ClientInfo)
	r.SetParam(req.CouponOuterParam)

	result, err := client.PostExecute(r.Request, req.Session)
	if err != nil {
		return 0, err
	}
	if len(result) == 0 {
		return 0, errors.New("no result.")
	}

	var response PromoCouponCreateResponse
	err = json.Unmarshal(result, &response)
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
