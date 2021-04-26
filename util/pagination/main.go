package pagination

import (
	"github.com/jinzhu/gorm"

	"github.com/zsy-cn/bms/protos"
)

// BuildPaginationQuery 根据分页参数构建gorm查询语句
func BuildPaginationQuery(query *gorm.DB, queryParams *protos.Pagination) *gorm.DB {
	if queryParams.PageSize == 0 {
		queryParams.PageSize = 10
	}
	query = query.Limit(queryParams.PageSize)

	if queryParams.Page == 0 {
		queryParams.Page = 1
	}
	query = query.Offset((queryParams.Page - 1) * queryParams.PageSize)

	// 默认按主键id排序
	orderStr := ""
	if queryParams.SortBy != "" {
		orderStr = queryParams.SortBy
	} else {
		orderStr = "id"
	}
	// 默认降序排序
	if queryParams.Order != false {
		orderStr = orderStr + " ASC"
	} else {
		orderStr = orderStr + " DESC"
	}
	query = query.Order(orderStr)
	return query
}
