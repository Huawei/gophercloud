package quotas

import (
	"github.com/Huawei/gophercloud"
)

func ListURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("quotas")
}
