package apiversions

import (
	"github.com/Huawei/gophercloud"
	"github.com/Huawei/gophercloud/pagination"
)

// List lists all the API versions available to end-users.
func List(c *gophercloud.ServiceClient) pagination.Pager {
	return pagination.NewPager(c, listURL(c), func(r pagination.PageResult) pagination.Page {
		return APIVersionPage{pagination.SinglePageBase(r)}
	})
}

// Get will get a specific API version, specified by major ID.
func Get(client *gophercloud.ServiceClient, v string) (r GetResult) {
	_, r.Err = client.Get(getURL(client, v), &r.Body, nil)
	return
}
