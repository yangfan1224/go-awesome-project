package pegasus
import (
	"context"
	"github.com/XiaoMi/pegasus-go-client/pegasus"
	"fmt"
)

func AccessPegasus() error {
	cfg := pegasus.Config {
		MetaServers: []string{"192.168.1.104:34601", "192.168.1.129:34601", "192.168.1.78:34601"},
	}

	client := pegasus.NewClient(cfg)

	tb, err := client.OpenTable(context.Background(), "userlabelTest")
	defer  client.Close()
	if err != nil{
		return fmt.Errorf("OpenTable ERROR: %v", err)
	}

	err = tb.Set(context.Background(), []byte("h1"), []byte("s1"), []byte("v1"))
	if err != nil{
		fmt.Errorf("table set error: %v", err)
	}

	value, err := tb.Get(context.Background(), []byte("h1"), []byte("s1"))
	if err != nil{
		return fmt.Errorf("get value error: %v", err)

	}

	println(string(value[:]))

	return nil
}
