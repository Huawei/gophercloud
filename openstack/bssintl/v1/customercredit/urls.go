package customercredit

import "github.com/Huawei/gophercloud"

//GET /v1.0/{partner_id}/partner/account-mgr/credit
func getQueryCreditURL(client *gophercloud.ServiceClient, domainId string) string {
	return client.ServiceURL(domainId, "partner/account-mgr/credit")
}

//POST /v1.0/{partner_id}/partner/account-mgr/credit
func getSetCreditURL(client *gophercloud.ServiceClient, domainId string) string {
	return client.ServiceURL(domainId, "partner/account-mgr/credit")
}
