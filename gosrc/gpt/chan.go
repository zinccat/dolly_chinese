package gpt

import (
	"fmt"
	"github.com/zinccat/dolly_chinese/gosrc/shared"
	"time"
)

var workableCh chan *Client
var brokenCh chan *Client
var outOfServiceCh chan *Client

func Init() {
	var gpts []*Client
	for _, token := range shared.Cfg.Tokens {
		gpts = append(gpts, New(token))
	}
	if len(gpts) == 0 {
		panic("no gpt3 token")
	}
	workableCh = make(chan *Client, len(gpts))
	brokenCh = make(chan *Client, len(gpts))
	outOfServiceCh = make(chan *Client, len(gpts))
	for _, gpt := range gpts {
		workableCh <- gpt
	}
	statistic()
	go daemon()
}

func daemon() {
	gpt := <-outOfServiceCh
	time.Sleep(time.Minute)
	gpt.Status = OK
	workableCh <- gpt
	fmt.Println("[GPT SCHEDULER]", gpt.Token, "-> OK")
	statistic()
}

func statistic() {
	fmt.Println("[GPT STATISTICS]", "Valid:", len(workableCh), "Broken:", len(brokenCh), "OOS:", len(outOfServiceCh))
}

func GetGPT() *Client {
	gpt := <-workableCh
	if gpt.IsOk() {
		workableCh <- gpt
		return gpt
	}
	if gpt.IsBanned() {
		fmt.Println("[GPT SCHEDULER]", gpt.Token, "-> BANNED")
		brokenCh <- gpt
	} else if gpt.IsOutOfService() {
		fmt.Println("[GPT SCHEDULER]", gpt.Token, "-> OOS")
		outOfServiceCh <- gpt
	} else {
		fmt.Println("[GPT SCHEDULER]", gpt.Token, "-> OOS")
		outOfServiceCh <- gpt
	}
	statistic()
	return GetGPT()
}
