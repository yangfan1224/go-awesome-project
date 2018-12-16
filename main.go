package main

import (
	"context"
	"fmt"
	"github.com/gocolly/colly"
	"sync"
)

type Person struct {
	Name string
	Age int
}

func worker(ctx context.Context, wg *sync.WaitGroup) error {
	defer wg.Done()

	for {
		select {
		default:
			fmt.Println("hello")
		case <-ctx.Done():
			fmt.Printf("Done err is: %v\n", ctx.Err())
			return ctx.Err()
		}
	}
}

type demo struct {

}

func (demo) Read(p []byte) (n int, err error) {
	panic("implement me")
}

func (demo) Write(p []byte) (n int, err error) {
	panic("implement me")
}

func TestDefer()(int, error){
	defer func() {
		fmt.Println("defer")
	}()
	return fmt.Println("return")
}
type PubgCompetition struct {
	username string `json:"playerName"`
	matchmode int8 `json:"matchModel"`
	matchid string `json:"id"`
	playat string `json:"playedAt"`
}

func main() {
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
