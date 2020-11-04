package dbopts

import (
	"MediaWeb/api"
	"database/sql"
	"fmt"
)

func checkReadConn()*sql.DB {
	if api.DBReadConn == nil {
		return nil
	}

	return api.DBReadConn.Conn
}

func exec(sql string, vals ...interface{})error {
	if api.DBWriteConn == nil {
		return api.ConnNilErr
	}

	conn := api.DBWriteConn.Conn
	stmt, err := conn.Prepare(sql)
	defer stmt.Close()

	if err != nil {
		return err
	}
	res, err := stmt.Exec(vals...)
	if err != nil {
		return err
	}

	n, err := res.RowsAffected()

	if err != nil {
		return err
	}

	if n <=0 {
		err = fmt.Errorf("sql exec is invaild!")
		return err
	}

	return nil
}

func QueryOne(sql string, vals ...interface{})(r *sql.Row, err error) {
	conn := checkReadConn()
	if conn == nil {
		err = api.ConnNilErr
		return
	}

	stmt, err := conn.Prepare(sql)
	defer stmt.Close()
	if err != nil {
		return
	}

	r = stmt.QueryRow(vals...)
return
}

func QueryMany(sql string, vals ...interface{})(rs *sql.Rows, err error) {
	conn := checkReadConn()
	if conn == nil {
		err = api.ConnNilErr
		return
	}

	stmt, err := conn.Prepare(sql)
	defer stmt.Close()
	if err != nil {
		return
	}

	rs, err = stmt.Query(vals...)
	return
}