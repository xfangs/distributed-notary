package service

import (
	"fmt"

	"crypto/ecdsa"

	"net/http"

	"encoding/json"

	"bytes"

	"github.com/SmartMeshFoundation/distributed-notary/api/userapi"
	"github.com/SmartMeshFoundation/distributed-notary/chain"
	spectrumevents "github.com/SmartMeshFoundation/distributed-notary/chain/spectrum/events"
	"github.com/SmartMeshFoundation/distributed-notary/cmd/nonce_server/nonceapi"
	"github.com/SmartMeshFoundation/distributed-notary/models"
	"github.com/ethereum/go-ethereum/common"
	"github.com/nkbai/log"
)

/*
	其他service回调DispatchService的入口
*/
type dispatchServiceBackend interface {
	getSelfPrivateKey() *ecdsa.PrivateKey
	getSelfNotaryInfo() *models.NotaryInfo
	getChainByName(chainName string) (c chain.Chain, err error)
	getLockinInfo(scTokenAddress common.Address, secretHash common.Hash) (lockinInfo *models.LockinInfo, err error)
	getLockoutInfo(scTokenAddress common.Address, secretHash common.Hash) (lockoutInfo *models.LockoutInfo, err error)
	getNotaryService() *NotaryService
	getSCTokenMetaInfoBySCTokenAddress(scTokenAddress common.Address) (scToken *models.SideChainTokenMetaInfo)
	/*
		发送http请求给nonce-sever,调用/api/1/apply-nonce接口申请可用某个账号的可用nonce,合约调用之前使用
	*/
	applyNonceFromNonceServer(chainName string, account common.Address) (nonce uint64, err error)

	/*
		notaryService在部署合约之后调用,原则上除此和启动时,其余地方不能调用
	*/
	registerNewSCToken(scTokenMetaInfo *models.SideChainTokenMetaInfo) (err error)

	/*
		notaryService在协商调用合约之后,更新lockinInfo中的NotaryIDInCharge字段,其余地方不应该调用
	*/
	updateLockinInfoNotaryIDInChargeID(scTokenAddress common.Address, secretHash common.Hash, notaryID int) (err error)
	/*
		notaryService在协商调用合约之后,更新lockinInfo中的NotaryIDInCharge字段,其余地方不应该调用
	*/
	updateLockoutInfoNotaryIDInChargeID(scTokenAddress common.Address, secretHash common.Hash, notaryID int) (err error)
}

func (ds *DispatchService) getSelfPrivateKey() *ecdsa.PrivateKey {
	return ds.notaryService.privateKey
}

func (ds *DispatchService) getSelfNotaryInfo() *models.NotaryInfo {
	return ds.notaryService.self
}

func (ds *DispatchService) getChainByName(chainName string) (c chain.Chain, err error) {
	var ok bool
	c, ok = ds.chainMap[chainName]
	if !ok {
		err = fmt.Errorf("can not find chain %s,something must wrong", chainName)
		return
	}
	return
}

func (ds *DispatchService) getNotaryService() *NotaryService {
	return ds.notaryService
}

func (ds *DispatchService) getLockinInfo(scTokenAddress common.Address, secretHash common.Hash) (lockinInfo *models.LockinInfo, err error) {
	ds.scToken2CrossChainServiceMapLock.Lock()
	defer ds.scToken2CrossChainServiceMapLock.Unlock()
	cs, ok := ds.scToken2CrossChainServiceMap[scTokenAddress]
	if !ok {
		panic("never happen")
	}
	return cs.lockinHandler.getLockin(secretHash)
}

func (ds *DispatchService) getLockoutInfo(scTokenAddress common.Address, secretHash common.Hash) (lockoutInfo *models.LockoutInfo, err error) {
	ds.scToken2CrossChainServiceMapLock.Lock()
	defer ds.scToken2CrossChainServiceMapLock.Unlock()
	cs, ok := ds.scToken2CrossChainServiceMap[scTokenAddress]
	if !ok {
		panic("never happen")
	}
	return cs.lockoutHandler.getLockout(secretHash)
}

