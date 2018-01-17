package capsules

import (
	"github.com/gophercloud/gophercloud"
)

type commonResult struct {
	gophercloud.Result
}

// Extract is a function that accepts a result and extracts a capsule resource.
func (r commonResult) Extract() (*Capsule, error) {
	var s *Capsule
	err := r.ExtractInto(&s)
	return s, err
}

// GetResult represents the result of a get operation.
type GetResult struct {
	commonResult
}

// Represents a Container Orchestration Engine Bay, i.e. a cluster
type Capsule struct {
	// UUID for the capsule
	UUID string `json:"uuid"`

	// ID for the capsule
	ID int `json:"id"`

	UserID string `json:"user_id"`

	ProjectID string `json:"project_id"`

	// cpu for the capsule
	CPU int `json:"cpu"`

	// cpu for the capsule
	Memory string `json:"memory"`

	Name string `json:"meta_name"`

	// Indicates whether capsule is currently operational. Possible values include:
	// Running,
	Status string `json:"status"`

	// The created time of the capsule.
	Created string `json:"created_at"`

	// Links includes HTTP references to the itself, useful for passing along to
	// other APIs that might want a server reference.
	Links []interface{} `json:"links"`

	// The capsule version
	CapsuleVersion string `json:"capsule_version"`

	RestartPolicy map[string]string `json:"restart_policy"`

	MetaLables map[string]string `json:"meta_labels"`

	ContainersUUIDs []string `json:"containers_uuids"`
}
