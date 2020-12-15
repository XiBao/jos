package center

import (
	"errors"

	"github.com/davecgh/go-spew/spew"

	"github.com/XiBao/jos/api/util"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/sdk"
	center "github.com/XiBao/jos/sdk/request/interact/center/evaluate"
	"github.com/daviddengcn/ljson"
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

type CreateEvaluateActivityData struct {
	Code      string `json:"code,omitempty" codec:"code,omitempty"`
	ErrorDesc string `json:"error_description,omitempty" codec:"error_description,omitempty"`
	Result    uint64 `json:"result,omitempty" codec:"result,omitempty"`
}

func CreateEvaluateActivity(req *CreateEvaluateActivityRequest) (uint64, error) {
	spew.Dump(req)
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := center.NewCreateEvaluateActivityRequest()
	r.SetClientSource(req.ClientSource)
	r.SetEvaluateActivity(req.EvaluateActivity)

	result, err := client.Execute(r.Request, req.Session)
	if err != nil {
		return 0, err
	}
	result = util.RemoveJsonSpace(result)

	var response CreateEvaluateActivityResponse
	err = ljson.Unmarshal(result, &response)
	if err != nil {
		return 0, err
	}
	if response.ErrorResp != nil {
		return 0, response.ErrorResp
	}
	if response.Data.Code != "0" {
		return 0, errors.New(response.Data.ErrorDesc)
	}

	return response.Data.Result, nil
}
