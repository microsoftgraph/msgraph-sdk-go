package models
import (
    "strings"
    "errors"
)
// Provides operations to manage the security singleton.
type LogonType int

const (
    UNKNOWN_LOGONTYPE LogonType = iota
    INTERACTIVE_LOGONTYPE
    REMOTEINTERACTIVE_LOGONTYPE
    NETWORK_LOGONTYPE
    BATCH_LOGONTYPE
    SERVICE_LOGONTYPE
    UNKNOWNFUTUREVALUE_LOGONTYPE
)

func (i LogonType) String() string {
    return []string{"UNKNOWN", "INTERACTIVE", "REMOTEINTERACTIVE", "NETWORK", "BATCH", "SERVICE", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseLogonType(v string) (interface{}, error) {
    result := UNKNOWN_LOGONTYPE
    switch strings.ToUpper(v) {
        case "UNKNOWN":
            result = UNKNOWN_LOGONTYPE
        case "INTERACTIVE":
            result = INTERACTIVE_LOGONTYPE
        case "REMOTEINTERACTIVE":
            result = REMOTEINTERACTIVE_LOGONTYPE
        case "NETWORK":
            result = NETWORK_LOGONTYPE
        case "BATCH":
            result = BATCH_LOGONTYPE
        case "SERVICE":
            result = SERVICE_LOGONTYPE
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_LOGONTYPE
        default:
            return 0, errors.New("Unknown LogonType value: " + v)
    }
    return &result, nil
}
func SerializeLogonType(values []LogonType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
