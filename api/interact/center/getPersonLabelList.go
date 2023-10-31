package center

import (
	"fmt"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/interact/center"
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

func (r GetPersonLabelListResponse) IsError() bool {
	return r.ErrorResp != nil || r.Data == nil || r.Data.IsError()
}

func (r GetPersonLabelListResponse) Error() string {
	if r.ErrorResp != nil {
		return r.ErrorResp.Error()
	}
	if r.Data != nil {
		return r.Data.Error()
	}
	return "no result data"
}

type GetPersonLabelListData struct {
	Code      string                    `json:"code,omitempty" codec:"code,omitempty"`
	ErrorDesc string                    `json:"error_description,omitempty" codec:"error_description,omitempty"`
	Result    *GetPersonLabelListResult `json:"Results,omitempty" codec:"Results,omitempty"`
}

func (r GetPersonLabelListData) IsError() bool {
	return r.Code != "0" || r.Result == nil || r.Result.IsError()
}

func (r GetPersonLabelListData) Error() string {
	if r.Result != nil && r.Result.IsError() {
		return r.Result.Error()
	}
	if r.Code != "0" {
		return fmt.Sprintf("code:%s, msg:%s", r.Code, r.ErrorDesc)
	}
	return "no result data"
}

type GetPersonLabelListResult struct {
	Code     int                       `json:"code,omitempty" codec:"code,omitempty"`
	Msg      string                    `json:"msg,omitempty" codec:"code,omitempty"`
	DataList []GetPersonLabelListLabel `json:"dataList,omitempty" codec:"dataList,omitempty"`
}

func (r GetPersonLabelListResult) IsError() bool {
	return r.Code != 200
}

func (r GetPersonLabelListResult) Error() string {
	return fmt.Sprintf("code:%d, msg:%s", r.Code, r.Msg)
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

	var response GetPersonLabelListResponse
	if err := client.Execute(r.Request, req.Session, &response); err != nil {
		return nil, err
	}
	return response.Data.Result.DataList, nil
}
