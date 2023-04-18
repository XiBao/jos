package points

import (
	"encoding/json"
	"errors"

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

type GetPointsRuleData struct {
	Code      string     `json:"code,omitempty" codec:"code,omitempty"`
	JsfResult *JsfResult `json:"jsfResult,omitempty" codec:"jsfResult,omitempty"`
}

type JsfResult struct {
	Result []*PointsRule `json:"jsfResult,omitempty" codec:"jsfResult,omitempty"`
	Code   string        `json:"code,omitempty" codec:"code,omitempty"`
	Desc   string        `json:"desc,omitempty" codec:"desc,omitempty"`
}

func GetPointsRule(req *GetPointsRuleRequest) ([]*PointsRule, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := points.NewGetPointsRuleRequest()
	result, err := client.Execute(r.Request, req.Session)
	if err != nil {
		return nil, err
	}
	if len(result) == 0 {
		return nil, errors.New("No order info.")
	}

	var response GetPointsRuleResponse
	err = json.Unmarshal(result, &response)
	if err != nil {
		return nil, err
	}
	if response.ErrorResp != nil {
		return nil, response.ErrorResp
	}
	if response.Data.JsfResult.Desc != "SUCCESS" {
		return nil, response.ErrorResp
	}

	return response.Data.JsfResult.Result, nil

}
