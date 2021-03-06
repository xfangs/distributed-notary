package mecdsa

import (
	"bytes"
	"errors"
	"fmt"
	"io"

	"github.com/SmartMeshFoundation/distributed-notary/params"

	"math/big"

	"time"

	"github.com/SmartMeshFoundation/distributed-notary/curv/feldman"
	"github.com/SmartMeshFoundation/distributed-notary/curv/proofs"
	"github.com/SmartMeshFoundation/distributed-notary/curv/share"
	"github.com/SmartMeshFoundation/distributed-notary/models"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/nkbai/log"
)

// MessageToSign 待签名的消息体
type MessageToSign interface {
	GetName() string
	GetTransportBytes() []byte
	GetSignBytes() []byte
	Parse(buf []byte) error
}

/*
DistributedSignMessage 由某个公证人主导发起,其他一组公证人参与的,相互协调最终生成有效签名的过程.
*/
type DistributedSignMessage struct {
	db                 *models.DB
	Key                common.Hash               //此次签名的唯一key,由签名主导公证人指定
	PrivateKey         common.Hash               //此次签名用到的分布式私钥, 在数据库中的key
	XI                 share.SPrivKey            //协商好的私钥片
	PaillierPubKeys    map[int]*proofs.PublicKey //其他公证人的同态加密公钥
	PaillierPrivateKey *proofs.PrivateKey        //我的同态加密私钥
	PublicKey          *share.SPubKey            //上次协商生成的总公钥
	Vss                *feldman.VerifiableSS     //上次协商生成的feldman vss
	L                  *models.SignMessage       //此次签名生成过程中需要保存到数据库的信息
	selfNotaryID       int
}

/*
NewDistributedSignMessage 一开始就要确定哪些公证人参与此次签名生成,
人数t > ThresholdCount && t <= ShareCount
指出要签名的交易,公证人应该对此交易做校验,是否是一个合法的交易
*/
func NewDistributedSignMessage(db *models.DB, notaryID int, message MessageToSign, key common.Hash, privateKey common.Hash, s []int) (l *DistributedSignMessage, err error) {
	if len(s) <= params.ThresholdCount {
		err = fmt.Errorf("candidates notary too less")
		return
	}
	l = &DistributedSignMessage{
		db:           db,
		Key:          key,
		PrivateKey:   privateKey,
		selfNotaryID: notaryID,
	}
	l2 := &models.SignMessage{
		Key:             l.Key,
		UsedPrivateKey:  l.PrivateKey,
		Message:         message.GetSignBytes(),
		MessageName:     message.GetName(),
		SignTime:        time.Now().Unix(),
		S:               s,
		Phase1BroadCast: make(map[int]*models.SignBroadcastPhase1),
		Phase2MessageB:  make(map[int]*models.MessageBPhase2),
		Phase3Delta:     make(map[int]*models.DeltaPhase3),
		Phase5A:         make(map[int]*models.Phase5A),
		Phase5C:         make(map[int]*models.Phase5C),
		Phase5D:         make(map[int]share.SPrivKey),
		AlphaGamma:      make(map[int]share.SPrivKey),
		AlphaWI:         make(map[int]share.SPrivKey),
		Delta:           make(map[int]share.SPrivKey),
	}
	err = l.db.NewSignMessage(l2)
	if err != nil {
		return nil, err
	}
	err = l.loadLockout()
	if err != nil {
		return nil, err
	}
	err = l.createSignKeys()
	return
}

// NewDistributedSignMessageFromDB :
func NewDistributedSignMessageFromDB(db *models.DB, selfNotaryID int, key common.Hash, privateKeyID common.Hash) (d *DistributedSignMessage, err error) {
	d = &DistributedSignMessage{
		db:           db,
		selfNotaryID: selfNotaryID,
		Key:          key,
		PrivateKey:   privateKeyID,
	}
	err = d.loadLockout()
	return
}

// 从数据库中载入相关信息,可以做缓存
func (l *DistributedSignMessage) loadLockout() error {
	//if l.L != nil {
	//	return nil
	//}
	p, err := l.db.LoadPrivateKeyInfo(l.PrivateKey)
	if err != nil {
		return err
	}
	l.XI = p.XI
	l.PaillierPrivateKey = p.PaillierPrivkey
	l.PaillierPubKeys = make(map[int]*proofs.PublicKey)
	for k, v := range p.PaillierKeysProof2 {
		l.PaillierPubKeys[k] = v.PaillierPubkey
	}
	l.PublicKey = &share.SPubKey{X: p.PublicKeyX, Y: p.PublicKeyY}
	l.Vss = p.SecretShareMessage3[l.selfNotaryID].Vss

	l.L, err = l.db.LoadSignMessage(l.Key)
	if err != nil {
		return err
	}
	return nil
}

