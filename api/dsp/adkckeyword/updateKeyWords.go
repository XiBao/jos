package adkckeyword

import (
	"errors"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/dsp/adkckeyword"
	"github.com/daviddengcn/ljson"
)

type UpdateKeyWordsRequest struct {
	api.BaseRequest
	Name        string `json:"name,omitempty" codec:"name,omitempty"`                    // 关键词名称
	Price       string `json:"price,omitempty" codec:"price,omitempty"`                  // 关键词出价
	Type        string `json:"type,omitempty" codec:"type,omitempty"`                    // 关键词类型:1精确匹配 4.短语匹配 8.切词包含
	MobilePrice string `json:"mobile_price,mobile_price" codec:"mobile_price,omitempty"` // 关键词无线出价
	AdGroupId   uint64 `json:"ad_group_id,omitempty" codec:"ad_group_id,omitempty"`      // 单元id
}

type UpdateKeyWordsResponse struct {
	ErrorResp *api.ErrorResponnse `json:"error_response,omitempty" codec:"error_response,omitempty"`
	Data      *UpdateKeyWordsData `json:"jingdong_dsp_adkckeyword_updateKeyWords_responce,omitempty" codec:"jingdong_dsp_adkckeyword_updateKeyWords_responce,omitempty"`
}

type UpdateKeyWordsData struct {
	Result UpdateKeyWordsResult `json:"updatekeywords_result,omitempty" codec:"updatekeywords_result,omitempty"`
}

type UpdateKeyWordsResult struct {
	Success    bool   `json:"success,omitempty" codec:"success,omitempty"`
	ResultCode string `json:"resultCode,omitempty" codec:"resultCode,omitempty"`
	ErrorMsg   string `json:"errorMsg,omitempty" codec:"errorMsg,omitempty"`
}

// 更新关键词状态
func UpdateKeyWords(req *UpdateKeyWordsRequest) (bool, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := adkckeyword.NewUpdateKeyWordsRequest()
	r.SetName(req.Name)
	r.SetPrice(req.Price)
	r.SetType(req.Type)
	if req.MobilePrice != "" {
		r.SetMobilePrice(req.MobilePrice)
	}
	r.SetAdGroupId(req.AdGroupId)

	result, err := client.Execute(r.Request, req.Session)
	if err != nil {
		return false, err
	}
	if len(result) == 0 {
		return false, errors.New("no result info")
	}
	var response UpdateKeyWordsResponse
	err = ljson.Unmarshal(result, &response)
	if err != nil {
		return false, err
	}
	if response.ErrorResp != nil {
		return false, response.ErrorResp
	}

	if !response.Data.Result.Success {
		return false, errors.New(response.Data.Result.ErrorMsg)
	}

	return true, nil

}
