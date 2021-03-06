package events

import (
	"github.com/SmartMeshFoundation/distributed-notary/chain"
	"github.com/SmartMeshFoundation/distributed-notary/chain/spectrum/contracts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// CancelLockinEvent :
type CancelLockinEvent struct {
	*chain.BaseEvent
	Account    common.Address `json:"account"` // lockin的用户地址
	SecretHash common.Hash    `json:"secret_hash"`
}

// CreateCancelLockinEvent :
func CreateCancelLockinEvent(log types.Log) (event CancelLockinEvent, err error) {
	e := &contracts.AtmosphereTokenCancelLockin{}
	err = unpackLog(&atmosphereTokenABI, e, AtmosphereTokenCancelLockinEventName, &log)
	if err != nil {
		return
	}
	event.BaseEvent = createBaseEventFromSpectrumLog(AtmosphereTokenCancelLockinEventName, log)
	// params
	event.Account = e.Account
	event.SecretHash = e.SecretHash
	return
}