/*
		//参数：自己的私钥片、系数点乘集合和我的多项式y集合、签名人的原始编号、所有签名人的编号
//==>每个公证人的公私钥、公钥片、{{{t,n},t+1个系数点乘G的结果(c1...c2)},y1...yn}
*/
func (l *DistributedSignMessage) createSignKeys() (err error) {
	lambdaI := l.Vss.MapShareToNewParams(l.selfNotaryID, l.L.S)  //lamda_i 解释：通过lamda_i对原所有证明人的群来映射出签名者群
	wi := share.ModMul(lambdaI, l.XI)                            //wi： 我原来的编号在对应签名群编号的映射关系 ，原来我是xi(私钥片) 现在是wi（我在新的签名群中的私钥片）
	gwiX, gwiY := share.S.ScalarBaseMult(wi.Bytes())             //我在签名群中的公钥片
	gammaI := share.RandomPrivateKey()                           //临时私钥
	gGammaIX, gGammaIY := share.S.ScalarBaseMult(gammaI.Bytes()) //临时公钥
	l.L.SignedKey = &models.SignedKey{
		WI:      wi,
		Gwi:     &share.SPubKey{X: gwiX, Y: gwiY},
		KI:      share.RandomPrivateKey(),
		GammaI:  gammaI,
		GGammaI: &share.SPubKey{X: gGammaIX, Y: gGammaIY},
	}
	err = l.db.UpdateSignMessage(l.L)
	return
}

/*
GeneratePhase1Broadcast 确定此次签名所用临时私钥,不能再换了.
*/
func (l *DistributedSignMessage) GeneratePhase1Broadcast() (msg *models.SignBroadcastPhase1, err error) {
	blindFactor := share.RandomBigInt()
	gGammaIX, _ := share.S.ScalarBaseMult(l.L.SignedKey.GammaI.Bytes())
	com := createCommitmentWithUserDefinedRandomNess(gGammaIX, blindFactor)
	msg = &models.SignBroadcastPhase1{
		Com:         com,
		BlindFactor: blindFactor,
	}
	l.L.Phase1BroadCast[l.selfNotaryID] = msg
	err = l.db.UpdateSignMessage(l.L)
	return
}

//ReceivePhase1Broadcast 收集此次签名中,其他公证人所用临时公钥,保证在后续步骤中不会被替换
func (l *DistributedSignMessage) ReceivePhase1Broadcast(msg *models.SignBroadcastPhase1, index int) (finish bool, err error) {
	err = l.loadLockout()
	if err != nil {
		return
	}
	_, ok := l.L.Phase1BroadCast[index]
	if ok {
		err = fmt.Errorf("phase1 broadcast for %d already exist", index)
		return
	}
	l.L.Phase1BroadCast[index] = msg
	err = l.db.UpdateSignMessage(l.L)
	finish = len(l.L.Phase1BroadCast) == len(l.L.S)
	return
}

//p2p告知E(ki) 签名发起人，告诉其他所有人 步骤1 生成ki， 保证后续ki不会发生变化
func newMessageA(ki share.SPrivKey, paillierPubKey *proofs.PublicKey) (*models.MessageA, error) {
	ca, err := proofs.Encrypt(paillierPubKey, ki.Bytes())
	if err != nil {
		return nil, err
	}
	return &models.MessageA{C: ca}, nil
}

/*
GeneratePhase2MessageA p2p告知E(ki) 签名发起人，告诉其他所有人 步骤1 生成ki， 保证后续ki不会发生变化,
同时其他人是无法获取到ki,只有自己知道自己的ki
*/
func (l *DistributedSignMessage) GeneratePhase2MessageA() (msg *models.MessageA, err error) {
	err = l.loadLockout()
	if err != nil {
		return
	}
	ma, err := newMessageA(l.L.SignedKey.KI, &l.PaillierPrivateKey.PublicKey)
	if err != nil {
		return
	}
	l.L.MessageA = ma
	err = l.db.UpdateSignMessage(l.L)
	return ma, err
}

/*
NewMessageB :
我(bob)签名key中的临时私钥gammaI、(alice)paillier公钥、(alice)paillier公钥加密ki的结果
//返回：cB和两个证明(其他人验证)
我（alice）收到其他人（bob)给我的ma以后， 立即计算mb和两个证明，发送给bob
gammaI: alice的临时私钥
paillierPubKey: bob的同态加密公钥
ca: bob发送给alice的E(Ki)
*/
func NewMessageB(gammaI share.SPrivKey, paillierPubKey *proofs.PublicKey, ca *models.MessageA) (*models.MessageB, error) {
	betaTagPrivateKey := share.RandomPrivateKey()
	cBetaTag, err := proofs.Encrypt(paillierPubKey, betaTagPrivateKey.Bytes()) //paillier加密一个随机数
	if err != nil {
		return nil, err
	}
	bca := proofs.Mul(paillierPubKey, ca.C, gammaI.Bytes()) //ca.C：加密ki的结果 gammaI:gammaI
	//cB=b * E(ca) + E(beta_tag)   (b:gammaI  ca:ca.C   )
	cb := proofs.AddCipher(paillierPubKey, bca, cBetaTag)
	//beta= -bata_tag mod q
	beta := share.ModSub(share.PrivKeyZero.Clone(), betaTagPrivateKey)

	//todo 提供证明 ：证明gammaI是我自己的,证明beta是我合法提供的 既然提供了beta,这里面的betatagproof应该是可以忽略的
	bproof := proofs.Prove(gammaI)
	betaTagProof := proofs.Prove(betaTagPrivateKey)
	return &models.MessageB{
		C:            cb,
		BProof:       bproof,
		BetaTagProof: betaTagProof,
		Beta:         beta,
	}, nil
}

