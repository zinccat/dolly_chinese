package model

type CfgModel struct {
	Tokens  []string `json:"tokens"`
	MaxConc int      `json:"max_conc,default:10"`
}
