package proxy

import (
	"context"

	"fmt"

	"math/big"

	"github.com/SmartMeshFoundation/distributed-notary/chain/spectrum/client"
	"github.com/SmartMeshFoundation/distributed-notary/chain/spectrum/contracts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// SideChainErc20TokenProxy :
type SideChainErc20TokenProxy struct {
	Contract *contracts.AtmosphereToken
	conn     *client.SafeEthClient
}

// NewSideChainErc20TokenProxy :
func NewSideChainErc20TokenProxy(conn *client.SafeEthClient, tokenAddress common.Address) (p *SideChainErc20TokenProxy, err error) {
	code, err := conn.CodeAt(context.Background(), tokenAddress, nil)
	if err == nil && len(code) > 0 {
		c, err2 := contracts.NewAtmosphereToken(tokenAddress, conn)
		if err = err2; err != nil {
			return
		}
		p = &SideChainErc20TokenProxy{
			Contract: c,
			conn:     conn,
		}
		return
	}
	err = fmt.Errorf("no code at %s", tokenAddress.String())
	return
}

// QueryLockin impl chain.ContractProxy
func (p *SideChainErc20TokenProxy) QueryLockin(accountHex string) (secretHash common.Hash, expiration uint64, amount *big.Int, err error) {
	account := common.HexToAddress(accountHex)
	var cExpiration *big.Int
	secretHash, cExpiration, amount, err = p.Contract.QueryLockin(nil, account)
	if err != nil {
		return
	}
	expiration = cExpiration.Uint64()
	return
}

// QueryLockout impl chain.ContractProxy
func (p *SideChainErc20TokenProxy) QueryLockout(accountHex string) (secretHash common.Hash, expiration uint64, amount *big.Int, err error) {
	account := common.HexToAddress(accountHex)
	var cExpiration *big.Int
	secretHash, cExpiration, amount, err = p.Contract.QueryLockout(nil, account)
	if err != nil {
		return
	}
	expiration = cExpiration.Uint64()
	return
}

// Lockin impl chain.ContractProxy
func (p *SideChainErc20TokenProxy) Lockin(opts *bind.TransactOpts, accountHex string, secret common.Hash) (err error) {
	account := common.HexToAddress(accountHex)
	var tx *types.Transaction
	tx, err = p.Contract.Lockin(opts, account, secret)
	if err != nil {
		return
	}
	ctx := context.Background()
	_, err = bind.WaitMined(ctx, p.conn, tx)
	return
}

// CancelLockin impl chain.ContractProxy
func (p *SideChainErc20TokenProxy) CancelLockin(opts *bind.TransactOpts, accountHex string) (err error) {
	account := common.HexToAddress(accountHex)
	var tx *types.Transaction
	tx, err = p.Contract.CancelLockin(opts, account)
	if err != nil {
		return
	}
	ctx := context.Background()
	_, err = bind.WaitMined(ctx, p.conn, tx)
	return
}
