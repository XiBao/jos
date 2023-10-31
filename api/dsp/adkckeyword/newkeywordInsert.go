package adkckeyword

import (
	"fmt"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/dsp/adkckeyword"
)

type KeywordInsertRequest struct {
	api.BaseRequest
	Name      string  `json:"name,omitempty" codec:"name,omitempty"`             //关键字
	Price     float64 `json:"price,omitempty" codec:"price,omitempty"`           //出价
	Type      uint8   `json:"type,omitempty" codec:"type,omitempty"`             //购买类型：1.精确匹配4.短语匹配8.切词匹配
	AdGroupId uint64  `json:"ad_groupId,omitempty" codec:"ad_groupId,omitempty"` //单元id
}

type KeywordInsertResponse struct {
	ErrorResp *api.ErrorResponnse `json:"error_response,omitempty" codec:"error_response,omitempty"`
	Data      *KeywordInsertData  `json:"jingdong_dsp_adkckeyword_newkeyword_insert_responce,omitempty" codec:"jingdong_dsp_adkckeyword_newkeyword_insert_responce,omitempty"`
}

func (r KeywordInsertResponse) IsError() bool {
	return r.ErrorResp != nil || r.Data == nil || r.Data.IsError()
}

func (r KeywordInsertResponse) Error() string {
	if r.ErrorResp != nil {
		return r.ErrorResp.Error()
	}
	if r.Data != nil {
		return r.Data.Error()
	}
	return "no result data"
}

type KeywordInsertData struct {
	Result KeywordInsertResult `json:"searchrecommendkeywords_result,omitempty" codec:"searchrecommendkeywords_result,omitempty"`
}

func (r KeywordInsertData) IsError() bool {
	return r.Result.IsError()
}

func (r KeywordInsertData) Error() string {
	return r.Result.Error()
}

type KeywordInsertResult struct {
	Success    bool   `json:"success,omitempty" codec:"success,omitempty"`
	ResultCode string `json:"resultCode,omitempty" codec:"resultCode,omitempty"`
	ErrorMsg   string `json:"errorMsg,omitempty" codec:"errorMsg,omitempty"`
}

func (r KeywordInsertResult) IsError() bool {
	return !r.Success
}

func (r KeywordInsertResult) Error() string {
	return fmt.Sprintf("code:%s, msg:%s", r.ResultCode, r.ErrorMsg)
}

// 插入关键词
func KeywordInsert(req *KeywordInsertRequest) (bool, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := adkckeyword.NewkeywordInsertRequest()
	r.SetName(req.Name)
	r.SetPrice(req.Price)
	r.SetType(req.Type)
	r.SetAdGroupId(req.AdGroupId)

	var response KeywordInsertResponse
	if err := client.Execute(r.Request, req.Session, &response); err != nil {
		return false, err
	}
	return true, nil
}
