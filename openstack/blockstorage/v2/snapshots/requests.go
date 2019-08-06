package snapshots

import (
	"fmt"

	"github.com/Huawei/gophercloud"
	"github.com/Huawei/gophercloud/pagination"
)

// CreateOptsBuilder allows extensions to add additional parameters to the
// Create request.
type CreateOptsBuilder interface {
	ToSnapshotCreateMap() (map[string]interface{}, error)
}

// CreateOpts contains options for creating a Snapshot. This object is passed to
// the snapshots.Create function. For more information about these parameters,
// see the Snapshot object.
type CreateOpts struct {
	//Create a snapshot source EVS UUID.
	VolumeID string `json:"volume_id" required:"true"`
	//Forced to create a snapshot, the default is false.
	Force bool `json:"force,omitempty"`
	//Snapshot name
	Name string `json:"name,omitempty"`
	//Snapshot description
	Description string `json:"description,omitempty"`
	//Metadata information of the cloud disk snapshot.
	Metadata map[string]string `json:"metadata,omitempty"`
}

// ToSnapshotCreateMap assembles a request body based on the contents of a
// CreateOpts.
func (opts CreateOpts) ToSnapshotCreateMap() (map[string]interface{}, error) {
	return gophercloud.BuildRequestBody(opts, "snapshot")
}

// Create will create a new Snapshot based on the values in CreateOpts. To
// extract the Snapshot object from the response, call the Extract method on the
// CreateResult.
func Create(client *gophercloud.ServiceClient, opts CreateOptsBuilder) (r CreateResult) {
	b, err := opts.ToSnapshotCreateMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Post(createURL(client), b, &r.Body, &gophercloud.RequestOpts{
		OkCodes: []int{202},
	})
	return
}

// Delete will delete the existing Snapshot with the provided ID.
func Delete(client *gophercloud.ServiceClient, id string) (r DeleteResult) {
	_, r.Err = client.Delete(deleteURL(client, id), nil)
	return
}

// Get retrieves the Snapshot with the provided ID. To extract the Snapshot
// object from the response, call the Extract method on the GetResult.
func Get(client *gophercloud.ServiceClient, id string) (r GetResult) {
	_, r.Err = client.Get(getURL(client, id), &r.Body, nil)
	return
}

// ListOptsBuilder allows extensions to add additional parameters to the List
// request.
type ListOptsBuilder interface {
	ToSnapshotListQuery() (string, error)
}

// ListOpts hold options for listing Snapshots. It is passed to the
// snapshots.List function.
type ListOpts struct {
	// AllTenants will retrieve snapshots of all tenants/projects.
	//AllTenants bool `q:"all_tenants"`

	// Name will filter by the specified snapshot name.
	Name string `q:"name"`

	// Status will filter by the specified status.
	Status string `q:"status"`

	// TenantID will filter by a specific tenant/project ID.
	// Setting AllTenants is required to use this.
	//TenantID string `q:"project_id"`

	// VolumeID will filter by a specified volume ID.
	VolumeID string `q:"volume_id"`

	//Used when paginating snapshots, used in conjunction with limit.
	Offset int `q:"offset"`

	//Returns the number of results limit, an integer greater than 0. The default is 1000.
	Limit int `q:"limit"`

	AvailabilityZone string `q:"availability_zone"`

	// SnapshotID will filter by a specified volume ID.
	//SnapshotID string `q:"volume_id"`
}

// ToSnapshotListQuery formats a ListOpts into a query string.
func (opts ListOpts) ToSnapshotListQuery() (string, error) {
	q, err := gophercloud.BuildQueryString(opts)
	return q.String(), err
}

// List returns Snapshots optionally limited by the conditions provided in
// ListOpts.
func List(client *gophercloud.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	url := listURL(client)
	if opts != nil {
		query, err := opts.ToSnapshotListQuery()
		if err != nil {
			return pagination.Pager{Err: err}
		}
		url += query
	}
	return pagination.NewPager(client, url, func(r pagination.PageResult) pagination.Page {
		return SnapshotPage{pagination.SinglePageBase(r)}
	})
}

// UpdateMetadataOptsBuilder allows extensions to add additional parameters to
// the Update request.
type UpdateMetadataOptsBuilder interface {
	ToSnapshotUpdateMetadataMap() (map[string]interface{}, error)
}

// UpdateMetadataOpts contain options for updating an existing Snapshot. This
// object is passed to the snapshots.Update function. For more information
// about the parameters, see the Snapshot object.
type UpdateMetadataOpts struct {
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

// ToSnapshotUpdateMetadataMap assembles a request body based on the contents of
// an UpdateMetadataOpts.
func (opts UpdateMetadataOpts) ToSnapshotUpdateMetadataMap() (map[string]interface{}, error) {
	return gophercloud.BuildRequestBody(opts, "")
}

// UpdateMetadata will update the Snapshot with provided information. To
// extract the updated Snapshot from the response, call the ExtractMetadata
// method on the UpdateMetadataResult.
func UpdateMetadata(client *gophercloud.ServiceClient, id string, opts UpdateMetadataOptsBuilder) (r UpdateMetadataResult) {
	b, err := opts.ToSnapshotUpdateMetadataMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Put(metadataURL(client, id), b, &r.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200},
	})
	return
}

