package models

import (
	"testing"

	"crypto/rand"

	"math/big"

	"fmt"

	"github.com/SmartMeshFoundation/distributed-notary/accounts"
	"github.com/SmartMeshFoundation/distributed-notary/curv/feldman"
	"github.com/SmartMeshFoundation/distributed-notary/curv/proofs"
	"github.com/SmartMeshFoundation/distributed-notary/curv/share"
	dutils "github.com/SmartMeshFoundation/distributed-notary/utils"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/nkbai/goutils"
	"github.com/stretchr/testify/assert"
)

func TestBTCPubKey(t *testing.T) {
	am := accounts.NewAccountManager("../testdata/keystore")
	addr := common.HexToAddress("0x201b20123b3c489b47fde27ce5b451a0fa55fd60")
	privateKeyBin, err := am.GetPrivateKey(addr, "123")
	if err != nil {
		return
	}
	pk, err := crypto.ToECDSA(privateKeyBin)
	if err != nil {
		return
	}
	p := &PrivateKeyInfo{
		PublicKeyX: pk.X,
		PublicKeyY: pk.Y,
	}
	fmt.Println(p.ToAddress().String())
	fmt.Println(p.ToBTCPubKeyAddress(&chaincfg.SimNetParams).EncodeAddress())
	//pk2 := (*btcec.PrivateKey)(pk)
}
func TestNewPrivateKeyInfo(t *testing.T) {
	db := SetupTestDB()
	defer db.Close()
	pprivKey, err := proofs.GenerateKey(rand.Reader, 1024)
	if err != nil {
		t.Error(err)
		return
	}
	r := dutils.NewRandomHash()
	p := &PrivateKeyInfo{
		Key:             r,
		Address:         dutils.NewRandomAddress(),
		UI:              share.RandomPrivateKey(),
		XI:              share.RandomPrivateKey(),
		PaillierPrivkey: pprivKey,
		PubKeysProof1: map[int]*KeyGenBroadcastMessage1{
			1: &KeyGenBroadcastMessage1{
				&proofs.DLogProof{
					ChallengeResponse: share.RandomPrivateKey(),
				},
			},
		},
	}
	err = db.NewPrivateKeyInfo(p)
	if err != nil {
		t.Error(err)
	}
	p2, err := db.LoadPrivateKeyInfo(p.Key)
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("p=%s", utils.StringInterface(p, 7))
	t.Logf("p2=%s", utils.StringInterface(p2, 7))
	assert.EqualValues(t, p, p2)

	p3, err := db.LoadPrivateKeyInfoByAccountAddress(p.Address)
	assert.Empty(t, err)
	assert.EqualValues(t, p, p3)

}

func TestDB_UpdateKeyGenStatus(t *testing.T) {
	db := SetupTestDB()
	defer db.Close()
	pprivKey, err := proofs.GenerateKey(rand.Reader, 1024)
	if err != nil {
		t.Error(err)
		return
	}
	r := dutils.NewRandomHash()
	p := &PrivateKeyInfo{
		Key:             r,
		UI:              share.RandomPrivateKey(),
		XI:              share.RandomPrivateKey(),
		PaillierPrivkey: pprivKey,
		PubKeysProof1: map[int]*KeyGenBroadcastMessage1{
			1: &KeyGenBroadcastMessage1{
				&proofs.DLogProof{
					ChallengeResponse: share.RandomPrivateKey(),
				},
			},
		},
	}
	err = db.NewPrivateKeyInfo(p)
	if err != nil {
		t.Error(err)
	}

	p2, err := db.LoadPrivateKeyInfo(p.Key)
	if err != nil {
		t.Error(err)
		return
	}
	assert.EqualValues(t, p, p2)

	p.PaillierKeysProof2 = make(map[int]*KeyGenBroadcastMessage2)
	p.PaillierKeysProof2[0] = &KeyGenBroadcastMessage2{
		PaillierPubkey: &p.PaillierPrivkey.PublicKey,
		Com:            big.NewInt(20),
	}
	err = db.KGUpdatePaillierKeysProof2(p)
	if err != nil {
		t.Error(err)
		return
	}
	p.Status = PrivateKeyNegotiateStatusPaillierPubKey
	p2, err = db.LoadPrivateKeyInfo(p.Key)
	assert.EqualValues(t, err, nil)
	t.Logf("p=%s", utils.StringInterface(p, 7))
	t.Logf("p2=%s", utils.StringInterface(p2, 7))
	assert.EqualValues(t, p2, p)

	p.SecretShareMessage3 = make(map[int]*KeyGenBroadcastMessage3)
	vss, _ := feldman.Share(1, 3, share.RandomPrivateKey())
	p.SecretShareMessage3[1] = &KeyGenBroadcastMessage3{
		Vss:   vss,
		Index: 1,
	}
	err = db.KGUpdateSecretShareMessage3(p)
	assert.EqualValues(t, err, nil)
	p.Status = PrivateKeyNegotiateStatusSecretShare
	p2, err = db.LoadPrivateKeyInfo(p.Key)
	assert.EqualValues(t, err, nil)
	t.Logf("p=%s", utils.StringInterface(p, 7))
	t.Logf("p2=%s", utils.StringInterface(p2, 7))
	assert.EqualValues(t, p2, p)

}
