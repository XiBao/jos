package coupon

import (
	"fmt"

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

func (r PromoCouponCreateResponse) IsError() bool {
	return r.ErrorResp != nil || r.Data == nil || r.Data.IsError()
}

func (r PromoCouponCreateResponse) Error() string {
	if r.ErrorResp != nil {
		return r.ErrorResp.Error()
	}
	if r.Data != nil {
		return r.Data.Error()
	}
	return "no result data"
}

type PromoCouponCreateData struct {
	Code      string                   `json:"code,omitempty" codec:"code,omitempty"`
	ErrorDesc string                   `json:"error_description,omitempty" codec:"error_description,omitempty"`
	Result    *PromoCouponCreateResult `json:"returnType,omitempty" codec:"returnType,omitempty"`
}

func (r PromoCouponCreateData) IsError() bool {
	return r.Code != "0" || r.Result == nil || r.Result.IsError()
}

func (r PromoCouponCreateData) Error() string {
	if r.Result != nil && r.Result.IsError() {
		return r.Result.Error()
	}
	return fmt.Sprintf("code:%s, msg:%s", r.Code, r.ErrorDesc)
}

type PromoCouponCreateResult struct {
	Code    string `json:"code,omitempty" codec:"code,omitempty"`
	Success bool   `json:"success,omitempty" codec:"success,omitempty"`
	Data    uint64 `json:"data,omitempty" codec:"data,omitempty"`
	Msg     string `json:"msg,omitempty" codec:"msg,omitempty"`
}

func (r PromoCouponCreateResult) IsError() bool {
	return r.Code != "200" || !r.Success
}

func (r PromoCouponCreateResult) Error() string {
	return fmt.Sprintf("code:%s, msg:%s", r.Code, r.Msg)
}

func PromoCouponCreate(req *PromoCouponCreateRequest) (uint64, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := coupon.NewSellerPromoCouponCreateRequest()

	r.SetClientInfo(req.ClientInfo)
	r.SetParam(req.CouponOuterParam)

	var response PromoCouponCreateResponse
	if err := client.PostExecute(r.Request, req.Session, &response); err != nil {
		return 0, err
	}
	return response.Data.Result.Data, nil
}