/*
ReceivePhase2MessageA alice 收到来自bob的E(ki), 向bob提供证明,自己持有着gammaI和WI
其中gammaI是临时私钥
WI包含着XI
*/
func (l *DistributedSignMessage) ReceivePhase2MessageA(msg *models.MessageA, index int) (mb *models.MessageBPhase2, err error) {
	err = l.loadLockout()
	if err != nil {
		return
	}
	mgGamma, err := NewMessageB(l.L.SignedKey.GammaI, l.PaillierPubKeys[index], msg)
	if err != nil {
		return
	}
	mbw, err := NewMessageB(l.L.SignedKey.WI, l.PaillierPubKeys[index], msg)
	if err != nil {
		return
	}
	mb = &models.MessageBPhase2{
		MessageBGamma: mgGamma,
		MessageBWi:    mbw,
	}
	//l.L.Phase2MessageB=make(map[int]*models.MessageBPhase2)
	///l.L.Phase2MessageB[l.selfNotaryID]=
	return
}

func verifyProofsGetAlpha(m *models.MessageB, dk *proofs.PrivateKey, a share.SPrivKey) (share.SPrivKey, error) {
	ashare, err := proofs.Decrypt(dk, m.C) //用dk解密cB
	if err != nil {
		return share.SPrivKey{}, err
	}
	alpha := new(big.Int).SetBytes(ashare)
	alphaKey := share.BigInt2PrivateKey(alpha)
	gAlphaX, gAlphaY := share.S.ScalarBaseMult(alphaKey.Bytes())
	babTagX, babTagY := share.S.ScalarMult(m.BProof.PK.X, m.BProof.PK.Y, a.Bytes())
	babTagX, babTagY = share.PointAdd(babTagX, babTagY, m.BetaTagProof.PK.X, m.BetaTagProof.PK.Y)
	if proofs.Verify(m.BProof) && proofs.Verify(m.BetaTagProof) &&
		babTagX.Cmp(gAlphaX) == 0 &&
		babTagY.Cmp(gAlphaY) == 0 {
		return alphaKey, nil
	}
	return share.SPrivKey{}, errors.New("invalid key")
}

//ReceivePhase2MessageB 收集并验证MessageB证明信息
func (l *DistributedSignMessage) ReceivePhase2MessageB(msg *models.MessageBPhase2, index int) (finish bool, err error) {
	err = l.loadLockout()
	if err != nil {
		return
	}
	alphaijGamma, err := verifyProofsGetAlpha(msg.MessageBGamma, l.PaillierPrivateKey, l.L.SignedKey.KI)
	if err != nil {
		return
	}
	alphaijWi, err := verifyProofsGetAlpha(msg.MessageBWi, l.PaillierPrivateKey, l.L.SignedKey.KI)
	if err != nil {
		return
	}
	//if !EqualGE(msg.MessageBWi.BProof.PK, l.L.SignedKey.Gwi) { 这里应该等于index的gwi,而不是l.index的gwi
	//	panic("not equal")
	//}
	l.L.Phase2MessageB[index] = msg
	l.L.AlphaGamma[index] = alphaijGamma
	l.L.AlphaWI[index] = alphaijWi
	err = l.db.UpdateSignMessage(l.L)
	if err != nil {
		return
	}
	finish = len(l.L.Phase2MessageB) == len(l.L.S)-1
	return
}

func (l *DistributedSignMessage) phase2DeltaI() share.SPrivKey {
	var k *models.SignedKey
	k = l.L.SignedKey
	if len(l.L.AlphaGamma) != len(l.L.S)-1 {
		panic("arg error")
	}
	//kiGammaI=ki * gammI+Sum(alpha_vec) +Sum(beta_vec)
	kiGammaI := k.KI.Clone()
	share.ModMul(kiGammaI, k.GammaI)
	for _, i := range l.L.S {
		if i == l.selfNotaryID {
			continue
		}
		share.ModAdd(kiGammaI, l.L.AlphaGamma[i])
		share.ModAdd(kiGammaI, l.L.Phase2MessageB[i].MessageBGamma.Beta)
	}
	return kiGammaI
}
func (l *DistributedSignMessage) phase2SigmaI() share.SPrivKey {
	if len(l.L.AlphaWI) != len(l.L.S)-1 {
		panic("length error")
	}
	kiwi := l.L.SignedKey.KI.Clone()
	share.ModMul(kiwi, l.L.SignedKey.WI)
	//todo vij=vji ?
	for _, i := range l.L.S {
		if i == l.selfNotaryID {
			continue
		}
		share.ModAdd(kiwi, l.L.AlphaWI[i])
		share.ModAdd(kiwi, l.L.Phase2MessageB[i].MessageBWi.Beta)
	}
	return kiwi
}

