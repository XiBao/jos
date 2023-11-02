package fullcoupon

import (
	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/fullcoupon"
)

// 分页查询满额返券活动列表，限制每页最多查询20条数据
type FullCouponGetPromoListInfoRequest struct {
	api.BaseRequest
	WareId    uint64 `json:"wareId,omitempty" codec:"wareId,omitempty"`       // 商品编码
	PageIndex int    `json:"pageIndex,omitempty" codec:"pageIndex,omitempty"` // 页码
	EvtStatus int    `json:"evtStatus" codec:"evtStatus"`                     // 促销状态 全部：-1 ；系统未审核：1；人工未审核：5；驳回：11；未开始：2；进行中：3；已暂停：4；已结束：6；即将结束：20
	EvtName   string `json:"evtName,omitempty" codec:"evtName,omitempty"`     // 促销名称
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

func (r FullCouponGetPromoListInfoResponse) IsError() bool {
	return r.ErrorResp != nil || r.Data == nil || r.Data.IsError()
}

func (r FullCouponGetPromoListInfoResponse) Error() string {
	if r.ErrorResp != nil {
		return r.ErrorResp.Error()
	}
	if r.Data != nil {
		return r.Data.Error()
	}
	return "no result data"
}

type FullCouponGetPromoListInfoResponseResult struct {
	Result *FullCouponGetPromoListInfoResponseData `json:"result,omitempty" codec:"result,omitempty"`
}

func (r FullCouponGetPromoListInfoResponseResult) IsError() bool {
	return r.Result == nil || r.Result.IsError()
}

func (r FullCouponGetPromoListInfoResponseResult) Error() string {
	if r.Result != nil {
		return r.Result.Error()
	}
	return "no result data"
}

type FullCouponGetPromoListInfoResponseData struct {
	Msg     string                                      `json:"msg,omitempty" codec:"msg,omitempty"`         // 状态码描述
	Code    string                                      `json:"code,omitempty" codec:"code,omitempty"`       // 状态码
	Success bool                                        `json:"success,omitempty" codec:"success,omitempty"` // 请求是否成功
	Data    *FullCouponGetPromoListInfoResponseDataList `json:"data,omitempty" codec:"data,omitempty"`
}

func (r FullCouponGetPromoListInfoResponseData) IsError() bool {
	return r.Data == nil
}

func (r FullCouponGetPromoListInfoResponseData) Error() string {
	return sdk.ErrorString(r.Code, r.Msg)
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

	var response FullCouponGetPromoListInfoResponse
	if err := client.Execute(r.Request, req.Session, &response); err != nil {
		return nil, 0, err
	}

	return response.Data.Result.Data.PromoList, response.Data.Result.Data.TotalPageCount, nil
}
