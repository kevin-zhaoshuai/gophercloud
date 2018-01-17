package testing

import (
	"fmt"
	"net/http"
	"testing"

	th "github.com/gophercloud/gophercloud/testhelper"
	fakeclient "github.com/gophercloud/gophercloud/testhelper/client"
)

type imageEntry struct {
	ID   string
	JSON string
}

// HandleImageGetSuccessfully test setup
func HandleCapsuleGetSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/capsules/cc654059-1a77-47a3-bfcf-715bde5aad9e", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fakeclient.TokenID)

		w.WriteHeader(http.StatusOK)
		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, `{
			"uuid": "cc654059-1a77-47a3-bfcf-715bde5aad9e",
			"status": "Running",
			"id": 1,
			"user_id": "d33b18c384574fd2a3299447aac285f0",
			"project_id": "6b8ffef2a0ac42ee87887b9cc98bdf68",
			"cpu": 1,
			"memory": "1024M",
			"meta_name": "test",
			"meta_labels": {"web": "app"},
			"created_at": "2018-01-12 09:37:25+00:00",
			"links": [
				{
				  "href": "http://10.10.10.10/v1/capsules/cc654059-1a77-47a3-bfcf-715bde5aad9e",
				  "rel": "self"
				},
				{
				  "href": "http://10.10.10.10/capsules/cc654059-1a77-47a3-bfcf-715bde5aad9e",
				  "rel": "bookmark"
				}
			],
			"capsule_version": "beta",
			"restart_policy":  {"MaximumRetryCount": "0", "Name": "no"},
			"containers_uuids": ["1739e28a-d391-4fd9-93a5-3ba3f29a4c9b", "d1469e8d-bcbc-43fc-b163-8b9b6a740930"]
		}`)
	})
}
