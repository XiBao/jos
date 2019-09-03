package campaign

import (
	"errors"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/dsp/kc/campaign"
	"github.com/daviddengcn/ljson"
)

type DeleteRequest struct {
	api.BaseRequest
	CompaignId string `json:"compaign_id,omitempty" codec:"compaign_id,omitempty"` //计划id
}

type DeleteResponse struct {
	ErrorResp *api.ErrorResponnse `json:"error_response,omitempty" codec:"error_response,omitempty"`
	Data      *DeleteData         `json:"jingdong_dsp_kc_campain_delete_responce,omitempty" codec:"jingdong_dsp_kc_campain_delete_responce,omitempty"`
}

type DeleteData struct {
	Result *DeleteResult `json:"deletekuaichecampaign_result"`
}

type DeleteResult struct {
	ResultCode string `json:"resultCode,omitempty" codec:"resultCode,omitempty"`
	ErrorMsg   string `json:"errorMsg,omitempty" codec:"errorMsg,omitempty"`
	Success    bool   `json:"success,omitempty" codec:"success,omitempty"`
}

//删除计划
func Delete(req *DeleteRequest) (bool, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := campaign.NewCampainDeleteRequest()
	r.SetCompaignId(req.CompaignId)

	result, err := client.Execute(r.Request, req.Session)
	if err != nil {
		return false, err
	}
	if len(result) == 0 {
		return false, errors.New("no result info")
	}
	var response DeleteResponse
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
