package isv

import (
	"errors"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/isv"
	"github.com/daviddengcn/ljson"
)

type UploadDBOperationLogRequest struct {
	api.BaseRequest
	UserIp    string `json:"user_ip" codec:"user_ip"`       // 该访问请求的客户端外网IP
	AppName   string `json:"app_name" codec:"app_name"`     // 日志产生的应用名称
	JosAppKey string `json:"josAppKey," codec:"josAppKey"`  // 宙斯开放平台颁发的app_key
	DeviceId  string `json:"device_id" codec:"device_id"`   // 用户设备唯一标识
	UserId    string `json:"user_id" codec:"user_id"`       // ISV帐号体系中的用户ID或者用户名
	Url       string `json:"url" codec:"url"`               // 客户端请求的URL客户端请求的URL
	Db        string `json:"db" codec:"db"`                 // 连接的数据库实例名称或IP
	Sql       string `json:"sql" codec:"sql"`               // sql语句
	Timestamp uint64 `json:"time_stamp" codec:"time_stamp"` // 整型时间戳，精确到毫秒，1970年01月01日0点中以来的毫秒数
}

type UploadDBOperationLogResponse struct {
	ErrorResp *api.ErrorResponnse         `json:"error_response,omitempty" codec:"error_response,omitempty"`
	Data      *UploadDBOperationLogResult `json:"jingdong_isv_uploadDBOperationLog_responce,omitempty" codec:"jingdong_isv_uploadDBOperationLog_responce,omitempty"`
}

type UploadDBOperationLogResult struct {
	Code string `json:"code,omitempty" codec:"code,omitempty"`     //返回码
	C    int    `json:"result,omitempty" codec:"result,omitempty"` //是否成功
}

func UploadDBOperationLog(req *UploadDBOperationLogRequest) (int, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := isv.NewIsvUploadDBOperationLogRequest()

	r.SetUserIp(req.UserIp)
	r.SetAppName(req.AppName)
	r.SetJosAppKey(req.JosAppKey)
	r.SetDeviceId(req.DeviceId)
	r.SetUserId(req.UserId)
	r.SetUrl(req.Url)
	r.SetDb(req.Db)
	r.SetSql(req.Sql)
	r.SetTimestamp(req.Timestamp)
	r.Request.IsLogGW = true

	result, err := client.Execute(r.Request, req.Session)
	if err != nil {
		return -1, err
	}
	if len(result) == 0 {
		return -1, errors.New("no result info")
	}
	var response UploadDBOperationLogResponse
	err = ljson.Unmarshal(result, &response)
	if err != nil {
		return -1, err
	}

	if response.ErrorResp != nil {
		return -1, response.ErrorResp
	}

	if response.Data.Code != "0" {
		return -1, errors.New(response.Data.Code)
	}

	return response.Data.C, nil
}
