package database

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type MySQLFilter struct {
	Where gin.H
	Not   gin.H
	Or    []gin.H
}

func BuildMySQLFilter(dbMySQL *gorm.DB, filter MySQLFilter) *gorm.DB {
	if len(filter.Where) > 0 {
		for k, v := range filter.Where {
			dbMySQL = dbMySQL.Where(k, v)
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

	return dbMySQL
}
