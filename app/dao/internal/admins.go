// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
)

// AdminsDao is the manager for logic model data accessing and custom defined data operations functions management.
type AdminsDao struct {
	Table   string        // Table is the underlying table name of the DAO.
	Group   string        // Group is the database configuration group name of current DAO.
	Columns AdminsColumns // Columns is the short type for Columns, which contains all the column names of Table for convenient usage.
}

// AdminsColumns defines and stores column names for table admins.
type AdminsColumns struct {
	Id        string //
	Username  string // 用户名
	Password  string // 密码
	CreatedAt string //
	UpdatedAt string //
}

//  adminsColumns holds the columns for table admins.
var adminsColumns = AdminsColumns{
	Id:        "id",
	Username:  "username",
	Password:  "password",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// NewAdminsDao creates and returns a new DAO object for table data access.
func NewAdminsDao() *AdminsDao {
	return &AdminsDao{
		Group:   "default",
		Table:   "admins",
		Columns: adminsColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *AdminsDao) DB() gdb.DB {
	return g.DB(dao.Group)
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *AdminsDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.Table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *AdminsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx *gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
