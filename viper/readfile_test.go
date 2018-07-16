package viper

import (
	"testing"
	"github.com/spf13/viper"
)

func TestReadConfig(t *testing.T) {
	author := ReadConfig()
	if author != "tom" {
		t.Error("author is not tom!")
	}

	layouts := viper.GetString("LayoutDir")
	if layouts != "layouts" {
		t.Error("LayoutDir is not layouts!")
	}

	//for i := 0; i < 10000000000; i++ {
	//	time.Sleep(10000)
	//}

}
