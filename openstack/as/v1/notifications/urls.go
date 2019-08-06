package notifications

import (
	"github.com/Huawei/gophercloud"
)

func DeleteURL(c *gophercloud.ServiceClient, scalingGroupId string, topicUrn string) string {
	return c.ServiceURL("scaling_notification", scalingGroupId, topicUrn)
}

func EnableURL(c *gophercloud.ServiceClient, scalingGroupId string) string {
	return c.ServiceURL("scaling_notification", scalingGroupId)
}

func ListURL(c *gophercloud.ServiceClient, scalingGroupId string) string {
	return c.ServiceURL("scaling_notification", scalingGroupId)
}
