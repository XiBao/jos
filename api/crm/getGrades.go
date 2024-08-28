package crm

import (
	"context"

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

func (r GetGradesResponse) IsError() bool {
	return r.ErrorResp != nil || r.Data == nil
}

func (r GetGradesResponse) Error() string {
	if r.ErrorResp != nil {
		return r.ErrorResp.Error()
	}
	return "no result data"
}

type GetGradesData struct {
	Code   string            `json:"code,omitempty" codec:"code,omitempty"`
	Result []GetGradesResult `json:"grade_promotions,omitempty" codec:"grade_promotions,omitempty"`
}

type GetGradesResult struct {
	CurGrade          string `json:"cur_grade,omitempty" codec:"cur_grade,omitempty"`
	CurGradeName      string `json:"cur_grade_name,omitempty" codec:"cur_grade_name,omitempty"`
	NextUpgradeCount  int    `json:"next_upgrade_count,omitempty" codec:"next_upgrade_count,omitempty"`
	NextUpgradeAmount int    `json:"next_upgrade_amount,omitempty" codec:"next_upgrade_amount,omitempty"`
	NextGrade         string `json:"next_grade,omitempty" codec:"next_grade,omitempty"`
	NextGradeName     string `json:"next_grade_name,omitempty" codec:"next_grade_name,omitempty"`
}

func GetGrades(ctx context.Context, req *GetGradesRequest) ([]GetGradesResult, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := crm.NewGetGradesRequest()

	var response GetGradesResponse
	if err := client.Execute(ctx, r.Request, req.Session, &response); err != nil {
		return nil, err
	}
	return response.Data.Result, nil
}
