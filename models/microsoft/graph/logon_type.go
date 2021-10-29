package graph
import (
    "strings"
    "errors"
)
// 
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
    switch strings.ToUpper(v) {
        case "UNKNOWN":
            return UNKNOWN_LOGONTYPE, nil
        case "INTERACTIVE":
            return INTERACTIVE_LOGONTYPE, nil
        case "REMOTEINTERACTIVE":
            return REMOTEINTERACTIVE_LOGONTYPE, nil
        case "NETWORK":
            return NETWORK_LOGONTYPE, nil
        case "BATCH":
            return BATCH_LOGONTYPE, nil
        case "SERVICE":
            return SERVICE_LOGONTYPE, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_LOGONTYPE, nil
    }
    return 0, errors.New("Unknown LogonType value: " + v)
}
func SerializeLogonType(values []LogonType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
