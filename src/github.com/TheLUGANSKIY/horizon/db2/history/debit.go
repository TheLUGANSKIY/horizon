package history

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/TheLUGANSKIY/horizon/db2"
)

// Debits provides a helper to filter rows from the `history_debits` table
// with pre-defined filters.
func (q *Q) Debits() *DebitsQ {
	return &DebitsQ{
		parent: q,
		sql: selectDebit,
	}
}

// DebitByID loads a row from `history_debits`, by id
func (q *Q) DebitByID(dest interface{}, id int64) error {
	sql := selectDebit.Limit(1).Where("deb.id = ?", id)
	return q.Get(dest, sql)
}

// Page specifies the paging constraints for the query being built by `q`.
func (q *DebitsQ) Page(page db2.PageQuery) *DebitsQ {
	if q.Err != nil {
		return q
	}

	q.sql, q.Err = page.ApplyTo(q.sql, "deb.id")
	return q
}

// Select loads the results of the query specified by `q` into `dest`.
func (q *DebitsQ) Select(dest interface{}) error {
	if q.Err != nil {
		return q.Err
	}

	q.Err = q.parent.Select(dest, q.sql)
	return q.Err
}

var selectDebit = sq.Select("deb.*").From("history_debits deb")