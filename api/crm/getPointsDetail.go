package crm

import (
	"time"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/crm"
)

type GetPointsDetailRequest struct {
	api.BaseRequest
	CustomerPin string `json:"customer_pin,omitempty" codec:"customer_pin,omitempty"` // 客户唯一号
	StartTime   string `json:"start_time,omitempty" codec:"start_time,omitempty"`     // 查询开始时间(yyyyMMddHHmmss)
	EndTime     string `json:"end_time,omitempty" codec:"end_time,omitempty"`         // 查询结束时间(yyyyMMddHHmmss)
	//StartPage   int    `json:"start_page,omitempty" codec:"start_page,omitempty"`       // 本次查询起始页码（大于等于1）
	//EndPage     int    `json:"end_page,omitempty" codec:"end_page,omitempty"`           // 本次查询终止页码（大于等于startRowkey）
	//Page        int    `json:"page,omitempty" codec:"page,omitempty"`                   // 本次查询页（要返回数据的页码，从1开始，必须在startPage和endPage之间（可以相等））
	//PageSize    int    `json:"page_size,omitempty" codec:"page_size,omitempty"`         // 页长
	//StartRowKey string `json:"start_row_key,omitempty" codec:"start_row_key,omitempty"` // 本次Hbase查询的endRowKey,做为下一查询的startRowKey（page对应页数的startRowkey，如果为空，需要page=1，从第一页开始查询）
}

type GetPointsDetailResponse struct {
	ErrorResp *api.ErrorResponnse  `json:"error_response,omitempty" codec:"error_response,omitempty"`
	Data      *GetPointsDetailData `json:"jingdong_pop_crm_getPointsDetail_responce,omitempty" codec:"jingdong_pop_crm_getPointsDetail_responce,omitempty"`
}

func (r GetPointsDetailResponse) IsError() bool {
	return r.ErrorResp != nil || r.Data == nil || r.Data.IsError()
}

func (r GetPointsDetailResponse) Error() string {
	if r.ErrorResp != nil {
		return r.ErrorResp.Error()
	}
	if r.Data != nil {
		return r.Data.Error()
	}
	return "no result data"
}

type GetPointsDetailData struct {
	Code   string                 `json:"code,omitempty" codec:"code,omitempty"`
	Result *GetPointsDetailResult `json:"getpointsdetail_result,omitempty" codec:"getpointsdetail_result,omitempty"`
}

func (r GetPointsDetailData) IsError() bool {
	return r.Result == nil
}

func (r GetPointsDetailData) Error() string {
	return "no result data"
}

type GetPointsDetailResult struct {
	PointsDetailViews []*PointsDetailView    `json:"pointsDetailViews,omitempty" codec:"pointsDetailViews,omitempty"`
	PageRowKeys       map[string]interface{} `json:"pageRowKeys,omitempty" codec:"pageRowKeys,omitempty"`
	HasNext           bool                   `json:"hasNext,omitempty" codec:"hasNext,omitempty"`
}

type PointsDetailView struct {
	OccurTime  string `json:"occurTime,omitempty" codec:"occurTime,omitempty"`
	SourceType int    `json:"sourceType,omitempty" codec:"sourceType,omitempty"`
	CurPoints  int64  `json:"curPoints,omitempty" codec:"curPoints,omitempty"`
	VenderId   uint64 `json:"venderId,omitempty" codec:"venderId,omitempty"`
	Points     int64  `json:"points,omitempty" codec:"points,omitempty"`
	BusinessId string `json:"businessId,omitempty" codec:"businessId,omitempty"`
	ResId      string `json:"resId,omitempty" codec:"resId,omitempty"`
	Msg        string `json:"msg,omitempty" codec:"msg,omitempty"`
}

func GetPointsDetail(req *GetPointsDetailRequest) ([]*PointsDetailView, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := crm.NewGetPointsDetailRequest()
	r.SetCustomerPin(req.CustomerPin)

	now := time.Now()
	if req.StartTime != "" {
		r.SetStartTime(req.StartTime)
	} else {
		r.SetStartTime(now.AddDate(-1, 0, 0).Format("20060102150405"))
	}
	if req.EndTime != "" {
		r.SetEndTime(req.EndTime)
	} else {
		r.SetEndTime(now.Format("20060102150405"))
	}
	r.SetStartPage(1)
	r.SetEndPage(10)
	r.SetPage(1)
	r.SetPageSize(50)
	//r.SetStartRowKey(req.StartRowKey)

	var response GetPointsDetailResponse
	if err := client.Execute(r.Request, req.Session, &response); err != nil {
		return nil, err
	}
	return response.Data.Result.PointsDetailViews, nil
}
