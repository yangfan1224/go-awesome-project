package pipe

import (
	"io"
	"os"
	"strconv"
	"time"
)

func PipeExample()(error){
	pipeR,pipeW := io.Pipe()
	go func() {
		for i :=0; i < 100; i++{
			pipeW.Write([]byte (strconv.Itoa(i) + "\n"))
			time.Sleep(100 * time.Millisecond)
		}
		pipeW.Close()
	}()

	io.Copy(os.Stdout, pipeR)
	return nil
}