// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"my-blog/app/dao/internal"
)

// categoriesDao is the manager for logic model data accessing and custom defined data operations functions management.
// You can define custom methods on it to extend its functionality as you wish.
type categoriesDao struct {
	*internal.CategoriesDao
}

var (
	// Category is globally public accessible object for table categories operations.
	Category categoriesDao
)

func init() {
	Category = categoriesDao{
		internal.NewCategoriesDao(),
	}
}

// Fill with you ideas below.
