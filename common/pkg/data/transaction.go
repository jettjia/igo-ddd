package data

import (
	"context"

	"gorm.io/gorm"
)

// Transaction 事务管理
type Transaction interface {
	ExecTx(context.Context, func(ctx context.Context) error) error
}

// 用来承载事务的上下文
type contextTxKey struct{}

// NewTransaction .
func NewTransaction(d *Data) Transaction {
	return d
}

// ExecTx gorm Transaction
func (d *Data) ExecTx(ctx context.Context, fn func(ctx context.Context) error) error {
	return d.Mysql.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		ctx = context.WithValue(ctx, contextTxKey{}, tx)
		return fn(ctx)
	})
}

// DB 根据此方法来判断当前的 db 是不是使用 事务的 DB
func (d *Data) DB(ctx context.Context) *gorm.DB {
	tx, ok := ctx.Value(contextTxKey{}).(*gorm.DB)
	if ok {
		return tx
	}
	return d.Mysql
}
