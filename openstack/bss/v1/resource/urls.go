package resource

import "github.com/Huawei/gophercloud"

func listURL(client *gophercloud.ServiceClient, domainId string) string {
	return client.ServiceURL(domainId, "common/order-mgr/resources/detail")
}
