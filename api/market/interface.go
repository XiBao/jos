package market

type CartSku struct {
	Count      uint   `json:"count"`
	SourceType string `json:"sourcetype"`
	Time       uint64 `json:"time"`
	SkuId      uint64 `json:"sku"`
}
