package code

var (
	ErrUserExist             = newError(10000101, "用户名已存在: %s")
	ErrUserNotExist          = newError(10000102, "用户名不存在: %s")
	ErrUserTokenGenFaild     = newError(10000103, "生成用户token失败: %s")
	ErrUserTokenDelFaild     = newError(10000104, "清理用户token失败: %s: %s")
	ErrUserTokenSetFaild     = newError(10000105, "设置用户token失败: %s: %s")
	ErrUserTokenInvalid      = newError(10000106, "用户token无效: %s")
	ErrUserPassInvalid       = newError(10000107, "用户密码无效: %s")
	ErrUserTokenNotFound     = newError(10000108, "用户token不存在")
	ErrUserTokenRefreshFaild = newError(10000109, "刷新用户token失败: %s")
)
