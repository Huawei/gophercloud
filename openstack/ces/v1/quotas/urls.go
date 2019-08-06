package quotas

import (
	"github.com/Huawei/gophercloud"
)

func getURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("quotas")
}
