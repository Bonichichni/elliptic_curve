package main

import (
	"crypto/elliptic"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"math/big"
)

type ECPoint struct {
	X *big.Int
	Y *big.Int
}

func main() {
	//перевірка рівняння
	c := CreateCurve("P224")
	G := BasePointGGet(c)

	k := SetRandom(256)
	d := SetRandom(256)

	H1 := ScalarMult(*d, G, c)
	H2 := ScalarMult(*k, H1, c)

	H3 := ScalarMult(*k, G, c)
	H4 := ScalarMult(*d, H3, c)

	result := IsEqual(H2, H4)

	if result {
		fmt.Println("Результат перевірки коректності: Вірно")
	} else {
		fmt.Println("Результат перевірки коректності: Невірно")
	}
	//перевірка серіаліації і десеріалізації
	s := ECPointToString(G)
	G1 := StringToECPoint(s)
	if IsEqual(G, G1) {
		fmt.Println("Серіалізація і десеріалізація пройшла успішно")
	} else {
		fmt.Println("Неправильна серіалізація або десеріалізація")
	}
}

func CreateCurve(p string) elliptic.Curve { //Функція для створення певного виду еліптичної кривої
	var curve elliptic.Curve

	switch p {
	case "P224":
		curve = elliptic.P224()
	case "P256":
		curve = elliptic.P256()
	case "P384":
		curve = elliptic.P384()
	case "P521":
		curve = elliptic.P521()
	default:
		fmt.Println("Дана бібліотека не підтримує ваш вид еліптичної кривої")
	}
	return curve
}

func BasePointGGet(c elliptic.Curve) ECPoint { //G-generator receiving
	return ECPoint{c.Params().Gx, c.Params().Gy}
}

func ECPointGen(x, y *big.Int) (point ECPoint) { //ECPoint creation
	return ECPoint{x, y}
}

func IsOnCurveCheck(a ECPoint, c elliptic.Curve) (b bool) { //DOES P ∈ CURVE?
	return c.IsOnCurve(a.X, a.Y)
}

func AddECPoints(a, b ECPoint, c elliptic.Curve) (p ECPoint) { //P + Q
	x, y := c.Add(a.X, a.Y, b.X, b.Y)
	return ECPoint{x, y}
}

func DoubleECPoints(a ECPoint, c elliptic.Curve) (p ECPoint) { //2P
	x, y := c.Double(a.X, a.Y)
	return ECPoint{x, y}
}

func ScalarMult(k big.Int, a ECPoint, c elliptic.Curve) (p ECPoint) { //k * P
	x, y := c.ScalarMult(a.X, a.Y, k.Bytes())
	return ECPoint{x, y}

}

func ECPointToString(point ECPoint) (s string) { //Serialize point
	data, err := json.Marshal(point)
	if err != nil {
		fmt.Println("Помилка при серіалізації точки:", err)
		return ""
	}
	return string(data)
}

func StringToECPoint(s string) (point ECPoint) { //Deserialize point
	err := json.Unmarshal([]byte(s), &point)
	if err != nil {
		fmt.Println("Помилка при десеріалізації рядка:", err)
	}
	return point
}

func PrintECPoint(point ECPoint) { //Print point
	fmt.Printf("X: %s\nY: %s\n", point.X.String(), point.Y.String())
}

func IsEqual(a, b ECPoint) bool { // IsEqual порівнює дві точки
	return a.X.Cmp(b.X) == 0 && a.Y.Cmp(b.Y) == 0
}

func SetRandom(bitSize int) *big.Int { // SetRandom генерує випадкове число заданого біту
	num, err := rand.Int(rand.Reader, new(big.Int).Lsh(big.NewInt(1), uint(bitSize)))
	if err != nil {
		fmt.Println("Помилка при генерації випадкового числа:", err)
	}
	return num
}
