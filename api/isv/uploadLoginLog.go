package isv

import (
	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/isv"
)

type UploadLoginLogRequest struct {
	api.BaseRequest
	Result    uint8  `json:"result" codec:"result"`         // 登录结果， 0:成功；1:失败
	UserIp    string `json:"user_ip" codec:"user_ip"`       // 该访问请求的客户端外网IP
	AppName   string `json:"app_name" codec:"app_name"`     // 日志产生的应用名称
	JosAppKey string `json:"josAppKey," codec:"josAppKey"`  // 宙斯开放平台颁发的app_key
	JdId      string `json:"jd_id" codec:"jd_id"`           // 和用户关联的京东帐号（如果没有帐号，设置可以关联到京东帐号的信息，如店铺名。如果关联多个帐号，用英文逗号分隔）
	DeviceId  string `json:"device_id" codec:"device_id"`   // 用户设备唯一标识
	UserId    string `json:"user_id" codec:"user_id"`       // ISV帐号体系中的用户ID或者用户名
	Message   string `json:"message" codec:"message"`       // 登录结果额外信息，比如失败原因
	Timestamp uint64 `json:"time_stamp" codec:"time_stamp"` // 整型时间戳，精确到毫秒，1970年01月01日0点中以来的毫秒数
}

type UploadLoginLogResponse struct {
	ErrorResp *api.ErrorResponnse   `json:"error_response,omitempty" codec:"error_response,omitempty"`
	Data      *UploadLoginLogResult `json:"jingdong_isv_uploadLoginLog_responce,omitempty" codec:"jingdong_isv_uploadLoginLog_responce,omitempty"`
}

func (r UploadLoginLogResponse) IsError() bool {
	return r.ErrorResp != nil || r.Data == nil || r.Data.IsError()
}

func (r UploadLoginLogResponse) Error() string {
	if r.ErrorResp != nil {
		return r.ErrorResp.Error()
	}
	if r.Data != nil {
		return r.Data.Error()
	}
	return "no result data"
}

type UploadLoginLogResult struct {
	Code string `json:"code,omitempty" codec:"code,omitempty"`     //返回码
	C    int    `json:"result,omitempty" codec:"result,omitempty"` //是否成功
}

func (r UploadLoginLogResult) IsError() bool {
	return r.Code != "0"
}

func (r UploadLoginLogResult) Error() string {
	return r.Code
}

func UploadLoginLog(req *UploadLoginLogRequest) (int, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := isv.NewIsvUploadLoginLogRequest()

	r.SetResult(req.Result)
	r.SetUserIp(req.UserIp)
	r.SetAppName(req.AppName)
	r.SetJosAppKey(req.JosAppKey)
	r.SetJdId(req.JdId)
	r.SetDeviceId(req.DeviceId)
	r.SetUserId(req.UserId)
	r.SetMessage(req.Message)
	r.SetTimestamp(req.Timestamp)
	r.Request.IsLogGW = true

	var response UploadLoginLogResponse
	if err := client.Execute(r.Request, req.Session, &response); err != nil {
		return -1, err
	}
	return response.Data.C, nil
}
