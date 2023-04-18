package campaign

import (
	"encoding/json"
	"errors"

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

type StatusUpdateData struct {
	Result *StatusUpdateResult `json:"updatestatus_result"`
}

type StatusUpdateResult struct {
	Status     int    `json:"status,omitempty" codec:"status,omitempty"`
	ResultCode string `json:"resultCode,omitempty" codec:"resultCode,omitempty"`
	ErrorMsg   string `json:"errorMsg,omitempty" codec:"errorMsg,omitempty"`
	Success    bool   `json:"success,omitempty" codec:"success,omitempty"`
}

// 修改计划状态
func StatusUpdate(req *StatusUpdateRequest) (bool, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := campaign.NewCampainStatusUpdateRequest()

	r.SetStatus(req.Status)
	r.SetCompaignId(req.CompaignId)

	result, err := client.Execute(r.Request, req.Session)
	if err != nil {
		return false, err
	}
	if len(result) == 0 {
		return false, errors.New("no result info")
	}
	var response StatusUpdateResponse
	err = json.Unmarshal(result, &response)
	if err != nil {
		return false, err
	}
	if response.ErrorResp != nil {
		return false, response.ErrorResp
	}

	if !response.Data.Result.Success {
		return false, errors.New(response.Data.Result.ErrorMsg)
	}

	return true, nil
}
