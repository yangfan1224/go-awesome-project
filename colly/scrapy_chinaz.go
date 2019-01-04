package colly

import (
	"encoding/json"
	"fmt"
	"github.com/gocolly/colly"
	"log"
	"os"
	"strconv"
	"strings"
)
type UrlDoc struct {
	Title string `json:"title"`
	Url string `json:"url"`
	IconUrl string `json:"icon"`
	IconType int `json:"icon_type"`
	Introduction string `json:"intro"`
	Weight int `json:"weight"`
}

func ScrapyChinaz(){
	fName := "urldoc.json"
	file, err := os.Create(fName)
	if err != nil {
		log.Fatalf("Cannot create file %q: %s\n", fName, err)
		return
	}
	defer file.Close()

	enc := json.NewEncoder(file)

	c := colly.NewCollector()
	c2 := c.Clone()

	UrlDocList := make([]UrlDoc,0)
	c.OnHTML("li.clearfix", func(e *colly.HTMLElement) {
		var title string
		var url string
		var intro string
		var weight int
		e.ForEach(".CentTxt", func(i int, e2 *colly.HTMLElement) {
			title = e2.ChildText("a.pr10")
			url = e2.ChildText("span.col-gray")
			if strings.HasPrefix(url,"www."){
				url = "http://" + url
			}else{
				if len(strings.Split(url,".")) > 2 {
					url = "http://" + url
				}else{
					url = "http://www." + url
				}
			}
			intro = e2.ChildText("p.RtCInfo")
			intro = strings.Replace(intro,`网站简介：`,"",1)

		})

		e.ForEach(".RtCRateCent", func(i int, e2 *colly.HTMLElement) {
			var err error
			score := e2.ChildText("span")
			if weight,err = strconv.Atoi(strings.Replace(score,"得分:","",1)); err != nil {
				fmt.Printf("strconv.Atoi error: %s", err)
			}
		})

		//fmt.Printf("title=%s, url=%s, weight=%d, intro=%s\n", title, url, weight, intro)
		urlDoc :=UrlDoc{Title:title,
			Url:url,
			Introduction:intro,
			Weight:weight,
			IconType: 1,
		}
		UrlDocList = append(UrlDocList,urlDoc)
		ctx := e.Request.Ctx
		ctx.Put("url",url)
		c2.Request("GET", url, nil, ctx, nil)
	})

	c.OnResponse(func(response *colly.Response) {
		//fmt.Println(string(response.Body))
	})

	c.OnHTML("div.ListPageWrap", func(e *colly.HTMLElement) {
		e.ForEach("a:contains('>')", func(i int, e2 *colly.HTMLElement) {
			href := e2.Attr("href")
			fmt.Println(e.Request.AbsoluteURL(href))
			c.Visit(e.Request.AbsoluteURL(href))
		})
	})

	c.OnRequest(func(request *colly.Request) {
		if request.URL.String() =="http://top.chinaz.com/all/index_34.html" {
			c.Visit("http://top.chinaz.com/all/index_35.html")
		}
	})

	c.OnError(func(r *colly.Response, e error) {
		fmt.Printf("err: %s, at request: %s Now Retrying\n", e, r.Request.URL.String() )
		var url string
		if url = r.Request.URL.String(); strings.HasPrefix(url,"https://"){
			url  = strings.Replace(url,"https","http",1)
		}else{
			url  = strings.Replace(url,"http","https",1)
		}
		r.Request.Visit(url)
	})

	var iconDownload = make (map[string]string)

	c2.OnHTML(`link[rel="shortcut icon"]`, func(e *colly.HTMLElement) {
		url := e.Request.Ctx.Get("url")
		if _, ok :=iconDownload[url]; ok {
			return
		}
		iconUrl := e.Request.AbsoluteURL(e.Attr("href"))
		fmt.Println(iconUrl)
		iconDownload[url] = iconUrl
	})

	c2.OnError(func(r *colly.Response, e error) {
		fmt.Printf("err: %s, at request: %s Now Retrying\n", e, r.Request.URL.String() )
		var url string
		if url = r.Request.URL.String(); strings.HasPrefix(url,"https://"){
			url  = strings.Replace(url,"https","http",1)
		}else{
			url  = strings.Replace(url,"http","https",1)
		}
		r.Request.Visit(url)
	})

	c.Visit("http://top.chinaz.com/all/index.html")

	for i := range UrlDocList{
		doc := UrlDocList[i]
		if icon, ok :=iconDownload[doc.Url]; ok {
			doc.IconUrl = icon
		}
		urlid := fmt.Sprintf("%s%d",Md5Hex(doc.Url), len(doc.Url))
		index := &IndexJson{ID{IndexID:urlid}}

		if err := enc.Encode(index); err != nil{
			fmt.Printf("encode index error %s", err)
		}
		if err := enc.Encode(&doc); err != nil{
			fmt.Printf("encode urlDoc error %s\n", err)
		}
	}
}
