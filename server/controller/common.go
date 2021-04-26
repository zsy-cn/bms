package controller

import (
	"os"

	"github.com/henrylee2cn/faygo"

	"github.com/zsy-cn/bms/util/log"
)

var logger = log.NewLogger(os.Stdout)

// Pagination 列表查询操作中的分页参数
type Pagination struct {
	Page     uint64 `param:"<in:query> <name:page>"`
	PageSize uint64 `param:"<in:query> <name:pageSize>"`
	SortBy   string `param:"<in:query> <name:sortBy>"`
	Order    bool   `param:"<in:query> <name:order>"`
}

// EmptyHandler ...
type EmptyHandler struct{}

// Serve ...
func (i *EmptyHandler) Serve(ctx *faygo.Context) (err error) {
	result := NewResult()
	defer ctx.JSON(200, result, true)
	return
}

// Result http响应体
type Result struct {
	Code int         `json:"code"` // return code, 0 for success
	Msg  string      `json:"msg"`  // message, "" while success
	Data interface{} `json:"data"` // data object
}

// NewResult ...
func NewResult() *Result {
	return &Result{
		Code: 0,
		Msg:  "",
		Data: nil,
	}
}
