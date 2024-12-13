package page

//
//import (
//	"math"
//	"net/http"
//	"strconv"
//
//	val "github.com/go-ozzo/ozzo-validation/v4"
//)
//
//const pageParamNumber = "page[number]"
//
//const (
//	// OrderTypeAsc means result should be sorted in ascending order.
//	OrderTypeAsc = "asc"
//	// OrderTypeDesc means result should be sorted in descending order.
//	OrderTypeDesc = "desc"
//)
//
//// OffsetParams is a wrapper around pgdb.OffsetPageParams with useful validation and rendering methods
//type OffsetParams struct {
//	Limit      uint64 `page:"limit" default:"15" json:"limit"`
//	Order      string `page:"order" default:"desc" json:"order"`
//	PageNumber uint64 `page:"number" json:"number"`
//}
//
//func (p *OffsetParams) Validate() error {
//	return val.Errors{
//		pageParamLimit:  val.Validate(p.Limit, val.Max(maxLimit)),
//		pageParamNumber: val.Validate(p.PageNumber, val.Max(uint64(math.MaxInt32))),
//		pageParamOrder:  val.Validate(p.Order, val.In(OrderTypeAsc, OrderTypeDesc)),
//	}.Filter()
//}
//
//func (p *OffsetParams) GetLinks(r *http.Request, resourceCount uint64) *models.Links {
//	result := models.Links{
//		Self: p.getLink(r, p.PageNumber),
//	}
//
//	if p.PageNumber != 0 {
//		result.Prev = p.getLink(r, p.PageNumber-1)
//	}
//
//	if p.Limit*p.PageNumber < resourceCount {
//		result.Next = p.getLink(r, p.PageNumber+1)
//	}
//	return &result
//}
//
//func (p *OffsetParams) getLink(r *http.Request, number uint64) string {
//	u := r.URL
//	query := u.Query()
//	query.Set(pageParamNumber, strconv.FormatUint(number, 10))
//	query.Set(pageParamLimit, strconv.FormatUint(p.Limit, 10))
//	query.Set(pageParamOrder, p.Order)
//	u.RawQuery = query.Encode()
//	return u.String()
//}
