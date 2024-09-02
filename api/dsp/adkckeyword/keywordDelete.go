package adkckeyword

import (
	"context"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/dsp/adkckeyword"
)

type KeywordDeleteRequest struct {
	api.BaseRequest
	AdGroupId    uint64 `json:"ad_group_id,omitempty" codec:"ad_group_id,omitempty"`       // 单元id
	KeyWordsName string `json:"key_words_name,omitempty" codec:"key_words_name,omitempty"` // 关键词
}

type KeywordDeleteResponse struct {
	ErrorResp *api.ErrorResponnse `json:"error_response,omitempty" codec:"error_response,omitempty"`
	Data      *KeywordDeleteData  `json:"jingdong_dsp_adkckeyword_keyword_delete_responce,omitempty" codec:"jingdong_dsp_adkckeyword_keyword_delete_responce,omitempty"`
}

func (r KeywordDeleteResponse) IsError() bool {
	return r.ErrorResp != nil || r.Data == nil || r.Data.IsError()
}

func (r KeywordDeleteResponse) Error() string {
	if r.ErrorResp != nil {
		return r.ErrorResp.Error()
	}
	if r.Data != nil {
		return r.Data.Error()
	}
	return "no result data"
}

type KeywordDeleteData struct {
	Result KeywordDeleteResult `json:"result,omitempty" codec:"result,omitempty"`
}

func (r KeywordDeleteData) IsError() bool {
	return r.Result.IsError()
}

func (r KeywordDeleteData) Error() string {
	return r.Result.Error()
}

type KeywordDeleteResult struct {
	Success    bool   `json:"success,omitempty" codec:"success,omitempty"`
	ResultCode string `json:"resultCode,omitempty" codec:"resultCode,omitempty"`
	ErrorMsg   string `json:"errorMsg,omitempty" codec:"errorMsg,omitempty"`
}

func (r KeywordDeleteResult) IsError() bool {
	return !r.Success
}

func (r KeywordDeleteResult) Error() string {
	return sdk.ErrorString(r.ResultCode, r.ErrorMsg)
}

// 删除关键词
func KeywordDelete(ctx context.Context, req *KeywordDeleteRequest) (bool, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := adkckeyword.NewKeywordDeleteRequest()
	r.SetAdGroupId(req.AdGroupId)
	r.SetKeyWordsName(req.KeyWordsName)

	var response KeywordDeleteResponse
	if err := client.Execute(ctx, r.Request, req.Session, &response); err != nil {
		return false, err
	}
	return true, nil
}
