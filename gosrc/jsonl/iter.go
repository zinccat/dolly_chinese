package jsonl

import (
	"fmt"
	"github.com/KevinZonda/GoX/pkg/iox"
	"github.com/zinccat/dolly_chinese/gosrc/text"
)

type DollyIter struct {
	Models []DollyModel
	Index  int
}

func InitDollyIter(set DollySet) *DollyIter {
	return &DollyIter{
		Models: set,
		Index:  0,
	}
}

func (i *DollyIter) Value() DollyModel {
	return i.Models[i.Index]
}

func (i *DollyIter) Reset() {
	i.Index = 0
}

func (i *DollyIter) HasNext() bool {
	return i.Index < len(i.Models)
}

func (i *DollyIter) Next() {
	//m := i.Value()
	i.Index++
}

func (i *DollyIter) Save() error {
	txt := ToText(i.Models)
	return iox.WriteAllText(text.JSONL_FILE, txt)
}

func (i *DollyIter) Translate(trans func(string) (string, error)) {
	var err error
	for i.HasNext() {
		fmt.Printf("正在翻译第 %d 条\n", i.Index+1)
		m := i.Value()
		if m.Translated {
			fmt.Println("已翻译，跳过")
			i.Next()
			continue
		}
		m.Instruction, err = trans(m.Instruction)
		if err != nil {
			fmt.Printf("[错误] %d.Instruction -> %v\n", i.Index, err)
			m.Instruction = "错误！" + m.Instruction + err.Error()
		}
		m.Context, err = trans(m.Context)
		if err != nil {
			fmt.Printf("[错误] %d.Context -> %v\n", i.Index, err)
			m.Context = "错误！" + m.Context + err.Error()
		}
		m.Response, err = trans(m.Response)
		if err != nil {
			fmt.Printf("[错误] %d.Response -> %v\n", i.Index, err)
			m.Response = "错误！" + m.Response + err.Error()
		}
		m.Translated = true
		i.Models[i.Index] = m
		i.Save()
		i.Next()
	}
}
