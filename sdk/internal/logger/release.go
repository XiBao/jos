package logger

import (
	"encoding/json"
	"io"
)

type DefaultLogger struct{}

func (l DefaultLogger) DebugPrintError(err error) {}

func (l DefaultLogger) DebugPrintStringResponse(str string) {}

func (l DefaultLogger) DebugPrintGetRequest(url string) {}

func (l DefaultLogger) DebugPrintPostJSONRequest(url string, body []byte) {}

func (l DefaultLogger) DebugPrintPostMultipartRequest(url string, body []byte) {}

func (l DefaultLogger) DecodeJSONHttpResponse(r io.Reader, v interface{}) error {
	return json.NewDecoder(r).Decode(v)
}
