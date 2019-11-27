package secret

import (
	"crypto/sha1"
	"encoding/hex"
	"io"

	uuid "github.com/nu7hatch/gouuid"
)

type ReportAttribute struct {
	Type    ReportType  `json:"type"`
	Host    string      `json:"host,omitempty"`
	Level   ReportLevel `json:"level,omitempty"`
	Service string      `json:"service,omitempty"`
	SdkVer  uint        `json:"sdk_ver"`
	Env     string      `json:"env,omitempty"`
	Ts      int64       `json:"ts"`
	Code    string      `json:"code,omitempty"`
	Msg     string      `json:"msg,omitempty"`
	Heap    string      `json:"heap,omitempty"`
	Enccnt  uint        `json:"enccnt,omitempty"`
	deccnt  uint        `json:"deccnt,omitempty"`
}

type ReportStatistcAttribute struct {
	Type      ReportType  `json:"type"`
	Host      string      `json:"host,omitempty"`
	Level     ReportLevel `json:"level,omitempty"`
	Service   string      `json:"service,omitempty"`
	SdkVer    uint        `json:"sdk_ver"`
	Env       string      `json:"env,omitempty"`
	Ts        int64       `json:"ts"`
	Enccnt    uint        `json:"enccnt"`
	Deccnt    uint        `json:"deccnt"`
	Encerrcnt uint        `json:"encerrcnt"`
	Decerrcnt uint        `json:"decerrcnt"`
}

func NewBusinessId() string {
	token, _ := uuid.NewV4()
	h := sha1.New()
	io.WriteString(h, token.String())
	return hex.EncodeToString(h.Sum(nil))
}
