package adkckeyword

import (
	"errors"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/dsp/adkckeyword"
	"github.com/daviddengcn/ljson"
)

type KeywordDeleteRequest struct {
	api.BaseRequest
	AdGroupId    uint64 `json:"ad_group_id,omitempty" codec:"ad_group_id,omitempty"`       //单元id
	KeyWordsName string `json:"key_words_name,omitempty" codec:"key_words_name,omitempty"` //关键词
}

type KeywordDeleteResponse struct {
	ErrorResp *api.ErrorResponnse      `json:"error_response,omitempty" codec:"error_response,omitempty"`
	Data      *RecommendkeywordGetData `json:"jingdong_dsp_adkckeyword_keyword_delete_responce,omitempty" codec:"jingdong_dsp_adkckeyword_keyword_delete_responce,omitempty"`
}

type KeywordDeleteData struct {
	Result KeywordInsertResult `json:"result,omitempty" codec:"result,omitempty"`
}

type KeywordDeleteResult struct {
	Success    bool   `json:"success,omitempty" codec:"success,omitempty"`
	ResultCode string `json:"resultCode,omitempty" codec:"resultCode,omitempty"`
	ErrorMsg   string `json:"errorMsg,omitempty" codec:"errorMsg,omitempty"`
}

//删除关键词
func KeywordDelete(req *KeywordDeleteRequest) (bool, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := adkckeyword.NewKeywordDeleteRequest()
	r.SetAdGroupId(req.AdGroupId)
	r.SetKeyWordsName(req.KeyWordsName)
	result, err := client.Execute(r.Request, req.Session)
	if err != nil {
		return false, err
	}
	if len(result) == 0 {
		return false, errors.New("no result info")
	}
	var response KeywordDeleteResponse
	err = ljson.Unmarshal(result, &response)
	if err != nil {
		return false, err
	}
	if response.ErrorResp != nil {
		return false, response.ErrorResp
	}

	if !response.Data.Result.Success {
		if response.Data.Result.ErrorMsg == `` {
			response.Data.Result.ErrorMsg = "新建关键词失败"
		}
		return false, errors.New(response.Data.Result.ErrorMsg)
	}

	return true, nil
}
