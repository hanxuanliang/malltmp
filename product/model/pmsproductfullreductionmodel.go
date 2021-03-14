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
	pmsProductFullReductionFieldNames          = builderx.RawFieldNames(&PmsProductFullReduction{})
	pmsProductFullReductionRows                = strings.Join(pmsProductFullReductionFieldNames, ",")
	pmsProductFullReductionRowsExpectAutoSet   = strings.Join(stringx.Remove(pmsProductFullReductionFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	pmsProductFullReductionRowsWithPlaceHolder = strings.Join(stringx.Remove(pmsProductFullReductionFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"
)

type (
	PmsProductFullReductionModel interface {
		Insert(data PmsProductFullReduction) (sql.Result, error)
		FindOne(id int64) (*PmsProductFullReduction, error)
		Update(data PmsProductFullReduction) error
		Delete(id int64) error
	}

	defaultPmsProductFullReductionModel struct {
		conn  sqlx.SqlConn
		table string
	}

	PmsProductFullReduction struct {
		ProductId   sql.NullInt64   `db:"product_id"`
		FullPrice   sql.NullFloat64 `db:"full_price"`
		ReducePrice sql.NullFloat64 `db:"reduce_price"`
		Id          int64           `db:"id"`
	}
)

func NewPmsProductFullReductionModel(conn sqlx.SqlConn) PmsProductFullReductionModel {
	return &defaultPmsProductFullReductionModel{
		conn:  conn,
		table: "`pms_product_full_reduction`",
	}
}

func (m *defaultPmsProductFullReductionModel) Insert(data PmsProductFullReduction) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?)", m.table, pmsProductFullReductionRowsExpectAutoSet)
	ret, err := m.conn.Exec(query, data.ProductId, data.FullPrice, data.ReducePrice)
	return ret, err
}

func (m *defaultPmsProductFullReductionModel) FindOne(id int64) (*PmsProductFullReduction, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", pmsProductFullReductionRows, m.table)
	var resp PmsProductFullReduction
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

func (m *defaultPmsProductFullReductionModel) Update(data PmsProductFullReduction) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, pmsProductFullReductionRowsWithPlaceHolder)
	_, err := m.conn.Exec(query, data.ProductId, data.FullPrice, data.ReducePrice, data.Id)
	return err
}

func (m *defaultPmsProductFullReductionModel) Delete(id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.Exec(query, id)
	return err
}
