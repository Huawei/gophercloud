package datastores

import "github.com/Huawei/gophercloud"

func listURL(sc *gophercloud.ServiceClient, databasename string) string {
	return sc.ServiceURL("datastores", databasename)
}
