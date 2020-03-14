package main

import (
	"encoding/json"
	"fmt"
	"github.com/Huawei/gophercloud"
	"github.com/Huawei/gophercloud/functiontest/common"
	"github.com/Huawei/gophercloud/openstack"
	"github.com/Huawei/gophercloud/openstack/bss/v1/payperuseresource"
)

func main() {
	fmt.Println("payperuseresource start...")

	//打开debug日志
	gophercloud.EnableDebug = true

	provider, err := common.AuthToken()
	//provider, err := common.AuthAKSK()
	if err != nil {
		fmt.Println("get provider client failed")
		fmt.Println(err.Error())
		return
	}
	fmt.Println("auth success!")

	// 初始化服务的client
	sc, err := openstack.NewBSSV1(provider, gophercloud.EndpointOpts{})
	if err != nil {
		fmt.Println("get bss client failed")
		fmt.Println(err.Error())
		return
	}

	TestQueryCustomerResource(sc)

	fmt.Println("payperuseresource end...")
}

func TestQueryCustomerResource(client *gophercloud.ServiceClient) {
	opts := payperuseresource.QueryCustomerResourceOpts{
		CustomerResourceId:   "",
		CustomerId:           "123",
		RegionCode:           "",
		CloudServiceTypeCode: "0",
		ResourceIds:          nil,
		ResourceName:         "",
		StartTimeBegin:       "",
		StartTimeEnd:         "",
		PageNo:               0,
		PageSize:             0,
	}
	detailRsp,err := payperuseresource.QueryCustomerResource(client, opts).Extract()

	if err != nil {
		fmt.Println("err:", err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	bytes, _ := json.MarshalIndent(detailRsp, "", " ")
	fmt.Println(string(bytes))
	fmt.Println("TestQueryResources success")
}