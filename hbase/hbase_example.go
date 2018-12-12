package hbase

import (
	"context"
	"fmt"
	"github.com/tsuna/gohbase"
	"github.com/tsuna/gohbase/filter"
	"github.com/tsuna/gohbase/hrpc"
)

// Get s specific cell
var client gohbase.Client

func init(){
	client = gohbase.NewClient("node03,node04,node05")
}

func GetAEntireRow(){

	// Values maps a ColumnFamily -> Qualifiers -> Values.
	getRequest, err := hrpc.NewGetStr(context.Background(), "DMP_IDMAPPING", "0$000735BF9765ACAB9F2C392C7674E4A7$1")
	if err != nil {
		fmt.Printf("hrpc NewGetStr error, err is: %v \n", err)
		return
	}
	getRsp, err := client.Get(getRequest)
	if err != nil {
		fmt.Printf("client.Get(getRequest) error, err is: %v \n", err)
		return
	}
	fmt.Println(getRsp)
}

func ScanWithFilter(){
	pFilter := filter.NewTimestampsFilter([]int64{1536422400000,1536508800000})
	scanRequest, err := hrpc.NewScanStr(context.Background(), "DMP_IDMAPPING",
		hrpc.Filters(pFilter))
	if err != nil{
		fmt.Printf("hrpc.NewScanStr error, err is: %v \n", err)
		return
	}
	scanRsp := client.Scan(scanRequest)
	for result, err := scanRsp.Next();err == nil;result, err = scanRsp.Next() {
		//fmt.Printf("resutls is %v", *result)
		for cell := range result.Cells {
			fmt.Println(cell)
		}
	}

}
