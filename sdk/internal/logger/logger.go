package logger

import "bytes"

type Logger interface {
	DebugPrintError(err error)
	DebugPrintStringResponse(str string)
	DebugPrintGetRequest(url string)
	DebugPrintPostJSONRequest(url string, body []byte)
	DebugPrintPostMultipartRequest(url string, body []byte)
	DecodeJSON(body []byte, v interface{}) error
}

var (
	Default = new(DefaultLogger)
	Debug   = new(DebugLogger)
)

func removeJsonSpace(data []byte) []byte {
	rs := bytes.Replace(data, []byte("\n"), []byte(""), -1)
	rs = bytes.Replace(rs, []byte("\t"), []byte(""), -1)
	return rs
}
