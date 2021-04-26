package pagination

import (
	"fmt"

	"github.com/zsy-cn/bms/protos"
)

// BuildPaginationQueryInString 根据分页参数构建gorm查询语句
func BuildPaginationQueryInString(sqlStr string, queryParams *protos.Pagination) string {
	// 默认按主键占用率排序
	orderStr := " order by "
	if queryParams.SortBy != "" {
		orderStr = orderStr + queryParams.SortBy
	} else {
		orderStr = orderStr + " main_tbl.id "
	}
	// 默认降序排序
	if queryParams.Order != false {
		orderStr = orderStr + " ASC "
	} else {
		orderStr = orderStr + " DESC "
	}
	sqlStr = sqlStr + orderStr

	// 貌似offset, limit要放在order后面, limit要放在offset后面?
	if queryParams.Page == 0 {
		queryParams.Page = 1
	}
	sqlStr = fmt.Sprintf(sqlStr+" offset %d ", (queryParams.Page-1)*queryParams.PageSize)
	if queryParams.PageSize == 0 {
		queryParams.PageSize = 10
	}
	sqlStr = fmt.Sprintf(sqlStr+" limit %d ", queryParams.PageSize)

	return sqlStr
}
