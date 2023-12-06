package ElGamal

import (
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"math/big"
)

var KeyLength = 2048
var bigOne = big.NewInt(1)

func GenKey() (p, g, a, b *big.Int) {
	p, _ = rand.Prime(rand.Reader, KeyLength)
	q := new(big.Int).Sub(p, bigOne)

	for {
		g, _ = rand.Int(rand.Reader, p)

		b := new(big.Int).Exp(g, q, p)
		if b.Cmp(bigOne) == 0 {
			// fmt.Println("g finded!")
			break
		}

	}

	a, _ = rand.Int(rand.Reader, q)
	b = new(big.Int).Exp(g, a, p)

	return p, g, a, b
}

func Sign(m string, p, g, a *big.Int) (r, s *big.Int) {
	var k, kr *big.Int
	hash := sha256.Sum256([]byte(m))
	hashBig := new(big.Int).SetBytes(hash[:])
	for {
		k, _ = rand.Int(rand.Reader, new(big.Int).Sub(p, bigOne))
		if new(big.Int).GCD(nil, nil, k, new(big.Int).Sub(p, bigOne)).String() == "1" {
			kr = new(big.Int).ModInverse(k, new(big.Int).Sub(p, bigOne))
			break
		}
	}

	// k, _ = rand.Int(rand.Reader, new(big.Int).Sub(p, bigOne))
	// kr = new(big.Int).ModInverse(k, p)
	r = new(big.Int).Exp(g, k, p)
	s = new(big.Int).Mod(new(big.Int).Mul(new(big.Int).Sub(hashBig, new(big.Int).Mul(a, r)), kr), new(big.Int).Sub(p, bigOne))

	return r, s
}

func Verify(m string, p, g, b, r, s *big.Int) bool {
	hash := sha256.Sum256([]byte(m))
	hashBig := new(big.Int).SetBytes(hash[:])

	// q := new(big.Int).Sub(p, bigOne)

	// y := new(big.Int).ModInverse(b, p)
	// sr := new(big.Int).ModInverse(s, p)
	// u1 := new(big.Int).Mod(new(big.Int).Mul(hashBig, sr), q)
	// u2 := new(big.Int).Mod(new(big.Int).Mul(r, sr), q)
	// v := new(big.Int).Mod(new(big.Int).Mul(new(big.Int).Exp(g, u1, p), new(big.Int).Exp(y, u2, p)), p)
	// fmt.Println(r.Text(16))
	// fmt.Println(v.Text(16))

	u1 := new(big.Int).Mod(new(big.Int).Mul(new(big.Int).Exp(b, r, p), new(big.Int).Exp(r, s, p)), p)
	u2 := new(big.Int).Exp(g, hashBig, p)
	fmt.Println(u1.Text(16))
	fmt.Println(u2.Text(16))
	if u1.Text(16) == u2.Text(16) {
		return true
	}
	return false
}
