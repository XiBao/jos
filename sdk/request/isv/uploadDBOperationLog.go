package isv

import (
	"github.com/XiBao/jos/sdk"
)

type IsvUploadDBOperationLogRequest struct {
	Request *sdk.Request
}

// create new request
func NewIsvUploadDBOperationLogRequest() (req *IsvUploadDBOperationLogRequest) {
	request := sdk.Request{MethodName: "jingdong.isv.uploadDBOperationLog", Params: make(map[string]interface{}, 10)}
	req = &IsvUploadDBOperationLogRequest{
		Request: &request,
	}
	return
}

func (req *IsvUploadDBOperationLogRequest) SetUserIp(userIp string) {
	req.Request.Params["user_ip"] = userIp
}

func (req *IsvUploadDBOperationLogRequest) GetUserIp() string {
	userIp, found := req.Request.Params["user_ip"]
	if found {
		return userIp.(string)
	}
	return ""
}

func (req *IsvUploadDBOperationLogRequest) SetAppName(appName string) {
	req.Request.Params["app_name"] = appName
}

func (req *IsvUploadDBOperationLogRequest) GetAppName() string {
	appName, found := req.Request.Params["app_name"]
	if found {
		return appName.(string)
	}
	return ""
}

func (req *IsvUploadDBOperationLogRequest) SetJosAppKey(josAppKey string) {
	req.Request.Params["josAppKey"] = josAppKey
}

func (req *IsvUploadDBOperationLogRequest) GetJosAppKey() string {
	josAppKey, found := req.Request.Params["josAppKey"]
	if found {
		return josAppKey.(string)
	}
	return ""
}

func (req *IsvUploadDBOperationLogRequest) SetDeviceId(deviceId string) {
	req.Request.Params["device_id"] = deviceId
}

func (req *IsvUploadDBOperationLogRequest) GetDeviceId() string {
	deviceId, found := req.Request.Params["device_id"]
	if found {
		return deviceId.(string)
	}
	return ""
}

func (req *IsvUploadDBOperationLogRequest) SetUserId(userId string) {
	req.Request.Params["user_id"] = userId
}

func (req *IsvUploadDBOperationLogRequest) GetUserId() string {
	userId, found := req.Request.Params["user_id"]
	if found {
		return userId.(string)
	}
	return ""
}

func (req *IsvUploadDBOperationLogRequest) SetUrl(url string) {
	req.Request.Params["url"] = url
}

func (req *IsvUploadDBOperationLogRequest) GetUrl() string {
	url, found := req.Request.Params["url"]
	if found {
		return url.(string)
	}
	return ""
}

func (req *IsvUploadDBOperationLogRequest) SetDb(db string) {
	req.Request.Params["db"] = db
}

func (req *IsvUploadDBOperationLogRequest) GetDb() string {
	db, found := req.Request.Params["db"]
	if found {
		return db.(string)
	}
	return ""
}

func (req *IsvUploadDBOperationLogRequest) SetSql(sql string) {
	req.Request.Params["sql"] = sql
}

func (req *IsvUploadDBOperationLogRequest) GetSql() string {
	sql, found := req.Request.Params["sql"]
	if found {
		return sql.(string)
	}
	return ""
}

func (req *IsvUploadDBOperationLogRequest) SetMessage(message string) {
	req.Request.Params["message"] = message
}

func (req *IsvUploadDBOperationLogRequest) GetMessage() string {
	message, found := req.Request.Params["message"]
	if found {
		return message.(string)
	}
	return ""
}

func (req *IsvUploadDBOperationLogRequest) SetTimestamp(timestamp uint64) {
	req.Request.Params["time_stamp"] = timestamp
}

func (req *IsvUploadDBOperationLogRequest) GetTimestamp() uint64 {
	timestamp, found := req.Request.Params["time_stamp"]
	if found {
		return timestamp.(uint64)
	}
	return 0
}
