package auth

import "time"

// CustomerInfo session中客户登录信息数据
type CustomerInfo struct {
	ID       uint64    // 客户ID
	Name     string    // 客户名称
	ReAuth   bool      // 双重认证
	ReAuthAt time.Time // 双重认证的时间, 用于中间件验证过期与否
}

// LoginParams ...
type LoginParams struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// ReAuthParams ...
type ReAuthParams struct {
	ID       uint64 `json:"id"`
	Password string `json:"password"` // 第二个密码
}
