package center

import (
	"errors"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/api/util"
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/interact/center"
	"github.com/daviddengcn/ljson"
)

type GetPersonLabelListRequest struct {
	api.BaseRequest
	AppName string `json:"appName,omitempty" codec:"appName,omitempty"` // 调用方应用名称，新接口接入必须联系产品，出现问题概不负责，且有权利追求责任及接口降级
	Channel uint8  `json:"channel,omitempty" codec:"channel,omitempty"`
}

type GetPersonLabelListResponse struct {
	ErrorResp *api.ErrorResponnse     `json:"error_response,omitempty" codec:"error_response,omitempty"`
	Data      *GetPersonLabelListData `json:"jingdong_interact_center_api_service_read_getPersonLabelList_responce,omitempty" codec:"jingdong_interact_center_api_service_read_getPersonLabelList_responce,omitempty"`
}

type GetPersonLabelListData struct {
	Code      string                    `json:"code,omitempty" codec:"code,omitempty"`
	ErrorDesc string                    `json:"error_description,omitempty" codec:"error_description,omitempty"`
	Result    *GetPersonLabelListResult `json:"Results,omitempty" codec:"Results,omitempty"`
}

type GetPersonLabelListResult struct {
	Code     int                       `json:"code,omitempty" codec:"code,omitempty"`
	Msg      string                    `json:"msg,omitempty" codec:"code,omitempty"`
	DataList []GetPersonLabelListLabel `json:"dataList,omitempty" codec:"dataList,omitempty"`
}

type GetPersonLabelListLabel struct {
	LabelDesc string                       `json:"labelDesc" codec:"labelDesc"`
	LabelName string                       `json:"labelName" codec:"labelName"`
	LabelId   string                       `json:"labelId" codec:"labelId"`
	LabelVal  []GetPersonLabelListLabelVal `json:"labelVal" codec:"labelVal"`
}

type GetPersonLabelListLabelVal struct {
	Level string `json:"level" codec:"level"`
	Desc  string `json:"desc" codec:"desc"`
}

func GetPersonLabelList(req *GetPersonLabelListRequest) ([]GetPersonLabelListLabel, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := center.NewGetPersonLabelListRequest()
	r.SetAppName(req.AppName)
	r.SetChannel(req.Channel)

	result, err := client.Execute(r.Request, req.Session)
	if err != nil {
		return nil, err
	}
	result = util.RemoveJsonSpace(result)

	var response GetPersonLabelListResponse
	err = ljson.Unmarshal(result, &response)
	if err != nil {
		return nil, err
	}
	if response.ErrorResp != nil {
		return nil, response.ErrorResp
	}
	if response.Data.Code != "0" {
		return nil, errors.New(response.Data.ErrorDesc)
	}
	if response.Data.Result == nil {
		return nil, errors.New("No find collect info result.")
	}
	if response.Data.Result.Code != 200 && response.Data.Result.Msg != "" {
		return nil, errors.New(response.Data.Result.Msg)
	}

	return response.Data.Result.DataList, nil
}
