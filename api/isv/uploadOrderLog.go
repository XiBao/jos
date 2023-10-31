package isv

import (
	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/isv"
)

type UploadOrderLogRequest struct {
	api.BaseRequest
	UserIp    string `json:"user_ip" codec:"user_ip"`       // 该访问请求的客户端外网IP
	AppName   string `json:"app_name" codec:"app_name"`     // 日志产生的应用名称
	JosAppKey string `json:"josAppKey," codec:"josAppKey"`  // 宙斯开放平台颁发的app_key
	JdId      string `json:"jd_id" codec:"jd_id"`           // 和用户关联的京东帐号（如果没有帐号，设置可以关联到京东帐号的信息，如店铺名。如果关联多个帐号，用英文逗号分隔）
	DeviceId  string `json:"device_id" codec:"device_id"`   // 用户设备唯一标识
	UserId    string `json:"user_id" codec:"user_id"`       // ISV帐号体系中的用户ID或者用户名
	FileMd5   string `json:"file_md5" codec:"file_md5"`     // 如对应订单操作为导出，导出的文件MD5值
	OrderIds  string `json:"order_ids" codec:"order_ids"`   // 订单号列表，用英文逗号分隔，每次最多100条记录。如果超过100条，拆分成多条请求。
	Operation int    `json:"operation" codec:"operation"`   // 对订单的操作，参考OrderOperationTypeEnum 1 批量查询订单列表 2 批量打印订单列表 3 批量导出订单列表 4 批量删除订单列表 5 查询订单 6 打印订单 7 导出订单 8 删除订单 9 修改订单
	Url       string `json:"url" codec:"url"`               // 客户端请求的URL,如果是服务器自行触发的订单读取操作（比如通过定时任务读取处理订单），url为空
	Timestamp uint64 `json:"time_stamp" codec:"time_stamp"` // 整型时间戳，精确到毫秒，1970年01月01日0点中以来的毫秒数
}

type UploadOrderLogResponse struct {
	ErrorResp *api.ErrorResponnse   `json:"error_response,omitempty" codec:"error_response,omitempty"`
	Data      *UploadOrderLogResult `json:"jingdong_isv_uploadOrderInfoLog_responce,omitempty" codec:"jingdong_isv_uploadOrderInfoLog_responce,omitempty"`
}

func (r UploadOrderLogResponse) IsError() bool {
	return r.ErrorResp != nil || r.Data == nil || r.Data.IsError()
}

func (r UploadOrderLogResponse) Error() string {
	if r.ErrorResp != nil {
		return r.ErrorResp.Error()
	}
	if r.Data != nil {
		return r.Data.Error()
	}
	return "no result data"
}

type UploadOrderLogResult struct {
	Code string `json:"code,omitempty" codec:"code,omitempty"`     //返回码
	C    int    `json:"result,omitempty" codec:"result,omitempty"` //是否成功
}

func (r UploadOrderLogResult) IsError() bool {
	return r.Code != "0"
}

func (r UploadOrderLogResult) Error() string {
	return r.Code
}

func UploadOrderLog(req *UploadOrderLogRequest) (int, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := isv.NewIsvUploadOrderLogRequest()

	r.SetUserIp(req.UserIp)
	r.SetAppName(req.AppName)
	r.SetJosAppKey(req.JosAppKey)
	r.SetJdId(req.JdId)
	r.SetDeviceId(req.DeviceId)
	r.SetUserId(req.UserId)
	r.SetFileMd5(req.FileMd5)
	r.SetOrderIds(req.OrderIds)
	r.SetOperation(req.Operation)
	r.SetUrl(req.Url)
	r.SetTimestamp(req.Timestamp)
	r.Request.IsLogGW = true

	var response UploadOrderLogResponse
	if err := client.Execute(r.Request, req.Session, &response); err != nil {
		return -1, err
	}
	return response.Data.C, nil
}
