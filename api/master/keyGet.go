package master

import (
	"errors"
	"fmt"
	"time"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/crypto"
	"github.com/XiBao/jos/sdk/request/master"
	"github.com/daviddengcn/ljson"
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

type MasterKeyResponse struct {
	Result *MasterKeyResult `json:"response,omitempty" codec:"response,omitempty"`
}

type MasterKeyResult struct {
	Code              int                `json:"status_code,omitempty" codec:"status_code,omitempty"`
	ErrorDesc         string             `json:"errorMsg,omitempty" codec:"errorMsg,omitempty"`
	Tid               string             `json:"tid,omitempty" codec:"tid,omitempty"`
	Ts                int64              `json:"ts,omitempty" codec:"ts,omitempty"`
	EncService        string             `json:"enc_service,omitempty" codec:"enc_service,omitempty"`
	KeyCacheDisabled  int                `json:"key_cache_disabled,omitempty" codec:"key_cache_disabled,omitempty"`
	KeyBackupDisabled int                `json:"key_backup_disabled,omitempty" codec:"key_backup_disabled,omitempty"`
	ServiceKeyList    []*crypto.KeyStore `json:"service_key_list,omitempty" codec:"service_key_list,omitempty"`
}

// 获取数据解密的密钥
func MasterKeyGet(req *MasterKeyGetRequest) (keyStore *crypto.KeyStore, err error) {
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
		return
	}
	r.SetSig(sig)
	r.SetSdkVer(req.SdkVer)
	r.SetTs(req.Ts)
	r.SetTid(req.Tid)

	result, err := client.Execute(r.Request, req.Session)
	if err != nil {
		return
	}
	if len(result) == 0 {
		err = errors.New("No result.")
		return
	}
	var response MasterKeyGetResponse
	err = ljson.Unmarshal(result, &response)
	if err != nil {
		return
	}
	if response.ErrorResp != nil {
		err = response.ErrorResp
		return
	}
	if response.Response.Result == nil {
		err = errors.New("No result.")
		return
	}
	if response.Response.Result.Code != 0 {
		err = errors.New(response.Response.Result.ErrorDesc)
		return
	}
	if len(response.Response.Result.ServiceKeyList) == 0 {
		err = errors.New("no result")
		return
	}
	return response.Response.Result.ServiceKeyList[0], nil
}
