package account

type AccountInfo struct {
	AccountId    uint64 `json:"accountId,omitempty" codec:"accountId,omitempty"`
	Status       uint8  `json:"status,omitempty" codec:"status,omitempty"`             // 账户状态
	MainNum      uint32 `json:"mainNum,omitempty" codec:"mainNum,omitempty"`           // 充值（主）账户京豆数量
	SubNum       uint32 `json:"subNum,omitempty" codec:"subNum,omitempty"`             // 充值（副）账户京豆数量
	Creator      string `json:"creator,omitempty" codec:"creator,omitempty"`           // 创建人
	Modified     uint64 `json:"modified,omitempty" codec:"modified,omitempty"`         // 最近一次修改时间
	AccountCode  string `json:"accountCode,omitempty" codec:"accountCode,omitempty"`   // 京豆账户,当accountType=1，则accountCode为商家id;当accountType=2，则accouontCode为供应商简码
	Created      uint64 `json:"created,omitempty" codec:"created,omitempty"`           // 创建时间
	FreezeNum    uint32 `json:"freezeNum,omitempty" codec:"freezeNum,omitempty"`       // 冻结京豆数量
	AccountType  uint8  `json:"accountType,omitempty" codec:"accountType,omitempty"`   // 京豆账户类型,1:商家账户,2:供应商账户,3:品牌商账户
	AvailableNum uint32 `json:"availableNum,omitempty" codec:"availableNum,omitempty"` // 可用京豆数量
}
