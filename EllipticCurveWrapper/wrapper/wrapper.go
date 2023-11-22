package "github.com/Zaptoss/Cryptography-for-Developers/EllipticCurveWrapper"

import (
	"crypto/elliptic"
	"crypto/rand"
	"fmt"
	"math/big"
	"strings"
	// "crypto/ecdh"
)

var curve elliptic.Curve = elliptic.P224()

type ECPoint struct {
	X *big.Int
	Y *big.Int
}

func BasePointGGet() (point ECPoint) {
	point = ECPoint{}
	_, point.X, point.Y, _ = elliptic.GenerateKey(curve, rand.Reader)
	return point
}

func ECPointGen(x, y *big.Int) (point ECPoint) {
	point.X = x
	point.Y = y
	return point
}

func IsOnCurveCheck(a ECPoint) (result bool) {
	result = curve.Params().IsOnCurve(a.X, a.Y)
	return result
}

func AddECPoints(a, b ECPoint) (c ECPoint) {
	c.X, c.Y = curve.Params().Add(a.X, a.Y, b.X, b.Y)
	return c
}

func DoubleECPoint(a ECPoint) (c ECPoint) {
	c.X, c.Y = curve.Params().Double(a.X, a.Y)
	return c
}

func ScalarMult(k *big.Int, a ECPoint) (c ECPoint) {
	c.X, c.Y = curve.Params().ScalarMult(a.X, a.Y, k.Bytes())
	return c
}

func ECPoint2String(point ECPoint) (s string) {
	s = point.X.String() + " " + point.Y.String()
	return s
}

func String2ECPoint(s string) (point ECPoint) {
	point.X = &big.Int{}
	point.Y = &big.Int{}
	stringsArr := strings.Split(s, " ")
	point.X.SetString(stringsArr[0], 10)
	point.Y.SetString(stringsArr[1], 10)
	return point
}

func PrintECPoint(point ECPoint) {
	fmt.Printf("X: %s;\nY: %s;\n", point.X, point.Y)
}
