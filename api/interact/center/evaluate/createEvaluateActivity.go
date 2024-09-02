package center

import (
	"context"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/sdk"
	center "github.com/XiBao/jos/sdk/request/interact/center/evaluate"
)

type CreateEvaluateActivityRequest struct {
	api.BaseRequest
	ClientSource     *center.CreateEvaluateActivityClientSource `json:"ClientSource"`
	EvaluateActivity *center.CreateEvaluateActivityBody         `json:"EvaluateActivity"`
}

type CreateEvaluateActivityResponse struct {
	ErrorResp *api.ErrorResponnse         `json:"error_response,omitempty" codec:"error_response,omitempty"`
	Data      *CreateEvaluateActivityData `json:"jingdong_com_jd_interact_center_api_service_write_EvaluateActivityWriteService_createActivity_responce,omitempty" codec:"jingdong_com_jd_interact_center_api_service_write_EvaluateActivityWriteService_createActivity_responce,omitempty"`
}

func (r CreateEvaluateActivityResponse) IsError() bool {
	return r.ErrorResp != nil || r.Data == nil || r.Data.IsError()
}

func (r CreateEvaluateActivityResponse) Error() string {
	if r.ErrorResp != nil {
		return r.ErrorResp.Error()
	}
	if r.Data != nil {
		return r.Data.Error()
	}
	return "no result data"
}

type CreateEvaluateActivityData struct {
	Code      string `json:"code,omitempty" codec:"code,omitempty"`
	ErrorDesc string `json:"error_description,omitempty" codec:"error_description,omitempty"`
	Result    uint64 `json:"result,omitempty" codec:"result,omitempty"`
}

func (r CreateEvaluateActivityData) IsError() bool {
	return r.Code != "0"
}

func (r CreateEvaluateActivityData) Error() string {
	return sdk.ErrorString(r.Code, r.ErrorDesc)
}

func CreateEvaluateActivity(ctx context.Context, req *CreateEvaluateActivityRequest) (uint64, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := center.NewCreateEvaluateActivityRequest()
	r.SetClientSource(req.ClientSource)
	r.SetEvaluateActivity(req.EvaluateActivity)

	var response CreateEvaluateActivityResponse
	if err := client.Execute(ctx, r.Request, req.Session, &response); err != nil {
		return 0, err
	}
	return response.Data.Result, nil
}
