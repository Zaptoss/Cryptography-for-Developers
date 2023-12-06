package main

import (
	// sign "ElGamal/sign"
	encrypt "ElGamal/encrypt"
)

func main() {
	// p, g, a, b := sign.GenKey()
	// fmt.Printf("p: %s\ng: %s\na: %s\nb: %s\n", p.Text(16), g.Text(16), a.Text(16), b.Text(16))
	// r, s := sign.Sign("Hello", p, g, a)
	// fmt.Printf("r: %s\ns: %s\n", r.Text(16), s.Text(16))
	// if sign.Verify("Hello", p, g, b, r, s) == true {
	// 	fmt.Println("Success!")
	// }
	p, g, a, b := encrypt.GenKey()
	encrypt.Encrypt([]byte("hello"), p, g, a)
	// z := encrypt.Decrypt(x, b, p)
	// fmt.Println(z)
}
