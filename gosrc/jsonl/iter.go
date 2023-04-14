package jsonl

import (
	"fmt"
	"github.com/KevinZonda/GoX/pkg/iox"
	"github.com/zinccat/dolly_chinese/gosrc/text"
	"sync"
	"time"
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

func (i *DollyIter) Save() error {
	fmt.Println("[保存] 保存到文件")
	txt := ToText(i.Models)
	err := iox.WriteAllText(text.JSONL_FILE, txt)
	if err != nil {
		fmt.Printf("[错误] 保存失败 -> %v\n", err)
		return err
	}
	return err
}

func (i *DollyIter) SaveDaemon(t time.Duration) {
	go func() {
		for {
			time.Sleep(t)
			err := i.Save()
			if err != nil {
				fmt.Printf("[错误] 保存失败 -> %v\n", err)
			}
		}
	}()
}

func (i *DollyIter) Translate(trans func(string) (string, error)) {
	var wg sync.WaitGroup
	count := 0
	for idx, mod := range i.Models {
		if !mod.NeedTranslate() {
			fmt.Println("[SKIP] 已翻译，跳过 IDX:", idx)
			continue
		}
		wg.Add(1)
		count++
		go func(_idx int) {
			defer wg.Done()
			mod := &i.Models[_idx]
			mod.translate(trans, _idx)

		}(idx)
		if count >= text.MAX_CONC {
			wg.Wait()
			i.Save()
			count = 0
			wg = sync.WaitGroup{}
		}
	}
	wg.Wait()
	i.Save()
}

func (m *DollyModel) translate(trans func(string) (string, error), idx int) {
	fmt.Printf("[翻译] 正在翻译第 %d 条\n", idx+1)
	var err error
	m.Instruction, err = trans(m.Instruction)
	if err != nil {
		fmt.Printf("[错误] %d.Instruction -> %v\n", idx, err)
		m.Instruction = "错误！" + m.Instruction + err.Error()
	}
	m.Context, err = trans(m.Context)
	if err != nil {
		fmt.Printf("[错误] %d.Context -> %v\n", idx, err)
		m.Context = "错误！" + m.Context + err.Error()
	}
	m.Response, err = trans(m.Response)
	if err != nil {
		fmt.Printf("[错误] %d.Response -> %v\n", idx, err)
		m.Response = "错误！" + m.Response + err.Error()
	}
	m.Translated = true
}
