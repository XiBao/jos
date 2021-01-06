package secret

const (
	DefaultServerUrl = "https://api.jd.com/routerjson"
	CIPHER_LEN       = 2
)

type ReportText = string

const (
	INIT      ReportText = "INIT"
	EXCEPTION ReportText = "EXCEPTION"
	STATISTIC ReportText = "STATISTIC"
	EVENT     ReportText = "EVENT"
)

type ReportType = uint

const (
	INIT_TYPE      ReportType = 1
	EXCEPTION_TYPE ReportType = 2
	STATISTIC_TYPE ReportType = 3
	EVENT_TYPE     ReportType = 4
)

type ReportLevel = uint

const (
	INFO_LEVEL   ReportLevel = 1
	WARN_LEVEL   ReportLevel = 2
	ERROR_LEVEL  ReportLevel = 3
	SEVERE_LEVEL ReportLevel = 4
)
