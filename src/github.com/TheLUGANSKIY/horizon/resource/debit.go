package resource

import (
	"github.com/TheLUGANSKIY/horizon/assets"
	"github.com/TheLUGANSKIY/horizon/db2/core"
	"golang.org/x/net/context"
)


func (this *Debit) Populate(ctx context.Context, row core.Debit) (err error) {
	this.Type, err = assets.String(row.Assettype)
	if err != nil {
		return
	}

	this.Owner = row.Owner
	this.Debitor = row.Debitor
	this.Issuer = row.Issuer
	this.Code = row.Assetcode
	return
}