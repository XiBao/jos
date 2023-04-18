package account

import (
	"encoding/json"
	"errors"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/api/util"
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/account"
)

type GetAccountInfoRequest struct {
	api.BaseRequest

	AccountType uint8  `json:"accountType" codec:"accountType"`
	AccountCode string `json:"accountCode" codec:"accountCode"`
}

type GetAccountInfoResponse struct {
	ErrorResp *api.ErrorResponnse `json:"error_response,omitempty" codec:"error_response,omitempty"`
	Data      *GetAccountInfoData `json:"jingdong_pop_account_getAccountInfo_responce,omitempty" codec:"jingdong_pop_account_getAccountInfo_responce,omitempty"`
}

type GetAccountInfoData struct {
	Code      string `json:"code,omitempty" codec:"code,omitempty"`
	ErrorDesc string `json:"error_description,omitempty" codec:"error_description,omitempty"`

	Result *AccountInfo `json:"beanAccount,omitempty" codec:"beanAccount,omitempty"`
}

func GetAccountInfo(req *GetAccountInfoRequest) (*AccountInfo, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := account.NewGetAccountInfoRequest()
	r.SetAccountType(req.AccountType)
	r.SetAccountCode(req.AccountCode)

	result, err := client.Execute(r.Request, req.Session)
	if err != nil {
		return nil, err
	}
	result = util.RemoveJsonSpace(result)

	var response GetAccountInfoResponse
	err = json.Unmarshal(result, &response)
	if err != nil {
		return nil, err
	}
	if response.ErrorResp != nil {
		return nil, response.ErrorResp
	}
	if response.Data.Code != "0" {
		return nil, errors.New(response.Data.ErrorDesc)
	}
	if response.Data.Result == nil {
		return nil, errors.New("No bean account info.")
	}

	return response.Data.Result, nil
}
