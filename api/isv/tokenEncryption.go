package isv

import (
	"context"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/isv"
)

type IsvTokenEncryptionRequest struct {
	api.BaseRequest
	TokenStr string `json:"tokenStr" codec:"tokenStr"` // 加密前的token字符串
}

type IsvTokenEncryptionResponse struct {
	ErrorResp *api.ErrorResponnse       `json:"error_response,omitempty" codec:"error_response,omitempty"`
	Result    *IsvTokenEncryptionResult `json:"jingdong_jos_isv_token_encryption_responce,omitempty" codec:"jingdong_jos_isv_token_encryption_responce,omitempty"`
}

func (r IsvTokenEncryptionResponse) IsError() bool {
	return r.ErrorResp != nil || r.Result == nil || r.Result.IsError()
}

func (r IsvTokenEncryptionResponse) Error() string {
	if r.ErrorResp != nil {
		return r.ErrorResp.Error()
	}
	if r.Result != nil {
		return r.Result.Error()
	}
	return "no result data"
}

type IsvTokenEncryptionResult struct {
	Code string                  `json:"code,omitempty" codec:"code,omitempty"`
	RT   *IsvTokenEncryptionData `json:"returnType,omitempty" codec:"returnType,omitempty"`
}

func (r IsvTokenEncryptionResult) IsError() bool {
	return r.RT == nil || r.RT.IsError()
}

func (r IsvTokenEncryptionResult) Error() string {
	if r.RT != nil {
		return r.RT.Error()
	}
	return "no result data"
}

type IsvTokenEncryptionData struct {
	Msg         string `json:"msg,omitempty" codec:"msg,omitempty"`                     // 描述信息
	Code        int    `json:"code,omitempty" codec:"code,omitempty"`                   // code码
	RequestId   string `json:"requestId,omitempty" codec:"requestId,omitempty"`         // 请求id，用以追踪
	Data        string `json:"data,omitempty" codec:"data,omitempty"`                   // 加密后的token信息
	OpenIdBuyer string `json:"open_id_buyer,omitempty" codec:"open_id_buyer,omitempty"` // 加密后的token信息
}

func (r IsvTokenEncryptionData) IsError() bool {
	return r.Code != 0
}

func (r IsvTokenEncryptionData) Error() string {
	return r.Msg
}

func IsvTokenEncryption(ctx context.Context, req *IsvTokenEncryptionRequest) (*IsvTokenEncryptionData, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := isv.NewIsvTokenEncryptionRequest()

	r.SetTokenStr(req.TokenStr)

	var response IsvTokenEncryptionResponse
	if err := client.Execute(ctx, r.Request, req.Session, &response); err != nil {
		return nil, err
	}
	if response.IsError() {
		return nil, response
	}
	return response.Result.RT, nil
}
