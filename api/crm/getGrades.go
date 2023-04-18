package crm

import (
	"encoding/json"
	"errors"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/crm"
)

type GetGradesRequest struct {
	api.BaseRequest
}

type GetGradesResponse struct {
	ErrorResp *api.ErrorResponnse `json:"error_response,omitempty" codec:"error_response,omitempty"`
	Data      *GetGradesData      `json:"jingdong_crm_grade_get_responce,omitempty" codec:"jingdong_crm_grade_get_responce,omitempty"`
}

type GetGradesData struct {
	Code   string             `json:"code,omitempty" codec:"code,omitempty"`
	Result []*GetGradesResult `json:"grade_promotions,omitempty" codec:"grade_promotions,omitempty"`
}

type GetGradesResult struct {
	CurGrade          string `json:"cur_grade,omitempty" codec:"cur_grade,omitempty"`
	CurGradeName      string `json:"cur_grade_name,omitempty" codec:"cur_grade_name,omitempty"`
	NextUpgradeCount  int    `json:"next_upgrade_count,omitempty" codec:"next_upgrade_count,omitempty"`
	NextUpgradeAmount int    `json:"next_upgrade_amount,omitempty" codec:"next_upgrade_amount,omitempty"`
	NextGrade         string `json:"next_grade,omitempty" codec:"next_grade,omitempty"`
	NextGradeName     string `json:"next_grade_name,omitempty" codec:"next_grade_name,omitempty"`
}

func GetGrades(req *GetGradesRequest) ([]*GetGradesResult, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := crm.NewGetGradesRequest()

	result, err := client.Execute(r.Request, req.Session)
	if err != nil {
		return nil, err
	}
	if len(result) == 0 {
		return nil, errors.New("No result info.")
	}
	var response GetGradesResponse
	err = json.Unmarshal(result, &response)
	if err != nil {
		return nil, err
	}

	if response.ErrorResp != nil {
		return nil, response.ErrorResp
	}

	return response.Data.Result, nil
}
