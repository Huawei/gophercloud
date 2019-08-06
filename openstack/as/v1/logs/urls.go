package logs

import (
	"github.com/Huawei/gophercloud"
)

func ListURL(c *gophercloud.ServiceClient, scalingGroupId string) string {
	return c.ServiceURL("scaling_activity_log", scalingGroupId)
}
