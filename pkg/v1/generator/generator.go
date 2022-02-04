package generator

import (
	"crypto/rand"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"strings"
)

// Generator of random passwords.
type Generator struct {
	io.Reader
}

// NewGeneratorFromReader creates a new generator with a given source.
func NewGeneratorFromReader(r io.Reader) *Generator {
	return &Generator{
		Reader: r,
	}
}

// NewCryptoRandGenerator creates a new generator with a crypto/rand source.
func NewCryptoRandGenerator() *Generator {
	return NewGeneratorFromReader(rand.Reader)
}

// GeneratePassword generates a password of a given size using given chars.
func (g Generator) GeneratePassword(chars string, size int) (string, error) {
	switch {
	case size <= 0:
		return "", errors.New("crypto: invalid size")
	case len(chars) == 0:
		return "", errors.New("crypto: no chars")
	case len(chars) == 1:
		return strings.Repeat(chars, size), nil
	}

	charRunes := []rune(chars)

	var sb strings.Builder

	for i := 0; i < size; i++ {
		v, err := g.RandInt64()
		if err != nil {
			return "", fmt.Errorf("crypto: rand int64: %w", err)
		}

		r := charRunes[v%uint64(len(charRunes))]

		_, err = sb.WriteRune(r)
		if err != nil {
			return "", fmt.Errorf("crypto: writing rune: %w", err)
		}
	}

	return sb.String(), nil
}

// RandInt64 generates a random int64 number.
func (g Generator) RandInt64() (v uint64, err error) {
	err = binary.Read(g.Reader, binary.LittleEndian, &v)

	return v, err
}
