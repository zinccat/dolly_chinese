package jsonl

import (
	"encoding/json"
	"strings"
)

type DollySet []DollyModel

type DollyModel struct {
	Instruction string `json:"instruction"`
	Context     string `json:"context"`
	Response    string `json:"response"`
	Category    string `json:"category"`
	Translated  bool   `json:"translated,default:false"`
}

func (m *DollyModel) NeedTranslate() bool {
	if !m.Translated {
		return true
	}
	if strings.HasPrefix(m.Instruction, "错误！") {
		return true
	}
	if strings.HasPrefix(m.Context, "错误！") {
		return true
	}
	if strings.HasPrefix(m.Response, "错误！") {
		return true
	}
	return false
}

func FromText(s string) []DollyModel {
	lines := strings.Split(s, "\n")
	models := make([]DollyModel, len(lines))
	for i, line := range lines {
		var m DollyModel
		if err := json.Unmarshal([]byte(line), &m); err != nil {
			panic(err)
		}
		models[i] = m
	}
	return models
}

func ToText(models []DollyModel) string {
	lines := make([]string, len(models))
	for i, m := range models {
		b, err := json.Marshal(m)
		if err != nil {
			panic(err)
		}
		lines[i] = string(b)
	}
	return strings.Join(lines, "\n")
}
