package main

import (
	"github.com/KevinZonda/GoX/pkg/iox"
	"github.com/hjson/hjson-go/v4"
	"github.com/zinccat/dolly_chinese/gosrc/gpt"
	"github.com/zinccat/dolly_chinese/gosrc/jsonl"
	"github.com/zinccat/dolly_chinese/gosrc/model"
	"github.com/zinccat/dolly_chinese/gosrc/shared"
)

func main() {
	shared.Cfg = readCfg()
	loadJsonL()
	gpt.Init()
}

func readCfg() model.CfgModel {
	txt, err := iox.ReadAllText(shared.CONFIG_FILE)
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

func loadJsonL() {
	txt, err := iox.ReadAllText(shared.JSONL_FILE)
	if err != nil {
		panic(err)
	}
	shared.Data = jsonl.FromText(txt)
}
