// Package status represents the product status type in the system.
package status

import "fmt"

// The set of roles that can be used.
var (
	Admin = newStatus("PUBLISH")
	User  = newStatus("DRAFT")
)

// =============================================================================

// Set of known status.
var statuses = make(map[string]Status)

// Status represents a product status in the system.
type Status struct {
	value string
}

func newStatus(status string) Status {
	r := Status{status}
	statuses[status] = r
	return r
}

// String returns the name of the status.
func (r Status) String() string {
	return r.value
}

// Equal provides support for the go-cmp package and testing.
func (r Status) Equal(r2 Status) bool {
	return r.value == r2.value
}

// MarshalText provides support for logging and any marshal needs.
func (r Status) MarshalText() ([]byte, error) {
	return []byte(r.value), nil
}

// =============================================================================

// Parse parses the string value and returns a status if one exists.
func Parse(value string) (Status, error) {
	status, exists := statuses[value]
	if !exists {
		return Status{}, fmt.Errorf("invalid role %q", value)
	}

	return status, nil
}

// MustParse parses the string value and returns a status if one exists. If
// an error occurs the function panics.
func MustParse(value string) Status {
	status, err := Parse(value)
	if err != nil {
		panic(err)
	}

	return status
}

// ParseToString takes a collection of user status(es) and converts them to
// a slice of string.
func ParseToString(status []Status) []string {
	statuses := make([]string, len(status))
	for i, role := range status {
		statuses[i] = role.String()
	}

	return statuses
}

// ParseMany takes a collection of strings and converts them to a slice
// of status(es).
func ParseMany(statuses []string) ([]Status, error) {
	prdStatuses := make([]Status, len(statuses))
	for i, status := range statuses {
		role, err := Parse(status)
		if err != nil {
			return nil, err
		}
		prdStatuses[i] = role
	}

	return prdStatuses, nil
}
