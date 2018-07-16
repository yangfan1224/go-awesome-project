package kafka

import (
	"testing"
	"os"
)

func TestConsumer(t *testing.T) {
	os.Setenv("PKG_CONFIG_PATH", "/usr/lib/pkgconfig")
	os.Setenv("LD_LIBRARY_PATH", "/usr/lib/")
	Consumer()
}
