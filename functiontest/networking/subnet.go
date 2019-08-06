package main

import (
	"fmt"
	"encoding/json"

	"github.com/Huawei/gophercloud"
	"github.com/Huawei/gophercloud/functiontest/common"
	"github.com/Huawei/gophercloud/openstack"
	"github.com/Huawei/gophercloud/openstack/networking/v2/subnets"
)

var subnetID string

func main() {
	fmt.Println("main start...")

	//provider, err := common.AuthAKSK()
	provider, err := common.AuthToken()
	if err != nil {
		fmt.Println("get provider client failed")
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	sc, err := openstack.NewNetworkV2(provider, gophercloud.EndpointOpts{})
	if err != nil {
		fmt.Println("get network client failed")
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	TestSubnetList(sc)
	TestSubnetCreate(sc)
	TestSubnetGet(sc)
	TestSubnetUpdate(sc)
	TestSubnetDelete(sc)

	fmt.Println("main end...")
}


func TestSubnetList(sc *gophercloud.ServiceClient) {
	allpages, err := subnets.List(sc, subnets.ListOpts{}).AllPages()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	subnet, err := subnets.ExtractSubnets(allpages)
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println("Test get subnet list success!")
	p, _ := json.MarshalIndent(subnet, "", " ")
	fmt.Println(string(p))
}

func TestSubnetGet(sc *gophercloud.ServiceClient) {
	subnet, err := subnets.Get(sc, subnetID).Extract()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println("Test get subnet detail success!")
	p, _ := json.MarshalIndent(subnet, "", " ")
	fmt.Println(string(p))
}

func TestSubnetCreate(sc *gophercloud.ServiceClient) {
	opts := subnets.CreateOpts{
		Name:"testsubnet",
		NetworkID:"021d431d-d430-41fc-a6df-9ce50b9e8169",
		CIDR:"192.168.1.0/24",
	}

	subnet, err := subnets.Create(sc, opts).Extract()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println("Test create subnet success!")
	subnetID=subnet.ID
	p, _ := json.MarshalIndent(subnet, "", " ")
	fmt.Println(string(p))
}

func TestSubnetUpdate(sc *gophercloud.ServiceClient) {
	opts := subnets.UpdateOpts{
		Name:"testsubnet2",
	}

	subnet,err := subnets.Update(sc, subnetID,opts).Extract()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println("Test update subnet success!")
	p, _ := json.MarshalIndent(subnet, "", " ")
	fmt.Println(string(p))
}

func TestSubnetDelete(sc *gophercloud.ServiceClient) {
	err := subnets.Delete(sc, subnetID).ExtractErr()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println("Test delete subnet success!")
}