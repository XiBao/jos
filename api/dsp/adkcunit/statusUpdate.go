package adkcunit

import (
	"context"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/dsp/adkcunit"
)

type AdkcunitStatusUpdateRequest struct {
	api.BaseRequest
	Status    uint8  `json:"status"`      // 0 1 2
	AdGroupId string `json:"ad_group_id"` // 支持批量修改  "id1,id2,id3"
}

type AdkcunitStatusUpdateResponse struct {
	ErrorResp *api.ErrorResponnse       `json:"error_response,omitempty" codec:"error_response,omitempty"`
	Data      *AdkcunitStatusUpdateData `json:"jingdong_dsp_adkcunit_status_update_responce,omitempty" codec:"jingdong_dsp_adkcunit_status_update_responce,omitempty"`
}

func (r AdkcunitStatusUpdateResponse) IsError() bool {
	return r.ErrorResp != nil || r.Data == nil || r.Data.IsError()
}

func (r AdkcunitStatusUpdateResponse) Error() string {
	if r.ErrorResp != nil {
		return r.ErrorResp.Error()
	}
	if r.Data != nil {
		return r.Data.Error()
	}
	return "no result data"
}

type AdkcunitStatusUpdateData struct {
	Result *AdkcunitStatusUpdateResult `json:"updatestatus_result,omitempty" codec:"updatestatus_result,omitempty"`
}

func (r AdkcunitStatusUpdateData) IsError() bool {
	return r.Result == nil || r.Result.IsError()
}

func (r AdkcunitStatusUpdateData) Error() string {
	if r.Result != nil {
		return r.Result.Error()
	}
	return "no result data"
}

type AdkcunitStatusUpdateResult struct {
	Status     uint8  `json:"status,omitempty" codec:"status,omitempty"`
	ErrorMsg   string `json:"errorMsg,omitempty" codec:"errorMsg,omitempty"`
	ResultCode string `json:"resultCode,omitempty" codec:"resultCode,omitempty"`
	Success    bool   `json:"success,omitempty" codec:"success,omitempty"`
}

func (r AdkcunitStatusUpdateResult) IsError() bool {
	return !r.Success
}

func (r AdkcunitStatusUpdateResult) Error() string {
	return sdk.ErrorString(r.ResultCode, r.ErrorMsg)
}

// 更新单元状态
func AdkcunitStatusUpdate(ctx context.Context, req *AdkcunitStatusUpdateRequest) (bool, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := adkcunit.NewAdkcunitStatusUpdateRequest()
	r.SetStatus(req.Status)
	r.SetAdGroupId(req.AdGroupId)

	var response AdkcunitStatusUpdateResponse
	if err := client.Execute(ctx, r.Request, req.Session, &response); err != nil {
		return false, err
	}
	return true, nil
}
