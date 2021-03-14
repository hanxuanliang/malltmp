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
		Sale                       sql.NullInt64   `db:"sale" json:"sale"`                     // 销量
		PreviewStatus              sql.NullInt64   `db:"preview_status" json:"preview_status"` // 是否为预告商品：0->不是；1->是
		Keywords                   sql.NullString  `db:"keywords" json:"keywords"`
		Note                       sql.NullString  `db:"note" json:"note"`
		PromotionStartTime         sql.NullTime    `db:"promotion_start_time" json:"promotion_start_time"`   // 促销开始时间
		ProductCategoryName        sql.NullString  `db:"product_category_name" json:"product_category_name"` // 商品分类名称
		PromotionPrice             sql.NullFloat64 `db:"promotion_price" json:"promotion_price"`             // 促销价格
		SubTitle                   sql.NullString  `db:"sub_title" json:"sub_title"`                         // 副标题
		OriginalPrice              sql.NullFloat64 `db:"original_price" json:"original_price"`               // 市场价
		ServiceIds                 sql.NullString  `db:"service_ids" json:"service_ids"`                     // 以逗号分割的产品服务：1->无忧退货；2->快速退款；3->免费包邮
		DetailTitle                sql.NullString  `db:"detail_title" json:"detail_title"`
		ProductSn                  string          `db:"product_sn" json:"product_sn"` // 货号
		Price                      sql.NullFloat64 `db:"price" json:"price"`
		Stock                      sql.NullInt64   `db:"stock" json:"stock"` // 库存
		DetailDesc                 string          `db:"detail_desc" json:"detail_desc"`
		DetailMobileHtml           string          `db:"detail_mobile_html" json:"detail_mobile_html"` // 移动端网页详情
		Id                         int64           `db:"id" json:"id"`
		FeightTemplateId           sql.NullInt64   `db:"feight_template_id" json:"feight_template_id"`
		ProductAttributeCategoryId sql.NullInt64   `db:"product_attribute_category_id" json:"product_attribute_category_id"`
		PublishStatus              sql.NullInt64   `db:"publish_status" json:"publish_status"` // 上架状态：0->下架；1->上架
		VerifyStatus               sql.NullInt64   `db:"verify_status" json:"verify_status"`   // 审核状态：0->未审核；1->审核通过
		Name                       string          `db:"name" json:"name"`
		Description                string          `db:"description" json:"description"`       // 商品描述
		PromotionType              sql.NullInt64   `db:"promotion_type" json:"promotion_type"` // 促销类型：0->没有促销使用原价;1->使用促销价；2->使用会员价；3->使用阶梯价格；4->使用满减价格；5->限时购
		Pic                        sql.NullString  `db:"pic" json:"pic"`
		GiftGrowth                 int64           `db:"gift_growth" json:"gift_growth"`                 // 赠送的成长值
		UsePointLimit              sql.NullInt64   `db:"use_point_limit" json:"use_point_limit"`         // 限制使用的积分数
		AlbumPics                  sql.NullString  `db:"album_pics" json:"album_pics"`                   // 画册图片，连产品图片限制为5张，以逗号分割
		PromotionPerLimit          sql.NullInt64   `db:"promotion_per_limit" json:"promotion_per_limit"` // 活动限购数量
		Sort                       sql.NullInt64   `db:"sort" json:"sort"`                               // 排序
		GiftPoint                  int64           `db:"gift_point" json:"gift_point"`                   // 赠送的积分
		LowStock                   sql.NullInt64   `db:"low_stock" json:"low_stock"`                     // 库存预警值
		BrandId                    sql.NullInt64   `db:"brand_id" json:"brand_id"`
		ProductCategoryId          sql.NullInt64   `db:"product_category_id" json:"product_category_id"`
		DeleteStatus               sql.NullInt64   `db:"delete_status" json:"delete_status"`           // 删除状态：0->未删除；1->已删除
		NewStatus                  sql.NullInt64   `db:"new_status" json:"new_status"`                 // 新品状态:0->不是新品；1->新品
		RecommandStatus            sql.NullInt64   `db:"recommand_status" json:"recommand_status"`     // 推荐状态；0->不推荐；1->推荐
		Unit                       sql.NullString  `db:"unit" json:"unit"`                             // 单位
		Weight                     sql.NullFloat64 `db:"weight" json:"weight"`                         // 商品重量，默认为克
		DetailHtml                 string          `db:"detail_html" json:"detail_html"`               // 产品详情网页内容
		PromotionEndTime           sql.NullTime    `db:"promotion_end_time" json:"promotion_end_time"` // 促销结束时间
		BrandName                  sql.NullString  `db:"brand_name" json:"brand_name"`                 // 品牌名称
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
	ret, err := m.conn.Exec(query, data.Sale, data.PreviewStatus, data.Keywords, data.Note, data.PromotionStartTime, data.ProductCategoryName, data.PromotionPrice, data.SubTitle, data.OriginalPrice, data.ServiceIds, data.DetailTitle, data.ProductSn, data.Price, data.Stock, data.DetailDesc, data.DetailMobileHtml, data.FeightTemplateId, data.ProductAttributeCategoryId, data.PublishStatus, data.VerifyStatus, data.Name, data.Description, data.PromotionType, data.Pic, data.GiftGrowth, data.UsePointLimit, data.AlbumPics, data.PromotionPerLimit, data.Sort, data.GiftPoint, data.LowStock, data.BrandId, data.ProductCategoryId, data.DeleteStatus, data.NewStatus, data.RecommandStatus, data.Unit, data.Weight, data.DetailHtml, data.PromotionEndTime, data.BrandName)
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
	_, err := m.conn.Exec(query, data.Sale, data.PreviewStatus, data.Keywords, data.Note, data.PromotionStartTime, data.ProductCategoryName, data.PromotionPrice, data.SubTitle, data.OriginalPrice, data.ServiceIds, data.DetailTitle, data.ProductSn, data.Price, data.Stock, data.DetailDesc, data.DetailMobileHtml, data.FeightTemplateId, data.ProductAttributeCategoryId, data.PublishStatus, data.VerifyStatus, data.Name, data.Description, data.PromotionType, data.Pic, data.GiftGrowth, data.UsePointLimit, data.AlbumPics, data.PromotionPerLimit, data.Sort, data.GiftPoint, data.LowStock, data.BrandId, data.ProductCategoryId, data.DeleteStatus, data.NewStatus, data.RecommandStatus, data.Unit, data.Weight, data.DetailHtml, data.PromotionEndTime, data.BrandName, data.Id)
	return err
}

func (m *defaultPmsProductModel) Delete(id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.Exec(query, id)
	return err
}
