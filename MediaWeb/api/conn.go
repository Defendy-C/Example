package api

import (
	"MediaWeb/common"
	"MediaWeb/tools/initformatter"
	"database/sql"
	"errors"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gomodule/redigo/redis"
	"os"
)

type dbConnInfo struct {
	user     string
	password string
	host     string
	dbname   string
	Conn     *sql.DB
}

type redisConnInfo struct {
	host string
	port string
	Conn redis.Conn
}

var (
	iniNotFindErr error = errors.New("don't find relative field in the ini file")
	QueryNilErr error = sql.ErrNoRows
	ConnNilErr  error = errors.New("this Conn is null")
	DBReadConn  *dbConnInfo
	DBWriteConn *dbConnInfo
	RedisReadConn *redisConnInfo
	RedisWriteConn *redisConnInfo

)

func newDBConn(pkg *initformatter.DataPkg)(ci *dbConnInfo, err error) {
	if pkg == nil {
		return nil, iniNotFindErr
	}

	ci = new(dbConnInfo)
	ci.user, err = pkg.GetAttrStr("user")
	ci.password, err = pkg.GetAttrStr("password")
	ci.host, err = pkg.GetAttrStr("host")
	ci.dbname, err = pkg.GetAttrStr("dbname")

	if err != nil {
		return nil, err
	}

	dsn := ci.user + ":" + ci.password + "@tcp(" + ci.host + ")/" + ci.dbname
	ci.Conn, err = sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	err = ci.Conn.Ping()
	if err != nil {
		return nil, err
	}

	return
}

func newRedisConn(pkg *initformatter.DataPkg)(ci *redisConnInfo, err error) {
	if pkg == nil {
		return nil, iniNotFindErr
	}

	ci = new(redisConnInfo)
	ci.host, err = pkg.GetAttrStr("host")
	ci.port, err = pkg.GetAttrStr("port")

	if err != nil {
		return nil, err
	} else {
		ci.Conn, err = redis.Dial("tcp", ci.host + ":" + ci.port)
		if err != nil {
			return nil ,err
		}
	}

	return
}

func (ci *dbConnInfo)Close() {
	ci.Conn.Close()
}

func init() {
	iniFilename := common.Path  + string(os.PathSeparator) + "dbconfig.ini"
	err := buildConn(iniFilename)
	if err != nil {
		panic(err)
	}
}

func buildConn(inifilename string)error {
	f, err := initformatter.New(inifilename)
	if err != nil {
		return err
	}

	if DBReadConn == nil {
		rcdp := f.GetDataPkgByName("DBReadConn")
		DBReadConn, err = newDBConn(rcdp)
	}

	if DBWriteConn == nil {
		wcdp := f.GetDataPkgByName("DBWriteConn")
		DBWriteConn, err = newDBConn(wcdp)
	}

	if RedisReadConn == nil {
		rcdp := f.GetDataPkgByName("RedisReadConn")
		RedisReadConn, err = newRedisConn(rcdp)
	}

	if RedisWriteConn == nil {
		wcdp := f.GetDataPkgByName("RedisWriteConn")
		RedisWriteConn, err = newRedisConn(wcdp)
	}

	return err
}
