package campaign

import (
	"fmt"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/api/dsp"
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/dsp/kc/campaign"
)

type ListRequest struct {
	api.BaseRequest
	PageNum  int `json:"page_num,omitempty" codec:"page_num,omitempty"`
	PageSize int `json:"page_size,omitempty" codec:"page_size,omitempty"`
}

type ListResponse struct {
	ErrorResp *api.ErrorResponnse `json:"error_response,omitempty" codec:"error_response,omitempty"`
	Data      *ListData           `json:"jingdong_dsp_kc_campain_list_responce,omitempty" codec:"jingdong_dsp_kc_campain_list_responce,omitempty"`
}

func (r ListResponse) IsError() bool {
	return r.ErrorResp != nil || r.Data == nil || r.Data.IsError()
}

func (r ListResponse) Error() string {
	if r.ErrorResp != nil {
		return r.ErrorResp.Error()
	}
	if r.Data != nil {
		return r.Data.Error()
	}
	return "no result data"
}

type ListData struct {
	Result *ListResult `json:"querylistbyparam_result"`
}

func (r ListData) IsError() bool {
	return r.Result == nil || r.Result.IsError()
}

func (r ListData) Error() string {
	if r.Result != nil {
		return r.Result.Error()
	}
	return "no result data"
}

type ListResult struct {
	ResultCode string     `json:"resultCode,omitempty" codec:"resultCode,omitempty"`
	ErrorMsg   string     `json:"errorMsg,omitempty" codec:"errorMsg,omitempty"`
	Success    bool       `json:"success,omitempty" codec:"success,omitempty"`
	Value      *ListValue `json:"value,omitempty" codec:"value,omitempty"`
}

func (r ListResult) IsError() bool {
	return !r.Success || r.Value == nil
}

func (r ListResult) Error() string {
	if !r.Success {
		return fmt.Sprintf("code:%s, msg:%s", r.ResultCode, r.ErrorMsg)
	}
	return "no result data"
}

type ListValue struct {
	Datas     []*Query       `json:"datas,omitempty" codec:"datas,omitempty"`
	Paginator *dsp.Paginator `json:"paginator,omitempty" codec:"paginator,omitempty"`
}

type Query struct {
	Id                 uint64  `json:"id,omitempty" codec:"id,omitempty"`                                 //计划ID
	Name               string  `json:"name,omitempty" codec:"name,omitempty"`                             //计划名称
	DayBudgetStr       string  `json:"dayBudgetStr,omitempty" codec:"dayBudgetStr,omitempty"`             //预算
	DayBudgetResult    float64 `json:"dayBudgetResult,omitempty" codec:"dayBudgetResult,omitempty"`       //预算
	StartTime          uint64  `json:"startTime,omitempty" codec:"startTime,omitempty"`                   //开始时间
	EneTime            uint64  `json:"eneTime,omitempty" codec:"eneTime,omitempty"`                       //结束时间
	TimeRangePriceCoef string  `json:"timeRangePriceCoef,omitempty" codec:"timeRangePriceCoef,omitempty"` //投放时间段
	Status             int     `json:"status,omitempty" codec:"status,omitempty"`                         //状态
	PutType            int     `json:"putType,omitempty" codec:"putType,omitempty"`                       //推广类型
	BusinessType       int     `json:"businessType,omitempty" codec:"businessType,omitempty"`             //业务类型
}

// 快车.计划信息（批量获取）
func List(req *ListRequest) ([]*Query, int, error) {

	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := campaign.NewCampainListRequest()
	r.SetPageNum(req.PageNum)
	r.SetPageSize(req.PageSize)

	var response ListResponse
	if err := client.Execute(r.Request, req.Session, &response); err != nil {
		return nil, 0, err
	}
	return response.Data.Result.Value.Datas, response.Data.Result.Value.Paginator.Items, nil

}
