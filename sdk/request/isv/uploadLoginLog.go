package isv

import (
	"github.com/XiBao/jos/sdk"
)

type IsvUploadLoginLogRequest struct {
	Request *sdk.Request
}

// create new request
func NewIsvUploadLoginLogRequest() (req *IsvUploadLoginLogRequest) {
	request := sdk.Request{MethodName: "jingdong.isv.uploadLoginLog", Params: make(map[string]interface{}, 9)}
	req = &IsvUploadLoginLogRequest{
		Request: &request,
	}
	return
}

func (req *IsvUploadLoginLogRequest) SetResult(result uint8) {
	req.Request.Params["result"] = result
}

func (req *IsvUploadLoginLogRequest) GetResult() uint8 {
	result, found := req.Request.Params["result"]
	if found {
		return result.(uint8)
	}
	return 0
}

func (req *IsvUploadLoginLogRequest) SetUserIp(userIp string) {
	req.Request.Params["user_ip"] = userIp
}

func (req *IsvUploadLoginLogRequest) GetUserIp() string {
	userIp, found := req.Request.Params["user_ip"]
	if found {
		return userIp.(string)
	}
	return ""
}

func (req *IsvUploadLoginLogRequest) SetAppName(appName string) {
	req.Request.Params["app_name"] = appName
}

func (req *IsvUploadLoginLogRequest) GetAppName() string {
	appName, found := req.Request.Params["app_name"]
	if found {
		return appName.(string)
	}
	return ""
}

func (req *IsvUploadLoginLogRequest) SetJosAppKey(josAppKey string) {
	req.Request.Params["josAppKey"] = josAppKey
}

func (req *IsvUploadLoginLogRequest) GetJosAppKey() string {
	josAppKey, found := req.Request.Params["josAppKey"]
	if found {
		return josAppKey.(string)
	}
	return ""
}

func (req *IsvUploadLoginLogRequest) SetJdId(jdId string) {
	req.Request.Params["jd_id"] = jdId
}

func (req *IsvUploadLoginLogRequest) GetJdId() string {
	jdId, found := req.Request.Params["jd_id"]
	if found {
		return jdId.(string)
	}
	return ""
}

func (req *IsvUploadLoginLogRequest) SetDeviceId(deviceId string) {
	req.Request.Params["device_id"] = deviceId
}

func (req *IsvUploadLoginLogRequest) GetDeviceId() string {
	deviceId, found := req.Request.Params["device_id"]
	if found {
		return deviceId.(string)
	}
	return ""
}

func (req *IsvUploadLoginLogRequest) SetUserId(userId string) {
	req.Request.Params["user_id"] = userId
}

func (req *IsvUploadLoginLogRequest) GetUserId() string {
	userId, found := req.Request.Params["user_id"]
	if found {
		return userId.(string)
	}
	return ""
}

func (req *IsvUploadLoginLogRequest) SetMessage(message string) {
	req.Request.Params["message"] = message
}

func (req *IsvUploadLoginLogRequest) GetMessage() string {
	message, found := req.Request.Params["message"]
	if found {
		return message.(string)
	}
	return ""
}

func (req *IsvUploadLoginLogRequest) SetTimestamp(timestamp uint64) {
	req.Request.Params["time_stamp"] = timestamp
}

func (req *IsvUploadLoginLogRequest) GetTimestamp() uint64 {
	timestamp, found := req.Request.Params["time_stamp"]
	if found {
		return timestamp.(uint64)
	}
	return 0
}
