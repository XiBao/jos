package campaign

import (
	"context"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/dsp/kc/campaign"
)

type StatusUpdateRequest struct {
	api.BaseRequest
	Status     int    `json:"status,omitempty" codec:"status,omitempty"`
	CompaignId string `json:"compaignId,omitempty" codec:"compaignId,omitempty"`
}

type StatusUpdateResponse struct {
	ErrorResp *api.ErrorResponnse `json:"error_response,omitempty" codec:"error_response,omitempty"`
	Data      *StatusUpdateData   `json:"jingdong_dsp_kc_campain_status_update_responce,omitempty" codec:"jingdong_dsp_kc_campain_status_update_response,omitempty"`
}

func (r StatusUpdateResponse) IsError() bool {
	return r.ErrorResp != nil || r.Data == nil || r.Data.IsError()
}

func (r StatusUpdateResponse) Error() string {
	if r.ErrorResp != nil {
		return r.ErrorResp.Error()
	}
	if r.Data != nil {
		return r.Data.Error()
	}
	return "no result data"
}

type StatusUpdateData struct {
	Result *StatusUpdateResult `json:"updatestatus_result"`
}

func (r StatusUpdateData) IsError() bool {
	return r.Result == nil || r.Result.IsError()
}

func (r StatusUpdateData) Error() string {
	if r.Result != nil {
		return r.Result.Error()
	}
	return "no result data"
}

type StatusUpdateResult struct {
	Status     int    `json:"status,omitempty" codec:"status,omitempty"`
	ResultCode string `json:"resultCode,omitempty" codec:"resultCode,omitempty"`
	ErrorMsg   string `json:"errorMsg,omitempty" codec:"errorMsg,omitempty"`
	Success    bool   `json:"success,omitempty" codec:"success,omitempty"`
}

func (r StatusUpdateResult) IsError() bool {
	return !r.Success
}

func (r StatusUpdateResult) Error() string {
	return sdk.ErrorString(r.ResultCode, r.ErrorMsg)
}

// 修改计划状态
func StatusUpdate(ctx context.Context, req *StatusUpdateRequest) (bool, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := campaign.NewCampainStatusUpdateRequest()

	r.SetStatus(req.Status)
	r.SetCompaignId(req.CompaignId)

	var response StatusUpdateResponse
	if err := client.Execute(ctx, r.Request, req.Session, &response); err != nil {
		return false, err
	}
	return true, nil
}
