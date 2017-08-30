package horizon

import (
	"errors"

	"fmt"

	"github.com/TheLUGANSKIY/go/xdr"
	"github.com/TheLUGANSKIY/horizon/db2"
	"github.com/TheLUGANSKIY/horizon/db2/history"
	"github.com/TheLUGANSKIY/horizon/render/hal"
	"github.com/TheLUGANSKIY/horizon/resource"
)

// TradeIndexAction renders a page of effect resources, filtered to include
// only trades, identified by a normal page query and optionally filtered by an account
// or order book
type TradeIndexAction struct {
	Action
	AccountFilter string
	Selling       xdr.Asset
	Buying        xdr.Asset
	PagingParams  db2.PageQuery
	Records       []history.Effect
	// LedgerRecords is a cache of loaded ledger data used to populate the time
	// when a trade occurred.
	LedgerRecords map[int32]history.Ledger
	Page          hal.Page
}

// JSON is a method for actions.JSON
func (action *TradeIndexAction) JSON() {
	action.Do(
		action.EnsureHistoryFreshness,
		action.loadParams,
		action.loadRecords,
		action.loadLedgers,
		action.loadPage,
		func() {
			hal.Render(action.W, action.Page)
		},
	)
}

// LoadQuery sets action.Query from the request params
func (action *TradeIndexAction) loadParams() {
	action.AccountFilter = action.GetString("account_id")
	action.PagingParams = action.GetPageQuery()

	// scott: It is unfortunate that we have this string guard below.  Instead, we
	// should probably add an alternative to `GetAsset` that returns a zero-value
	// xdr.Asset when not provided by the request.  Perhaps `MaybeGetAsset`?.
	if action.GetString("selling_asset_type") != "" {
		action.Selling = action.GetAsset("selling_")
		action.Buying = action.GetAsset("buying_")
	}
}

// loadRecords populates action.Records
func (action *TradeIndexAction) loadRecords() {
	trades := action.HistoryQ().Effects().OfType(history.EffectTrade)

	if action.AccountFilter != "" {
		trades = trades.ForAccount(action.AccountFilter)
	}

	if (action.Selling != xdr.Asset{} || action.Buying != xdr.Asset{}) {
		trades = trades.ForOrderBook(action.Selling, action.Buying)
	}

	action.Err = trades.Page(action.PagingParams).Select(&action.Records)
}

// loadLedgers collects the unique ledgers referenced in the loaded trades and loads the details for each.
func (action *TradeIndexAction) loadLedgers() {
	if len(action.Records) == 0 {
		return
	}

	ledgerSequences := make([]interface{}, len(action.Records))

	// populate the unique sequences
	for i, trade := range action.Records {
		ledgerSequences[i] = trade.LedgerSequence()
	}

	var ledgers []history.Ledger
	action.Err = action.HistoryQ().LedgersBySequence(
		&ledgers,
		ledgerSequences...,
	)
	if action.Err != nil {
		return
	}

	action.LedgerRecords = map[int32]history.Ledger{}
	for _, l := range ledgers {
		action.LedgerRecords[l.Sequence] = l
	}
}

// loadPage populates action.Page
func (action *TradeIndexAction) loadPage() {
	for _, record := range action.Records {
		var res resource.Trade

		ledger, found := action.LedgerRecords[record.LedgerSequence()]
		if !found {
			msg := fmt.Sprintf("could not find ledger data for sequence %d", record.LedgerSequence())
			action.Err = errors.New(msg)
			return
		}

		action.Err = res.Populate(action.Ctx, record, ledger)
		if action.Err != nil {
			return
		}

		action.Page.Add(res)
	}

	action.Page.BaseURL = action.BaseURL()
	action.Page.BasePath = action.Path()
	action.Page.Limit = action.PagingParams.Limit
	action.Page.Cursor = action.PagingParams.Cursor
	action.Page.Order = action.PagingParams.Order
	action.Page.PopulateLinks()
}
