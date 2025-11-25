package sequence

import (
	"database/sql"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/logx"
)


type Mysql struct {
	conn sqlx.SqlConn
}

const sqlReplaceIntoStub = `REPLACE INTO sequence (stub) VALUES ('a')`

func NewMysql(dsn string) *Mysql {
	conn := sqlx.NewMysql(dsn)
	return &Mysql{
		conn: conn,
	}
}

// 取下一个号
func (m *Mysql) Next() (seq uint64, err error){
	// 准备语句
	var stmt sqlx.StmtSession
	stmt, err = m.conn.Prepare(sqlReplaceIntoStub)
	if err != nil {
		logx.Error("conn.Perpare failed", logx.LogField{Key: "err", Value: err.Error()})
		return 0, err
	}
	defer stmt.Close()

	// 执行
	var rest sql.Result
	rest, err = stmt.Exec()
	if err != nil {
		logx.Error("stmt.Exec failed", logx.LogField{Key: "err", Value: err.Error()})
		return 0, err
	}

	// 获取插入ID
	var lid int64
	lid, err = rest.LastInsertId()
	if err != nil {
		logx.Error("rest.LastInsertId failed", logx.LogField{Key: "err", Value: err.Error()})
		return 0, err
	}

	return uint64(lid), nil
}