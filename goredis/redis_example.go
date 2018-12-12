package goredis

import (
	"fmt"
	"github.com/go-redis/redis"
)

var client *redis.Client

func init(){
	client = redis.NewClient(&redis.Options{
		Addr:     "192.168.1.78:6380",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}

func ExampleNewClient() {
	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
	// Output: PONG <nil>
}

func ExampleClient() {
	err := client.Set("key", "value", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := client.Get("key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)

	val2, err := client.Get("key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}
}

func ExampleFifo(){
	err := client.LPush("fifo", "1").Err()
	if err != nil {
		panic(err)
	}
	err = client.LPush("fifo", "2").Err()
	if err != nil {
		panic(err)
	}
	err = client.LPush("fifo", "3").Err()
	if err != nil {
		panic(err)
	}
	ret, err := client.RPop("fifo").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(ret)
	ret, err = client.RPop("fifo").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(ret)
	ret, err = client.RPop("fifo").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(ret)
	ret, err = client.RPop("fifo").Result()
	if err == redis.Nil {
		fmt.Println(err)
	}else{
		fmt.Println(ret)
	}
	//if ret != nil {
	//	fmt.Println(ret)
	//}


}