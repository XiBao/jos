package fullcoupon

import (
	"errors"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/fullcoupon"
	"github.com/daviddengcn/ljson"
)

// 分页查询满额返券活动列表，限制每页最多查询20条数据
type FullCouponGetPromoListInfoRequest struct {
	api.BaseRequest
	WareId    uint64 `json:"wareId,omitempty" codec:"wareId,omitempty"`       // 商品编码
	PageIndex int    `json:"pageIndex,omitempty" codec:"pageIndex,omitempty"` // 页码
	EvtStatus int    `json:"evtStatus", codec:"evtStatus"`                    // 促销状态 全部：-1 ；系统未审核：1；人工未审核：5；驳回：11；未开始：2；进行中：3；已暂停：4；已结束：6；即将结束：20
	EvtName   string `json:"evtName,omitempty", codec:"evtName,omitempty"`    // 促销名称
	PageSize  int    `json:"pageSize,omitempty" codec:"pageSize,omitempty"`   // 促销名称
	StartTime string `json:"startTime,omitempty" codec:"startTime,omitempty"` // 促销开始时间
	PromoId   uint64 `json:"promoId,omitempty" codec:"promoId,omitempty"`     // 促销编码
	EndTime   string `json:"endTime,omitempty" codec:"endTime,omitempty"`     // 促销结束时间
	SkuId     uint64 `json:"skuId,omitempty" codec:"skuId,omitempty"`         // 商品ID
	AppKey    string `json:"appKey,omitempty" codec:"appKey,omitempty"`       // ISV渠道key
}

type FullCouponGetPromoListInfoResponse struct {
	ErrorResp *api.ErrorResponnse                       `json:"error_response,omitempty" codec:"error_response,omitempty"`
	Data      *FullCouponGetPromoListInfoResponseResult `json:"jingdong_fullCoupon_getPromoListInfo_responce,omitempty" codec:"jingdong_fullCoupon_getPromoListInfo_responce,omitempty"`
}

type FullCouponGetPromoListInfoResponseResult struct {
	Result *FullCouponGetPromoListInfoResponseData `json:"result,omitempty" codec:"result,omitempty"`
}

type FullCouponGetPromoListInfoResponseData struct {
	Msg     string                                      `json:"msg,omitempty" codec:"msg,omitempty"`         // 状态码描述
	Code    string                                      `json:"code,omitempty" codec:"code,omitempty"`       // 状态码
	Success bool                                        `json:"success,omitempty" codec:"success,omitempty"` // 请求是否成功
	Data    *FullCouponGetPromoListInfoResponseDataList `json:"data,omitempty" codec:"data,omitempty"`
}

type FullCouponGetPromoListInfoResponseDataList struct {
	Total          int             `json:"total" codec:"total"`
	TotalPageCount int             `json:"totalPageCount" codec:"totalPageCount"`
	PageIndex      int             `json:"pageIndex" codec:"pageIndex"`
	PageSize       int             `json:"pageSize" codec:"pageSize"`
	PromoList      []PromoListInfo `json:"dataList,omitempty" codec:"dataList,omitempty"`
}

func GetPromoListInfo(req *FullCouponGetPromoListInfoRequest) ([]PromoListInfo, int, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := fullcoupon.NewFullCouponGetPromoListInfoRequest()
	r.SetAppKey(req.AppKey)
	if req.WareId > 0 {
		r.SetWareId(req.WareId)
	}
	if req.EvtStatus != 0 {
		r.SetEvtStatus(req.EvtStatus)
	}
	if req.EvtName != "" {
		r.SetEvtName(req.EvtName)
	}
	if req.StartTime != "" {
		r.SetStartTime(req.StartTime)
	}
	if req.EndTime != "" {
		r.SetEndTime(req.EndTime)
	}
	if req.PromoId > 0 {
		r.SetPromoId(req.PromoId)
	}
	if req.SkuId > 0 {
		r.SetSkuId(req.SkuId)
	}
	r.SetPageIndex(req.PageIndex)
	r.SetPageSize(req.PageSize)

	result, err := client.Execute(r.Request, req.Session)
	if err != nil {
		return nil, 0, err
	}
	if len(result) == 0 {
		return nil, 0, errors.New("no result.")
	}

	var response FullCouponGetPromoListInfoResponse
	err = ljson.Unmarshal(result, &response)
	if err != nil {
		return nil, 0, err
	}
	if response.ErrorResp != nil {
		return nil, 0, response.ErrorResp
	}

	if response.Data == nil || response.Data.Result == nil || response.Data.Result.Data == nil || response.Data.Result.Data.PromoList == nil {
		return nil, 0, errors.New("no promo list.")
	}

	return response.Data.Result.Data.PromoList, response.Data.Result.Data.TotalPageCount, nil
}
