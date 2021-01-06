package master

import (
	"github.com/XiBao/jos/sdk"
)

type MasterKeyGetRequest struct {
	Request *sdk.Request
}

// create new request
func NewMasterKeyGet() (req *MasterKeyGetRequest) {
	request := sdk.Request{MethodName: "jingdong.jos.master.key.get", Params: make(map[string]interface{}, 4)}
	req = &MasterKeyGetRequest{
		Request: &request,
	}
	return
}

func (req *MasterKeyGetRequest) SetSig(sig string) {
	req.Request.Params["sig"] = sig
}

func (req *MasterKeyGetRequest) GetSig() string {
	sig, found := req.Request.Params["sig"]
	if found {
		return sig.(string)
	}
	return ""
}

func (req *MasterKeyGetRequest) SetSdkVer(sdkVer int) {
	req.Request.Params["sdk_ver"] = sdkVer
}

func (req *MasterKeyGetRequest) GetSdkVer() int {
	sdkVer, found := req.Request.Params["sdk_ver"]
	if found {
		return sdkVer.(int)
	}
	return 0
}

func (req *MasterKeyGetRequest) SetTs(ts int64) {
	req.Request.Params["ts"] = ts
}

func (req *MasterKeyGetRequest) GetTs() int64 {
	ts, found := req.Request.Params["ts"]
	if found {
		return ts.(int64)
	}
	return 0
}

func (req *MasterKeyGetRequest) SetTid(tid string) {
	req.Request.Params["tid"] = tid
}

func (req *MasterKeyGetRequest) GetTid() string {
	tid, found := req.Request.Params["tid"]
	if found {
		return tid.(string)
	}
	return ""
}
