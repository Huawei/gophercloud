package main

import (
	"encoding/json"
	"fmt"
	"github.com/Huawei/gophercloud"
	"github.com/Huawei/gophercloud/auth/aksk"
	"github.com/Huawei/gophercloud/openstack"
	"github.com/Huawei/gophercloud/openstack/bssintl/v1/utilities"
)

func main() {
	//AKSK auth，initial parameter.
	opts := aksk.AKSKOptions{
		IdentityEndpoint: "https://iam.xxx.yyy.com/v3",
		AccessKey:        "{your AK string}",
		SecretKey:        "{your SK string}",
		Cloud:            "yyy.com",
		DomainID:         "{domainID}",
	}

	//initial provider client。
	provider, errAuth := openstack.AuthenticatedClient(opts)
	if errAuth != nil {
		fmt.Println("get provider client failed")
		fmt.Println(errAuth.Error())
		return
	}

	// initial client
	sc, err := openstack.NewBSSIntlV1(provider, gophercloud.EndpointOpts{})
	if err != nil {
		fmt.Println("get bss client failed")
		fmt.Println(err.Error())
		return
	}

	 SendVerificationCode(sc)

	fmt.Println("utilities end...")
}

func SendVerificationCode(client *gophercloud.ServiceClient) {
	var a = 1
	opts := utilities.SendVerificationCodeOpts{
		ReceiverType: &a,
		MobilePhone:  "123",
	}
	detailRsp,err := utilities.SendVerificationCode(client, opts).Extract()

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
	fmt.Println("VerificationCode success")
}