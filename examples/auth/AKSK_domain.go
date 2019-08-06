package examples

import (
	"fmt"
	"github.com/Huawei/gophercloud/auth/aksk"
	"github.com/Huawei/gophercloud"
	"github.com/Huawei/gophercloud/openstack"
	"github.com/Huawei/gophercloud/openstack/identity/v3/users"
)

// AuthAKSKUserList using AKSK auth method ,list users .
func AuthAKSKUserList() {

	fmt.Println("main start...")

	// init AKSK auth options
	akskOptions := aksk.AKSKOptions{
		IdentityEndpoint: "https://iam.cn-north-1.myhuaweicloud.com/v3",
		DomainID:         "replace-your-domainID",
		Cloud:           "myhuaweicloud.com",
		Region:           "cn-north-1",
		AccessKey:        "replace-your-ak",
		SecretKey:        "replace-your-sk",
	}

	//init provider client
	provider, err := openstack.AuthenticatedClient(akskOptions)
	if err != nil {
		panic(err)
	}

	// init IAM client
	iamClient, err := openstack.NewIdentityV3(provider, gophercloud.EndpointOpts{})
	if err != nil {
		fmt.Println("get IAM v3 client failed")
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	// build http request，list all users belong to this domain
	page, err := users.List(iamClient, users.ListOpts{}).AllPages()
	if err != nil {

		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	// extract http response body
	userList, err := users.ExtractUsers(page)

	if err != nil {
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	// print result
	for _, d := range userList {

		fmt.Println("user id is :", d.ID)
		fmt.Println("user name is :", d.Name)

	}

	fmt.Println("main end...")
}
