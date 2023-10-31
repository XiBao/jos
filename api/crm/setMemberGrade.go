package crm

import (
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

func (r SetMemberGradeResponse) IsError() bool {
	return r.ErrorResp != nil || r.Data == nil || r.Data.IsError()
}

func (r SetMemberGradeResponse) Error() string {
	if r.ErrorResp != nil {
		return r.ErrorResp.Error()
	}
	if r.Data != nil {
		return r.Data.Error()
	}
	return "no result data"
}

type SetMemberGradeData struct {
	ReturnType *ReturnType `json:"returnType,omitempty" codec:"returnType,omitempty"`
}

func (r SetMemberGradeData) IsError() bool {
	return r.ReturnType == nil || r.ReturnType.IsError()
}

func (r SetMemberGradeData) Error() string {
	if r.ReturnType != nil {
		return r.ReturnType.Error()
	}
	return "no result data"
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

	var response SetMemberGradeResponse
	if err := client.Execute(r.Request, req.Session, &response); err != nil {
		return false, err
	}
	return response.Data.ReturnType.Data, nil

}
