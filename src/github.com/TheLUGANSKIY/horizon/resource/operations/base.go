package operations

import (
	"fmt"
	"github.com/TheLUGANSKIY/horizon/db2/history"
	"github.com/TheLUGANSKIY/horizon/httpx"
	"github.com/TheLUGANSKIY/horizon/render/hal"
	"golang.org/x/net/context"
)

// PagingToken implements hal.Pageable
func (this Base) PagingToken() string {
	return this.PT
}

// Populate fills out this resource using `row` as the source.
func (this *Base) Populate(ctx context.Context, row history.Operation) {
	this.ID = fmt.Sprintf("%d", row.ID)
	this.PT = row.PagingToken()
	this.SourceAccount = row.SourceAccount
	this.populateType(row)

	lb := hal.LinkBuilder{httpx.BaseURL(ctx)}
	self := fmt.Sprintf("/operations/%d", row.ID)
	this.Links.Self = lb.Link(self)
	this.Links.Succeeds = lb.Linkf("/effects?order=desc&cursor=%s", this.PT)
	this.Links.Precedes = lb.Linkf("/effects?order=asc&cursor=%s", this.PT)
	this.Links.Transaction = lb.Linkf("/transactions/%s", row.TransactionHash)
	this.Links.Effects = lb.Link(self, "effects")
}

func (this *Base) populateType(row history.Operation) {
	var ok bool
	this.TypeI = int32(row.Type)
	this.Type, ok = TypeNames[row.Type]

	if !ok {
		this.Type = "unknown"
	}
}
