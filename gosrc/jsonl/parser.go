package jsonl

import (
	"encoding/json"
	"strings"
)

type Dollymodel struct {
	Instruction string `json:"instruction"`
	Context     string `json:"context"`
	Response    string `json:"response"`
	Category    string `json:"category"`
}

func FromText(s string) []Dollymodel {
	lines := strings.Split(s, "\n")
	models := make([]Dollymodel, len(lines))
	for i, line := range lines {
		var m Dollymodel
		if err := json.Unmarshal([]byte(line), &m); err != nil {
			panic(err)
		}
		models[i] = m
	}
	return models
}

func ToText(models []Dollymodel) string {
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
