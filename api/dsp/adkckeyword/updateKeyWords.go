package adkckeyword

import (
	"context"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/dsp/adkckeyword"
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

func (r UpdateKeyWordsResponse) IsError() bool {
	return r.ErrorResp != nil || r.Data == nil || r.Data.IsError()
}

func (r UpdateKeyWordsResponse) Error() string {
	if r.ErrorResp != nil {
		return r.ErrorResp.Error()
	}
	if r.Data != nil {
		return r.Data.Error()
	}
	return "no result data"
}

type UpdateKeyWordsData struct {
	Result UpdateKeyWordsResult `json:"updatekeywords_result,omitempty" codec:"updatekeywords_result,omitempty"`
}

func (r UpdateKeyWordsData) IsError() bool {
	return r.Result.IsError()
}

func (r UpdateKeyWordsData) Error() string {
	return r.Result.Error()
}

type UpdateKeyWordsResult struct {
	Success    bool   `json:"success,omitempty" codec:"success,omitempty"`
	ResultCode string `json:"resultCode,omitempty" codec:"resultCode,omitempty"`
	ErrorMsg   string `json:"errorMsg,omitempty" codec:"errorMsg,omitempty"`
}

func (r UpdateKeyWordsResult) IsError() bool {
	return !r.Success
}

func (r UpdateKeyWordsResult) Error() string {
	return sdk.ErrorString(r.ResultCode, r.ErrorMsg)
}

// 更新关键词状态
func UpdateKeyWords(ctx context.Context, req *UpdateKeyWordsRequest) (bool, error) {
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

	var response UpdateKeyWordsResponse
	if err := client.Execute(ctx, r.Request, req.Session, &response); err != nil {
		return false, err
	}
	return true, nil
}
