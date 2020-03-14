package customermanagement

import "github.com/Huawei/gophercloud"


// POST /v1.0/{partner_id}/partner/customer-mgr/check-user
func getCheckCustomerRegisterInfoURL(client *gophercloud.ServiceClient, domainId string) string {
	return client.ServiceURL(domainId, "partner/customer-mgr/check-user")
}

// POST /v1.0/{partner_id}/partner/customer-mgr/customer
func getCreateCustomerURL(client *gophercloud.ServiceClient, domainId string) string {
	return client.ServiceURL(domainId, "partner/customer-mgr/customer")
}

// POST /v1.0/{partner_id}/partner/customer-mgr/customer
func getQueryCustomerURL(client *gophercloud.ServiceClient, domainId string) string {
	return client.ServiceURL(domainId, "partner/customer-mgr/query")
}

// POST /v1.0/{partner_id}/partner/customer-mgr/frozens
func getFrozenCustomerURL(client *gophercloud.ServiceClient, domainId string) string {
	return client.ServiceURL(domainId, "partner/customer-mgr/frozens")
}

// POST /v1.0/{partner_id}/partner/customer-mgr/unfrozens
func getUnFrozenCustomerURL(client *gophercloud.ServiceClient, domainId string) string {
	return client.ServiceURL(domainId, "partner/customer-mgr/unfrozens")
}