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
	pmsBrandFieldNames          = builderx.RawFieldNames(&PmsBrand{})
	pmsBrandRows                = strings.Join(pmsBrandFieldNames, ",")
	pmsBrandRowsExpectAutoSet   = strings.Join(stringx.Remove(pmsBrandFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	pmsBrandRowsWithPlaceHolder = strings.Join(stringx.Remove(pmsBrandFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"
)

type (
	PmsBrandModel interface {
		Insert(data PmsBrand) (sql.Result, error)
		FindOne(id int64) (*PmsBrand, error)
		Update(data PmsBrand) error
		Delete(id int64) error
	}

	defaultPmsBrandModel struct {
		conn  sqlx.SqlConn
		table string
	}

	PmsBrand struct {
		Sort                sql.NullInt64  `db:"sort"`
		ShowStatus          sql.NullInt64  `db:"show_status"`
		ProductCount        sql.NullInt64  `db:"product_count"` // 产品数量
		Logo                sql.NullString `db:"logo"`          // 品牌logo
		BrandStory          string         `db:"brand_story"`   // 品牌故事
		Name                sql.NullString `db:"name"`
		FirstLetter         sql.NullString `db:"first_letter"`          // 首字母
		ProductCommentCount sql.NullInt64  `db:"product_comment_count"` // 产品评论数量
		BigPic              sql.NullString `db:"big_pic"`               // 专区大图
		Id                  int64          `db:"id"`
		FactoryStatus       sql.NullInt64  `db:"factory_status"` // 是否为品牌制造商：0->不是；1->是
	}
)

func NewPmsBrandModel(conn sqlx.SqlConn) PmsBrandModel {
	return &defaultPmsBrandModel{
		conn:  conn,
		table: "`pms_brand`",
	}
}

func (m *defaultPmsBrandModel) Insert(data PmsBrand) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, pmsBrandRowsExpectAutoSet)
	ret, err := m.conn.Exec(query, data.Sort, data.ShowStatus, data.ProductCount, data.Logo, data.BrandStory, data.Name, data.FirstLetter, data.ProductCommentCount, data.BigPic, data.FactoryStatus)
	return ret, err
}

func (m *defaultPmsBrandModel) FindOne(id int64) (*PmsBrand, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", pmsBrandRows, m.table)
	var resp PmsBrand
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

func (m *defaultPmsBrandModel) Update(data PmsBrand) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, pmsBrandRowsWithPlaceHolder)
	_, err := m.conn.Exec(query, data.Sort, data.ShowStatus, data.ProductCount, data.Logo, data.BrandStory, data.Name, data.FirstLetter, data.ProductCommentCount, data.BigPic, data.FactoryStatus, data.Id)
	return err
}

func (m *defaultPmsBrandModel) Delete(id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.Exec(query, id)
	return err
}
