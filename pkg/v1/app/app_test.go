package app_test

import (
	"bytes"
	"io"
	"strconv"
	"strings"
	"testing"
	"unicode/utf8"

	"github.com/hedhyw/go-psw/pkg/v1/app"
	"github.com/hedhyw/go-psw/pkg/v1/consts"

	"github.com/hedhyw/gherkingen/pkg/v1/bdd"
	"github.com/stretchr/testify/assert"
)

func run(out io.Writer, positionalArgs ...string) (err error) {
	return app.Run(
		append([]string{"psw"}, positionalArgs...),
		out,
	)
}

func TestGeneratorOfPasswords(t *testing.T) {
	f := bdd.NewFeature(t, "Generator of passwords")

	f.Scenario("User wants to generate a password", func(t *testing.T, f *bdd.Feature) {
		var out bytes.Buffer
		f.When("app is called without arguments", func() {
			assert.NoError(t, run(&out))
		})
		f.Then("a random password is printed", func() {
			assert.Greater(t, out.Len(), 0)
		})
	})

	f.Scenario("User wants to generate a password of given length", func(t *testing.T, f *bdd.Feature) {
		type testCase struct {
			PasswordLength int `field:"<password_length>"`
		}

		testCases := map[string]testCase{
			"10":  {10},
			"100": {100},
		}

		f.TestCases(testCases, func(t *testing.T, f *bdd.Feature, tc testCase) {
			var out bytes.Buffer
			f.When("app is called with a single argument <password_length>", func() {
				assert.NoError(t, run(&out, strconv.Itoa(tc.PasswordLength)))
			})
			f.Then("a random password of length <password_length> is printed", func() {
				assert.Equal(t, tc.PasswordLength, len(strings.TrimSpace(out.String())))
			})
		})
	})

	f.Scenario("User gives an invalid password length argument", func(t *testing.T, f *bdd.Feature) {
		type testCase struct {
			PasswordLength string `field:"<password_length>"`
		}

		testCases := map[string]testCase{
			"0":       {"0"},
			"-1":      {"-1"},
			"invalid": {"invalid"},
		}

		f.TestCases(testCases, func(t *testing.T, f *bdd.Feature, tc testCase) {
			var err error
			f.When("app is called with a single argument <password_length>", func() {
				err = run(&bytes.Buffer{}, tc.PasswordLength)
			})
			f.Then("app returns an error", func() {
				assert.Error(t, err)
			})
		})
	})

	f.Scenario("User wants to use only specific groups of chars", func(t *testing.T, f *bdd.Feature) {
		type testCase struct {
			Chars string `field:"<chars>"`
			Group string `field:"group"`
		}

		testCases := map[string]testCase{
			"a_LowerLetters": {"a", consts.LowerLetters},
			"A_UpperLetters": {"A", consts.UpperLetters},
			"1_Digits":       {"1", consts.Digits},
			"._Symbols":      {".", consts.Symbols},
			"ё_Single":       {"ё", "ё"},
		}
		const passwordLength = 10

		f.TestCases(testCases, func(t *testing.T, f *bdd.Feature, tc testCase) {
			var password string

			f.When("app is called with <password_length> = 10 and <chars>", func() {
				var out bytes.Buffer
				assert.NoError(t, run(&out, strconv.Itoa(passwordLength), tc.Chars))

				password = strings.TrimSpace(out.String())
			})
			f.Then("a random password of length <password_length> is printed", func() {
				assert.Equal(t, utf8.RuneCountInString(password), passwordLength)
			})
			f.And("password contains only chars of a <group>.", func() {
				for _, p := range password {
					assert.Contains(t, tc.Group, string(p))
				}
			})
		})
	})
}
