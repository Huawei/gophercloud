package metrics

import "github.com/Huawei/gophercloud"

func getMetricsURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("metrics")
}
