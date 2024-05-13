package jm

type PurchaseInfo struct {
	Pin         string `json:"pin"`           // 当前用户pin
	EndDate     int64  `json:"endDate"`       // 结束时间
	ItemCode    string `json:"itemCode"`      // 服务市场售卖服务的版本，itemCode可以通过服务市场后台获取
	VersionNo   int    `json:"versionNo"`     // 版本
	AppKey      string `json:"appKey"`        // appKey
	IsModule    int    `json:"isModule"`      // 1 按版本 2 按模块 （ 发布服务时计费方式为：按周期 计费中的 对应的 版本和模块计费 ）
	OpenIdBuyer string `json:"open_id_buyer"` // 当前用户pin
}
