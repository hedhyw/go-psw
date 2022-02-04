package generator_test

import (
	"bytes"
	"crypto/rand"
	"io"
	"strings"
	"testing"

	"github.com/hedhyw/go-psw/pkg/v1/consts"
	"github.com/hedhyw/go-psw/pkg/v1/generator"

	"github.com/stretchr/testify/assert"
)

// getTestSource should be unchaged, it is required for retrieving of
// deterministic results.
func getTestSource() io.Reader {
	return strings.NewReader(
		strings.Repeat("hello world", 100),
	)
}

func TestGeneratorFromReader_RandInt64(t *testing.T) {
	t.Parallel()

	t.Run("ok", func(t *testing.T) {
		t.Parallel()

		g := generator.NewGeneratorFromReader(getTestSource())

		v1, err := g.RandInt64()
		assert.NoError(t, err)

		v2, err := g.RandInt64()
		assert.NoError(t, err)

		assert.NotEqual(t, v1, v2)
	})

	t.Run("read_failed", func(t *testing.T) {
		t.Parallel()

		g := generator.NewGeneratorFromReader(bytes.NewReader(nil))

		_, err := g.RandInt64()
		assert.Error(t, err)
	})
}

func TestGeneratorFromReader_GeneratePassword(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		Name  string
		Chars string
		Size  int
		Ok    bool
	}{{
		Name:  "digits_10",
		Chars: consts.Digits,
		Size:  10,
		Ok:    true,
	}, {
		Name:  "upper_letters_20",
		Chars: consts.UpperLetters,
		Size:  20,
		Ok:    true,
	}, {
		Name:  "single_char_10",
		Chars: "1",
		Size:  10,
		Ok:    true,
	}, {
		Name:  "size_negative",
		Chars: consts.Digits,
		Size:  -1,
		Ok:    false,
	}, {
		Name:  "size_is_zero",
		Chars: consts.Digits,
		Size:  0,
		Ok:    false,
	}}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()

			g := generator.NewGeneratorFromReader(getTestSource())

			gotPassword, err := g.GeneratePassword(tc.Chars, tc.Size)
			if tc.Ok {
				assert.NoError(t, err)
				assert.Len(t, gotPassword, tc.Size)

				for _, p := range gotPassword {
					assert.Contains(t, tc.Chars, string(p))
				}
			} else {
				assert.Error(t, err)
			}
		})
	}
}

func TestCryptoRandGenerator_GeneratePassword(t *testing.T) {
	t.Parallel()

	g := generator.NewCryptoRandGenerator()
	assert.Equal(t, g.Reader, rand.Reader)

	const chars = consts.Digits
	const size = 10

	gotPassword, err := g.GeneratePassword(consts.Digits, size)
	assert.NoError(t, err)

	for _, p := range gotPassword {
		assert.Contains(t, chars, string(p))
	}
}
