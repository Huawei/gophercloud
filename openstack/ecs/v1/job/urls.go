package job

import "github.com/Huawei/gophercloud"

// Querying task statuses URL
func jobURL(sc *gophercloud.ServiceClient, jobId string) string {
	return sc.ServiceURL("jobs", jobId)
}