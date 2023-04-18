package shorturl

import (
	"encoding/json"
	"errors"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/shorturl"
)

type GenerateURLFastestRequest struct {
	api.BaseRequest
	Domain      string `json:"domain"`      // 3.cn
	Length      uint   `json:"length"`      // 短码长度，最小8位：100天，9位365天
	RealUrl     string `json:"realUrl"`     // 长域名
	ExpiredDays uint   `json:"expiredDays"` // 0为有访问自动续期，8位最大100天，9位最大365天
}

type GenerateURLFastestResponse struct {
	ErrorResp *api.ErrorResponnse     `json:"error_response,omitempty"`
	Data      *GenerateURLFastestData `json:"jingdong_shorturl_generateURLFastest_responce,omitempty"`
}

type GenerateURLFastestData struct {
	Result *GenerateURLFastestResult `json:"generatejdurl_result,omitempty"`
}

type GenerateURLFastestResult struct {
	Code     string `json:"code,omitempty"`
	Message  string `json:"codeText,omitempty"`
	ShortUrl string `json:"shortUrl,omitempty"`
	RealUrl  string `json:"realUrl,omitempty"`
	Username string `json:"username,omitempty"`
	Ts       int64  `json:"ts,omitempty"`
}

// 生成短域新的api接口
func GenerateURLFastest(req *GenerateURLFastestRequest) (*GenerateURLFastestResult, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := shorturl.NewGenerateURLFastestRequest()
	r.SetDomain(req.Domain)
	r.SetLength(req.Length)
	r.SetRealUrl(req.RealUrl)
	r.SetExpiredDays(req.ExpiredDays)

	result, err := client.Execute(r.Request, req.Session)
	if err != nil {
		return nil, err
	}
	var response GenerateURLFastestResponse
	err = json.Unmarshal(result, &response)
	if err != nil {
		return nil, err
	}

	if response.Data == nil {
		return nil, errors.New("no data")
	}

	if response.Data.Result.Code != "200" {
		return nil, &api.ErrorResponnse{Code: response.Data.Result.Code, ZhDesc: response.Data.Result.Message}
	}
	return response.Data.Result, nil
}
