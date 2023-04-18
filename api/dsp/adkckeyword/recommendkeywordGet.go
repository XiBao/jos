package adkckeyword

import (
	"encoding/json"
	"errors"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/dsp/adkckeyword"
)

type RecommendkeywordGetRequest struct {
	api.BaseRequest
	SkuId       uint64 `json:"sku_id,omitempty" codec:"sku_id,omitempty"`
	SearchType  uint8  `json:"search_type,omitempty" codec:"search_type,omitempty"`     //查询范围 1.按整体查询/2.按无线端查询/3.按PC端查询
	Order       uint8  `json:"order,omitempty" codec:"order,omitempty"`                 //	排序方式 0.正序/1.倒序
	SortType    uint8  `json:"sort_type,omitempty" codec:"sort_type,omitempty"`         // 排序方式 1.按搜索量排序/2.按平均出价排序/3.按竞争激烈程度排序
	KeyWordType uint8  `json:"key_word_type,omitempty" codec:"key_word_type,omitempty"` // 	关键词来源 1.商品关键词/2.相似商品关键词/3.行业热词
}

type RecommendkeywordGetResponse struct {
	ErrorResp *api.ErrorResponnse      `json:"error_response,omitempty" codec:"error_response,omitempty"`
	Data      *RecommendkeywordGetData `json:"jingdong_dsp_adkckeyword_recommendkeyword_get_responce,omitempty" codec:"jingdong_dsp_adkckeyword_recommendkeyword_get_responce,omitempty"`
}

type RecommendkeywordGetData struct {
	Result *RecommendkeywordGetResult `json:"searchrecommendkeywords_result,omitempty" codec:"searchrecommendkeywords_result,omitempty"`
}

type RecommendkeywordGetResult struct {
	ErrorMsg   string                    `json:"errorMsg,omitempty" codec:"errorMsg,omitempty"`
	ResultCode string                    `json:"resultCode,omitempty" codec:"resultCode,omitempty"`
	Success    bool                      `json:"success,omitempty" codec:"success,omitempty"`
	Value      *RecommendkeywordGetValue `json:"data,omitempty" codec:"data,omitempty"`
}

type RecommendkeywordGetValue struct {
	Datas []*KeyWordRecommendQuery `json:"datas,omitempty" codec:"datas,omitempty"`
}

type KeyWordRecommendQuery struct {
	AvgBigPrice float64 `json:"avgBigPrice,omitempty" codec:"avgBigPrice,omitempty"` //平均出家
	StarCount   uint8   `json:"starCount,omitempty" codec:"starCount,omitempty"`     //星星数量
	Pv          uint64  `json:"pv,omitempty" codec:"pv,omitempty"`                   //三十访问量
	KeyWord     string  `json:"keyWord,omitempty" codec:"keyWord,omitempty"`         //三十访问量
}

// 获取推荐关键词
func RecommendkeywordGet(req *RecommendkeywordGetRequest) ([]*KeyWordRecommendQuery, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := adkckeyword.NewRecommendkeywordGetRequest()
	r.SetSkuId(req.SkuId)
	r.SetSearchType(req.SearchType)
	r.SetOrder(req.Order)
	r.SetSortType(req.SortType)
	r.SetKeyWordType(req.KeyWordType)
	result, err := client.Execute(r.Request, req.Session)
	if err != nil {
		return nil, err
	}
	if len(result) == 0 {
		return nil, errors.New("no result info")
	}
	var response RecommendkeywordGetResponse
	err = json.Unmarshal(result, &response)
	if err != nil {
		return nil, err
	}
	if response.ErrorResp != nil {
		return nil, response.ErrorResp
	}

	if !response.Data.Result.Success {
		return nil, errors.New(response.Data.Result.ErrorMsg)
	}

	return response.Data.Result.Value.Datas, nil
}