/*
GeneratePhase3DeltaI 依据上一步协商信息,生成我自己的DeltaI,然后广播给所有其他人,需要这些参与公证人得到完整的的Delta
但是生成的SigmaI自己保留,在生成自己的签名片的时候使用
*/
func (l *DistributedSignMessage) GeneratePhase3DeltaI() (msg *models.DeltaPhase3, err error) {
	err = l.loadLockout()
	if err != nil {
		return
	}
	deltaI := l.phase2DeltaI()
	sigmaI := l.phase2SigmaI()
	l.L.Sigma = sigmaI
	l.L.Delta[l.selfNotaryID] = deltaI
	err = l.db.UpdateSignMessage(l.L)
	if err != nil {
		return
	}
	msg = &models.DeltaPhase3{Delta: deltaI}
	return
}

//ReceivePhase3DeltaI 收集所有的deltaI
func (l *DistributedSignMessage) ReceivePhase3DeltaI(msg *models.DeltaPhase3, index int) (finish bool, err error) {
	err = l.loadLockout()
	if err != nil {
		return
	}
	_, ok := l.L.Delta[index]
	if ok {
		err = fmt.Errorf("ReceivePhase3DeltaI for %d already exist", index)
		return
	}
	l.L.Delta[index] = msg.Delta
	err = l.db.UpdateSignMessage(l.L)
	finish = len(l.L.Delta) == len(l.L.S)
	return
}

func phase4(delta share.SPrivKey,
	gammaIProof map[int]*models.MessageBPhase2,
	phase1Broadcast map[int]*models.SignBroadcastPhase1) (*share.SPubKey, error) {
	if len(gammaIProof) != len(phase1Broadcast) {
		panic("length must equal")
	}
	for i, p := range gammaIProof {
		//校验第一步广播的参数没有发生变化
		if createCommitmentWithUserDefinedRandomNess(p.MessageBGamma.BProof.PK.X,
			phase1Broadcast[i].BlindFactor).Cmp(phase1Broadcast[i].Com) == 0 {
			continue
		}
		return nil, errors.New("invliad key")
	}

	//tao_i=g^^gamma
	sumx, sumy := new(big.Int), new(big.Int)
	for _, p := range gammaIProof {
		if sumx.Cmp(big.NewInt(0)) == 0 && sumy.Cmp(big.NewInt(0)) == 0 {
			sumx = p.MessageBGamma.BProof.PK.X
			sumy = p.MessageBGamma.BProof.PK.Y
		} else {
			sumx, sumy = share.PointAdd(sumx, sumy, p.MessageBGamma.BProof.PK.X, p.MessageBGamma.BProof.PK.Y)
		}

	}
	rx, ry := share.S.ScalarMult(sumx, sumy, delta.Bytes())
	return &share.SPubKey{X: rx, Y: ry}, nil
}

//phase3 计算：inverse(delta) mod q
// all parties broadcast delta_i and compute delta_i ^(-1)
func phase3ReconstructDelta(delta map[int]share.SPrivKey) share.SPrivKey {
	sum := share.PrivKeyZero.Clone()
	for _, deltaI := range delta {
		share.ModAdd(sum, deltaI)
	}
	return share.InvertN(sum)
}

//GeneratePhase4R 所有公证人都应该得到相同的R,其中R.X就是最后签名(r,s,v)中的r
func (l *DistributedSignMessage) GeneratePhase4R() (R *share.SPubKey, err error) {
	err = l.loadLockout()
	if err != nil {
		return
	}
	delta := phase3ReconstructDelta(l.L.Delta)
	//少一个自己的MessageB todo fixme 这里面需要有更好的方式来生成数据,后续必须优化
	mgGamma, err := NewMessageB(l.L.SignedKey.GammaI, &l.PaillierPrivateKey.PublicKey, l.L.MessageA)
	if err != nil {
		return
	}
	l.L.Phase2MessageB[l.selfNotaryID] = &models.MessageBPhase2{MessageBGamma: mgGamma, MessageBWi: nil}
	R, err = phase4(delta, l.L.Phase2MessageB, l.L.Phase1BroadCast)
	if err != nil {
		return
	}
	l.L.R = R
	err = l.db.UpdateSignMessage(l.L)
	return
}

//const
func phase5LocalSignature(ki share.SPrivKey, message *big.Int,
	R *share.SPubKey, sigmaI share.SPrivKey,
	pubkey *share.SPubKey) *models.LocalSignature {
	m := share.BigInt2PrivateKey(message)
	r := share.BigInt2PrivateKey(R.X)
	si := share.ModMul(m, ki)
	share.ModMul(r, sigmaI)
	share.ModAdd(si, r) //si=m * k_i + r * sigma_i
	return &models.LocalSignature{
		LI:   share.RandomPrivateKey(),
		RhoI: share.RandomPrivateKey(),
		//li:   big.NewInt(71),
		//rhoi: big.NewInt(73),
		R: &share.SPubKey{
			X: new(big.Int).Set(R.X),
			Y: new(big.Int).Set(R.Y),
		},
		SI: si, //签名片
		M:  new(big.Int).Set(message),
		Y: &share.SPubKey{
			X: new(big.Int).Set(pubkey.X),
			Y: new(big.Int).Set(pubkey.Y),
		},
	}

}

