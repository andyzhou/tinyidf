package main

import (
	"github.com/andyzhou/tinyidf"
	"log"
	"sort"
)

func main()  {
	//try load relate data
	dt := tinyidf.NewTokenizer("./dict/dict.txt")
	tfIdf := tinyidf.NewTfIdf(dt, "./dict/idf.txt")

	//setup source string
	s := "此外，公司拟对全资子公司吉林欧亚置业有限公司增资4.3亿元，增资后，吉林欧亚" +
		"置业注册资本由7000万元增加到5亿元。吉林欧亚置业主要经营范围为房地产开发及" +
		"百货零售等业务。目前在建吉林欧亚城市商业综合体项目。2013年，实现营业收入" +
		"0万元，实现净利润-139.13万元。"

	//try add stop words
	tfIdf.AddStopWord("")

	//try extract key words
	kws := tfIdf.Extract(s, 20)

	got := make([]string, 0, len(kws))
	for _, kw := range kws {
		log.Printf("word:%v, score:%v\n", kw.GetWord(), kw.GetScore())
		got = append(got, kw.GetWord())
	}

	want := []string{"欧亚", "吉林", "置业", "万元", "增资", "4.3", "7000", "2013",
		"139.13", "实现", "综合体", "经营范围", "亿元", "在建", "全资", "注册资本",
		"百货", "零售", "子公司", "营业"}

	sort.Strings(got)
	sort.Strings(want)
	log.Println("got:", got)
	log.Println("want:", want)
}
