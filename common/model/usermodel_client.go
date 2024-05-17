// Code generated by goctl. DO NOT EDIT.

package model

import (
	"context"
	"fmt"
    "strings"
	"time"

    "github.com/openui-backend-go/common/database"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	userFieldNames          = builder.RawFieldNames(&User{})
	userRows                = strings.Join(userFieldNames, ",")
	userRowsExpectAutoSet   = strings.Join(stringx.Remove(userFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	userRowsWithPlaceHolder = strings.Join(stringx.Remove(userFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cacheUserIdPrefix     = "cache:user:id:"
	cacheUserEmailPrefix = "cache:user:email:"
)

type (
	userModel interface {
		Insert(ctx context.Context, data *User) error
		FindOne(ctx context.Context, id int64) (*User, error)
		FindOneByEmail(ctx context.Context, email string) (*User, error)
		Update(ctx context.Context, data *User) error
		Delete(ctx context.Context, id int64) error
	}

	defaultUserModel struct {
		db *database.GormDao
        cache *database.DcRedisClient
		table string
	}

	User struct {
		Id         int64     `db:"id"`
		Name       string    `db:"name"`     // 用户姓名
		Email string `db:"email"`
		Password   string    `db:"password"` // 用户密码
		CreateTime time.Time `db:"create_time"`
		UpdateTime time.Time `db:"update_time"`
	}
)

func newUserModel(conn *database.GormDao, c *database.DcRedisClient) *defaultUserModel {
    // 创建缓存
	return &defaultUserModel{
		db:   conn,
        cache: c,
		table: "user",
	}
}

func (m *defaultUserModel) Delete(ctx context.Context, id int64) error {
    user := &User{Id: id}
    err := m.db.Delete(ctx, m.table, user, false)
    if err != nil {
        return err
    }
	return nil
}

func (m *defaultUserModel) FindOne(ctx context.Context, id int64) (*User, error) {
    user := User{}
    sql := "id = ?"
    err := m.db.First(ctx, m.table, &user, sql, id)
    if err != nil {
        return nil, err
    }
	return &user, nil
}

func (m *defaultUserModel) FindOneByEmail(ctx context.Context, email string) (*User, error) {
    user := User{}
    sql := "email = ?"
    err := m.db.First(ctx, m.table, &user, sql, email)
    if err != nil {
        return nil, err
    }
	return &user, nil
}

func (m *defaultUserModel) Insert(ctx context.Context, data *User) error{
    err := m.db.Create(ctx, m.table, data)
    if err != nil {
        return err
    }
	return nil
}

func (m *defaultUserModel) Update(ctx context.Context, newData *User) error {
    err := m.db.Update(ctx, m.table, uint(newData.Id), newData)
    if err != nil {
        return err
    }
	return nil
}

func (m *defaultUserModel) formatPrimary(primary any) string {
	return fmt.Sprintf("%s%v", cacheUserIdPrefix, primary)
}

func (m *defaultUserModel) tableName() string {
	return m.table
}
