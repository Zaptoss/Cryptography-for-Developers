package main

import (
	sha1 "DigitSign/sha1"
	"crypto/rand"
	"fmt"
	"math/big"
)

var keyLength = 2048
var bigOne = big.NewInt(1)
var bigNegOne = big.NewInt(-1)

func KeyGen() (pb, pv, n *big.Int) {
	var p, q, Fn, e, d *big.Int

	for true {
		e, _ = rand.Prime(rand.Reader, keyLength)
		p, _ = rand.Prime(rand.Reader, keyLength/2)
		q, _ = rand.Prime(rand.Reader, keyLength/2)
		Fn = new(big.Int).Mul(new(big.Int).Sub(p, bigOne), new(big.Int).Sub(q, bigOne))

		if new(big.Int).GCD(nil, nil, e, Fn).String() == "1" && e.Cmp(Fn) == -1 {
			break
		}
	}

	n = new(big.Int).Mul(p, q)
	d = new(big.Int).ModInverse(e, Fn)

	return e, d, n
}

func SignFile(filePath string, pv, n *big.Int) (s *big.Int) {
	hash := sha1.NewHash([]byte{})
	hash.SHA1File(filePath)
	hashSum, _ := new(big.Int).SetString(hash.GetHex(), 16)
	return new(big.Int).Exp(hashSum, pv, n)
}

func VerifyFile(filePath string, signature, pb, n *big.Int) bool {
	hash := sha1.NewHash([]byte{})
	hash.SHA1File(filePath)
	announcedHash := new(big.Int).Exp(signature, pb, n)
	if announcedHash.Text(16) == hash.GetHex() {
		return true
	}
	return false
}

func SignMessage(msg string, pv, n *big.Int) (s *big.Int) {
	hash := sha1.NewHash([]byte(msg))
	hash.SHA1()
	hashSum, _ := new(big.Int).SetString(hash.GetHex(), 16)
	return new(big.Int).Exp(hashSum, pv, n)
}

func VerifyMessage(msg string, signature, pb, n *big.Int) bool {
	hash := sha1.NewHash([]byte(msg))
	hash.SHA1()
	announcedHash := new(big.Int).Exp(signature, pb, n)
	if announcedHash.Text(16) == hash.GetHex() {
		return true
	}
	return false
}

func main() {
	pb, pv, n := KeyGen()
	// fmt.Printf("Public Key: %s\nPrivate Key: %s\nn: %s\n", pb.Text(16), pv.Text(16), n.Text(16))
	enc := SignMessage("test.txt", pv, n)
	dec := VerifyMessage("test.txt", enc, pb, n)
	fmt.Println(dec)
}
