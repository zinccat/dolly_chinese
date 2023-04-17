package gpt

import (
	"errors"
	"fmt"
	"github.com/zinccat/dolly_chinese/gosrc/gpt/prompt"
	"strings"
	"time"
)

var ErrNotAvail = errors.New("not available")

func (g *Client) statusAdjust(e error) {
	// return fmt.Errorf("error, status code: %d, message: %w", res.StatusCode, errRes.Error)
	if e == nil {
		return
	}
	err := e.Error()
	if strings.Contains(err, "status code: 200") {
		return
	}
	if strings.Contains(err, "status code: 429") {
		if strings.Contains(err, "Your access was terminated due to violation of our policies") {
			g.Status = Banned
		} else {
			g.Status = OutOfService
		}
	} else if strings.Contains(err, "status code: 403") {
		g.Status = Banned
	} else if strings.Contains(err, "status code: 401") {
		g.Status = Banned
	} else if strings.Contains(err, "status code: 503") {
		g.Status = OutOfService
	} else {
		g.Status = OutOfService
	}
}
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
		fmt.Println("[GPT错误]", c.Token, "->", err.Error())
		c.statusAdjust(err)
		if c.Status != OK {
			fmt.Println("[GPT INSTANCE]", c.Token, "->", c.Status.String())
		}
		time.Sleep(20 * time.Second)
	}
	if err == nil {
		return "错误！", ErrNotAvail
	} else {
		return "错误！", err
	}
}
