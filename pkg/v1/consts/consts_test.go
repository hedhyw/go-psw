package consts_test

import (
	"testing"
	"unicode"

	"github.com/hedhyw/go-psw/pkg/v1/consts"

	"github.com/stretchr/testify/assert"
)

func TestLowerLetters(t *testing.T) {
	t.Parallel()

	for _, r := range consts.LowerLetters {
		assert.True(t, unicode.IsLower(r), string(r))
	}
}

func TestUpperLetters(t *testing.T) {
	t.Parallel()

	for _, r := range consts.UpperLetters {
		assert.True(t, unicode.IsUpper(r), string(r))
	}
}

func TestDigits(t *testing.T) {
	t.Parallel()

	for _, r := range consts.Digits {
		assert.True(t, unicode.IsDigit(r), string(r))
	}
}

func TestSymbols(t *testing.T) {
	for _, r := range consts.Symbols {
		assert.True(t, !unicode.IsDigit(r), string(r))
		assert.True(t, !unicode.IsUpper(r), string(r))
		assert.True(t, !unicode.IsLower(r), string(r))
		assert.True(t, !unicode.IsControl(r), string(r))
	}
}
