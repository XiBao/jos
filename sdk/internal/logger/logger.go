package logger

import "io"

type Logger interface {
	DebugPrintError(err error)
	DebugPrintStringResponse(str string)
	DebugPrintGetRequest(url string)
	DebugPrintPostJSONRequest(url string, body []byte)
	DebugPrintPostMultipartRequest(url string, body []byte)
	DecodeJSONHttpResponse(r io.Reader, v interface{}) error
}

var (
	Default = new(DefaultLogger)
	Debug   = new(DebugLogger)
)
