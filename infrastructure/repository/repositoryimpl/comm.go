package repositoryimpl

import (
	"fmt"
	"strconv"

	"gorm.io/gorm"

	"github.com/jettjia/go-ddd-demo/types"
)

// GenerateQueryCondition 组装 搜索
func GenerateQueryCondition(conditions []*types.Query) string {
	var condition string
	for k, v := range conditions {
		if k > 0 {
			condition += " and "
		}

		if v.Operator == types.Operator_LIKE {
			condition += fmt.Sprintf("%v%s'%%%v%%'", v.Key, types.OperatorMap[v.Operator], v.Value)
		} else if v.Operator == types.Operator_IN {
			condition += fmt.Sprintf(` %s %s (%s)`, v.Key, types.OperatorMap[v.Operator], v.Value)
		} else {
			//bool string int
			_, err := strconv.ParseBool(v.Value)
			if err != nil {
				condition += fmt.Sprintf("%v%s'%v'", v.Key, types.OperatorMap[v.Operator], v.Value)
			} else {
				condition += fmt.Sprintf("%v%s%v", v.Key, types.OperatorMap[v.Operator], v.Value)
			}
		}
	}

	return condition
}

// Paginate gorm 分页
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
