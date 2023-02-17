package main

import (
	crypto "crypto/rand"
	"fmt"
	"math/big"
	"math/rand"
)

/*
	генератор случайных чисел

Случайные числа, сгенерированные через math/rand
считаются криптографически ненадежными, так как из-за использования определенного сида последовательности повторяются.

Второй вариант является криптографически надежным, он использует пакет crypto/rand. API использует Reader для предоставления
экземпляра криптографически сильного генератора псевдо-случайных чисел. Данный пакет сам по себе имеет свой Reader по умолчанию,
который основан на системном генераторе случайных чисел
*/
func main() {
	meaning1 := rand.New(rand.NewSource(10))
	meaning2 := rand.New(rand.NewSource(10))

	for i := 0; i < 5; i++ {
		ran1 := meaning1.Int()
		ran2 := meaning2.Int()

		if ran1 != ran2 {

			fmt.Println("Сгенерированная случайным образом неравная последовательность")
			break
		} else {
			fmt.Println("math/ran1:%d, math/ran2:%d\n", ran1, ran2)
		}

	}

	for i := 0; i < 5; i++ {
		safeNum1 := NewCryptoRand()
		safeNum2 := NewCryptoRand()
		if safeNum1 == safeNum2 {
			fmt.Println("сгенерированы одинаковые значения ")
			break
		} else {
			fmt.Printf("Crypto/Rand1: %d, Crypto/Rand2: %d\n", safeNum1, safeNum2)
		}

	}
}

func NewCryptoRand() int64 {
	safeNum, err := crypto.Int(crypto.Reader, big.NewInt(100234))
	if err != nil {
		panic(err)

	}
	return safeNum.Int64()
}
