package graph
import (
    "strings"
    "errors"
)
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
    switch strings.ToUpper(v) {
        case "UNKNOWN":
            return UNKNOWN_VPPTOKENSTATE, nil
        case "VALID":
            return VALID_VPPTOKENSTATE, nil
        case "EXPIRED":
            return EXPIRED_VPPTOKENSTATE, nil
        case "INVALID":
            return INVALID_VPPTOKENSTATE, nil
        case "ASSIGNEDTOEXTERNALMDM":
            return ASSIGNEDTOEXTERNALMDM_VPPTOKENSTATE, nil
    }
    return 0, errors.New("Unknown VppTokenState value: " + v)
}
func SerializeVppTokenState(values []VppTokenState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
