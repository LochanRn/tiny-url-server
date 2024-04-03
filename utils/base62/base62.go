package base62

import (
	"math/big"
	"strings"
)

const base62Chars = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"

func Base62Encode(input string) string {
	var result strings.Builder
	num := StringToBigInt(input)

	for num.Sign() > 0 {
		remainder := new(big.Int)
		num.DivMod(num, big.NewInt(62), remainder)
		result.WriteByte(base62Chars[remainder.Int64()])
	}

	return ReverseString(result.String())
}

func Base62Decode(encoded string) string {
	var result strings.Builder
	num := new(big.Int)

	for _, char := range ReverseString(encoded) {
		index := strings.IndexByte(base62Chars, byte(char))
		num.Mul(num, big.NewInt(62))
		num.Add(num, big.NewInt(int64(index)))
	}

	for num.Sign() > 0 {
		remainder := new(big.Int)
		num.DivMod(num, big.NewInt(256), remainder)
		result.WriteByte(byte(remainder.Int64()))
	}

	return ReverseString(result.String())
}

func ReverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func StringToBigInt(s string) *big.Int {
	num, _ := new(big.Int).SetString(s, 10)
	return num
}
