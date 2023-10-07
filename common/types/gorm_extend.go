package types

import (
	"fmt"
	"strconv"

	"gorm.io/gorm"
)

// GenerateQueryCondition assembly search, currently only supports the and method
func GenerateQueryCondition(conditions []*Query) string {
	var condition string
	for k, v := range conditions {
		if k > 0 {
			condition += " and "
		}

		if v.Operator == Operator_LIKE {
			condition += fmt.Sprintf("%v%s'%%%v%%'", v.Key, OperatorMap[v.Operator], v.Value)
		} else if v.Operator == Operator_IN {
			condition += fmt.Sprintf(` %s %s (%s)`, v.Key, OperatorMap[v.Operator], v.Value)
		} else {
			//bool string int
			_, err := strconv.ParseBool(v.Value)
			if err != nil {
				condition += fmt.Sprintf("%v%s'%v'", v.Key, OperatorMap[v.Operator], v.Value)
			} else {
				condition += fmt.Sprintf("%v%s%v", v.Key, OperatorMap[v.Operator], v.Value)
			}
		}
	}

	return condition
}

// Paginate gorm page
func Paginate(page, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page == 0 {
			page = 1
		}

		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}

		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}
