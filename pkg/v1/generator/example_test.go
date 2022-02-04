package generator_test

import (
	"fmt"
	"log"
	"strings"

	"github.com/hedhyw/go-psw/pkg/v1/consts"
	"github.com/hedhyw/go-psw/pkg/v1/generator"
)

func ExampleGenerator_GeneratePassword() {
	g := generator.NewCryptoRandGenerator()

	chars, err := g.GeneratePassword(consts.Digits+consts.LowerLetters, 10)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(chars)
}

func ExampleGenerator_RandInt64() {
	g := generator.NewCryptoRandGenerator()

	v, err := g.RandInt64()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(v)
}

func ExampleNewGeneratorFromReader() {
	g := generator.NewGeneratorFromReader(
		strings.NewReader(
			strings.Repeat("hello world", 10),
		),
	)

	v, err := g.GeneratePassword(consts.Digits, 5)
	if err != nil {
		log.Fatal(err)
	}

	// Output: 44468
	fmt.Print(v)
}
