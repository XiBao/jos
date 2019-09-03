package adkckeyword

import (
	"errors"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/dsp/adkckeyword"
	"github.com/daviddengcn/ljson"
)

type KeywordpricesuggestQueryRequest struct {
	api.BaseRequest
	Key        string `json:"key,omitempty" codec:"key,omitempty"`                 // 关键字
	MobileType uint8  `json:"mobile_type,omitempty" codec:"mobile_type,omitempty"` // 设备类型 0.PC 1.无线
}

type KeywordpricesuggestQueryResponse struct {
	ErrorResp *api.ErrorResponnse           `json:"error_response,omitempty" codec:"error_response,omitempty"`
	Data      *KeywordpricesuggestQueryData `json:"jingdong_dsp_adkckeyword_keywordpricesuggest_query_responce,omitempty" codec:"jingdong_dsp_adkckeyword_keywordpricesuggest_query_responce,omitempty"`
}

type KeywordpricesuggestQueryData struct {
	Result KeywordpricesuggestQueryResult `json:"getPriceForeCast_result,omitempty" codec:"getPriceForeCast_result,omitempty"`
}

type KeywordpricesuggestQueryResult struct {
	Success    bool                           `json:"success,omitempty" codec:"success,omitempty"`
	ResultCode string                         `json:"resultCode,omitempty" codec:"resultCode,omitempty"`
	ErrorMsg   string                         `json:"errorMsg,omitempty" codec:"errorMsg,omitempty"`
	Value      *KeywordpricesuggestQueryValue `json:"value,omitempty" codec:"value,omitempty"`
}
type KeywordpricesuggestQueryValue struct {
	HourHigh   []*DspDayForeCast `json:"hourHigh,omitempty" codec:"hourHigh,omitempty"`
	DayLow     []*DspDayForeCast `json:"dayLow,omitempty" codec:"dayLow,omitempty"`
	HourLow    []*DspDayForeCast `json:"hourLow,omitempty" codec:"hourLow,omitempty"`
	HourMiddle []*DspDayForeCast `json:"hourMiddle,omitempty" codec:"hourMiddle,omitempty"`
	DayMiddle  []*DspDayForeCast `json:"dayMiddle,omitempty" codec:"dayMiddle,omitempty"`
	DayHigh    []*DspDayForeCast `json:"dayHigh,omitempty" codec:"dayHigh,omitempty"`
}

// 查询.快车.关键词出价建议
func KeywordpricesuggestQuery(req *KeywordpricesuggestQueryRequest) (*KeywordpricesuggestQueryValue, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := adkckeyword.NewKeywordpricesuggestQueryRequest()
	r.SetKey(req.Key)
	r.SetMobileType(req.MobileType)
	result, err := client.Execute(r.Request, req.Session)
	if err != nil {
		return nil, err
	}
	if len(result) == 0 {
		return nil, errors.New("no result info")
	}
	var response KeywordpricesuggestQueryResponse
	err = ljson.Unmarshal(result, &response)

	if err != nil {
		return nil, err
	}
	if response.ErrorResp != nil {
		return nil, response.ErrorResp
	}

	if !response.Data.Result.Success {
		if response.Data.Result.ErrorMsg == `` {
			response.Data.Result.ErrorMsg = "新建关键词失败"
		}
		return nil, errors.New(response.Data.Result.ErrorMsg)
	}

	return response.Data.Result.Value, nil
}
