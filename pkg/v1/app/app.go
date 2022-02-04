package app

import (
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/hedhyw/go-psw/pkg/v1/consts"
	"github.com/hedhyw/go-psw/pkg/v1/generator"
)

const (
	defaultSize = 18
)

func Run(args []string, out io.Writer) (err error) {
	chars := consts.LowerLetters + consts.UpperLetters + consts.Digits
	size := defaultSize

	if len(args) > 1 {
		size, err = strconv.Atoi(args[1])
		if err != nil {
			return fmt.Errorf("parsing length: %w", err)
		}
	}

	if len(args) > 2 {
		chars, err = prepareChars(args[2])
		if err != nil {
			return fmt.Errorf("preparing chars: %w", err)
		}
	}

	g := generator.NewCryptoRandGenerator()
	password, err := g.GeneratePassword(chars, size)
	if err != nil {
		return fmt.Errorf("generating password: %w", err)
	}

	_, err = fmt.Fprintln(out, password)
	if err != nil {
		return fmt.Errorf("printing: %w", err)
	}

	return nil
}

func prepareChars(in string) (string, error) {
	var sb strings.Builder

	var err error
	for _, r := range in {
		switch {
		case strings.ContainsRune(consts.LowerLetters, r):
			_, err = sb.WriteString(consts.LowerLetters)
		case strings.ContainsRune(consts.UpperLetters, r):
			_, err = sb.WriteString(consts.UpperLetters)
		case strings.ContainsRune(consts.Digits, r):
			_, err = sb.WriteString(consts.Digits)
		case strings.ContainsRune(consts.Symbols, r):
			_, err = sb.WriteString(consts.Symbols)
		default:
			_, err = sb.WriteRune(r)
		}
		if err != nil {
			return "", err
		}
	}

	return sb.String(), nil
}
