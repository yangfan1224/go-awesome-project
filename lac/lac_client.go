package lac

import (
	"fmt"
	"github.com/TarsCloud/TarsGo/tars"
	"github.com/yangfan1224/go-awesome-project/lac/LacWordTag"
)
var comm *tars.Communicator

func lacTaging() int {
	comm = tars.NewCommunicator()
	obj := "WordSegmentation.LacTagServer.LacServiceObj@tcp -h 192.168.1.78 -p 25152 -t 60000"
	app := new(LacWordTag.LacService)
	/*
	 // if your service has been registered at tars registry
	 comm = tars.NewCommunicator()
	 obj := "TestApp.HelloGo.SayHelloObj"
	 // tarsregistry service at 192.168.1.1:17890
	 comm.SetProperty("locator", "tars.tarsregistry.QueryObj@tcp -h 192.168.1.1 -p 17890")
	*/

	comm.StringToProxy(obj, app)
	var resp []LacWordTag.TAG
	ret, err := app.LacTag("无人驾驶，自主实现目标引导；障碍规避，合理进行路径选择；智能操作，自动完成货物装卸……不久前，由比亚迪与新加坡科技工程有限公司联合研发的无人驾驶纯电动叉车，在新加坡正式与公众亮相，在物流运输等行业，掀起了一股新能源的绿色变革风暴", &resp)
	if err != nil {
		fmt.Println(err)
		return -1
	}
	fmt.Println("ret: ", ret, "resp: ", resp)
	return 0
}