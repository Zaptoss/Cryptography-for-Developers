package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

var keyLength = 2048
var bigOne = big.NewInt(1)
var bigNegOne = big.NewInt(-1)

func KeyGen() (pv, pb *big.Int) {
	var p, q *big.Int
	// var p1 = big.NewInt(0)
	// var q1 = big.NewInt(0)
	// var n = big.NewInt(0)
	var Fn = big.NewInt(0)
	var e = big.NewInt(0xfffe)
	var d = big.NewInt(1)
	p, _ = rand.Prime(rand.Reader, keyLength/2)
	q, _ = rand.Prime(rand.Reader, keyLength/2)

	// n = new(big.Int).Mul(p, q)
	// p1.Sub(p, bigOne)
	// q1.Sub(q, bigOne)

	// fmt.Println(e)
	Fn = new(big.Int).Mul(new(big.Int).Sub(p, bigOne), new(big.Int).Sub(q, bigOne))

	fmt.Println(Fn)
	d = new(big.Int).ModInverse(e, Fn)
	fmt.Println(d)

	// fmt.Println(e, d)
	// e.Mul(e, d)
	// fmt.Println(e)
	// e.Mod(e, n)
	// fmt.Println(e)

	return e, d
}

func main() {
	KeyGen()
	// a := big.NewInt(20)
	// b := big.NewInt(30)
	// a.Mul(a, b)
	// fmt.Println(a)

	// n := big.NewInt(0xffff)
	// mod, _ := rand.Prime(rand.Reader, keyLength/2)
	// inv := new(big.Int).ModInverse(n, mod)
	// fmt.Println(inv)

}
