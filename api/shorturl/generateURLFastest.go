package shorturl

import (
	"context"

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

func (r GenerateURLFastestResponse) IsError() bool {
	return r.ErrorResp != nil || r.Data == nil || r.Data.IsError()
}

func (r GenerateURLFastestResponse) Error() string {
	if r.ErrorResp != nil {
		return r.ErrorResp.Error()
	}
	if r.Data != nil {
		return r.Data.Error()
	}
	return "no result data"
}

type GenerateURLFastestData struct {
	Result *GenerateURLFastestResult `json:"generatejdurl_result,omitempty"`
}

func (r GenerateURLFastestData) IsError() bool {
	return r.Result == nil || r.Result.IsError()
}

func (r GenerateURLFastestData) Error() string {
	if r.Result != nil {
		return r.Result.Error()
	}
	return "no result data"
}

type GenerateURLFastestResult struct {
	Code     string `json:"code,omitempty"`
	Message  string `json:"codeText,omitempty"`
	ShortUrl string `json:"shortUrl,omitempty"`
	RealUrl  string `json:"realUrl,omitempty"`
	Username string `json:"username,omitempty"`
	Ts       int64  `json:"ts,omitempty"`
}

func (r GenerateURLFastestResult) IsError() bool {
	return r.Code != "200"
}

func (r GenerateURLFastestResult) Error() string {
	return sdk.ErrorString(r.Code, r.Message)
}

// 生成短域新的api接口
func GenerateURLFastest(ctx context.Context, req *GenerateURLFastestRequest) (*GenerateURLFastestResult, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := shorturl.NewGenerateURLFastestRequest()
	r.SetDomain(req.Domain)
	r.SetLength(req.Length)
	r.SetRealUrl(req.RealUrl)
	r.SetExpiredDays(req.ExpiredDays)

	var response GenerateURLFastestResponse
	if err := client.Execute(ctx, r.Request, req.Session, &response); err != nil {
		return nil, err
	}
	return response.Data.Result, nil
}
