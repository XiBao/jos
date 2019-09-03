package helper

import (
	"errors"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/sdk"
	"github.com/daviddengcn/ljson"
)

type GetAccessTokenNewRequest struct {
	api.BaseRequest
	Code string `json:"code,omitempty" codec:"code,omitempty"`
}

type GetAccessTokenNewResponse struct {
	Code int    `json:"code,omitempty" codec:"code,omitempty"`
	Msg  string `json:"msg,omitempty" codec:"msg,omitempty"`

	AccessToken  string `json:"access_token,omitempty" codec:"access_token,omitempty"`
	ExpiresIn    int64  `json:"expires_in,omitempty" codec:"expires_in,omitempty"`
	RefreshToken string `json:"refresh_token,omitempty" codec:"refresh_token,omitempty"`
	Time         uint64 `json:"time,omitempty" codec:"time,omitempty"`
	TokenType    string `json:"token_type,omitempty" codec:"token_type,omitempty"`
	UserNick     string `json:"user_nick,omitempty" codec:"user_nick,omitempty"`
	Uid          string `json:"uid,omitempty" codec:"uid,omitempty"`
	OpenId       string `json:"open_id,omitempty" codec:"open_id,omitempty"`
}

func GetAccessTokenNew(req *GetAccessTokenNewRequest) (*GetAccessTokenNewResponse, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug

	result, err := client.GetAccessTokenNew(req.Code)
	if err != nil {
		return nil, err
	}
	if result == "" {
		return nil, errors.New("No result.")
	}
	var response GetAccessTokenNewResponse
	err = ljson.Unmarshal([]byte(result), &response)
	if err != nil {
		return nil, err
	}
	if response.Code != 0 {
		return nil, errors.New(response.Msg)
	}

	return &response, nil
}
