package fullcoupon

import (
	"errors"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/fullcoupon"
	"github.com/daviddengcn/ljson"
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

type FullCouponGetPromoWaresResponseResult struct {
	Result *FullCouponGetPromoWaresResponseData `json:"returnType,omitempty" codec:"returnType,omitempty"`
}

type FullCouponGetPromoWaresResponseData struct {
	Msg     string                                   `json:"msg,omitempty" codec:"msg,omitempty"`         // 状态码描述
	Code    string                                   `json:"code,omitempty" codec:"code,omitempty"`       // 状态码
	Success bool                                     `json:"success,omitempty" codec:"success,omitempty"` // 请求是否成功
	Data    *FullCouponGetPromoWaresResponseDataList `json:"data,omitempty" codec:"data,omitempty"`
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

	result, err := client.Execute(r.Request, req.Session)
	if err != nil {
		return nil, 0, err
	}
	if len(result) == 0 {
		return nil, 0, errors.New("no result.")
	}

	var response FullCouponGetPromoWaresResponse
	err = ljson.Unmarshal(result, &response)
	if err != nil {
		return nil, 0, err
	}
	if response.ErrorResp != nil {
		return nil, 0, response.ErrorResp
	}

	if response.Data == nil || response.Data.Result == nil || response.Data.Result.Data == nil || response.Data.Result.Data.WareList == nil {
		return nil, 0, errors.New("no ware list.")
	}

	return response.Data.Result.Data.WareList, response.Data.Result.Data.TotalPageCount, nil
}
