package utils

import (
	"crypto/rand"

	blst "github.com/supranational/blst/bindings/go"
)

// For minimal-pubkey-size operations:
// type PublicKey = blst.P1Affine
// type Signature = blst.P2Affine
// type AggregateSignature = blst.P2Aggregate
// type AggregatePublicKey = blst.P1Aggregate

// For minimal-signature-size operations:
type PublicKey = blst.P2Affine
type Signature = blst.P1Affine
type AggregateSignature = blst.P1Aggregate
type AggregatePublicKey = blst.P2Aggregate

var dst = []byte("BLS_SIG_BLS12381G1_XMD:SHA-256_SSWU_RO_NUL_")

func GenerateBatchTestKeyPairs(n int) ([]*blst.SecretKey, []*PublicKey) {
	sks := make([]*blst.SecretKey, n)
	pubks := make([]*PublicKey, n)
	for i := 0; i < n; i++ {
		sk := genRandomKey()
		sks[i] = sk
		pubks[i] = new(PublicKey).From(sk)
	}
	return sks, pubks
}

func SignMsg(sks []*blst.SecretKey, msg []byte) []*Signature {
	sigs := make([]*Signature, len(sks))
	for i := 0; i < len(sks); i++ {
		sigs[i] = new(Signature).Sign(sks[i], msg, dst)
	}
	return sigs
}

func AggSig(sigs []*Signature) (*Signature, bool) {
	aggSig := new(AggregateSignature)
	if !aggSig.Aggregate(sigs, false) {
		return nil, false
	}
	return aggSig.ToAffine(), true
}

func VerifyMultiSig(sig *Signature, pubks []*PublicKey, msg []byte) bool {
	return sig.FastAggregateVerify(false, pubks, msg, dst)
}

func genRandomKey() *blst.SecretKey {
	// Generate 32 bytes of randomness
	var ikm [32]byte
	_, err := rand.Read(ikm[:])

	if err != nil {
		return nil
	}
	return blst.KeyGen(ikm[:])
}