func Detail(client *gophercloud.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	url := detailURL(client)
	if opts != nil {
		query, err := opts.ToSnapshotListQuery()
		if err != nil {
			return pagination.Pager{Err: err}
		}
		url += query
	}

	return pagination.NewPager(client, url, func(r pagination.PageResult) pagination.Page {
		return SnapshotPage{pagination.SinglePageBase(r)}
	})
}

type UpdateOptsBuilder interface {
	ToSnapshotUpdateMap() (map[string]interface{}, error)
}

type UpdateOpts struct {
	Name               string `json:"name,omitempty"`
	Description        string `json:"description,omitempty"`
	DisplayName        string `json:"display_name,omitempty"`
	DisplayDescription string `json:"display_description,omitempty"`
}

func (opts UpdateOpts) ToSnapshotUpdateMap() (map[string]interface{}, error) {
	return gophercloud.BuildRequestBody(opts, "snapshot")
}

func Update(client *gophercloud.ServiceClient, id string, opts UpdateOptsBuilder) (r UpdateResult) {
	b, err := opts.ToSnapshotUpdateMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Put(updateURL(client, id), b, &r.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200},
	})
	return
}

// MetadataOptsBuilder allows extensions to add additional parameters to
// the meatadata requests.
type MetadataOptsBuilder interface {
	ToSnapshotMetadataMap() (map[string]interface{}, error)
}

// MetadataOpts contain options for creating or updating an existing Voulme. This
// object is passed to the volumes create and update function. For more information
// about the parameters, see the Snapshot object.
type MetadataOpts struct {
	Metadata map[string]string `json:"metadata,omitempty"`
}

// ToSnapshotMetadataMap assembles a request body based on the contents of
// an MetadataOpts.
func (opts MetadataOpts) ToSnapshotMetadataMap() (map[string]interface{}, error) {
	return gophercloud.BuildRequestBody(opts, "")
}

// CreateMetadata create metadata for Snapshot.
func CreateMetadata(client *gophercloud.ServiceClient, id string, opts MetadataOptsBuilder) (r MetadataResult) {
	b, err := opts.ToSnapshotMetadataMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Post(metadataURL(client, id), b, &r.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200},
	})
	return
}

// GetMetadata returns exist metadata of Snapshot.
func GetMetadata(client *gophercloud.ServiceClient, id string) (r MetadataResult) {
	_, r.Err = client.Get(metadataURL(client, id), &r.Body, nil)
	return
}

type MetadataKeyOptsBuilder interface {
	ToSnapshotMetadataKeyMap() (map[string]interface{}, error)
}

type MetadataKeyOpts struct {
	Metadata map[string]string `json:"meta,omitempty"`
}

func (opts MetadataKeyOpts) ToSnapshotMetadataKeyMap() (map[string]interface{}, error) {
	return gophercloud.BuildRequestBody(opts, "")
}

// GetMetadataKey return specific key value in metadata.
func GetMetadataKey(client *gophercloud.ServiceClient, id, key string) (r MetadataKeyResult) {
	_, r.Err = client.Get(metadataKeyURL(client, id, key), &r.Body, nil)
	return
}

// UpdateMetadataKey update sepcific key to the given map key value.
func UpdateMetadataKey(client *gophercloud.ServiceClient, id, key string, opts MetadataKeyOptsBuilder) (r MetadataKeyResult) {
	b, err := opts.ToSnapshotMetadataKeyMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Put(metadataKeyURL(client, id, key), b, &r.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200},
	})
	return
}

// DeleteMetadataKey delete specific key in metadata
func DeleteMetadataKey(client *gophercloud.ServiceClient, id, key string) (r DeleteMetadataKeyResult) {
	_, r.Err = client.Delete(metadataKeyURL(client, id, key), &gophercloud.RequestOpts{OkCodes: []int{200}}, )
	return
}

// IDFromName is a convienience function that returns a snapshot's ID given its name.
func IDFromName(client *gophercloud.ServiceClient, name string) (string, error) {
	count := 0
	id := ""
	pages, err := List(client, nil).AllPages()
	if err != nil {
		return "", err
	}

	all, err := ExtractSnapshots(pages)
	if err != nil {
		return "", err
	}

	for _, s := range all.Snapshots {
		if s.Name == name {
			count++
			id = s.ID
		}
	}

	switch count {
	case 0:
		//return "", gophercloud.ErrResourceNotFound{Name: name, ResourceType: "snapshot"}

		message := fmt.Sprintf(gophercloud.CE_ResourceNotFoundMessage, "snapshot", name)
		err := gophercloud.NewSystemCommonError(gophercloud.CE_ResourceNotFoundCode, message)
		return "", err
	case 1:
		return id, nil
	default:
		//return "", gophercloud.ErrMultipleResourcesFound{Name: name, Count: count, ResourceType: "snapshot"}

		message := fmt.Sprintf(gophercloud.CE_MultipleResourcesFoundMessage, count, "snapshot", name)
		err := gophercloud.NewSystemCommonError(gophercloud.CE_MultipleResourcesFoundCode, message)
		return "", err
	}
}
