package main

import (
	rsa "DigitSign/rsa"
	"bufio"
	"fmt"
	"math/big"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	// bytes, _, _ := reader.ReadLine()
	// fmt.Printf(string(bytes))
	// fmt.Scanln()

	var bytes []byte
	var state = true
	for state {
		fmt.Printf("Enter you command(h - help, q - exit): ")
		bytes, _, _ = reader.ReadLine()
		switch string(bytes) {
		case "h":
			fmt.Println("## kg - key generation; this command generate p and q, one of this can used as private key and other as public key, and n - modul")
			fmt.Println("## sf - sign file; it take your private key, n as hex and file path as string for sign file")
			fmt.Println("## vf - verify file; it take signature, private key and n as hex and file path as string for verify file")
			fmt.Println("## sm - sign message; it take your private key, n as hex and message as string for sign file")
			fmt.Println("## vf - verify message; it take signature, private key and n as hex and file path as string for verify message")
		case "q":
			state = false
		case "kg":
			fmt.Println("It can take some time...")
			pb, pv, n := rsa.KeyGen()
			fmt.Printf("-- Public Key: %s\n-- Private Key: %s\n-- n: %s\n", pb.Text(16), pv.Text(16), n.Text(16))
		case "sf":
			var filePath, pvS, nS string
			fmt.Printf("-- Enter filepath: ")
			bytes, _, _ = reader.ReadLine()
			filePath = string(bytes)
			fmt.Printf("-- Enter private key: ")
			fmt.Scanln(&pvS)
			fmt.Printf("-- Enter n: ")
			fmt.Scanln(&nS)

			pv := big.NewInt(0)
			n := big.NewInt(0)

			pv.SetString(pvS, 16)
			n.SetString(nS, 16)
			signature := rsa.SignFile(filePath, pv, n)
			fmt.Printf("-- Signature is: %s\n", signature.Text(16))

		case "sm":
			var message, pvS, nS string
			fmt.Printf("-- Enter message: ")
			bytes, _, _ = reader.ReadLine()
			message = string(bytes)
			fmt.Printf("-- Enter private key: ")
			fmt.Scanln(&pvS)
			fmt.Printf("-- Enter n: ")
			fmt.Scanln(&nS)

			pv := big.NewInt(0)
			n := big.NewInt(0)

			pv.SetString(pvS, 16)
			n.SetString(nS, 16)
			signature := rsa.SignMessage(message, pv, n)
			fmt.Printf("-- Signature is: %s\n", signature.Text(16))
		case "vf":
			var filePath, signatureS, pbS, nS string
			fmt.Printf("-- Enter filepath: ")
			fmt.Scanln(&filePath)
			fmt.Printf("-- Enter signature: ")
			fmt.Scanln(&signatureS)
			fmt.Printf("-- Enter public key: ")
			fmt.Scanln(&pbS)
			fmt.Printf("-- Enter n: ")
			fmt.Scanln(&nS)

			signature := big.NewInt(0)
			pb := big.NewInt(0)
			n := big.NewInt(0)
			signature.SetString(signatureS, 16)
			pb.SetString(pbS, 16)
			n.SetString(nS, 16)
			if rsa.VerifyFile(filePath, signature, pb, n) {
				fmt.Println("Sign verified!")
			} else {
				fmt.Println("Sign was corrupted!")
			}
		case "vm":
			var message, signatureS, pbS, nS string
			fmt.Printf("-- Enter message: ")
			bytes, _, _ = reader.ReadLine()
			message = string(bytes)
			fmt.Printf("-- Enter signature: ")
			fmt.Scanln(&signatureS)
			fmt.Printf("-- Enter public key: ")
			fmt.Scanln(&pbS)
			fmt.Printf("-- Enter n: ")
			fmt.Scanln(&nS)

			signature := big.NewInt(0)
			pb := big.NewInt(0)
			n := big.NewInt(0)

			signature.SetString(signatureS, 16)
			pb.SetString(pbS, 16)
			n.SetString(nS, 16)
			if rsa.VerifyMessage(message, signature, pb, n) {
				fmt.Println("Sign verified!")
			} else {
				fmt.Println("Sign was corrupted!")
			}
		}

	}
}
