package jsonl

import "github.com/zinccat/dolly_chinese/gosrc/gpt"

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

func (i *DollyIter) Translate() {
	var err error
	for i.HasNext() {
		m := i.Value()
		m.Context, err = gpt.Translate(m.Context)
		if err != nil {
			m.Context = "错误！" + m.Context + err.Error()
		}
		m.Response, err = gpt.Translate(m.Response)
		if err != nil {
			m.Response = "错误！" + m.Response + err.Error()
		}
		m.Translated = true
		i.Models[i.Index] = m
		i.Next()
	}
}
