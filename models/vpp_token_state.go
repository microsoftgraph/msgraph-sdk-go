package models
import (
    "strings"
    "errors"
)
// Provides operations to manage the deviceAppManagement singleton.
type VppTokenState int

const (
    UNKNOWN_VPPTOKENSTATE VppTokenState = iota
    VALID_VPPTOKENSTATE
    EXPIRED_VPPTOKENSTATE
    INVALID_VPPTOKENSTATE
    ASSIGNEDTOEXTERNALMDM_VPPTOKENSTATE
)

func (i VppTokenState) String() string {
    return []string{"UNKNOWN", "VALID", "EXPIRED", "INVALID", "ASSIGNEDTOEXTERNALMDM"}[i]
}
func ParseVppTokenState(v string) (interface{}, error) {
    result := UNKNOWN_VPPTOKENSTATE
    switch strings.ToUpper(v) {
        case "UNKNOWN":
            result = UNKNOWN_VPPTOKENSTATE
        case "VALID":
            result = VALID_VPPTOKENSTATE
        case "EXPIRED":
            result = EXPIRED_VPPTOKENSTATE
        case "INVALID":
            result = INVALID_VPPTOKENSTATE
        case "ASSIGNEDTOEXTERNALMDM":
            result = ASSIGNEDTOEXTERNALMDM_VPPTOKENSTATE
        default:
            return 0, errors.New("Unknown VppTokenState value: " + v)
    }
    return &result, nil
}
func SerializeVppTokenState(values []VppTokenState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
