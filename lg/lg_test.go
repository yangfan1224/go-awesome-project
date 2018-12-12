package lg

import (
	"log"
	"os"
	"testing"
)

func TestLogf(t *testing.T) {
	logger :=  log.New(os.Stderr, "[lg_test]", log.Ldate|log.Ltime)
	Logf(logger, DEBUG,INFO, "time is %d",10)

}
