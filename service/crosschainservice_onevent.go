package service

import (
	"errors"

	"fmt"

	"bytes"

	"github.com/SmartMeshFoundation/distributed-notary/chain"
	ethevents "github.com/SmartMeshFoundation/distributed-notary/chain/ethereum/events"
	smcevents "github.com/SmartMeshFoundation/distributed-notary/chain/spectrum/events"
	"github.com/SmartMeshFoundation/distributed-notary/models"
	"github.com/SmartMeshFoundation/distributed-notary/utils"
	"github.com/nkbai/log"
)

// OnEvent 链上事件逻辑处理 TODO
func (cs *CrossChainService) OnEvent(e chain.Event) {
	var err error
	switch event := e.(type) {
	/*
		events about block number
	*/
	case ethevents.NewBlockEvent:
		err = cs.onMCNewBlockEvent(event)
	case smcevents.NewBlockEvent:
		err = cs.onSCNewBlockEvent(event)
	/*
		events about lockin
	*/
	case ethevents.PrepareLockinEvent: // MCPLI
		log.Info(SCTokenLogMsg(cs.meta, "Receive MC PrepareLockinEvent :\n%s", utils.ToJSONStringFormat(event)))
		err = cs.onMCPrepareLockin(event)
	case smcevents.PrepareLockinEvent: // SCPLI
		log.Info(SCTokenLogMsg(cs.meta, "Receive SC PrepareLockinEvent :\n%s", utils.ToJSONStringFormat(event)))
		err = cs.onSCPrepareLockin(event)
	case smcevents.LockinSecretEvent: //  SCLIS
		log.Info(SCTokenLogMsg(cs.meta, "Receive SC LockinSecretEvent :\n%s", utils.ToJSONStringFormat(event)))
		err = cs.onSCLockinSecret(event)
	case ethevents.LockinEvent: // MCLI
		log.Info(SCTokenLogMsg(cs.meta, "Receive MC LockinEvent :\n%s", utils.ToJSONStringFormat(event)))
		err = cs.onMCLockin(event)
	case ethevents.CancelLockinEvent: // MCCancelLI
		log.Info(SCTokenLogMsg(cs.meta, "Receive MC CancelLockinEvent :\n%s", utils.ToJSONStringFormat(event)))
		err = cs.onMCCancelLockin(event)
	case smcevents.CancelLockinEvent: // SCCancelLI
		log.Info(SCTokenLogMsg(cs.meta, "Receive SC CancelLockinEvent :\n%s", utils.ToJSONStringFormat(event)))
		err = cs.onSCCancelLockin(event)
	/*
		events about lockout
	*/
	case smcevents.PrepareLockoutEvent: // SCPLO
	case ethevents.PrepareLockoutEvent: // MCPLO
	case ethevents.LockoutSecretEvent: // MCLOS
	case smcevents.LockoutEvent: // SCLO
	case ethevents.CancelLockoutEvent: // MCCancelLO
	case smcevents.CancelLockoutEvent: // SCCancelLO
	default:
		err = errors.New("unknow event")
	}
	if err != nil {
		log.Error(SCTokenLogMsg(cs.meta, "%s event deal err =%s", e.GetEventName(), err.Error()))
	}
	return
}

func (cs *CrossChainService) onMCNewBlockEvent(event ethevents.NewBlockEvent) (err error) {
	// TODO
	return
}

func (cs *CrossChainService) onSCNewBlockEvent(event smcevents.NewBlockEvent) (err error) {
	// TODO
	return
}