//com:commit的Ci(广播)
//Phase5ADecom1 :commit的Di(广播)
func phase5aBroadcast5bZkproof(l *models.LocalSignature) (*models.Phase5Com1, *models.Phase5ADecom1, *proofs.HomoELGamalProof) {
	blindFactor := share.RandomBigInt()
	//Ai=g^^rho_i
	aix, aiy := share.S.ScalarBaseMult(l.RhoI.Bytes())
	lIRhoI := l.LI.Clone()
	share.ModMul(lIRhoI, l.RhoI)
	//Bi=G*lIRhoI
	bix, biy := share.S.ScalarBaseMult(lIRhoI.Bytes())
	//vi=R*si+G*li
	tx, ty := share.S.ScalarMult(l.R.X, l.R.Y, l.SI.Bytes()) //R^^si
	vix, viy := share.S.ScalarBaseMult(l.LI.Bytes())         //g^^li
	vix, viy = share.PointAdd(vix, viy, tx, ty)

	inputhash := proofs.CreateHashFromGE([]*share.SPubKey{
		{X: vix, Y: viy}, {X: aix, Y: aiy}, {X: bix, Y: biy},
	})
	com := createCommitmentWithUserDefinedRandomNess(inputhash.D, blindFactor)

	//proof是5b的zkp构造
	witness := proofs.NewHomoElGamalWitness(l.LI, l.SI) //li si
	delta := &proofs.HomoElGamalStatement{
		G: share.NewGE(aix, aiy),               //Ai
		H: share.NewGE(l.R.X, l.R.Y),           //R
		Y: share.NewGE(share.S.Gx, share.S.Gy), //g
		D: share.NewGE(vix, viy),               //Vi
		E: share.NewGE(bix, biy),               //Bi
	}
	//证明提供的是正确的si???
	proof := proofs.CreateHomoELGamalProof(witness, delta)
	return &models.Phase5Com1{Com: com},
		&models.Phase5ADecom1{
			Vi:          share.NewGE(vix, viy),
			Ai:          share.NewGE(aix, aiy),
			Bi:          share.NewGE(bix, biy),
			BlindFactor: blindFactor,
		},
		proof
}

/*
GeneratePhase5a5bZkProof 从此步骤开始,互相不断交换信息来判断对方能够生成正确的Si(也就是签名片),
如果所有参与者都能生成最终的签名片,那么我才能把自己的签名片告诉对方.
si的累加和就是签名(r,s,v)中的s
*/
func (l *DistributedSignMessage) GeneratePhase5a5bZkProof() (msg *models.Phase5A, err error) {
	err = l.loadLockout()
	if err != nil {
		return
	}
	//messageHash := utils.Sha256(l.L.Message)
	messageBN := new(big.Int).SetBytes(l.L.Message[:])
	localSignature := phase5LocalSignature(l.L.SignedKey.KI, messageBN, l.L.R, l.L.Sigma, l.PublicKey)
	phase5Com, phase5ADecom, helgamalProof := phase5aBroadcast5bZkproof(localSignature)
	msg = &models.Phase5A{Phase5Com1: phase5Com, Phase5ADecom1: phase5ADecom, Proof: helgamalProof}
	l.L.Phase5A[l.selfNotaryID] = msg
	l.L.LocalSignature = localSignature
	err = l.db.UpdateSignMessage(l.L)
	return
}

// ReceivePhase5A5BProof :
func (l *DistributedSignMessage) ReceivePhase5A5BProof(msg *models.Phase5A, index int) (finish bool, err error) {
	err = l.loadLockout()
	if err != nil {
		return
	}
	_, ok := l.L.Phase5A[index]
	if ok {
		err = fmt.Errorf("ReceivePhase5A5BProof already exist for %d", index)
		return
	}
	delta := &proofs.HomoElGamalStatement{
		G: msg.Phase5ADecom1.Ai,
		H: l.L.R,
		Y: share.NewGE(share.S.Gx, share.S.Gy),
		D: msg.Phase5ADecom1.Vi,
		E: msg.Phase5ADecom1.Bi,
	}
	inputhash := proofs.CreateHashFromGE([]*share.SPubKey{
		msg.Phase5ADecom1.Vi,
		msg.Phase5ADecom1.Ai,
		msg.Phase5ADecom1.Bi,
	})
	e := createCommitmentWithUserDefinedRandomNess(inputhash.D, msg.Phase5ADecom1.BlindFactor)
	if e.Cmp(msg.Phase5Com1.Com) == 0 &&
		msg.Proof.Verify(delta) {

	} else {
		err = errors.New("invalid com")
		return
	}
	l.L.Phase5A[index] = msg
	finish = len(l.L.Phase5A) == len(l.L.S)
	err = l.db.UpdateSignMessage(l.L)
	return
}

