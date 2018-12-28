package utils

import (
	"fmt"
	"io"
	"os"

	"crypto/ecdsa"
	rand2 "crypto/rand"

	"encoding/hex"
	"encoding/json"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

//EmptyHash all zero,invalid
var EmptyHash = common.Hash{}

//EmptyAddress all zero,invalid
var EmptyAddress = common.Address{}

//func Bytes2Hash(data []byte) common.Hash {
//	var h common.Hash
//	h.SetBytes(data)
//	return h
//}

//NewRandomHash generate random hash,for testonly
func NewRandomHash() common.Hash {
	return Sha3(Random(64))
}

//NewRandomAddress generate a address,there maybe no corresponding priv key
func NewRandomAddress() common.Address {
	hash := Sha3([]byte(Random(10)))
	return common.BytesToAddress(hash[12:])
}

func readFullOrPanic(r io.Reader, v []byte) int {
	n, err := io.ReadFull(r, v)
	if err != nil {
		panic(err)
	}
	return n
}

// Random takes a parameter (int) and returns random slice of byte
// ex: var randomstrbytes []byte; randomstrbytes = utils.Random(32)
func Random(n int) []byte {
	v := make([]byte, n)
	readFullOrPanic(rand2.Reader, v)
	return v
}

//Sha3 is short for Keccak256Hash
func Sha3(data ...[]byte) common.Hash {
	return crypto.Keccak256Hash(data...)
}

//PublicKeyToAddress convert public key bin to address
func PublicKeyToAddress(pubkey []byte) common.Address {
	return common.BytesToAddress(crypto.Keccak256(pubkey[1:])[12:])
}

//SignData sign with ethereum format
func SignData(privKey *ecdsa.PrivateKey, data []byte) (sig []byte, err error) {
	hash := Sha3(data)
	//why add 27 for the last byte?
	sig, err = crypto.Sign(hash[:], privKey)
	if err == nil {
		sig[len(sig)-1] += byte(27)
	}
	return
}

//Ecrecover is a wrapper for crypto.Ecrecover
func Ecrecover(hash common.Hash, signature []byte) (addr common.Address, err error) {
	if len(signature) != 65 {
		err = fmt.Errorf("signature errr, len=%d,signature=%s", len(signature), hex.EncodeToString(signature))
		return
	}
	signature[len(signature)-1] -= 27 //why?
	pubkey, err := crypto.Ecrecover(hash[:], signature)
	if err != nil {
		signature[len(signature)-1] += 27
		return
	}
	addr = PublicKeyToAddress(pubkey)
	signature[len(signature)-1] += 27
	return
}

// ToJSONStringFormat :
func ToJSONStringFormat(v interface{}) string {
	buf, err := json.MarshalIndent(v, "", "\t")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// Exists returns true if directory||file exists
func Exists(dir string) bool {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		return false
	}
	return true
}
