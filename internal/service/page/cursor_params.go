package page

import (
	"math"
	"net/http"
	"strconv"

	val "github.com/go-ozzo/ozzo-validation/v4"
)

const (
	pageParamLimit  = "page[limit]"
	pageParamCursor = "page[cursor]"
	pageParamOrder  = "page[order]"

	maxLimit uint64 = 100
)

// CursorParams is a wrapper around pgdb.CursorPageParams with useful validation and rendering methods
type CursorParams struct {
	Cursor uint64 `page:"cursor" json:"cursor"`
	Order  string `page:"order" json:"order"`
	Limit  uint64 `page:"limit" json:"limit"`
}

func (p *CursorParams) Validate() error {
	return val.Errors{
		pageParamLimit:  val.Validate(p.Limit, val.Max(maxLimit)),
		pageParamOrder:  val.Validate(p.Order, val.In(OrderTypeAsc, OrderTypeDesc)),
		pageParamCursor: val.Validate(p.Cursor, val.Max(uint64(math.MaxInt32))),
	}.Filter()
}

func (p *CursorParams) GetLinks(r *http.Request, last int32) *models.Links {
	result := models.Links{
		Self: p.getLink(r, p.Cursor),
	}
	if last != 0 {
		result.Next = p.getLink(r, uint64(last))
	}
	return &result
}

func (p *CursorParams) getLink(r *http.Request, cursor uint64) string {
	u := r.URL
	query := u.Query()
	query.Set(pageParamCursor, strconv.FormatUint(cursor, 10))
	query.Set(pageParamLimit, strconv.FormatUint(p.Limit, 10))
	query.Set(pageParamOrder, p.Order)
	u.RawQuery = query.Encode()
	return u.String()
}
