package campaign

import (
	"context"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/dsp/kc/campaign"
)

type DaybudgetUpdateRequest struct {
	api.BaseRequest
	CampaignId int    `json:"campaignId,omitempty" codec:"campaignId,omitempty"` // 计划id
	DayBudget  uint64 `json:"day_budget,omitempty" codec:"day_budget,omitempty"` // 预算
}

type DaybudgetUpdateResponse struct {
	ErrorResp *api.ErrorResponnse  `json:"error_response,omitempty" codec:"error_response,omitempty"`
	Data      *DaybudgetUpdateData `json:"jingdong_dsp_kc_campain_daybudget_update_responce,omitempty" codec:"jingdong_dsp_kc_campain_daybudget_update_responce,omitempty"`
}

func (r DaybudgetUpdateResponse) IsError() bool {
	return r.ErrorResp != nil || r.Data == nil || r.Data.IsError()
}

func (r DaybudgetUpdateResponse) Error() string {
	if r.ErrorResp != nil {
		return r.ErrorResp.Error()
	}
	if r.Data != nil {
		return r.Data.Error()
	}
	return "no result data"
}

type DaybudgetUpdateData struct {
	Result *DaybudgetUpdateResult `json:"updatecampaigndaybudget_result"`
}

func (r DaybudgetUpdateData) IsError() bool {
	return r.Result == nil || r.Result.IsError()
}

func (r DaybudgetUpdateData) Error() string {
	if r.Result != nil {
		return r.Result.Error()
	}
	return "no result data"
}

type DaybudgetUpdateResult struct {
	CampaignId int    `json:"campaignId,omitempty" codec:"campaignId,omitempty"`
	ResultCode string `json:"resultCode,omitempty" codec:"resultCode,omitempty"`
	ErrorMsg   string `json:"errorMsg,omitempty" codec:"errorMsg,omitempty"`
	Success    bool   `json:"success,omitempty" codec:"success,omitempty"`
}

func (r DaybudgetUpdateResult) IsError() bool {
	return !r.Success
}

func (r DaybudgetUpdateResult) Error() string {
	return sdk.ErrorString(r.ResultCode, r.ErrorMsg)
}

// 修改计划日限额
func DaybudgetUpdate(ctx context.Context, req *DaybudgetUpdateRequest) (bool, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := campaign.NewCampainDaybudgetUpdateRequest()

	r.SetCampaignId(req.CampaignId)
	r.SetDayBudget(req.DayBudget)

	var response DaybudgetUpdateResponse
	if err := client.Execute(ctx, r.Request, req.Session, &response); err != nil {
		return false, err
	}
	return true, nil
}
