package ElGamal

import (
	"crypto/rand"
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

func EncryptPart(M []byte, p, g, y, q *big.Int) (a, b *big.Int) {
	k, _ := rand.Int(rand.Reader, q)

	a = new(big.Int).Exp(g, k, p)
	bp := new(big.Int).Exp(y, k, p)
	fmt.Println(M)
	b = new(big.Int).Mod(new(big.Int).Mul(bp, new(big.Int).SetBytes(M)), p)
	return a, b
}

func Encrypt(M []byte, p, g, y *big.Int) (enc [][2]*big.Int) {
	msgLength := len(M)
	msgAlignment := msgLength % 64
	// if msgAlignment != 0 {
	// 	for i := 0; i < 64-msgAlignment; i++ {
	// 		M = append(M, 0)
	// 	}
	// }
	var a, b *big.Int
	q := new(big.Int).Sub(p, bigOne)

	for i := 0; i < msgLength/64; i++ {
		a, b = EncryptPart(M[i*64:(i+1)*64], p, g, y, q)
		encPart := [2]*big.Int{a, b}
		enc = append(enc, encPart)
	}

	if msgAlignment != 0 {
		a, b = EncryptPart(M[msgLength-msgAlignment:msgLength], p, g, y, q)
		// fmt.Println(DecryptPart(a, b, x, p, q))
		encPart := [2]*big.Int{a, b}
		enc = append(enc, encPart)
	}

	return enc
}

func DecryptPart(a, b, x, p, q *big.Int) []byte {
	s := new(big.Int).Exp(a, x, p)
	sr := new(big.Int).ModInverse(s, p)
	M := new(big.Int).Mod(new(big.Int).Mul(b, sr), p)
	// M := new(big.Int).Mod(new(big.Int).Mul(new(big.Int).Exp(a, new(big.Int).Sub(q, x), p), b), p)
	return M.Bytes()
}

func Decrypt(enc [][2]*big.Int, x, p *big.Int) []byte {
	var result []byte
	q := new(big.Int).Sub(p, bigOne)
	for _, v := range enc {
		result = append(result, DecryptPart(v[0], v[1], x, p, q)...)
	}
	return result
}
