package gpt

import (
	"errors"
	"fmt"
	"github.com/sashabaranov/go-openai"
	"github.com/zinccat/dolly_chinese/gosrc/gpt/prompt"
	"strings"
)

var ErrNotAvail = errors.New("not available")

func Translate(s string) (string, error) {
	if strings.TrimSpace(s) == "" {
		return "", nil
	}
	val := ""
	var err error
	for i := 0; i < 3; i++ {
		c := GetGPT()
		val, err = prompt.Translate(c.cli, s)
		if err == nil {
			return val, nil
		}

		apiE, ok := err.(*openai.APIError)
		if ok {
			switch apiE.StatusCode {
			case 429:
				if strings.Contains(apiE.Message, "Your access was terminated due to violation of our policies") {
					c.Status = Banned
				} else {
					c.Status = OutOfService
				}
			case 403, 401:
				c.Status = Banned
			case 503:
				c.Status = OutOfService
			default:
				c.Status = OutOfService
			}
			if c.Status != OK {
				fmt.Println("[GPT INSTANCE]", c.Token, "->", c.Status.String())
			}
		}
	}
	if err == nil {
		return "错误！", ErrNotAvail
	} else {
		return "错误！", err
	}
}
