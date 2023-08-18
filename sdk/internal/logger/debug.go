package logger

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
)

type DebugLogger struct{}

func (l DebugLogger) DebugPrintError(err error) {
	log.Println("[JOS_DEBUG] [ERROR]", err)
}

func (l DebugLogger) DebugPrintStringResponse(str string) {
	log.Println("[JOS_DEBUG] [RESPONSE]", str)
}

func (l DebugLogger) DebugPrintGetRequest(url string) {
	log.Println("[JOS_DEBUG] [API] GET", url)
}

func (l DebugLogger) DebugPrintPostJSONRequest(url string, body []byte) {
	const format = "[JOS_DEBUG] [API] JSON POST %s\n" +
		"http request body:\n%s\n"

	buf := bytes.NewBuffer(make([]byte, 0, len(body)+1024))
	if err := json.Indent(buf, body, "", "    "); err == nil {
		body = buf.Bytes()
	}
	log.Printf(format, url, body)
}

func (l DebugLogger) DebugPrintPostMultipartRequest(url string, body []byte) {
	log.Println("[JOS_DEBUG] [API] multipart/form-data POST", url)
}

func (l DebugLogger) DecodeJSONHttpResponse(r io.Reader, v interface{}) error {
	body, err := io.ReadAll(r)
	if err != nil {
		return err
	}

	body2 := body
	buf := bytes.NewBuffer(make([]byte, 0, len(body2)+1024))
	if err := json.Indent(buf, body2, "", "    "); err == nil {
		body2 = buf.Bytes()
	}
	log.Printf("[JOS_DEBUG] [API] http response body:\n%s\n", body2)

	return json.Unmarshal(body, v)
}
