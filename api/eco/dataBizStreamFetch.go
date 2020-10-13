package eco

import (
	"errors"
	"fmt"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/eco"
	"github.com/daviddengcn/ljson"
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
}

type BizStreamFetchResponse struct {
	ErrorResp *api.ErrorResponnse `json:"error_response,omitempty" codec:"error_response,omitempty"`
	Res       *Response           `json:"jingdong_eco_data_biz_stream_fetch_responce,omitempty" codec:"jingdong_eco_data_biz_stream_fetch_responce,omitempty"`
}

type Response struct {
	Code string      `json:"code,omitempty" codec:"code,omitempty"`
	RT   *ReturnType `json:"returnType,omitempty" codec:"returnType,omitempty"`
}

type ReturnType struct {
	Code string `json:"code,omitempty" codec:"code,omitempty"`
	Desc string `json:"desc,omitempty" codec:"desc,omitempty"`
	Data string `json:"data,omitempty" codec:"data,omitempty"`
}

type ReturnTypeData struct {
	Header *ReturnTypeDataHeader `json:"header,omitempty" codec:"header,omitempty"`
	Body   *ReturnTypeDataBody   `json:"body,omitempty" codec:"body,omitempty"`
}

type ReturnTypeDataHeader struct {
	Code       string `json:"code,omitempty" codec:"code,omitempty"`
	DataStatus string `json:"dataStatus,omitempty" codec:"dataStatus,omitempty"`
	Desc       string `json:"desc,omitempty" codec:"desc,omitempty"`
}

type ReturnTypeDataBody struct {
	Data [][]string `json:"data,omitempty" codec:"code,omitempty"`
	Size uint       `json:"size,omitempty" codec:"code,omitempty"`
	Meta *Meta      `json:"meta,omitempty" codec:"code,omitempty"`
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
	fmt.Println(fmt.Sprintf("%#v", req))
	result, err := client.Execute(r.Request, req.Session)
	fmt.Println(fmt.Sprintf("%s", result))
	if err != nil {
		return nil, err
	}
	if len(result) == 0 {
		return nil, errors.New("No data.")
	}

	var response BizStreamFetchResponse
	err = ljson.Unmarshal([]byte(result), &response)
	if err != nil {
		return nil, err
	}
	if response.ErrorResp != nil {
		return nil, response.ErrorResp
	}
	if response.Res == nil || response.Res.Code != "0" {
		return nil, errors.New("unknow error")
	}
	if response.Res.RT == nil || response.Res.RT.Code != "200" {
		return nil, errors.New(response.Res.RT.Desc)
	}

	if response.Res.RT.Data == "" {
		return nil, errors.New("no return type data")
	}

	var returnTypeData ReturnTypeData
	err = ljson.Unmarshal([]byte(response.Res.RT.Data), &returnTypeData)
	if err != nil {
		return nil, err
	}
	if returnTypeData.Header == nil {
		return nil, errors.New("no return type data header")
	}
	if returnTypeData.Header.Code != "200" {
		return nil, errors.New(returnTypeData.Header.Desc)
	}

	if returnTypeData.Body.Meta == nil {
		return nil, errors.New("no return type data body meta")
	}
	metaIndex := make(map[int]string)
	for key, index := range returnTypeData.Body.Meta.MetaIndex {
		metaIndex[index] = key
	}
	metaSize := len(metaIndex)

	rs := []map[string]interface{}{}
	for _, data := range returnTypeData.Body.Data {
		dataMap := make(map[string]interface{})
		for i := 0; i < metaSize; i++ {
			dataMap[metaIndex[i]] = data[i]
		}
		rs = append(rs, dataMap)
	}

	return rs, nil
}
