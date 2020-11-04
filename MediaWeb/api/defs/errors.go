package defs

type Err struct {
	ErrCode int `json:"err_code"`
	ErrInfo string `json:"err_info"`
}

type HttpErr struct {
	StatusCode int
	Err
}

var (
	ErrNotPage = HttpErr{StatusCode:404, Err:Err{ErrInfo:"Don't find this page!", ErrCode:1}}
	ErrParseRequest = HttpErr{StatusCode:400, Err: Err{ErrInfo:"Request parse error!", ErrCode:2}}
	ErrTokenVerify = HttpErr{StatusCode:401, Err: Err{ErrInfo:"invalid token", ErrCode:4}}
)

func ErrDBOperator(err string)HttpErr {
	return HttpErr{StatusCode:500, Err: Err{ErrInfo:err, ErrCode:3}}
}
