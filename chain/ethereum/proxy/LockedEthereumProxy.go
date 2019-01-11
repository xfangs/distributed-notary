package proxy

import (
	"context"
	"fmt"
	"math/big"

	"github.com/SmartMeshFoundation/distributed-notary/chain/ethereum/client"
	"github.com/SmartMeshFoundation/distributed-notary/chain/ethereum/contracts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// LockedEthereumProxy :
type LockedEthereumProxy struct {
	Contract *contracts.LockedEthereum
	conn     *client.SafeEthClient
}

// NewLockedEthereumProxy :
func NewLockedEthereumProxy(conn *client.SafeEthClient, contractAddress common.Address) (p *LockedEthereumProxy, err error) {
	code, err := conn.CodeAt(context.Background(), contractAddress, nil)
	if err == nil && len(code) > 0 {
		c, err2 := contracts.NewLockedEthereum(contractAddress, conn)
		if err = err2; err != nil {
			return
		}
		p = &LockedEthereumProxy{
			Contract: c,
			conn:     conn,
		}
		return
	}
	err = fmt.Errorf("no code at %s", contractAddress.String())
	return
}

// QueryLockin impl chain.ContractProxy
func (p *LockedEthereumProxy) QueryLockin(accountHex string) (secretHash common.Hash, expiration uint64, amount *big.Int, err error) {
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
func (p *LockedEthereumProxy) QueryLockout(accountHex string) (secretHash common.Hash, expiration uint64, amount *big.Int, err error) {
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
func (p *LockedEthereumProxy) Lockin(opts *bind.TransactOpts, accountHex string, secret common.Hash) (err error) {
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
func (p *LockedEthereumProxy) CancelLockin(opts *bind.TransactOpts, accountHex string) (err error) {
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
