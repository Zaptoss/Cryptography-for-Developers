package main

import (
	wrapper "github.com/Zaptoss/Cryptography-for-Developers/EllipticCurveWrapper/wrapper"
	"crypto/rand"
	"fmt"
)

func main() {

	G := wrapper.BasePointGGet()
	k, _ := rand.Prime(rand.Reader, 256)
	d, _ := rand.Prime(rand.Reader, 256)

	H1 := wrapper.ScalarMult(d, G)
	H2 := wrapper.ScalarMult(k, H1)

	H3 := wrapper.ScalarMult(k, G)
	H4 := wrapper.ScalarMult(d, H3)
	fmt.Println(H2)
	fmt.Println(H4)
	if H2.X.String() == H4.X.String() && H2.Y.String() == H4.Y.String() {
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}

}
