package campaign

import (
	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/dsp/kc/campaign"
)

type DeleteRequest struct {
	api.BaseRequest
	CompaignId string `json:"compaign_id,omitempty" codec:"compaign_id,omitempty"` //计划id
}

type DeleteResponse struct {
	ErrorResp *api.ErrorResponnse `json:"error_response,omitempty" codec:"error_response,omitempty"`
	Data      *DeleteData         `json:"jingdong_dsp_kc_campain_delete_responce,omitempty" codec:"jingdong_dsp_kc_campain_delete_responce,omitempty"`
}

func (r DeleteResponse) IsError() bool {
	return r.ErrorResp != nil || r.Data == nil || r.Data.IsError()
}

func (r DeleteResponse) Error() string {
	if r.ErrorResp != nil {
		return r.ErrorResp.Error()
	}
	if r.Data != nil {
		return r.Data.Error()
	}
	return "no result data"
}

type DeleteData struct {
	Result *DeleteResult `json:"deletekuaichecampaign_result"`
}

func (r DeleteData) IsError() bool {
	return r.Result == nil || r.Result.IsError()
}

func (r DeleteData) Error() string {
	if r.Result != nil {
		return r.Result.Error()
	}
	return "no result data"
}

type DeleteResult struct {
	ResultCode string `json:"resultCode,omitempty" codec:"resultCode,omitempty"`
	ErrorMsg   string `json:"errorMsg,omitempty" codec:"errorMsg,omitempty"`
	Success    bool   `json:"success,omitempty" codec:"success,omitempty"`
}

func (r DeleteResult) IsError() bool {
	return !r.Success
}

func (r DeleteResult) Error() string {
	return sdk.ErrorString(r.ResultCode, r.ErrorMsg)
}

// 删除计划
func Delete(req *DeleteRequest) (bool, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := campaign.NewCampainDeleteRequest()
	r.SetCompaignId(req.CompaignId)

	var response DeleteResponse
	if err := client.Execute(r.Request, req.Session, &response); err != nil {
		return false, err
	}
	return true, nil

}
