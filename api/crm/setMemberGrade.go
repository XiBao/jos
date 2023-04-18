package crm

import (
	"encoding/json"
	"errors"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/crm"
)

type SetMemberGradeRequest struct {
	api.BaseRequest
	Pin   string `json:"pin"`   //用户Pin
	Grade uint8  `json:"grade"` //等级
}

type SetMemberGradeResponse struct {
	ErrorResp *api.ErrorResponnse `json:"error_response,omitempty" codec:"error_response,omitempty"`
	Data      *SetMemberGradeData `json:"jingdong_pop_crm_setMemberGrade_responce,omitempty" codec:"jingdong_pop_crm_setMemberGrade_responce,omitempty"`
}

type SetMemberGradeData struct {
	ReturnType *ReturnType `json:"returnType,omitempty" codec:"returnType,omitempty"`
}

// TODO 修改会员等级
func SetMemberGrade(req *SetMemberGradeRequest) (bool, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := crm.NewSetMemberGradeRequest()

	if len(req.Pin) > 0 {
		r.SetPin(req.Pin)
	}

	if req.Grade > 0 {
		r.SetGrade(req.Grade)
	}

	result, err := client.Execute(r.Request, req.Session)
	if err != nil {
		return false, err
	}
	if len(result) == 0 {
		return false, errors.New("no result info")
	}
	var response SetMemberGradeResponse
	err = json.Unmarshal(result, &response)
	if err != nil {
		return false, err
	}
	if response.ErrorResp != nil {
		return false, response.ErrorResp
	}

	if response.Data.ReturnType.Code != "200" {
		return false, errors.New(response.Data.ReturnType.Desc)
	}

	return response.Data.ReturnType.Data, nil

}
