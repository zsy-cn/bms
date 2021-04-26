package model

// Customer 客户表
type Customer struct {
	Base
	Name    string `json:"name" gorm:"unique;not null"`     // 登录用户名
	Passwd1 string `json:"passwd1" gorm:"size:50;not null"` // 登录密码
	Passwd2 string `json:"passwd2" gorm:"size:50;not null"` // 管理密码
	Title   string `json:"title" gorm:"size:100;not null"`  // 客户Title
	Address string `json:"address" grom:"size:100"`
	Path    string `json:"path" gorm:"size:50;unique; not null"` // 用户访问路径, 一般是客户公司的拼音或简写
	Enable  bool   `json:"enable" gorm:"default:true"`           // 是否可用

	// 将客户信息添加到lora-app-server时返回组织ID和客户ID, 都要记录一下, 方便日后修改和删除
	LoraUserID int64 `gorm:"not null;unique"`
	LoraOrgID  int64 `gorm:"not null;unique"`
}

// CustomerJSON ...
type CustomerJSON struct {
	ID         uint64     `json:"id"`
	Name       string     `json:"name"`    // 登录用户名
	Passwd1    string     `json:"passwd1"` // 登录密码(只在请求中带有, 响应操作不返回)
	Passwd2    string     `json:"passwd2"` // 管理密码
	Title      string     `json:"title"`   // 客户Title
	Address    string     `json:"address"`
	Path       string     `json:"path"`   // 用户访问路径
	Enable     bool       `json:"enable"` // 是否可用
	Contacts   []*Contact `json:"contacts"`
	ContactIds []uint64   `json:"contactIds"`
	CreateAt   string     `json:"createAt"`
	UpdateAt   string     `json:"updateAt"`
}
