package events

import (
	"github.com/Huawei/gophercloud"
)

func createURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("events")
}
