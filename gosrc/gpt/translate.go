package gpt

import (
	"errors"
	"github.com/zinccat/dolly_chinese/gosrc/gpt/prompt"
)

var ErrNotAvail = errors.New("not available")

func Translate(s string) (string, error) {
	for i := 0; i < 3; i++ {
		val, err := prompt.Translate(GetGPT().cli, s)
		// TODO: handle error
		if err == nil {
			return val, nil
		}
	}
	return "错误！", ErrNotAvail
}
