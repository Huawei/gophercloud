package groups

import "github.com/Huawei/gophercloud"

const rootPath = "security-groups"

func rootURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL(rootPath)
}

