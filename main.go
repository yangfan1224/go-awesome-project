package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"os"
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
	//h := md5.New()
	//h.Write(([]byte)("https://golang.org/pkg/crypto/md5/#example_Sum"))
	//md5out := h.Sum(nil)
	//dst := make([]byte, hex.EncodedLen(len(md5out)))
	//hex.Encode(dst, md5out)
	//fmt.Printf("%s\n", dst)
	//
	//url, _ := url.Parse("https://golang.org/pkg/crypto/md5/#example_Sum")
	//fmt.Println(url.Host)

	csvStr := `{"_id":{"$oid":"5a7d2c04fb47f88a33d451c4"},"playedAt":{"$numberLong":"1518073392"},"playerName":"gogo-gaga","assists":0,"buffed":4,"damage":1481.0,"headShot":1,"id":"Svg7Thpx__N3LAxphvieh6Rqz5dUjLhS4Hnsqo8gu2aZjmfucAQa5f8eV-oYcj84CWo1ysLewqHx0W77Zk0BrkOQkwySi93QzmeMOSzQ4PC0gYRi32olS9zpnUDEjUAp2B2S-3dHZjk=","kills":13,"matchModel":0,"moveDistance":5471.873820000001,"offset":"eyJfaWQiOiI1YTdjNmUxZjVkZDRhNjAwMDExZmE3ZDIiLCJzdGFydGVkX2F0IjoxNTE4MTAyMTkyLCJzZWFzb24iOiIyMDE4LTAyIn0=","rank":1,"rankChange":23.319350840000002,"resurrection":0,"rideDistance":1838.11182,"season":"","serverName":0,"stuns":0,"timeSurvival":1801.987,"totalUser":89,"treated":3,"walkDistance":3633.762}`
	var competition PubgCompetition
	var buf bytes.Buffer
	buf.Write([]byte(csvStr))
	if err := json.NewDecoder(&buf).Decode(&competition); err != nil {
		fmt.Printf("json.NewDecoder error: %s", err)
		os.Exit(1)
	}
	fmt.Printf("competition is %v", competition)


}
