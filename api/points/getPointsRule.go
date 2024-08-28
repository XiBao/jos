package points

import (
	"context"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/points"
)

type GetPointsRuleRequest struct {
	api.BaseRequest
}

type GetPointsRuleResponse struct {
	ErrorResp *api.ErrorResponnse `json:"error_response,omitempty" codec:"error_response,omitempty"`
	Data      *GetPointsRuleData  `json:"jingdong_points_jos_getPointsRule_responce,omitempty" codec:"jingdong_points_jos_getPointsRule_responce,omitempty"`
}

func (r GetPointsRuleResponse) IsError() bool {
	return r.ErrorResp != nil || r.Data == nil || r.Data.IsError()
}

func (r GetPointsRuleResponse) Error() string {
	if r.ErrorResp != nil {
		return r.ErrorResp.Error()
	}
	if r.Data != nil {
		return r.Data.Error()
	}
	return "no result data"
}

type GetPointsRuleData struct {
	Code      string     `json:"code,omitempty" codec:"code,omitempty"`
	JsfResult *JsfResult `json:"jsfResult,omitempty" codec:"jsfResult,omitempty"`
}

func (r GetPointsRuleData) IsError() bool {
	return r.JsfResult == nil || r.JsfResult.IsError()
}

func (r GetPointsRuleData) Error() string {
	if r.JsfResult != nil {
		return r.JsfResult.Error()
	}
	return "no result data"
}

type JsfResult struct {
	Result []PointsRule `json:"jsfResult,omitempty" codec:"jsfResult,omitempty"`
	Code   string       `json:"code,omitempty" codec:"code,omitempty"`
	Desc   string       `json:"desc,omitempty" codec:"desc,omitempty"`
}

func (r JsfResult) IsError() bool {
	return r.Desc != "SUCCESS"
}

func (r JsfResult) Error() string {
	return sdk.ErrorString(r.Code, r.Desc)
}

func GetPointsRule(ctx context.Context, req *GetPointsRuleRequest) ([]PointsRule, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := points.NewGetPointsRuleRequest()

	var response GetPointsRuleResponse
	if err := client.Execute(ctx, r.Request, req.Session, &response); err != nil {
		return nil, err
	}
	return response.Data.JsfResult.Result, nil
}
