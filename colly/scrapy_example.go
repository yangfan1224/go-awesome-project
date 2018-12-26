package colly

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/axgle/mahonia"
	"github.com/gocolly/colly"
	"log"
	"os"
	"strconv"
)

type IndexJson struct {
	Index ID `json:"index"`
}
type ID struct {
	IndexID string `json:"_id"`
}

type HotWord struct {
	Category string `json:"category"`
	Buzz string `json:"buzz"`
	Keyword string `json:"keyword"`
	KeywordSuggest Suggest `json:"keyword_suggest"`
	Url string `json:"url"`
}


type Suggest struct{
		Input []string `json:"input"`
		Weight int `json:"weight"`
}


func Md5Hex(keyword string) string {
	h := md5.New()
	h.Write([]byte(keyword))
	md5out := h.Sum(nil)
	dst := make([]byte, hex.EncodedLen(len(md5out)))
	hex.Encode(dst, md5out)
	return fmt.Sprintf("%s", dst)
}

func getTitleDescriptionKeywords(){
	//colly.Debugger(&debug.LogDebugger{})
	count := 0
	catcount := 0

	fName := "hotword.json"
	file, err := os.Create(fName)
	if err != nil {
		log.Fatalf("Cannot create file %q: %s\n", fName, err)
		return
	}
	defer file.Close()

	enc := json.NewEncoder(file)
	//enc.SetIndent("", "  ")

	c := colly.NewCollector()
	c2 := c.Clone()
	// Find and visit all links
	c.OnHTML("title", func(e *colly.HTMLElement) {
		fmt.Printf("title=%s\n",mahonia.NewDecoder("gbk").ConvertString(e.Text))
	})

	c.OnHTML("meta[name=description]", func(e *colly.HTMLElement) {
		fmt.Printf("content=%s\n",e.Attr("content"))
	})

	c.OnHTML("meta[name=keywords]", func(e *colly.HTMLElement) {
		fmt.Printf("keywords=%s\n",e.Attr("content"))
	})

	c.OnHTML("div.all-list", func(e *colly.HTMLElement) {
		//fmt.Println(mahonia.NewDecoder("gbk").ConvertString(e.Text))
		//fmt.Println(e.Attr("href"))
		var category string
		var buzz string
		var url string

		e.ForEach("h3.title", func(i int, e2 *colly.HTMLElement) {
			category = mahonia.NewDecoder("gbk").ConvertString(e2.ChildText("a"))
		})

		e.ForEach("div.links", func(i int, e2 *colly.HTMLElement) {
			e2.ForEach("a", func(i int, e3 *colly.HTMLElement) {
				buzz = mahonia.NewDecoder("gbk").ConvertString(e3.Text)
				url = e.Request.AbsoluteURL(e3.Attr("href"))
				//fmt.Printf("category=%s, buzz=%s, href=%s\n", category, buzz, url)

				ctx := e.Request.Ctx
				ctx.Put("category",category)
				ctx.Put("buzz",buzz)
				c2.Request("GET", url, nil, ctx, nil)
				catcount++
			})
		})
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.Host)
	})

	c.OnResponse(func(response *colly.Response) {
		//fmt.Println(string(response.Body))
	})

	c2.OnHTML("tr", func(e *colly.HTMLElement) {
		var url string
		var keyword string
		var weight string
		e.ForEach("td.keyword", func(i int, e2 *colly.HTMLElement) {
			url = e.Request.AbsoluteURL(e2.ChildAttr("a.list-title","href"))
			keyword = mahonia.NewDecoder("gbk").ConvertString(e2.ChildText("a.list-title"))
		})
		e.ForEach("td.last", func(i int, e2 *colly.HTMLElement) {
			weight = e2.ChildText("span.icon-rise, span.icon-fall, span.icon-fair")
		})

		if keyword != "" && weight != ""{
			count ++
			weightInt,err := strconv.Atoi(weight)
			if err != nil{
				fmt.Printf("strconv.Atoi ERROR, err is %s", err)
				weightInt = 1
			}
			//fmt.Printf("category=%s, buzz=%s, keyword=%s, weight=%s, url=%s\n", e.Request.Ctx.Get("category"), e.Request.Ctx.Get("buzz"), keyword, weight, url)

			index := &IndexJson{ID{IndexID:Md5Hex(keyword)}}
			content := &HotWord{
				Category:e.Request.Ctx.Get("category"),
				Buzz:e.Request.Ctx.Get("buzz"),
				Keyword:keyword,
				KeywordSuggest:Suggest{
					Input:[]string{keyword},
					Weight:weightInt,
				},
				Url: url,
			}
			if err := enc.Encode(index); err != nil{
				fmt.Errorf("encode index error %s", err)
			}
			if err := enc.Encode(content); err != nil{
				fmt.Errorf("encode content error %s", err)
			}
		}
	})
	c2.OnResponse(func(response *colly.Response) {
		//fmt.Println(string(response.Body))
	})

	c.Visit("http://top.baidu.com/boards?fr=topindex")
	fmt.Println(count)
	fmt.Println(catcount)
}