func phase5c(l *models.LocalSignature, deCommitments []*models.Phase5ADecom1,
	vi *share.SPubKey) (*models.Phase5Com2, *models.Phase5DDecom2, error) {

	//从广播的commit(Ci,Di)得到vi,ai
	v := vi.Clone()
	for i := 0; i < len(deCommitments); i++ {
		v.X, v.Y = share.PointAdd(v.X, v.Y, deCommitments[i].Vi.X, deCommitments[i].Vi.Y)
	}
	a := deCommitments[0].Ai.Clone()
	for i := 1; i < len(deCommitments); i++ {
		a.X, a.Y = share.PointAdd(a.X, a.Y, deCommitments[i].Ai.X, deCommitments[i].Ai.Y)
	}
	r := share.BigInt2PrivateKey(l.R.X)
	yrx, yry := share.S.ScalarMult(l.Y.X, l.Y.Y, r.Bytes())
	m := share.BigInt2PrivateKey(l.M)
	//Vi之积×g^(-m)*y^(-r)
	gmx, gmy := share.S.ScalarBaseMult(m.Bytes())
	v.X, v.Y = share.PointSub(v.X, v.Y, gmx, gmy)
	v.X, v.Y = share.PointSub(v.X, v.Y, yrx, yry)
	//UI=V * rhoi
	uix, uiy := share.S.ScalarMult(v.X, v.Y, l.RhoI.Bytes())
	//Ti=A * li
	tix, tiy := share.S.ScalarMult(a.X, a.Y, l.LI.Bytes())

	//commit(UI ,Ti)，广播出去
	inputhash := proofs.CreateHashFromGE([]*share.SPubKey{
		{X: uix, Y: uiy},
		{X: tix, Y: tiy},
	})
	blindFactor := share.RandomBigInt()
	com := createCommitmentWithUserDefinedRandomNess(inputhash.D, blindFactor)
	return &models.Phase5Com2{Com: com},
		&models.Phase5DDecom2{
			UI:          &share.SPubKey{X: uix, Y: uiy}, //Ci
			Ti:          &share.SPubKey{X: tix, Y: tiy}, //Di
			BlindFactor: blindFactor,
		},
		nil
}

//GeneratePhase5CProof  fixme 提供一个好的注释
func (l *DistributedSignMessage) GeneratePhase5CProof() (msg *models.Phase5C, err error) {
	err = l.loadLockout()
	if err != nil {
		return
	}
	if len(l.L.Phase5A) != len(l.L.S) {
		panic("cannot genrate5c until all 5b proof received")
	}
	var decomVec []*models.Phase5ADecom1
	for i, m := range l.L.Phase5A {
		if i == l.selfNotaryID { //phase5c 不应该包括自己的decommitment
			continue
		}
		decomVec = append(decomVec, m.Phase5ADecom1)
	}
	phase5com2, phase5decom2, err := phase5c(l.L.LocalSignature, decomVec, l.L.Phase5A[l.selfNotaryID].Phase5ADecom1.Vi)
	if err != nil {
		return
	}
	msg = &models.Phase5C{Phase5Com2: phase5com2, Phase5DDecom2: phase5decom2}
	l.L.Phase5C[l.selfNotaryID] = msg
	err = l.db.UpdateSignMessage(l.L)
	return
}

//ReceivePhase5cProof   fixme 暂时没有好的解释
func (l *DistributedSignMessage) ReceivePhase5cProof(msg *models.Phase5C, index int) (finish bool, err error) {
	err = l.loadLockout()
	if err != nil {
		return
	}
	//验证5c的hash(ui,ti)=5c的Ci

	inputhash := proofs.CreateHashFromGE([]*share.SPubKey{msg.Phase5DDecom2.UI, msg.Phase5DDecom2.Ti})
	inputhash.D = createCommitmentWithUserDefinedRandomNess(inputhash.D, msg.Phase5DDecom2.BlindFactor)
	if inputhash.D.Cmp(msg.Phase5Com2.Com) != 0 {
		err = errors.New("invalid com")
		return
	}
	_, ok := l.L.Phase5C[index]
	if ok {
		err = fmt.Errorf("ReceivePhase5cProof for %d already exist", index)
		return
	}
	l.L.Phase5C[index] = msg
	err = l.db.UpdateSignMessage(l.L)
	finish = len(l.L.Phase5C) == len(l.L.S)
	return
}

