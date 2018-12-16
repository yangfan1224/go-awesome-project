	package benchmark

import (
	"time"
	"fmt"
	"strings"
)

func SimpleJoin(){
	str := "yangfann1234567890,yangfann1234567890,yangfann1234567890,yangfann1234567890,yangfann1234567890"
	var strArray []string
	for i := 0; i < 100000; i++ {
		strArray = append(strArray,str)
	}

	var joinStr string
	var sep string

	start := time.Now()
	for _,str := range  strArray{
		joinStr += sep + str
	}

	secs := time.Since(start).Seconds()
	fmt.Printf("string + string: %fs", secs)

	start = time.Now()
	strings.Join(strArray,"")
	secs = time.Since(start).Seconds()
	fmt.Printf("strings.Join: %fs", secs)
}
