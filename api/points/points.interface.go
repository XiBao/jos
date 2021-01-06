package points

type PointsRule struct {
	RuleId   string  `json:"ruleId"`
	VenderId int64   `json:"venderId"`
	Points   int64   `json:"points"`
	Type     int8    `json:"type"`
	Multiple float64 `json:"multiple"`
}