/*
Generate5dProof 接受所有签名人的si的广播，有可能某个公证人会保留信息，最终生成有效的签名，私自保留下来,但是不告诉其他人自己的si是多少.
但是这种情况其他公证人可以知道,没有收到某个公证人的si
*/
func (l *DistributedSignMessage) Generate5dProof() (si share.SPrivKey, err error) {
	err = l.loadLockout()
	if err != nil {
		return
	}
	var deCommitments2 []*models.Phase5DDecom2
	var commitments2 []*models.Phase5Com2
	var deCommitments1 []*models.Phase5ADecom1
	for i, m := range l.L.Phase5C {
		deCommitments2 = append(deCommitments2, m.Phase5DDecom2)
		commitments2 = append(commitments2, m.Phase5Com2)
		deCommitments1 = append(deCommitments1, l.L.Phase5A[i].Phase5ADecom1)
	}
	//校验收到的关于si的信息,都是真实的,正确的,只有通过了,才能把自己的si告诉给其他公证人
	si, err = phase5d(l.L.LocalSignature, deCommitments2, commitments2, deCommitments1)
	if err != nil {
		return
	}
	l.L.Phase5D[l.selfNotaryID] = si
	err = l.db.UpdateSignMessage(l.L)
	return
}

/*
 校验收到的关于si的信息,都是真实的,正确的,只有通过了,才能把自己的si告诉给其他公证人
*/
func phase5d(l *models.LocalSignature, deCommitments2 []*models.Phase5DDecom2,
	commitments2 []*models.Phase5Com2,
	deCommitments1 []*models.Phase5ADecom1) (share.SPrivKey, error) {
	if len(deCommitments1) != len(deCommitments2) ||
		len(deCommitments2) != len(commitments2) {
		panic("arg error")
	}

	biasedSumTbX := new(big.Int).Set(share.S.Gx)
	biasedSumTbY := new(big.Int).Set(share.S.Gy)

	for i := 0; i < len(commitments2); i++ {
		//(5c的ti + 5a的bi)连加
		biasedSumTbX, biasedSumTbY = share.PointAdd(biasedSumTbX, biasedSumTbY,
			deCommitments2[i].Ti.X, deCommitments2[i].Ti.Y)
		biasedSumTbX, biasedSumTbY = share.PointAdd(biasedSumTbX, biasedSumTbY,
			deCommitments1[i].Bi.X, deCommitments1[i].Bi.Y)
	}
	//用于比较 UI 和 (5c的ti + 5a的bi)连加 是否相等
	for i := 0; i < len(commitments2); i++ {
		biasedSumTbX, biasedSumTbY = share.PointSub(
			biasedSumTbX, biasedSumTbY,
			deCommitments2[i].UI.X, deCommitments2[i].UI.Y,
		)
	}
	//log.Trace(fmt.Sprintf("(gx,gy)=(%s,%s)", share.S.Gx.Text(16), share.S.Gy.Text(16)))
	//log.Trace(fmt.Sprintf("(tbx,tby)=(%s,%s)", biasedSumTbX.Text(16), biasedSumTbY.Text(16)))
	if share.S.Gx.Cmp(biasedSumTbX) == 0 &&
		share.S.Gy.Cmp(biasedSumTbY) == 0 {
		return l.SI.Clone(), nil
	}
	return share.PrivKeyZero, errors.New("invalid key")
}

//RecevieSI 收集签名片,收集齐以后就可以得到完整的签名.所有公证人都应该得到有效的签名.
func (l *DistributedSignMessage) RecevieSI(si share.SPrivKey, index int) (signature []byte, finish bool, err error) {
	err = l.loadLockout()
	if err != nil {
		return
	}
	if _, ok := l.L.Phase5D[index]; ok {
		err = fmt.Errorf("si for %d already received", index)
		return
	}
	l.L.Phase5D[index] = si
	err = l.db.UpdateSignMessage(l.L)
	if err != nil {
		return
	}
	return l.GetFinalSignature()
}

// GetFinalSignature : 获取最终签名
func (l *DistributedSignMessage) GetFinalSignature() (signature []byte, finish bool, err error) {
	finish = len(l.L.Phase5D) == len(l.L.S)
	if !finish {
		return
	}
	s := l.L.LocalSignature.SI.Clone()
	//所有人的的si，包括自己
	for i, si := range l.L.Phase5D {
		if i == l.selfNotaryID {
			continue
		}
		share.ModAdd(s, si)
	}
	//r := share.BigInt2PrivateKey(l.L.R.X)
	signature, verifyResult := verify(s, l.L.R, l.L.LocalSignature.Y, l.L.LocalSignature.M)
	if !verifyResult {
		err = errors.New("invilad signature")
	}
	return
}

