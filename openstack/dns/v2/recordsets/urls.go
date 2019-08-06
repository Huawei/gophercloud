package recordsets

import "github.com/Huawei/gophercloud"

func baseURL(c *gophercloud.ServiceClient, zoneID string) string {
	return c.ServiceURL("zones", zoneID, "recordsets")
}

func rrsetURL(c *gophercloud.ServiceClient, zoneID string, rrsetID string) string {
	return c.ServiceURL("zones", zoneID, "recordsets", rrsetID)
}

func listURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("recordsets")
}