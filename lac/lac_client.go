package lac

import (
	"fmt"
	"github.com/TarsCloud/TarsGo/tars"
	"github.com/yangfan1224/go-awesome-project/lac/LacWordTag"
)
var comm *tars.Communicator

func lacTaging() int {
	comm = tars.NewCommunicator()
	obj := "WordSegmentation.LacTagServer.LacServiceObj"
	app := new(LacWordTag.LacService)
	/*
	 // if your service has been registered at tars registry
	 comm = tars.NewCommunicator()
	 obj := "TestApp.HelloGo.SayHelloObj"
	 // tarsregistry service at 192.168.1.1:17890
	 comm.SetProperty("locator", "tars.tarsregistry.QueryObj@tcp -h 192.168.1.1 -p 17890")
	*/
	comm.SetProperty("locator", "tars.tarsregistry.QueryObj@tcp -h 192.168.1.78 -p 17890")
	comm.StringToProxy(obj, app)
	var resp []LacWordTag.TAG
	ret, err := app.LacTag("四川深度在线广告传媒有限公司", &resp)
	if err != nil {
		fmt.Println(err)
		return -1
	}
	fmt.Println("ret: ", ret, "resp: ", resp)
	return 0
}