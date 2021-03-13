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
	pmsProductFieldNames          = builderx.RawFieldNames(&PmsProduct{})
	pmsProductRows                = strings.Join(pmsProductFieldNames, ",")
	pmsProductRowsExpectAutoSet   = strings.Join(stringx.Remove(pmsProductFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	pmsProductRowsWithPlaceHolder = strings.Join(stringx.Remove(pmsProductFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"
)

type (
	PmsProductModel interface {
		Insert(data PmsProduct) (sql.Result, error)
		FindOne(id int64) (*PmsProduct, error)
		Update(data PmsProduct) error
		Delete(id int64) error
	}

	defaultPmsProductModel struct {
		conn  sqlx.SqlConn
		table string
	}

	PmsProduct struct {
		BrandId                    sql.NullInt64   `db:"brand_id"`
		ProductSn                  string          `db:"product_sn"`         // 货号
		Description                string          `db:"description"`        // 商品描述
		Weight                     sql.NullFloat64 `db:"weight"`             // 商品重量，默认为克
		BrandName                  sql.NullString  `db:"brand_name"`         // 品牌名称
		UsePointLimit              sql.NullInt64   `db:"use_point_limit"`    // 限制使用的积分数
		ServiceIds                 sql.NullString  `db:"service_ids"`        // 以逗号分割的产品服务：1->无忧退货；2->快速退款；3->免费包邮
		AlbumPics                  sql.NullString  `db:"album_pics"`         // 画册图片，连产品图片限制为5张，以逗号分割
		DetailHtml                 string          `db:"detail_html"`        // 产品详情网页内容
		DetailMobileHtml           string          `db:"detail_mobile_html"` // 移动端网页详情
		Pic                        sql.NullString  `db:"pic"`
		RecommandStatus            sql.NullInt64   `db:"recommand_status"` // 推荐状态；0->不推荐；1->推荐
		Price                      sql.NullFloat64 `db:"price"`
		PromotionPrice             sql.NullFloat64 `db:"promotion_price"` // 促销价格
		Stock                      sql.NullInt64   `db:"stock"`           // 库存
		Keywords                   sql.NullString  `db:"keywords"`
		PromotionPerLimit          sql.NullInt64   `db:"promotion_per_limit"` // 活动限购数量
		Sale                       sql.NullInt64   `db:"sale"`                // 销量
		FeightTemplateId           sql.NullInt64   `db:"feight_template_id"`
		ProductAttributeCategoryId sql.NullInt64   `db:"product_attribute_category_id"`
		Name                       string          `db:"name"`
		PublishStatus              sql.NullInt64   `db:"publish_status"` // 上架状态：0->下架；1->上架
		VerifyStatus               sql.NullInt64   `db:"verify_status"`  // 审核状态：0->未审核；1->审核通过
		Sort                       sql.NullInt64   `db:"sort"`           // 排序
		SubTitle                   sql.NullString  `db:"sub_title"`      // 副标题
		PromotionType              sql.NullInt64   `db:"promotion_type"` // 促销类型：0->没有促销使用原价;1->使用促销价；2->使用会员价；3->使用阶梯价格；4->使用满减价格；5->限时购
		Id                         int64           `db:"id"`
		ProductCategoryId          sql.NullInt64   `db:"product_category_id"`
		NewStatus                  sql.NullInt64   `db:"new_status"`  // 新品状态:0->不是新品；1->新品
		GiftGrowth                 int64           `db:"gift_growth"` // 赠送的成长值
		LowStock                   sql.NullInt64   `db:"low_stock"`   // 库存预警值
		Note                       sql.NullString  `db:"note"`
		DetailTitle                sql.NullString  `db:"detail_title"`
		GiftPoint                  int64           `db:"gift_point"`     // 赠送的积分
		OriginalPrice              sql.NullFloat64 `db:"original_price"` // 市场价
		DetailDesc                 string          `db:"detail_desc"`
		ProductCategoryName        sql.NullString  `db:"product_category_name"` // 商品分类名称
		DeleteStatus               sql.NullInt64   `db:"delete_status"`         // 删除状态：0->未删除；1->已删除
		Unit                       sql.NullString  `db:"unit"`                  // 单位
		PreviewStatus              sql.NullInt64   `db:"preview_status"`        // 是否为预告商品：0->不是；1->是
		PromotionStartTime         sql.NullTime    `db:"promotion_start_time"`  // 促销开始时间
		PromotionEndTime           sql.NullTime    `db:"promotion_end_time"`    // 促销结束时间
	}
)

func NewPmsProductModel(conn sqlx.SqlConn) PmsProductModel {
	return &defaultPmsProductModel{
		conn:  conn,
		table: "`pms_product`",
	}
}

func (m *defaultPmsProductModel) Insert(data PmsProduct) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, pmsProductRowsExpectAutoSet)
	ret, err := m.conn.Exec(query, data.BrandId, data.ProductSn, data.Description, data.Weight, data.BrandName, data.UsePointLimit, data.ServiceIds, data.AlbumPics, data.DetailHtml, data.DetailMobileHtml, data.Pic, data.RecommandStatus, data.Price, data.PromotionPrice, data.Stock, data.Keywords, data.PromotionPerLimit, data.Sale, data.FeightTemplateId, data.ProductAttributeCategoryId, data.Name, data.PublishStatus, data.VerifyStatus, data.Sort, data.SubTitle, data.PromotionType, data.ProductCategoryId, data.NewStatus, data.GiftGrowth, data.LowStock, data.Note, data.DetailTitle, data.GiftPoint, data.OriginalPrice, data.DetailDesc, data.ProductCategoryName, data.DeleteStatus, data.Unit, data.PreviewStatus, data.PromotionStartTime, data.PromotionEndTime)
	return ret, err
}

func (m *defaultPmsProductModel) FindOne(id int64) (*PmsProduct, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", pmsProductRows, m.table)
	var resp PmsProduct
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

func (m *defaultPmsProductModel) Update(data PmsProduct) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, pmsProductRowsWithPlaceHolder)
	_, err := m.conn.Exec(query, data.BrandId, data.ProductSn, data.Description, data.Weight, data.BrandName, data.UsePointLimit, data.ServiceIds, data.AlbumPics, data.DetailHtml, data.DetailMobileHtml, data.Pic, data.RecommandStatus, data.Price, data.PromotionPrice, data.Stock, data.Keywords, data.PromotionPerLimit, data.Sale, data.FeightTemplateId, data.ProductAttributeCategoryId, data.Name, data.PublishStatus, data.VerifyStatus, data.Sort, data.SubTitle, data.PromotionType, data.ProductCategoryId, data.NewStatus, data.GiftGrowth, data.LowStock, data.Note, data.DetailTitle, data.GiftPoint, data.OriginalPrice, data.DetailDesc, data.ProductCategoryName, data.DeleteStatus, data.Unit, data.PreviewStatus, data.PromotionStartTime, data.PromotionEndTime, data.Id)
	return err
}

func (m *defaultPmsProductModel) Delete(id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.Exec(query, id)
	return err
}
