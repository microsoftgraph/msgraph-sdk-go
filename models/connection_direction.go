package models
import (
    "strings"
    "errors"
)
// Provides operations to manage the security singleton.
type ConnectionDirection int

const (
    UNKNOWN_CONNECTIONDIRECTION ConnectionDirection = iota
    INBOUND_CONNECTIONDIRECTION
    OUTBOUND_CONNECTIONDIRECTION
    UNKNOWNFUTUREVALUE_CONNECTIONDIRECTION
)

func (i ConnectionDirection) String() string {
    return []string{"UNKNOWN", "INBOUND", "OUTBOUND", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseConnectionDirection(v string) (interface{}, error) {
    result := UNKNOWN_CONNECTIONDIRECTION
    switch strings.ToUpper(v) {
        case "UNKNOWN":
            result = UNKNOWN_CONNECTIONDIRECTION
        case "INBOUND":
            result = INBOUND_CONNECTIONDIRECTION
        case "OUTBOUND":
            result = OUTBOUND_CONNECTIONDIRECTION
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_CONNECTIONDIRECTION
        default:
            return 0, errors.New("Unknown ConnectionDirection value: " + v)
    }
    return &result, nil
}
func SerializeConnectionDirection(values []ConnectionDirection) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
