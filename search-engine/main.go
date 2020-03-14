package main

import (
	"log"

	"github.com/go-ego/riot"
	"github.com/go-ego/riot/types"
)

var (
	// searcher is coroutine safe
	searcher = riot.Engine{}
)

func main() {
	// Init
	searcher.Init(types.EngineOpts{
		// Using:             4,
		NotUseGse: true,
	})
	defer searcher.Close()

	text := "新手学NLP，该从哪里入手？CMU 自然语言处理公开课 及课程辅助开源项目：NLP关键概念集 - 知乎"
	text1 := `人类行为时空特性的统计力学（一）——认识幂律分布 - 知乎`
	text2 := `专题 | Edmond Lau：如何创建良好的工程师文化？ - 知乎`

	// Add the document to the index, docId starts at 1
	searcher.Index("1", types.DocData{Content: text})
	searcher.Index("2", types.DocData{Content: text1}, false)
	searcher.IndexDoc("3", types.DocData{Content: text2}, true)

	// Wait for the index to refresh
	searcher.Flush()
	// engine.FlushIndex()

	// The search output format is found in the types.SearchResp structure
	log.Print(searcher.Search(types.SearchReq{Text: "知乎"}))
}
