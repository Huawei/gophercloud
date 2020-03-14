// +build acceptance blockstorage

package extensions

import (
	"testing"

	"github.com/Huawei/gophercloud/acceptance/clients"
	"github.com/Huawei/gophercloud/acceptance/tools"
	"github.com/Huawei/gophercloud/openstack/blockstorage/extensions/services"
)

func TestServicesList(t *testing.T) {
	blockClient, err := clients.NewBlockStorageV3Client()
	if err != nil {
		t.Fatalf("Unable to create a blockstorage client: %v", err)
	}

	allPages, err := services.List(blockClient, services.ListOpts{}).AllPages()
	if err != nil {
		t.Fatalf("Unable to list services: %v", err)
	}

	allServices, err := services.ExtractServices(allPages)
	if err != nil {
		t.Fatalf("Unable to extract services")
	}

	for _, service := range allServices {
		tools.PrintResource(t, service)
	}
}
