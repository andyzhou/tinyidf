package main

import (
	"github.com/andyzhou/tinyidf"
	"log"
	"reflect"
	"strings"
)

func main() {
	var got, want []string

	dt := tinyidf.NewTokenizer("./dict/dict.txt")

	got = dt.CutAll("我来到北京清华大学", true)
	want = []string{"我", "来到", "北京", "清华", "清华大学", "华大", "大学"}
	logResult("Full mode", got)
	checkResult(got, want)

	got = dt.Cut("我来到北京清华大学", true)
	gotForSearch := dt.CutForSearch("我来到北京清华大学", true)
	want = []string{"我", "来到", "北京", "清华大学"}
	logResult("Accurate mode", got)
	logResult("Accurate mode gotForSearch", gotForSearch)
	checkResult(got, want)

	got = dt.Cut("他来到了网易杭研大厦", true)
	want = []string{"他", "来到", "了", "网易", "杭研", "大厦"}
	logResult("New word '杭研'", got)
	checkResult(got, want)

	got = dt.Cut("小明硕士毕业于中国科学院计算所，后在日本京都大学深造", true)
	want = []string{"小明", "硕士", "毕业", "于", "中国科学院", "计算所", "，",
		"后", "在", "日本京都大学", "深造"}
	logResult("Accurate mode", got)
	checkResult(got, want)

	got = dt.CutForSearch("小明硕士毕业于中国科学院计算所，后在日本京都大学深造", true)
	want = []string{"小明", "硕士", "毕业", "于", "中国", "科学", "学院", "科学院",
		"中国科学院", "计算", "计算所", "，", "后", "在", "日本", "京都", "大学",
		"日本京都大学", "深造"}
	logResult("Search mode", got)
	checkResult(got, want)

	got = dt.Cut("但是有考据癖的人也当然不肯错过索隐的杨会、放弃附会的权利的。", true)
	want = []string{"但是", "有", "考据", "癖", "的", "人", "也", "当然", "不肯",
		"错过", "索隐", "的", "杨会", "、", "放弃", "附会", "的", "权利", "的", "。"}
	logResult("Accurate mode", got)
	checkResult(got, want)

	got = dt.CutAll("署名Ｐａｔｒｉｃ　Ｍａｈｏｎｅｙ", true)
	want = []string{"署名", "Ｐａｔｒｉｃ", "\u3000", "Ｍａｈｏｎｅｙ"}
	logResult("Accurate mode", got)
	checkResult(got, want)
}


func logResult(prefix string, res []string) {
	buf := strings.Builder{}
	if len(prefix) > 0 {
		buf.WriteString(prefix)
		buf.WriteString(": ")
	}
	buf.WriteString("[")
	n := len(res)
	for i, word := range res {
		buf.WriteString(word)
		if i != n-1 {
			buf.WriteString(" / ")
		}
	}
	buf.WriteString("]\n")
	log.Printf("%v", buf.String())
}

func checkResult(got, want []string) {
	if !reflect.DeepEqual(got, want) {
		log.Printf("got != want, got: %v, want: %v\n", got, want)
	}else{
		log.Printf("check result pass!\n\n")
	}
}