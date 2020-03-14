package storagetype

import "github.com/Huawei/gophercloud"

func listURL(sc *gophercloud.ServiceClient, databasename string) string {
	return sc.ServiceURL("storage-type", databasename)
}
