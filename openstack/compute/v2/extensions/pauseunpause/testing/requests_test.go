package testing

import (
	"testing"

	"github.com/Huawei/gophercloud/openstack/compute/v2/extensions/pauseunpause"
	th "github.com/Huawei/gophercloud/testhelper"
	"github.com/Huawei/gophercloud/testhelper/client"
)

const serverID = "{serverId}"

func TestPause(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	mockPauseServerResponse(t, serverID)

	err := pauseunpause.Pause(client.ServiceClient(), serverID).ExtractErr()
	th.AssertNoErr(t, err)
}

func TestUnpause(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	mockUnpauseServerResponse(t, serverID)

	err := pauseunpause.Unpause(client.ServiceClient(), serverID).ExtractErr()
	th.AssertNoErr(t, err)
}
