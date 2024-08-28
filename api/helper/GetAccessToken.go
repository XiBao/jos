package helper

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/sdk"
)

type GetAccessTokenRequest struct {
	api.BaseRequest
	Code        string `json:"code,omitempty" codec:"code,omitempty"`
	State       string `json:"state,omitempty" codec:"state,omitempty"`
	RedirectUri string `json:"redirect_uri,omitempty" codec:"redirect_uri,omitempty"`
}

type GetAccessTokenResponse struct {
	Code         int    `json:"code,omitempty" codec:"code,omitempty"`
	ErrorDesc    string `json:"error_description,omitempty" codec:"error_description,omitempty"`
	AccessToken  string `json:"access_token,omitempty" codec:"access_token,omitempty"`
	ExpiresIn    int64  `json:"expires_in,omitempty" codec:"expires_in,omitempty"`
	RefreshToken string `json:"refresh_token,omitempty" codec:"refresh_token,omitempty"`
	Time         uint64 `json:"time,omitempty" codec:"time,omitempty"`
	TokenType    string `json:"token_type,omitempty" codec:"token_type,omitempty"`
	UserNick     string `json:"user_nick,omitempty" codec:"user_nick,omitempty"`
	Uid          string `json:"uid,omitempty" codec:"uid,omitempty"`
	OpenId       string `json:"open_id,omitempty" codec:"open_id,omitempty"`
}

func GetAccessToken(ctx context.Context, req *GetAccessTokenRequest) (*GetAccessTokenResponse, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug

	result, err := client.GetAccessToken(req.Code, req.State, req.RedirectUri)
	if err != nil {
		return nil, err
	}
	if result == "" {
		return nil, errors.New("no result.")
	}
	var response GetAccessTokenResponse
	err = json.Unmarshal([]byte(result), &response)
	if err != nil {
		return nil, err
	}
	if response.Code != 0 {
		return nil, errors.New(response.ErrorDesc)
	}

	return &response, nil
}
