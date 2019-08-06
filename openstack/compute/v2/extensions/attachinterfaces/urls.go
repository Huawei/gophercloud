package attachinterfaces

import "github.com/Huawei/gophercloud"

func listInterfaceURL(client *gophercloud.ServiceClient, serverID string) string {
	return client.ServiceURL("servers", serverID, "os-interface")
}

func getInterfaceURL(client *gophercloud.ServiceClient, serverID, portID string) string {
	return client.ServiceURL("servers", serverID, "os-interface", portID)
}

func createInterfaceURL(client *gophercloud.ServiceClient, serverID string) string {
	return client.ServiceURL("servers", serverID, "os-interface")
}
func deleteInterfaceURL(client *gophercloud.ServiceClient, serverID, portID string) string {
	return client.ServiceURL("servers", serverID, "os-interface", portID)
}