/*
主链PrepareLockin(MCPLI)
该事件为一个LockinInfo的生命周期开端,发生于用户调用时
事件为已确认事件,直接构造LockinInfo并保存,等待后续调用
*/
func (cs *CrossChainService) onMCPrepareLockin(event ethevents.PrepareLockinEvent) (err error) {
	// 1. 查询
	secretHash, mcExpiration, amount, data, err := cs.mcProxy.QueryLockin(event.Account)
	if err != nil {
		err = fmt.Errorf("mcProxy.QueryLockin err = %s", err.Error())
		return
	}
	// 2. 构造LockinInfo
	lockinInfo := &models.LockinInfo{
		SecretHash:         secretHash,
		Secret:             utils.EmptyHash,
		UserAddress:        event.Account,
		SCTokenAddress:     cs.meta.SCToken,
		Amount:             amount,
		MCExpiration:       mcExpiration,
		SCExpiration:       0,
		MCLockStatus:       models.LockStatusLock,
		SCLockStatus:       models.LockStatusNone,
		Data:               data,
		NotaryIDInCharge:   models.UnknownNotaryIDInCharge,
		StartTime:          event.Time.Unix(),
		StartMCBlockNumber: event.BlockNumber,
	}
	// 3. 保存至db
	err = cs.db.NewLockinInfo(lockinInfo)
	if err != nil {
		err = fmt.Errorf("db.NewLockinInfo err = %s", err.Error())
		return
	}
	return
}

/*
侧链PrepareLockin(SCPLI)
该事件的发起方为公证人,可能为自己
事件为已确认事件,修改LockinInfo状态
*/
func (cs *CrossChainService) onSCPrepareLockin(event smcevents.PrepareLockinEvent) (err error) {
	// 1. 查询
	secretHash, scExpiration, amount, data, err := cs.scTokenProxy.QueryLockin(event.Account)
	if err != nil {
		err = fmt.Errorf("scTokenProxy.QueryLockin err = %s", err.Error())
		return
	}
	// 2. 获取本地LockinInfo信息
	lockinInfo, err := cs.db.GetLockinInfo(secretHash)
	if err != nil {
		err = fmt.Errorf("db.GetLockinInfo err = %s", err.Error())
		return
	}
	// 3.　校验 TODO
	if lockinInfo.UserAddress != event.Account {
		err = fmt.Errorf("userAddress does't match")
		return
	}
	if lockinInfo.SCExpiration != 0 || lockinInfo.MCLockStatus != models.LockStatusLock || lockinInfo.SCLockStatus != models.LockStatusNone || lockinInfo.Secret != utils.EmptyHash {
		err = fmt.Errorf("local lockinInfo status does't right,something must wrong, local lockinInfo:\n%s", utils.ToJSONStringFormat(lockinInfo))
		return
	}
	if secretHash != lockinInfo.SecretHash {
		err = fmt.Errorf("secretHash does't match")
		return
	}
	if lockinInfo.Amount.Cmp(amount) != 0 {
		err = fmt.Errorf("amount does't match")
		return
	}
	if bytes.Compare(lockinInfo.Data, data) != 0 {
		err = fmt.Errorf("data does't match")
		return
	}
	// 主链Expiration　必须大于　5倍侧链Expiration TODO
	if lockinInfo.MCExpiration < 5*scExpiration {
		err = fmt.Errorf("mcExpiration must bigger than scExpiration *  5")
		return
	}
	// 4. 修改状态,等待后续调用
	lockinInfo.SCExpiration = scExpiration
	lockinInfo.SCLockStatus = models.LockStatusLock
	err = cs.db.UpdateLockinInfo(lockinInfo)
	if err != nil {
		err = fmt.Errorf("db.UpdateLockinInfo err = %s", err.Error())
		return
	}
	return
}

