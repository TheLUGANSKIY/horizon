package horizon

import (
	"github.com/TheLUGANSKIY/horizon/ledger"
	"github.com/TheLUGANSKIY/horizon/render/hal"
	"github.com/TheLUGANSKIY/horizon/resource"
)

// RootAction provides a summary of the horizon instance and links to various
// useful endpoints
type RootAction struct {
	Action
}

// JSON renders the json response for RootAction
func (action *RootAction) JSON() {
	action.App.UpdateStellarCoreInfo()

	var res resource.Root
	res.Populate(
		action.Ctx,
		ledger.CurrentState(),
		action.App.horizonVersion,
		action.App.coreVersion,
		action.App.networkPassphrase,
		action.App.protocolVersion,
	)

	hal.Render(action.W, res)
}
