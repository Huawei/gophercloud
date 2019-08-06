package availabilityzones

import (
	"github.com/Huawei/gophercloud"
	"github.com/Huawei/gophercloud/pagination"
)

// List will return the existing availability zones.
func List(client *gophercloud.ServiceClient) pagination.Pager {
	return pagination.NewPager(client, listURL(client), func(r pagination.PageResult) pagination.Page {
		return AvailabilityZonePage{pagination.SinglePageBase(r)}
	})
}

// ListDetail will return the existing availability zones with detailed information.
func ListDetail(client *gophercloud.ServiceClient) pagination.Pager {
	return pagination.NewPager(client, listDetailURL(client), func(r pagination.PageResult) pagination.Page {
		return AvailabilityZonePage{pagination.SinglePageBase(r)}
	})
}
