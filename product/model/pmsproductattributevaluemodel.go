package model

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/tal-tech/go-zero/core/stores/sqlc"
	"github.com/tal-tech/go-zero/core/stores/sqlx"
	"github.com/tal-tech/go-zero/core/stringx"
	"github.com/tal-tech/go-zero/tools/goctl/model/sql/builderx"
)

var (
	pmsProductAttributeValueFieldNames          = builderx.RawFieldNames(&PmsProductAttributeValue{})
	pmsProductAttributeValueRows                = strings.Join(pmsProductAttributeValueFieldNames, ",")
	pmsProductAttributeValueRowsExpectAutoSet   = strings.Join(stringx.Remove(pmsProductAttributeValueFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	pmsProductAttributeValueRowsWithPlaceHolder = strings.Join(stringx.Remove(pmsProductAttributeValueFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"
)

type (
	PmsProductAttributeValueModel interface {
		Insert(data PmsProductAttributeValue) (sql.Result, error)
		FindOne(id int64) (*PmsProductAttributeValue, error)
		Update(data PmsProductAttributeValue) error
		Delete(id int64) error
	}

	defaultPmsProductAttributeValueModel struct {
		conn  sqlx.SqlConn
		table string
	}

	PmsProductAttributeValue struct {
		Id                 int64          `db:"id"`
		ProductId          sql.NullInt64  `db:"product_id"`
		ProductAttributeId sql.NullInt64  `db:"product_attribute_id"`
		Value              sql.NullString `db:"value"` // 手动添加规格或参数的值，参数单值，规格有多个时以逗号隔开
	}
)

func NewPmsProductAttributeValueModel(conn sqlx.SqlConn) PmsProductAttributeValueModel {
	return &defaultPmsProductAttributeValueModel{
		conn:  conn,
		table: "`pms_product_attribute_value`",
	}
}

func (m *defaultPmsProductAttributeValueModel) Insert(data PmsProductAttributeValue) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?)", m.table, pmsProductAttributeValueRowsExpectAutoSet)
	ret, err := m.conn.Exec(query, data.ProductId, data.ProductAttributeId, data.Value)
	return ret, err
}

func (m *defaultPmsProductAttributeValueModel) FindOne(id int64) (*PmsProductAttributeValue, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", pmsProductAttributeValueRows, m.table)
	var resp PmsProductAttributeValue
	err := m.conn.QueryRow(&resp, query, id)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultPmsProductAttributeValueModel) Update(data PmsProductAttributeValue) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, pmsProductAttributeValueRowsWithPlaceHolder)
	_, err := m.conn.Exec(query, data.ProductId, data.ProductAttributeId, data.Value, data.Id)
	return err
}

func (m *defaultPmsProductAttributeValueModel) Delete(id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.Exec(query, id)
	return err
}
