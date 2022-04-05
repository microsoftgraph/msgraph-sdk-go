package models
import (
    "strings"
    "errors"
)
// Provides operations to manage the cloudCommunications singleton.
type CallDirection int

const (
    INCOMING_CALLDIRECTION CallDirection = iota
    OUTGOING_CALLDIRECTION
)

func (i CallDirection) String() string {
    return []string{"INCOMING", "OUTGOING"}[i]
}
func ParseCallDirection(v string) (interface{}, error) {
    result := INCOMING_CALLDIRECTION
    switch strings.ToUpper(v) {
        case "INCOMING":
            result = INCOMING_CALLDIRECTION
        case "OUTGOING":
            result = OUTGOING_CALLDIRECTION
        default:
            return 0, errors.New("Unknown CallDirection value: " + v)
    }
    return &result, nil
}
func SerializeCallDirection(values []CallDirection) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
