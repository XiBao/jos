package adkckeyword

import (
	"encoding/json"
	"errors"

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
	ErrorResp *api.ErrorResponnse      `json:"error_response,omitempty" codec:"error_response,omitempty"`
	Data      *RecommendkeywordGetData `json:"jingdong_dsp_adkckeyword_newkeyword_insert_responce,omitempty" codec:"jingdong_dsp_adkckeyword_newkeyword_insert_responce,omitempty"`
}

type KeywordInsertData struct {
	Result KeywordInsertResult `json:"searchrecommendkeywords_result,omitempty" codec:"searchrecommendkeywords_result,omitempty"`
}

type KeywordInsertResult struct {
	Success    bool   `json:"success,omitempty" codec:"success,omitempty"`
	ResultCode string `json:"resultCode,omitempty" codec:"resultCode,omitempty"`
	ErrorMsg   string `json:"errorMsg,omitempty" codec:"errorMsg,omitempty"`
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
	result, err := client.Execute(r.Request, req.Session)
	if err != nil {
		return false, err
	}
	if len(result) == 0 {
		return false, errors.New("no result info")
	}
	var response KeywordInsertResponse
	err = json.Unmarshal(result, &response)
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
