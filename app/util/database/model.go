package database

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type MySQLFilter struct {
	Where   gin.H
	Like    gin.H
	Not     gin.H
	Or      gin.H
	Preload []string
	Limit   int64
	Offset  int64
	Order   string
	Sort    string
}

func BuildMySQLFilter(dbMySQL *gorm.DB, filter MySQLFilter) *gorm.DB {
	if len(filter.Where) > 0 {
		for k, v := range filter.Where {
			dbMySQL = dbMySQL.Where(k, v)
		}
	}

	if len(filter.Like) > 0 {
		for k, v := range filter.Like {
			dbMySQL = dbMySQL.Where(fmt.Sprintf("%s LIKE ?", k), fmt.Sprintf("%%%s%%", v))
		}
	}

	if len(filter.Not) > 0 {
		for k, v := range filter.Not {
			dbMySQL = dbMySQL.Not(k, v)
		}
	}

	if len(filter.Or) > 0 {
		for _, cond := range filter.Or {
			dbMySQL = dbMySQL.Or(cond)
		}
	}

	if len(filter.Preload) > 0 {
		for _, v := range filter.Preload {
			dbMySQL = dbMySQL.Preload(v)
		}
	}

	if filter.Limit > 0 {
		dbMySQL = dbMySQL.Limit(int(filter.Limit))
	}

	if filter.Offset > 0 {
		dbMySQL = dbMySQL.Offset(int(filter.Offset))
	}

	if filter.Order != "" && filter.Sort != "" {
		dbMySQL = dbMySQL.Order(fmt.Sprintf("%s %s", filter.Order, filter.Sort))
	}

	return dbMySQL
}
