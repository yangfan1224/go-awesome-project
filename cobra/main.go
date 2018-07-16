package main

import (
	"fmt"
	"os"
	"github.com/spf13/viper"
	"os/signal"
	"syscall"
	"time"
	"go-awesome-project/cobra/cmd"
)

func main() {
	rootCmd := cmd.NewRootCmd()

	if err:= rootCmd.Execute(); err != nil{
		fmt.Println(err)
		os.Exit(1)
	}

	author ,_:= rootCmd.PersistentFlags().GetString("author")
	println("author is ", author)
	fmt.Printf("author is %s \n", viper.GetString("author"))

	c := make(chan os.Signal)
	//监听指定信号 ctrl+c kill
	signal.Notify(c, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGUSR1, syscall.SIGUSR2)
	go func() {
		for s := range c {
			switch s {
			case syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT:
				fmt.Println("退出", s)
				ExitFunc()
			case syscall.SIGUSR1:
				fmt.Println("usr1", s)
			case syscall.SIGUSR2:
				fmt.Println("usr2", s)
			default:
				fmt.Println("other", s)
			}
		}
	}()

	fmt.Println("进程启动...")
	sum := 0
	for {
		sum++
		fmt.Println("sum:", sum)
		time.Sleep(time.Second)
	}
}

func ExitFunc()  {
	fmt.Println("开始退出...")
	fmt.Println("执行清理...")
	fmt.Println("结束退出...")
	os.Exit(0)
}
