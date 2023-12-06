package main

import (
	encrypt "ElGamal/encrypt"
	sign "ElGamal/sign"
	"fmt"
)

func main() {
	p, g, a, b := sign.GenKey()

	r, s := sign.Sign("Hello", p, g, a)

	if sign.Verify("Hello", p, g, b, r, s) == true {
		fmt.Println("Success!")
	} else {
		fmt.Println("Don't verify!")
	}

	if sign.Verify("Hellou", p, g, b, r, s) == true {
		fmt.Println("Success!")
	} else {
		fmt.Println("Don't verify!")
	}

	p, g, a, b = encrypt.GenKey()
	x := encrypt.Encrypt([]byte("It some text with length more then 65 bytes, because alghorithm slice msg by 64 byte blocks"), p, g, b)
	fmt.Println(string(encrypt.Decrypt(x, a, p)))

}
