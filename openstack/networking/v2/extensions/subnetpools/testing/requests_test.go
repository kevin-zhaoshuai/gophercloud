package testing

import (
	"fmt"
	"net/http"
	"testing"

	fake "github.com/gophercloud/gophercloud/openstack/networking/v2/common"
	"github.com/gophercloud/gophercloud/openstack/networking/v2/extensions/subnetpools"
	"github.com/gophercloud/gophercloud/pagination"
	th "github.com/gophercloud/gophercloud/testhelper"
)

func TestList(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/v2.0/subnetpools", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, SubnetPoolsListResult)
	})

	count := 0

	subnetpools.List(fake.ServiceClient(), subnetpools.ListOpts{}).EachPage(func(page pagination.Page) (bool, error) {
		count++
		actual, err := subnetpools.ExtractSubnetPools(page)
		if err != nil {
			t.Errorf("Failed to extract subnetpools: %v", err)
			return false, nil
		}

		expected := []subnetpools.SubnetPool{
			SubnetPool1,
			SubnetPool2,
			SubnetPool3,
		}

		th.CheckDeepEquals(t, expected, actual)

		return true, nil
	})

	if count != 1 {
		t.Errorf("Expected 1 page, got %d", count)
	}
}

func TestGet(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/v2.0/subnetpools/0a738452-8057-4ad3-89c2-92f6a74afa76", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, SubnetPoolGetResult)
	})

	s, err := subnetpools.Get(fake.ServiceClient(), "0a738452-8057-4ad3-89c2-92f6a74afa76").Extract()
	th.AssertNoErr(t, err)

	th.AssertEquals(t, s.ID, "0a738452-8057-4ad3-89c2-92f6a74afa76")
	th.AssertEquals(t, s.Name, "my-ipv6-pool")
	th.AssertEquals(t, s.DefaultQuota, 2)
	th.AssertEquals(t, s.TenantID, "1e2b9857295a4a3e841809ef492812c5")
	th.AssertEquals(t, s.ProjectID, "1e2b9857295a4a3e841809ef492812c5")
	th.AssertEquals(t, s.CreatedAt, "2018-01-01T00:00:01")
	th.AssertEquals(t, s.UpdatedAt, "2018-01-01T00:10:10")
	th.AssertDeepEquals(t, s.Prefixes, []string{
		"2001:db8::a3/64",
	})
	th.AssertEquals(t, s.DefaultPrefixLen, 64)
	th.AssertEquals(t, s.MinPrefixLen, 64)
	th.AssertEquals(t, s.MaxPrefixLen, 128)
	th.AssertEquals(t, s.AddressScopeID, "")
	th.AssertEquals(t, s.IPversion, 6)
	th.AssertEquals(t, s.Shared, false)
	th.AssertEquals(t, s.Description, "ipv6 prefixes")
	th.AssertEquals(t, s.IsDefault, true)
	th.AssertEquals(t, s.RevisionNumber, 2)
}
func TestCreate(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/v2.0/subnetpools", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		th.TestHeader(t, r, "Content-Type", "application/json")
		th.TestHeader(t, r, "Accept", "application/json")
		th.TestJSONRequest(t, r, SubnetPoolCreateRequest)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)

		fmt.Fprintf(w, SubnetPoolCreateResult)
	})

	opts := subnetpools.CreateOpts{
		Name: "my_ipv4_pool",
		Prefixes: []string{
			"10.10.0.0/16",
			"10.11.11.0/24",
		},
		MinPrefixLen:   25,
		MaxPrefixLen:   30,
		AddressScopeID: "3d4e2e2a-552b-42ad-a16d-820bbf3edaf3",
		Description:    "ipv4 prefixes",
	}
	s, err := subnetpools.Create(fake.ServiceClient(), opts).Extract()
	th.AssertNoErr(t, err)

	th.AssertEquals(t, s.Name, "my_ipv4_pool")
	th.AssertDeepEquals(t, s.Prefixes, []string{
		"10.10.0.0/16",
		"10.11.11.0/24",
	})
	th.AssertEquals(t, s.MinPrefixLen, 25)
	th.AssertEquals(t, s.MaxPrefixLen, 30)
	th.AssertEquals(t, s.AddressScopeID, "3d4e2e2a-552b-42ad-a16d-820bbf3edaf3")
	th.AssertEquals(t, s.Description, "ipv4 prefixes")
}