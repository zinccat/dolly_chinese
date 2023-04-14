package main

import (
	"github.com/KevinZonda/GoX/pkg/iox"
	"github.com/hjson/hjson-go/v4"
	"github.com/zinccat/dolly_chinese/gosrc/gpt"
	"github.com/zinccat/dolly_chinese/gosrc/jsonl"
	"github.com/zinccat/dolly_chinese/gosrc/model"
	"github.com/zinccat/dolly_chinese/gosrc/shared"
	"github.com/zinccat/dolly_chinese/gosrc/text"
)

func main() {
	shared.Cfg = readCfg()
	loadJsonL()
	gpt.Init()
	shared.Data.Translate(gpt.Translate)
}

func readCfg() model.CfgModel {
	txt, err := iox.ReadAllText(text.CONFIG_FILE)
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
	txt, err := iox.ReadAllText(text.JSONL_FILE)
	if err != nil {
		panic(err)
	}
	jl := jsonl.FromText(txt)
	shared.Data = jsonl.InitDollyIter(jl)
}
