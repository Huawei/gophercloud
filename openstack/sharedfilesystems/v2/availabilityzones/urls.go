package availabilityzones

import "github.com/Huawei/gophercloud"

func listURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("os-availability-zone")
}
