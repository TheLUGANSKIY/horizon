package core

import sq "github.com/Masterminds/squirrel"

// DebitsByAddress loads all debits for `addy`
func (q *Q) DebitsByAddress(dest interface{}, addy string) error {
	sql := selectDebit.Where("owner = ?", addy)
	return q.Select(dest, sql)
}

var selectDebit = sq.Select(
	"deb.owner",
	"deb.debitor",
	"deb.assettype",
	"deb.issuer",
	"deb.assetcode",
).From("debits deb")