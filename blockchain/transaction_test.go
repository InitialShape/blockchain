package blockchain

import (
	"testing"
	"golang.org/x/crypto/ed25519"
	"crypto/rand"
	"github.com/stretchr/testify/assert"
	"github.com/mr-tron/base58/base58"
)

const PUBLIC_KEY = "mVHLEtFHLYQE7mwvkhkUp9uKqq5VDCMLvjYtePtMix5"

func TestCoinbaseTransaction(t *testing.T) {
	publicKey, _, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		t.Error(err)
	}
	transaction := GenerateCoinbase(publicKey, 100)

	outputs := []Output{Output{publicKey, 100}}
	inputs := []Input{Input{[]byte{}, "", 0}}
	expected := Transaction{[]byte{}, inputs, outputs}

	assert.Equal(t, transaction, expected)
}

func TestTransactionGetBase58Hash(t *testing.T) {
	publicKey, err := base58.Decode(PUBLIC_KEY)
	if err != nil {
		t.Error(err)
	}
	outputs := []Output{Output{publicKey, 100}}
	inputs := []Input{Input{[]byte{}, "", 0}}
	transaction := Transaction{[]byte{}, inputs, outputs}
	hash, err := transaction.GetBase58Hash()
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, "3cqbiEKWBZpTYc3D6GTrguBRevhhcowbgdxZAdVApd5U", hash)
}

func TestTransactionSign(t *testing.T) {
	publicKey, privateKey, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		t.Error(err)
	}
	outputs := []Output{Output{publicKey, 100}}
	inputs := []Input{Input{[]byte{}, "", 0}}
	transaction := Transaction{[]byte{}, inputs, outputs}
	hash, err := transaction.GetHash()
	signature := ed25519.Sign(privateKey, hash)
	transaction.Sign(privateKey, 0)

	assert.Equal(t, transaction.Inputs[0].Signature, signature)
}