/*
侧链LockinSecret(SCLIS)
该事件由用户发起,已确认
*/
func (cs *CrossChainService) onSCLockinSecret(event smcevents.LockinSecretEvent) (err error) {
	// 1.根据密码hash查询LockinInfo
	secretHash := utils.ShaSecret(event.Secret[:])
	lockinInfo, err := cs.db.GetLockinInfo(secretHash)
	if err != nil {
		err = fmt.Errorf("db.GetLockinInfoBySecretHash err = %s", err.Error())
		return
	}
	// 2. 重复校验 TODO
	if lockinInfo.Secret != utils.EmptyHash {
		err = fmt.Errorf("receive repeat SCLockinSecret, ignore")
		return
	}
	// 3. 校验状态,好像没啥用,用户都拿走钱了,就算状态不对,也需要继续操作,让负责人尝试去主链lockin
	if lockinInfo.MCLockStatus != models.LockStatusLock || lockinInfo.SCLockStatus != models.LockStatusLock {
		err = fmt.Errorf("local lockinInfo status does't right,something must wrong, local lockinInfo:\n%s", utils.ToJSONStringFormat(lockinInfo))
	}
	// 3. 更新状态
	lockinInfo.Secret = event.Secret
	lockinInfo.SCLockStatus = models.LockStatusUnlock
	err = cs.db.UpdateLockinInfo(lockinInfo)
	if err != nil {
		err = fmt.Errorf("db.UpdateLockinInfo err = %s", err.Error())
		return
	}
	// 4. 如果自己是负责人,发起主链Lockin
	if lockinInfo.NotaryIDInCharge == cs.self.ID {
		err = cs.callMCLockin()
		if err != nil {
			err = fmt.Errorf("db.callMCLockin err = %s", err.Error())
			return
		}
	}
	return
}

/*
主链Lockin
收到该事件,说明一次Lockin完整结束,合约上已经清楚该UserAccount的lockin信息
*/
func (cs *CrossChainService) onMCLockin(event ethevents.LockinEvent) (err error) {
	// 1. 获取本地LockinInfo信息
	lockinInfo, err := cs.db.GetLockinInfo(event.SecretHash)
	if err != nil {
		err = fmt.Errorf("db.GetLockinInfo err = %s", err.Error())
		return
	}
	// 2. 校验 TODO
	if lockinInfo.UserAddress != event.Account {
		err = fmt.Errorf("userAddress does't match")
		return
	}
	// 3. 更新本地信息
	lockinInfo.MCLockStatus = models.LockStatusUnlock
	lockinInfo.EndTime = event.Time.Unix()
	lockinInfo.EndMCBlockNumber = event.BlockNumber
	err = cs.db.UpdateLockinInfo(lockinInfo)
	if err != nil {
		err = fmt.Errorf("db.UpdateLockinInfo err = %s", err.Error())
		return
	}
	return
}

/*
主链取消
*/
func (cs *CrossChainService) onMCCancelLockin(event ethevents.CancelLockinEvent) (err error) {
	// 1. 获取本地LockinInfo信息
	lockinInfo, err := cs.db.GetLockinInfo(event.SecretHash)
	if err != nil {
		err = fmt.Errorf("db.GetLockinInfo err = %s", err.Error())
		return
	}
	// 2. 校验 TODO
	if lockinInfo.UserAddress != event.Account {
		err = fmt.Errorf("userAddress does't match")
		return
	}
	// 3. 更新本地信息,endTime哪个链取消在后面取哪个
	lockinInfo.MCLockStatus = models.LockStatusCancel
	lockinInfo.EndTime = event.Time.Unix()
	lockinInfo.EndMCBlockNumber = event.BlockNumber
	err = cs.db.UpdateLockinInfo(lockinInfo)
	if err != nil {
		err = fmt.Errorf("db.UpdateLockinInfo err = %s", err.Error())
		return
	}
	return
}

/*
侧链取消
*/
func (cs *CrossChainService) onSCCancelLockin(event smcevents.CancelLockinEvent) (err error) {
	// 1. 获取本地LockinInfo信息
	lockinInfo, err := cs.db.GetLockinInfo(event.SecretHash)
	if err != nil {
		err = fmt.Errorf("db.GetLockinInfo err = %s", err.Error())
		return
	}
	// 2. 校验 TODO
	if lockinInfo.UserAddress != event.Account {
		err = fmt.Errorf("userAddress does't match")
		return
	}
	// 3. 更新本地信息,endTime哪个链取消在后面取哪个
	lockinInfo.SCLockStatus = models.LockStatusCancel
	lockinInfo.EndTime = event.Time.Unix()
	lockinInfo.EndMCBlockNumber = event.BlockNumber
	err = cs.db.UpdateLockinInfo(lockinInfo)
	if err != nil {
		err = fmt.Errorf("db.UpdateLockinInfo err = %s", err.Error())
		return
	}
	return
}
