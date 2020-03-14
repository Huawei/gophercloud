package bandwidths

import (
	"github.com/Huawei/gophercloud"
	"github.com/Huawei/gophercloud/pagination"
)

func Get(client *gophercloud.ServiceClient, bandwidthId string) (r GetResult) {
	url := GetURL(client, bandwidthId)
	_, r.Err = client.Get(url, &r.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200},
	})
	return
}

type ListOpts struct {
	// Specifies the resource ID of pagination query. If the parameter
	// is left blank, only resources on the first page are queried.
	Marker string `q:"marker"`

	// Specifies the number of records returned on each page.
	Limit int `q:"limit"`

	// enterprise_project_id
	// You can use this field to filter the bandwidth under an enterprise project.
	EnterpriseProjectId string `q:"enterprise_project_id"`
}

type ListOptsBuilder interface {
	ToListQuery() (string, error)
}

func (opts ListOpts) ToListQuery() (string, error) {
	q, err := gophercloud.BuildQueryString(opts)
	return q.String(), err
}

func List(client *gophercloud.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	url := ListURL(client)
	if opts != nil {
		query, err := opts.ToListQuery()
		if err != nil {
			return pagination.Pager{Err: err}
		}
		url += query
	}

	return pagination.NewPager(client, url,
		func(r pagination.PageResult) pagination.Page {
			return BandWidthPage{pagination.LinkedPageBase{PageResult: r}}

		})
}

type UpdateOpts struct {
	// Specifies the bandwidth name. The value is a string of 1 to 64
	// characters that can contain letters, digits, underscores (_), and hyphens (-).
	Name string `json:"name,omitempty"`

	// Specifies the bandwidth size. The value ranges from 1 Mbit/s to
	// 300 Mbit/s.
	Size int `json:"size,omitempty"`
}

type UpdateOptsBuilder interface {
	ToBandwidthsUpdateMap() (map[string]interface{}, error)
}

func (opts UpdateOpts) ToBandwidthsUpdateMap() (map[string]interface{}, error) {
	b, err := gophercloud.BuildRequestBody(&opts, "bandwidth")
	if err != nil {
		return nil, err
	}
	return b, nil
}

func Update(client *gophercloud.ServiceClient, bandwidthId string, opts UpdateOptsBuilder) (r UpdateResult) {
	b, err := opts.ToBandwidthsUpdateMap()
	if err != nil {
		r.Err = err
		return
	}

	_, r.Err = client.Put(UpdateURL(client, bandwidthId), b, &r.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200},
	})
	return
}
