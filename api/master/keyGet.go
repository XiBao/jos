package master

import (
	"context"
	"fmt"
	"time"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/crypto"
	"github.com/XiBao/jos/sdk/request/master"
)

type MasterKeyGetRequest struct {
	api.BaseRequest
	SdkVer int    `json:"sdk_ver,omitempty" codec:"sdk_ver,omitempty"`
	Ts     int64  `json:"ts,omitempty" codec:"ts,omitempty"`
	Key    string `json:"key,omitempty" codec:"key,omitempty"`
	Tid    string `json:"tid,omitempty" codec:"tid,omitempty"`
}

type MasterKeyGetResponse struct {
	ErrorResp *api.ErrorResponnse `json:"error_response,omitempty" codec:"error_response,omitempty"`
	Response  *MasterKeyResponse  `json:"jingdong_jos_master_key_get_responce,omitempty" codec:"jingdong_jos_master_key_get_responce,omitempty"`
}

func (r MasterKeyGetResponse) IsError() bool {
	return r.ErrorResp != nil || r.Response == nil || r.Response.IsError()
}

func (r MasterKeyGetResponse) Error() string {
	if r.ErrorResp != nil {
		return r.ErrorResp.Error()
	}
	if r.Response != nil {
		return r.Response.Error()
	}
	return "no result data"
}

type MasterKeyResponse struct {
	Result *MasterKeyResult `json:"response,omitempty" codec:"response,omitempty"`
}

func (r MasterKeyResponse) IsError() bool {
	return r.Result == nil || r.Result.IsError()
}

func (r MasterKeyResponse) Error() string {
	if r.Result != nil {
		return r.Result.Error()
	}
	return "no result data"
}

type MasterKeyResult struct {
	Code              int               `json:"status_code,omitempty" codec:"status_code,omitempty"`
	ErrorDesc         string            `json:"errorMsg,omitempty" codec:"errorMsg,omitempty"`
	Tid               string            `json:"tid,omitempty" codec:"tid,omitempty"`
	Ts                int64             `json:"ts,omitempty" codec:"ts,omitempty"`
	EncService        string            `json:"enc_service,omitempty" codec:"enc_service,omitempty"`
	KeyCacheDisabled  int               `json:"key_cache_disabled,omitempty" codec:"key_cache_disabled,omitempty"`
	KeyBackupDisabled int               `json:"key_backup_disabled,omitempty" codec:"key_backup_disabled,omitempty"`
	ServiceKeyList    []crypto.KeyStore `json:"service_key_list,omitempty" codec:"service_key_list,omitempty"`
}

func (r MasterKeyResult) IsError() bool {
	return r.Code != 0 || len(r.ServiceKeyList) == 0
}

func (r MasterKeyResult) Error() string {
	if r.Code != 0 {
		return sdk.ErrorString(r.Code, r.ErrorDesc)
	}
	return "no result data"
}

// 获取数据解密的密钥
func MasterKeyGet(ctx context.Context, req *MasterKeyGetRequest) ([]crypto.KeyStore, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := master.NewMasterKeyGet()
	if req.SdkVer == 0 {
		req.SdkVer = 2
	}
	if req.Ts == 0 {
		req.Ts = time.Now().UnixNano() / 1000000
	}
	js := fmt.Sprintf(`{"sdk_ver":%d,"ts":%d,"tid":"%s"}`, req.SdkVer, req.Ts, req.Tid)
	sig, err := crypto.HmacSha256(req.Key, js)
	if err != nil {
		return nil, err
	}
	r.SetSig(sig)
	r.SetSdkVer(req.SdkVer)
	r.SetTs(req.Ts)
	r.SetTid(req.Tid)

	var response MasterKeyGetResponse
	if err := client.Execute(ctx, r.Request, req.Session, &response); err != nil {
		return nil, err
	}
	return response.Response.Result.ServiceKeyList, nil
}
