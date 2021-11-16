package graph
import (
    "strings"
    "errors"
)
// 
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
    switch strings.ToUpper(v) {
        case "CONFIGURED":
            return CONFIGURED_CONNECTEDORGANIZATIONSTATE, nil
        case "PROPOSED":
            return PROPOSED_CONNECTEDORGANIZATIONSTATE, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_CONNECTEDORGANIZATIONSTATE, nil
    }
    return 0, errors.New("Unknown ConnectedOrganizationState value: " + v)
}
func SerializeConnectedOrganizationState(values []ConnectedOrganizationState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
