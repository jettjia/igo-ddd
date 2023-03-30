package types

// Query 搜索
type Query struct {
	Key      string   `json:"key"`      // 搜索关键词的键
	Value    string   `json:"value"`    // 搜索关键词的值
	Operator Operator `json:"operator"` // 判断条件
}

type Operator int32

const (
	Operator_GT    Operator = 0 //大于
	Operator_EQUAL Operator = 1 //等于
	Operator_LT    Operator = 2 //小于
	Operator_NEQ   Operator = 3 //不等于
	Operator_LIKE  Operator = 4 //模糊查询
	Operator_GTE   Operator = 5 // 大于等于
	Operator_LTE   Operator = 6 // 小于等于
	Operator_IN    Operator = 7 // in
)

// Enum value maps for Operator.
var (
	Operator_name = map[int32]string{
		0: "GT",
		1: "EQUAL",
		2: "LT",
		3: "NEQ",
		4: "LIKE",
		5: "GTE",
		6: "LTE",
		7: "IN",
	}
	Operator_value = map[string]int32{
		"GT":    0,
		"EQUAL": 1,
		"LT":    2,
		"NEQ":   3,
		"LIKE":  4,
		"GTE":   5,
		"LTE":   6,
		"IN":    7,
	}
)

var OperatorMap = map[Operator]string{
	Operator_GT:    " > ",
	Operator_EQUAL: " = ",
	Operator_LT:    " < ",
	Operator_NEQ:   " != ",
	Operator_LIKE:  " like ",
	Operator_GTE:   " >= ",
	Operator_LTE:   " <= ",
	Operator_IN:    " in ",
}
