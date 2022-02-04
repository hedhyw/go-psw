# go-psw

![Version](https://img.shields.io/github/v/tag/hedhyw/go-psw)
[![Build Status](https://app.travis-ci.com/hedhyw/go-psw.svg?branch=main)](https://app.travis-ci.com/github/hedhyw/go-psw)
[![Go Report Card](https://goreportcard.com/badge/github.com/hedhyw/go-psw)](https://goreportcard.com/report/github.com/hedhyw/go-psw)
[![Coverage Status](https://coveralls.io/repos/github/hedhyw/go-psw/badge.svg?branch=main)](https://coveralls.io/github/hedhyw/go-psw?branch=main)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/hedhyw/go-psw)](https://pkg.go.dev/github.com/hedhyw/go-psw?tab=doc)


**A tiny golang tool for generating a crypto-random password using the terminal.**

<img alt="Golang logo" src="https://raw.githubusercontent.com/rfyiamcool/golang_logo/master/png/golang_68.png" height="200" />

## Installation

```sh
go install github.com/hedhyw/go-psw/cmd/psw@latest
```

## Usage

```sh
psw [LENGTH] [CHARS_PATTERN]

# Just generate a secure-random password:
psw

# Generate a password of length 10:
psw 10

# Generate only with given pattern (next example will print [a-zA-Z]):
psw 10 aA

# Special case (next example will print only [ёàは]):
psw 10 ёàは
```

## Using as a library

```go
package main

import (
	"fmt"
	"log"

	"github.com/hedhyw/go-psw/pkg/v1/consts"
	"github.com/hedhyw/go-psw/pkg/v1/generator"
)

func main() {
	g := generator.NewCryptoRandGenerator()

	chars, err := g.GeneratePassword(
		consts.Digits+consts.LowerLetters, // Chars to use.
		10,                                // Password length.
	)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(chars)
}
```

## License

See [License](License).
