package experimental

import (
	"testing"

	"github.com/gophercloud/gophercloud/acceptance/clients"
	"github.com/gophercloud/gophercloud/openstack/container/experimental/capsules"
	th "github.com/gophercloud/gophercloud/testhelper"
)

func TestCapsuleGet(t *testing.T) {
	client, err := clients.NewContainerExperimentalClient()
	if err != nil {
		t.Fatalf("Unable to create an container experimental client: %v", err)
	}
	th.AssertNoErr(t, err)
	capsuleUUID := "b7628d1b-366f-48dd-8893-4d8529544ab6"
	if capsuleUUID == "" {
		t.Fatalf("In order to retrieve a capsule, the CapsuleUUID must be set")
	}
	capsule, err := capsules.Get(client, capsuleUUID).Extract()
	// Get a capsule

	th.AssertNoErr(t, err)
	th.AssertEquals(t, capsule.Status, "Running")
	th.AssertEquals(t, capsule.Name, "template")
	th.AssertEquals(t, capsule.CPU, 2)
}
