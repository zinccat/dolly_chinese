package main

import (
	"github.com/KevinZonda/GoX/pkg/iox"
	"github.com/hjson/hjson-go/v4"
	"github.com/zinccat/dolly_chinese/gosrc/gpt"
	"github.com/zinccat/dolly_chinese/gosrc/model"
	"github.com/zinccat/dolly_chinese/gosrc/shared"
)

func main() {
	shared.Cfg = readCfg()
	gpt.Init()
}

func readCfg() model.CfgModel {
	txt, err := iox.ReadAllText("config.hjson")
	if err != nil {
		panic(err)
	}
	cfg := model.CfgModel{}
	err = hjson.Unmarshal([]byte(txt), &cfg)
	if err != nil {
		panic(err)
	}
	return cfg
}
