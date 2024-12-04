package code

var codeMap = map[int]*AppError{}

// 通用错误
var (
	ErrSuccess        = newError(0, "success")
	ErrServer         = newError(10000000, "服务器内部错误")
	ErrDBUnknow       = newError(10000001, "数据库未知异常")
	ErrParams         = newError(10000002, "参数错误: %s")
	ErrTooManyRequest = newError(10000003, "请求太频繁: %s")
)
