package publicips

import "github.com/Huawei/gophercloud"

func CreateURL(c *gophercloud.ServiceClient)string{
	return c.ServiceURL("publicips")
}