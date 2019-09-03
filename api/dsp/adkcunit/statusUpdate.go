package adkcunit

import (
	"errors"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/dsp/adkcunit"
	"github.com/daviddengcn/ljson"
)

type AdkcunitStatusUpdateRequest struct {
	api.BaseRequest
	Status    uint8  `json:"status"`      // 0 1 2
	AdGroupId string `json:"ad_group_id"` //支持批量修改  "id1,id2,id3"
}

type AdkcunitStatusUpdateResponse struct {
	ErrorResp *api.ErrorResponnse       `json:"error_response,omitempty" codec:"error_response,omitempty"`
	Data      *AdkcunitStatusUpdateData `json:"jingdong_dsp_adkcunit_status_update_responce,omitempty" codec:"jingdong_dsp_adkcunit_status_update_responce,omitempty"`
}

type AdkcunitStatusUpdateData struct {
	Result *AdkcunitStatusUpdateResult `json:"updatestatus_result,omitempty" codec:"updatestatus_result,omitempty"`
}

type AdkcunitStatusUpdateResult struct {
	Status     uint8  `json:"status,omitempty" codec:"status,omitempty"`
	ErrorMsg   string `json:"errorMsg,omitempty" codec:"errorMsg,omitempty"`
	ResultCode string `json:"resultCode,omitempty" codec:"resultCode,omitempty"`
	Success    bool   `json:"success,omitempty" codec:"success,omitempty"`
}

//更新单元状态
func AdkcunitStatusUpdate(req *AdkcunitStatusUpdateRequest) (bool, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := adkcunit.NewAdkcunitStatusUpdateRequest()
	r.SetStatus(req.Status)
	r.SetAdGroupId(req.AdGroupId)

	result, err := client.Execute(r.Request, req.Session)

	if err != nil {
		return false, err
	}
	if len(result) == 0 {
		return false, errors.New("no result info")
	}
	var response AdkcunitStatusUpdateResponse
	err = ljson.Unmarshal(result, &response)
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
