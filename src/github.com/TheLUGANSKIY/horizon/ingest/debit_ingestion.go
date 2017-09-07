package ingest

import (
	"github.com/TheLUGANSKIY/go/xdr"
)

// Add writes a debit to the database while automatically tracking the index
// to use.
func(di *DebitIngestion) Add(owner xdr.AccountId, details interface{}) bool {
	if di.err != nil {
		return false
	}

	di.added++

	di.err = di.Dest.Debit(owner.Address(), details)
	if di.err != nil {
		return false
	}

	return true
}

// Finish marks this ingestion as complete, returning any error that was recorded.
func (di *DebitIngestion) Finish() error {
	err := di.err
	di.err = nil
	return err
}