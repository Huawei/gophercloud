package quotasets

import "github.com/Huawei/gophercloud"

const resourcePath = "os-quota-sets"

//func resourceURL(c *gophercloud.ServiceClient) string {
//	return c.ServiceURL(resourcePath)
//}

func getURL(c *gophercloud.ServiceClient, tenantID string) string {
	return c.ServiceURL(resourcePath, tenantID)
}

func getDetailURL(c *gophercloud.ServiceClient, tenantID string) string {
	return c.ServiceURL(resourcePath, tenantID, "detail")
}

func getLimitURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL( "limits")
}

func getDefaultURL(c *gophercloud.ServiceClient, tenantID string) string {
	return c.ServiceURL(resourcePath, tenantID, "defaults")
}

func updateURL(c *gophercloud.ServiceClient, tenantID string) string {
	return getURL(c, tenantID)
}

func deleteURL(c *gophercloud.ServiceClient, tenantID string) string {
	return getURL(c, tenantID)
}
