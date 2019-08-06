package quotas

import (
	"github.com/Huawei/gophercloud"
)

func ListURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("quotas")
}

func ListWithInstancesURL(c *gophercloud.ServiceClient, scalingGroupId string) string {
	return c.ServiceURL("quotas", scalingGroupId)
}