/*
使用SignatureNormalize来对签名进行处理,符合EIP155签名要求.
https://ethereum.stackexchange.com/questions/42455/during-ecdsa-signing-how-do-i-generate-the-recovery-id
I never found any proper documentation about the Recovery ID but I did talk with somebody on Reddit and they gave me my answer:

id = y1 & 1; // Where (x1,y1) = k x G;
if (s > curve.n / 2) id = id ^ 1; // Invert id if s of signature is over half the n
I had to modify the mbedtls library to pass back the Recovery ID but when I did I could generate transactions that Geth accepted 100% of the time.

The long explanation:

During signing, a point is generated (X, Y) called R and a number called S. R's X goes on to become r and S becomes s. In order to generate the Recovery ID you take the one's bit from Y. If S is bigger than half the curve's N parameter you invert that bit. That bit is the Recovery ID. Ethereum goes on to manipulate it to indicate compressed or uncompressed addresses as well as indicate what chain the transaction was signed for (so the transaction can't be replayed on another Ethereum chain that the private key might be present on). These modifications to the Recovery ID become v.

There's also a super rare chance that you need to set the second bit of the recovery id meaning the recovery id could in theory be 0, 1, 2, or 3. But there's a 0.000000000000000000000000000000000000373% of needing to set the second bit according to a question on Bitcoin.SE.


*/
func verify(s share.SPrivKey, R, y *share.SPubKey, message *big.Int) ([]byte, bool) {
	r := share.BigInt2PrivateKey(R.X)
	b := share.InvertN(s)
	a := share.BigInt2PrivateKey(message)
	u1 := a.Clone()
	u1 = share.ModMul(u1, b)
	u2 := r.Clone()
	u2 = share.ModMul(u2, b)

	gu1x, gu1y := share.S.ScalarBaseMult(u1.Bytes())
	yu2x, yu2y := share.S.ScalarMult(y.X, y.Y, u2.Bytes())
	gu1x, gu1y = share.PointAdd(gu1x, gu1y, yu2x, yu2y)
	//已经确认是一个有效的签名
	if share.BigInt2PrivateKey(gu1x).D.Cmp(r.D) == 0 {
		//return true
	} else {
		return nil, false
	}
	sig := make([]byte, 65)
	copy(sig[:32], bigIntTo32Bytes(r.D))
	copy(sig[32:64], bigIntTo32Bytes(s.D))
	//leave sig[64]=0

	//按照以太坊EIP155的规范来处理签名
	h := common.Hash{}
	var err error
	h.SetBytes(message.Bytes())
	if s.D.Cmp(halfN) >= 0 {
		/* //所谓的normalize就是如果s>n/2,s=n-s ,保证s的唯一性
		sig, err = secp256k1.SignatureNormalize(sig)
		if err != nil {
			log.Error(fmt.Sprintf("SignatureNormalize err %s\n,r=%s,s=%s", err, r.D, s.D))
			return nil, false
		}*/
		s2 := new(big.Int)
		s2 = s2.Sub(share.S.N, s.D)
		copy(sig[32:64], bigIntTo32Bytes(s2))
		tmp := readBigInt(bytes.NewBuffer(sig[32:64]))
		tmp = tmp.Add(tmp, s.D)
		if tmp.Cmp(share.S.N) != 0 {
			panic("must equal")
		}
		//log.Info(fmt.Sprintf("s=%s,normals=%s,n=fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd0364141", s.D.Text(16), .Text(16)))
	}

	key, err := crypto.GenerateKey()
	pubkey := key.PublicKey
	pubkey.X = y.X
	pubkey.Y = y.Y
	addr := crypto.PubkeyToAddress(pubkey)

	/*
		id = y1 & 1; // Where (x1,y1) = k x G; x1,y1 is R
		if (s > curve.n / 2) id = id ^ 1; // Invert id if s of signature is over half the n
	*/
	v := new(big.Int).Mod(R.Y, big2).Int64()
	if s.D.Cmp(halfN) > 0 {
		v = v ^ 1
	}
	sig[64] = byte(v)
	//try v=0
	pubkeybin, err := crypto.Ecrecover(h[:], sig)
	if err == nil {
		pubkey2 := crypto.ToECDSAPub(pubkeybin)
		addr2 := crypto.PubkeyToAddress(*pubkey2)
		if addr2 == addr {
			restoreAndCheckRS(sig)
			return sig, true
		}

	} else {
		log.Error(fmt.Sprintf("Ecrecover err %s", err))
	}
	return nil, false
}

//bigIntTo32Bytes convert a big int to bytes
func bigIntTo32Bytes(i *big.Int) []byte {
	data := i.Bytes()
	buf := make([]byte, 32)
	for i := 0; i < 32-len(data); i++ {
		buf[i] = 0
	}
	for i := 32 - len(data); i < 32; i++ {
		buf[i] = data[i-32+len(data)]
	}
	return buf
}

//readBigInt read big.Int from buffer
func readBigInt(reader io.Reader) *big.Int {
	bi := new(big.Int)
	tmpbuf := make([]byte, 32)
	_, err := reader.Read(tmpbuf)
	if err != nil {
		log.Error(fmt.Sprintf("read BigInt error %s", err))
	}
	bi.SetBytes(tmpbuf)
	return bi
}

func restoreAndCheckRS(sig []byte) {
	buf := bytes.NewBuffer(sig)
	readBigInt(buf)
	s := readBigInt(buf)
	v, err := buf.ReadByte()
	if err != nil {
		panic(err)
	}
	if v != 0 && v != 1 {
		panic("wrong v")
	}

	if s.Cmp(halfN) >= 0 {
		panic("wrong s")
	}
}

var halfN *big.Int
var big2 *big.Int

func init() {
	halfN = new(big.Int).Set(share.S.N)
	halfN.Div(halfN, big.NewInt(2))
	big2 = big.NewInt(2)
}
