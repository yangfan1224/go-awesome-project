package colly

import (
	"fmt"
	"github.com/gocolly/colly"
)

func getTitleDescriptionKeywords(){
	//colly.Debugger(&debug.LogDebugger{})
	c := colly.NewCollector()

	// Find and visit all links
	c.OnHTML("title", func(e *colly.HTMLElement) {
		fmt.Printf("title=%s\n",e.Text)
	})

	c.OnHTML("meta[name=description]", func(e *colly.HTMLElement) {
		fmt.Printf("content=%s\n",e.Attr("content"))
	})

	c.OnHTML("meta[name=keywords]", func(e *colly.HTMLElement) {
		fmt.Printf("keywords=%s\n",e.Attr("content"))
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.OnResponse(func(response *colly.Response) {
		fmt.Println("OnResponse is called.")
	})

	c.Visit("http://news.cctv.com/2018/12/15/ARTIzQ2H84mAvbUbI7weq0UH181215.shtml")
}