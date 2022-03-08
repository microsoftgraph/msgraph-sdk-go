package graph
import (
    "strings"
    "errors"
)
// Provides operations to manage the identityGovernance singleton.
type ConnectedOrganizationState int

const (
    CONFIGURED_CONNECTEDORGANIZATIONSTATE ConnectedOrganizationState = iota
    PROPOSED_CONNECTEDORGANIZATIONSTATE
    UNKNOWNFUTUREVALUE_CONNECTEDORGANIZATIONSTATE
)

func (i ConnectedOrganizationState) String() string {
    return []string{"CONFIGURED", "PROPOSED", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseConnectedOrganizationState(v string) (interface{}, error) {
    result := CONFIGURED_CONNECTEDORGANIZATIONSTATE
    switch strings.ToUpper(v) {
        case "CONFIGURED":
            result = CONFIGURED_CONNECTEDORGANIZATIONSTATE
        case "PROPOSED":
            result = PROPOSED_CONNECTEDORGANIZATIONSTATE
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_CONNECTEDORGANIZATIONSTATE
        default:
            return 0, errors.New("Unknown ConnectedOrganizationState value: " + v)
    }
    return &result, nil
}
func SerializeConnectedOrganizationState(values []ConnectedOrganizationState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
