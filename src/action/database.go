package action

import (
	"context"
	"log"

	"github.com/go-pg/pg/v9"
)

var DB *pg.DB

// QueryHook 查询钩子
type QueryHook struct{}

// BeforeQuery 查询前钩子
func (QueryHook) BeforeQuery(ctx context.Context, _ *pg.QueryEvent) (context.Context, error) {
	// 连接数据库
	if err := SetDatabase(); err != nil {
		log.Println(err.Error())
	}
	return ctx, nil
}

// AfterQuery 查询后钩子
func (QueryHook) AfterQuery(_ context.Context, qe *pg.QueryEvent) error {
	// 记录SQL语句
	stmt, err := qe.FormattedQuery()
	if err != nil {
		log.Println(err.Error())
	}

	log.Println("服务端执行SQL：", stmt)

	return nil
}

// SetDatabase 设置数据库
func SetDatabase() error {
	// 判断是否需要重启赋值
	if DB != nil {
		return nil
	}

	DB = pg.Connect(&pg.Options{
		Addr:     "127.0.0.1:5432",
		User:     "postgres",
		Password: "123456",
		Database: "mytest",
	})

	// 注册查询钩子
	DB.AddQueryHook(QueryHook{})

	return nil
}
