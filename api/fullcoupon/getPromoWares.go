package fullcoupon

import (
	"fmt"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/fullcoupon"
)

// 查询活动SPU信息 每页只支持查10条
type FullCouponGetPromoWaresRequest struct {
	api.BaseRequest
	PageIndex uint   `json:"pageIndex" codec:"pageIndex"` // 页码
	AppKey    string `json:"appKey" codec:"appKey"`       // ISV渠道key
	PromoId   uint64 `json:"promoId" codec:"promoId"`     // 促销ID
}

type FullCouponGetPromoWaresResponse struct {
	ErrorResp *api.ErrorResponnse                    `json:"error_response,omitempty" codec:"error_response,omitempty"`
	Data      *FullCouponGetPromoWaresResponseResult `json:"jingdong_fullCoupon_getPromoWares_responce,omitempty" codec:"jingdong_fullCoupon_getPromoWares_responce,omitempty"`
}

func (r FullCouponGetPromoWaresResponse) IsError() bool {
	return r.ErrorResp != nil || r.Data == nil || r.Data.IsError()
}

func (r FullCouponGetPromoWaresResponse) Error() string {
	if r.ErrorResp != nil {
		return r.ErrorResp.Error()
	}
	if r.Data != nil {
		return r.Data.Error()
	}
	return "no result data"
}

type FullCouponGetPromoWaresResponseResult struct {
	Result *FullCouponGetPromoWaresResponseData `json:"returnType,omitempty" codec:"returnType,omitempty"`
}

func (r FullCouponGetPromoWaresResponseResult) IsError() bool {
	return r.Result == nil || r.Result.IsError()
}

func (r FullCouponGetPromoWaresResponseResult) Error() string {
	if r.Result != nil {
		return r.Result.Error()
	}
	return "no result data"
}

type FullCouponGetPromoWaresResponseData struct {
	Msg     string                                   `json:"msg,omitempty" codec:"msg,omitempty"`         // 状态码描述
	Code    string                                   `json:"code,omitempty" codec:"code,omitempty"`       // 状态码
	Success bool                                     `json:"success,omitempty" codec:"success,omitempty"` // 请求是否成功
	Data    *FullCouponGetPromoWaresResponseDataList `json:"data,omitempty" codec:"data,omitempty"`
}

func (r FullCouponGetPromoWaresResponseData) IsError() bool {
	return r.Data == nil
}

func (r FullCouponGetPromoWaresResponseData) Error() string {
	return fmt.Sprintf("code:%s, msg:%s", r.Code, r.Msg)
}

type FullCouponGetPromoWaresResponseDataList struct {
	Total          int         `json:"total" codec:"total"`
	TotalPageCount int         `json:"totalPageCount" codec:"totalPageCount"`
	PageIndex      int         `json:"pageIndex" codec:"pageIndex"`
	PageSize       int         `json:"pageSize" codec:"pageSize"`
	WareList       []PromoWare `json:"dataList,omitempty" codec:"dataList,omitempty"`
}

func GetPromoWares(req *FullCouponGetPromoWaresRequest) ([]PromoWare, int, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := fullcoupon.NewFullCouponGetPromoWaresRequest()
	r.SetAppKey(req.AppKey)
	r.SetPromoId(req.PromoId)
	r.SetPageIndex(req.PageIndex)

	var response FullCouponGetPromoWaresResponse
	if err := client.Execute(r.Request, req.Session, &response); err != nil {
		return nil, 0, err
	}
	return response.Data.Result.Data.WareList, response.Data.Result.Data.TotalPageCount, nil
}
