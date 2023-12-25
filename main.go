package main

import (
	"crypto/rand"
	"math/big"
	"syscall/js"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

func main() {
	c := make(chan struct{}, 0)

	js.Global().Set("calcBulk", js.FuncOf(calcBulk))

	<-c
}

func calcBulk(this js.Value, p []js.Value) interface{} {
	contractAddr := p[0].String()
	senderAddr := p[1].String()
	difficulty := new(big.Int)
	difficulty.SetString(p[2].String(), 10)
	count := p[3].Int()

	target := new(big.Int).Div(new(big.Int).Lsh(big.NewInt(1), 256), difficulty)
	nonce := generateNonce()

	cAddr := common.HexToAddress(contractAddr)
	sAddr := common.HexToAddress(senderAddr)

	for i := 0; i <= count; i++ {
		hash := calcHash(cAddr, sAddr, nonce, difficulty)
		hashNum := new(big.Int).SetBytes(hash.Bytes())
		if hashNum.Cmp(target) == -1 {
			return nonce.String() // Return nonce as a string
		}
		nonce.Add(nonce, big.NewInt(1))
	}

	return nil // Return nil if no nonce is found
}

func calcHash(contractAddr, senderAddr common.Address, nonce *big.Int, difficulty *big.Int) common.Hash {
	data := []byte{}
	data = append(data, contractAddr.Bytes()...)
	data = append(data, senderAddr.Bytes()...)
	data = append(data, common.LeftPadBytes(nonce.Bytes(), 32)...)
	data = append(data, common.LeftPadBytes(difficulty.Bytes(), 32)...)

	return crypto.Keccak256Hash(data)
}

func generateNonce() *big.Int {
	max := new(big.Int)
	max.Exp(big.NewInt(2), big.NewInt(224), nil)

	n, _ := rand.Int(rand.Reader, max)

	if n.Sign() == 0 {
		n = big.NewInt(1)
	}

	return n
}
