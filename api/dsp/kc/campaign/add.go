package campaign

import (
	"context"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/dsp/kc/campaign"
)

type AddRequest struct {
	api.BaseRequest
	Name      string `json:"name,omitempty" codec:"name,omitempty"`             // 计划名称
	DayBudget int    `json:"day_budget,omitempty" codec:"day_budget,omitempty"` // 预算
}

type AddResponse struct {
	ErrorResp *api.ErrorResponnse `json:"error_response,omitempty" codec:"error_response,omitempty"`
	Data      *AddData            `json:"jingdong_dsp_kc_campainShop_add_responce,omitempty" codec:"jingdong_dsp_kc_campainShop_add_responce,omitempty"`
}

func (r AddResponse) IsError() bool {
	return r.ErrorResp != nil || r.Data == nil || r.Data.IsError()
}

func (r AddResponse) Error() string {
	if r.ErrorResp != nil {
		return r.ErrorResp.Error()
	}
	if r.Data != nil {
		return r.Data.Error()
	}
	return "no result data"
}

type AddData struct {
	Result AddResult `json:"addcampain_result,omitempty" codec:"addcampain_result,omitempty"`
}

func (r AddData) IsError() bool {
	return r.Result.IsError()
}

func (r AddData) Error() string {
	return r.Result.Error()
}

type AddResult struct {
	CampaignId int    `json:"campaignId,omitempty" codec:"campaignId,omitempty"`
	Success    bool   `json:"success,omitempty" codec:"success,omitempty"`
	ResultCode string `json:"resultCode,omitempty" codec:"resultCode,omitempty"`
	ErrorMsg   string `json:"errorMsg,omitempty" codec:"errorMsg,omitempty"`
}

func (r AddResult) IsError() bool {
	return !r.Success
}

func (r AddResult) Error() string {
	return sdk.ErrorString(r.ResultCode, r.ErrorMsg)
}

// 新建计划
func Add(ctx context.Context, req *AddRequest) (bool, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := campaign.NewCampainShopAddRequest()

	r.SetName(req.Name)
	r.SetDayBudget(req.DayBudget)

	var response AddResponse
	if err := client.Execute(ctx, r.Request, req.Session, &response); err != nil {
		return false, err
	}
	return true, nil
}
