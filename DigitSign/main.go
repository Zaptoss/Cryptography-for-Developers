package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"DigitSign/sha1"
)

var keyLength = 2048
var bigOne = big.NewInt(1)
var bigNegOne = big.NewInt(-1)

func KeyGen() (pb, pv, n *big.Int) {
	var p, q, Fn, d *big.Int
	var e = big.NewInt(0xffff)

	for true {
		p, _ = rand.Prime(rand.Reader, keyLength/2)
		q, _ = rand.Prime(rand.Reader, keyLength/2)
		Fn = new(big.Int).Mul(new(big.Int).Sub(p, bigOne), new(big.Int).Sub(q, bigOne))

		if new(big.Int).GCD(nil, nil, e, Fn).String() == "1" {
			break
		}
	}

	d = new(big.Int).ModInverse(e, Fn)
	n = new(big.Int).Mul(p, q)

	return e, d, n
}

func Sign(fileName string, pv, n *big.Int) (r *big.Int) {
	return new(big.Int).Exp(fileName, pv, n)
}

func Verify(fileName string, signature, pb, n *big.Int) (bool) {

}

func main() {
	pb, pv, n := KeyGen()
	fmt.Printf("Public Key: %s\nPrivate Key: %s\nn: %s\n", pb.String(), pv.String(), n.String())
	enc := EncryptDecrypt(big.NewInt(10), pb, n)
	dec := EncryptDecrypt(enc, pv, n)
	fmt.Println(dec)

}
