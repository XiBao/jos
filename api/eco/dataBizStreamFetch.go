package eco

import (
	"fmt"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/eco"
)

type BizStreamFetchRequest struct {
	api.BaseRequest
	ServiceId   string `json:"serviceId" codec:"serviceId"`
	TimeMin     string `json:"timeMin,omitempty" codec:"timeMin,omitempty"`
	TimeMax     string `json:"timeMax,omitempty" codec:"timeMax,omitempty"`
	Time        string `json:"TIME,omitempty" codec:"TIME,omitempty"`
	Timestamp   string `json:"TIMESTAMP,omitempty" codec:"TIMESTAMP,omitempty"`
	AdProStat   string `json:"ADPROSTAT,omitempty" codec:"ADPROSTAT,omitempty"`
	AdProId     string `json:"ADPROID,omitempty" codec:"ADPROID,omitempty"`
	Sku         string `json:"SKU,omitempty" codec:"SKU,omitempty"`
	AdType      string `json:"ADTYPE,omitempty" codec:"ADTYPE,omitempty"`
	ActEffectId string `json:"ACTEFFECTID,omitempty" codec:"ACTEFFECTID,omitempty"`
	TicketType  string `json:"TICKETTYPE,omitempty" codec:"TICKETTYPE,omitempty"`
}

type BizStreamFetchResponse struct {
	ErrorResp *api.ErrorResponnse `json:"error_response,omitempty" codec:"error_response,omitempty"`
	Res       *Response           `json:"jingdong_eco_data_biz_stream_fetch_responce,omitempty" codec:"jingdong_eco_data_biz_stream_fetch_responce,omitempty"`
}

func (r BizStreamFetchResponse) IsError() bool {
	return r.ErrorResp != nil || r.Res == nil || r.Res.IsError()
}

func (r BizStreamFetchResponse) Error() string {
	if r.ErrorResp != nil {
		return r.ErrorResp.Error()
	}
	if r.Res != nil {
		return r.Res.Error()
	}
	return "no result data"
}

type Response struct {
	Code string      `json:"code,omitempty" codec:"code,omitempty"`
	RT   *ReturnType `json:"returnType,omitempty" codec:"returnType,omitempty"`
}

func (r Response) IsError() bool {
	return r.RT == nil || r.RT.IsError()
}

func (r Response) Error() string {
	if r.RT != nil {
		return r.RT.Error()
	}
	return "no result data"
}

type ReturnType struct {
	Code string `json:"code,omitempty" codec:"code,omitempty"`
	Desc string `json:"desc,omitempty" codec:"desc,omitempty"`
	Data string `json:"data,omitempty" codec:"data,omitempty"`
}

func (r ReturnType) IsError() bool {
	return r.Code != "200" || r.Data == ""
}

func (r ReturnType) Error() string {
	if r.Code != "200" {
		return fmt.Sprintf("code:%s, msg:%s", r.Code, r.Desc)
	}
	return "no result data"
}

type ReturnTypeData struct {
	Header *ReturnTypeDataHeader `json:"header,omitempty" codec:"header,omitempty"`
	Body   *ReturnTypeDataBody   `json:"body,omitempty" codec:"body,omitempty"`
}

func (r ReturnTypeData) IsError() bool {
	return r.Header == nil || r.Header.IsError() || r.Body == nil || r.Body.IsError()
}

func (r ReturnTypeData) Error() string {
	if r.Header != nil {
		return r.Header.Error()
	}
	if r.Body != nil {
		return r.Body.Error()
	}
	return "missing header/body"
}

type ReturnTypeDataHeader struct {
	Code       string `json:"code,omitempty" codec:"code,omitempty"`
	DataStatus string `json:"dataStatus,omitempty" codec:"dataStatus,omitempty"`
	Desc       string `json:"desc,omitempty" codec:"desc,omitempty"`
}

func (r ReturnTypeDataHeader) IsError() bool {
	return r.Code != "200"
}

func (r ReturnTypeDataHeader) Error() string {
	return fmt.Sprintf("code:%s, msg:%s", r.Code, r.Desc)
}

type ReturnTypeDataBody struct {
	Data [][]string `json:"data,omitempty" codec:"code,omitempty"`
	Size uint       `json:"size,omitempty" codec:"code,omitempty"`
	Meta *Meta      `json:"meta,omitempty" codec:"code,omitempty"`
}

func (r ReturnTypeDataBody) IsError() bool {
	return r.Meta == nil
}

func (r ReturnTypeDataBody) Error() string {
	return "no return type data body meta"
}

type Meta struct {
	MetaIndex map[string]int `json:"metaIndex" codec:"metaIndex"`
}

// 数据数据开放接口
func BizStreamFetch(req *BizStreamFetchRequest) ([]map[string]interface{}, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := eco.NewBizStreamFetchRequest()
	r.SetServiceId(req.ServiceId)
	if req.TimeMin != "" {
		r.SetTimeMin(req.TimeMin)
	}
	if req.TimeMax != "" {
		r.SetTimeMax(req.TimeMax)
	}
	if req.Time != "" {
		r.SetTime(req.Time)
	}
	if req.Timestamp != "" {
		r.SetTimestamp(req.Timestamp)
	}
	if req.AdProStat != "" {
		r.SetAdProStat(req.AdProStat)
	}
	if req.AdProId != "" {
		r.SetAdProId(req.AdProId)
	}
	if req.Sku != "" {
		r.SetSku(req.Sku)
	}
	if req.AdType != "" {
		r.SetAdType(req.AdType)
	}
	if req.ActEffectId != "" {
		r.SetActEffectId(req.ActEffectId)
	}
	if req.TicketType != "" {
		r.SetTicketType(req.TicketType)
	}
	var response BizStreamFetchResponse
	if err := client.Execute(r.Request, req.Session, &response); err != nil {
		return nil, err
	}
	var returnTypeData ReturnTypeData
	if err := client.Logger().DecodeJSON([]byte(response.Res.RT.Data), &returnTypeData); err != nil {
		return nil, err
	}
	if returnTypeData.IsError() {
		return nil, returnTypeData
	}
	metaIndex := make(map[int]string, len(returnTypeData.Body.Meta.MetaIndex))
	for key, index := range returnTypeData.Body.Meta.MetaIndex {
		metaIndex[index] = key
	}
	metaSize := len(metaIndex)

	rs := make([]map[string]interface{}, len(returnTypeData.Body.Data))
	for _, data := range returnTypeData.Body.Data {
		dataMap := make(map[string]interface{})
		for i := 0; i < metaSize; i++ {
			dataMap[metaIndex[i]] = data[i]
		}
		rs = append(rs, dataMap)
	}

	return rs, nil
}
