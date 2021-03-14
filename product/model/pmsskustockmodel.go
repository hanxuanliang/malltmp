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
	pmsSkuStockFieldNames          = builderx.RawFieldNames(&PmsSkuStock{})
	pmsSkuStockRows                = strings.Join(pmsSkuStockFieldNames, ",")
	pmsSkuStockRowsExpectAutoSet   = strings.Join(stringx.Remove(pmsSkuStockFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	pmsSkuStockRowsWithPlaceHolder = strings.Join(stringx.Remove(pmsSkuStockFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"
)

type (
	PmsSkuStockModel interface {
		Insert(data PmsSkuStock) (sql.Result, error)
		FindOne(id int64) (*PmsSkuStock, error)
		Update(data PmsSkuStock) error
		Delete(id int64) error
	}

	defaultPmsSkuStockModel struct {
		conn  sqlx.SqlConn
		table string
	}

	PmsSkuStock struct {
		Id             int64           `db:"id"`
		ProductId      sql.NullInt64   `db:"product_id"`
		LowStock       sql.NullInt64   `db:"low_stock"`       // 预警库存
		Pic            sql.NullString  `db:"pic"`             // 展示图片
		Sale           sql.NullInt64   `db:"sale"`            // 销量
		PromotionPrice sql.NullFloat64 `db:"promotion_price"` // 单品促销价格
		LockStock      int64           `db:"lock_stock"`      // 锁定库存
		SpData         sql.NullString  `db:"sp_data"`         // 商品销售属性，json格式
		SkuCode        string          `db:"sku_code"`        // sku编码
		Price          sql.NullFloat64 `db:"price"`
		Stock          int64           `db:"stock"` // 库存
	}
)

func NewPmsSkuStockModel(conn sqlx.SqlConn) PmsSkuStockModel {
	return &defaultPmsSkuStockModel{
		conn:  conn,
		table: "`pms_sku_stock`",
	}
}

func (m *defaultPmsSkuStockModel) Insert(data PmsSkuStock) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, pmsSkuStockRowsExpectAutoSet)
	ret, err := m.conn.Exec(query, data.ProductId, data.LowStock, data.Pic, data.Sale, data.PromotionPrice, data.LockStock, data.SpData, data.SkuCode, data.Price, data.Stock)
	return ret, err
}

func (m *defaultPmsSkuStockModel) FindOne(id int64) (*PmsSkuStock, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", pmsSkuStockRows, m.table)
	var resp PmsSkuStock
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

func (m *defaultPmsSkuStockModel) Update(data PmsSkuStock) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, pmsSkuStockRowsWithPlaceHolder)
	_, err := m.conn.Exec(query, data.ProductId, data.LowStock, data.Pic, data.Sale, data.PromotionPrice, data.LockStock, data.SpData, data.SkuCode, data.Price, data.Stock, data.Id)
	return err
}

func (m *defaultPmsSkuStockModel) Delete(id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.Exec(query, id)
	return err
}