func (ds *DispatchService) getSCTokenMetaInfoBySCTokenAddress(scTokenAddress common.Address) (scToken *models.SideChainTokenMetaInfo) {
	ds.scToken2CrossChainServiceMapLock.Lock()
	cs, ok := ds.scToken2CrossChainServiceMap[scTokenAddress]
	if !ok {
		panic("never happen")
	}
	ds.scToken2CrossChainServiceMapLock.Unlock()
	scToken = cs.meta
	return
}

func (ds *DispatchService) registerNewSCToken(scTokenMetaInfo *models.SideChainTokenMetaInfo) (err error) {
	// 注册侧链合约:
	err = ds.chainMap[spectrumevents.ChainName].RegisterEventListenContract(scTokenMetaInfo.SCToken)
	if err != nil {
		log.Error("RegisterEventListenContract on side chain err : %s", err.Error())
		return
	}
	// 注册主链合约:
	mc, ok := ds.chainMap[scTokenMetaInfo.MCName]
	if !ok {
		log.Error("can not find chain %s,something must wrong", scTokenMetaInfo.MCName)
		return
	}
	err = mc.RegisterEventListenContract(scTokenMetaInfo.MCLockedContractAddress)
	if err != nil {
		log.Error("RegisterEventListenContract on main chain %s err : %s", scTokenMetaInfo.MCName, err.Error())
		return
	}
	// 6. 构造CrossChainService开始提供服务
	ds.scToken2CrossChainServiceMapLock.Lock()
	ds.scToken2CrossChainServiceMap[scTokenMetaInfo.SCToken] = NewCrossChainService(ds.db, ds, scTokenMetaInfo)
	ds.scToken2CrossChainServiceMapLock.Unlock()
	return
}

func (ds *DispatchService) updateLockinInfoNotaryIDInChargeID(scTokenAddress common.Address, secretHash common.Hash, notaryID int) (err error) {
	ds.scToken2CrossChainServiceMapLock.Lock()
	lh := ds.scToken2CrossChainServiceMap[scTokenAddress].lockinHandler
	ds.scToken2CrossChainServiceMapLock.Unlock()
	lockinInfo, err := lh.getLockin(secretHash)
	if err != nil {
		return
	}
	lockinInfo.NotaryIDInCharge = notaryID
	return lh.updateLockin(lockinInfo)
}
func (ds *DispatchService) updateLockoutInfoNotaryIDInChargeID(scTokenAddress common.Address, secretHash common.Hash, notaryID int) (err error) {
	ds.scToken2CrossChainServiceMapLock.Lock()
	lh := ds.scToken2CrossChainServiceMap[scTokenAddress].lockoutHandler
	ds.scToken2CrossChainServiceMapLock.Unlock()
	lockinInfo, err := lh.getLockout(secretHash)
	if err != nil {
		return
	}
	lockinInfo.NotaryIDInCharge = notaryID
	return lh.updateLockout(lockinInfo)
}

func (ds *DispatchService) applyNonceFromNonceServer(chainName string, account common.Address) (nonce uint64, err error) {
	url := ds.nonceServerHost + nonceapi.APIName2URLMap[nonceapi.APINameApplyNonce]
	req := nonceapi.NewApplyNonceReq(chainName, account, ds.notaryService.self.Host+userapi.APIName2URLMap[userapi.APIAdminNameCancelNonce])
	payload, err := json.Marshal(req)
	if err != nil {
		return
	}
	/* #nosec */
	resp, err := http.Post(url, "application/json", bytes.NewReader(payload))
	if err != nil {
		return
	}
	var buf [4096 * 1024]byte
	n := 0
	n, err = resp.Body.Read(buf[:])
	if err != nil && err.Error() == "EOF" {
		err = nil
	}
	var applyNonceResponse nonceapi.ApplyNonceResponse
	err = json.Unmarshal(buf[:n], &applyNonceResponse)
	if err != nil {
		return
	}
	nonce = applyNonceResponse.Nonce
	return
}
