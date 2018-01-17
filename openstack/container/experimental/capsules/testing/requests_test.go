package testing

import (
	"testing"

	"github.com/gophercloud/gophercloud/openstack/container/experimental/capsules"
	th "github.com/gophercloud/gophercloud/testhelper"
	fakeclient "github.com/gophercloud/gophercloud/testhelper/client"
)

func TestGetCapsule(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	HandleCapsuleGetSuccessfully(t)

	actualCapsule, err := capsules.Get(fakeclient.ServiceClient(), "cc654059-1a77-47a3-bfcf-715bde5aad9e").Extract()

	th.AssertNoErr(t, err)

	uuid := "cc654059-1a77-47a3-bfcf-715bde5aad9e"
	status := "Running"
	id := 1
	userID := "d33b18c384574fd2a3299447aac285f0"
	projectID := "6b8ffef2a0ac42ee87887b9cc98bdf68"
	cpu := 1
	memory := "1024M"
	name := "test"

	created := "2018-01-12 09:37:25+00:00"
	links :=  []interface{}{
		map[string]interface{}{
			"href": "http://10.10.10.10/v1/capsules/cc654059-1a77-47a3-bfcf-715bde5aad9e",
			"rel":  "self",
		},
		map[string]interface{}{
			"href": "http://10.10.10.10/capsules/cc654059-1a77-47a3-bfcf-715bde5aad9e",
			"rel":  "bookmark",
		},
	}
	capsuleVersion := "beta"
	restartPolicy := map[string]string{
		"Name": "no",
		"MaximumRetryCount": "0",
	}
	metaLables := map[string]string{
		"web": "app",
	}
	containersUUIDs := []string{
		"1739e28a-d391-4fd9-93a5-3ba3f29a4c9b",
		"d1469e8d-bcbc-43fc-b163-8b9b6a740930",
	}

	expectedCapsule := capsules.Capsule{
		UUID: uuid,
		ID: id,
		UserID: userID,
		ProjectID: projectID,
		CPU: cpu,
		Status: status,
		Memory: memory,
		Name: name,
		Created: created,
		Links: links,
		CapsuleVersion: capsuleVersion,
		RestartPolicy: restartPolicy,
		MetaLables: metaLables,
		ContainersUUIDs: containersUUIDs,
	}

	th.AssertDeepEquals(t, &expectedCapsule, actualCapsule)
}
