package graph
import (
    "strings"
    "errors"
)
// 
type ConnectionStatus int

const (
    UNKNOWN_CONNECTIONSTATUS ConnectionStatus = iota
    ATTEMPTED_CONNECTIONSTATUS
    SUCCEEDED_CONNECTIONSTATUS
    BLOCKED_CONNECTIONSTATUS
    FAILED_CONNECTIONSTATUS
    UNKNOWNFUTUREVALUE_CONNECTIONSTATUS
)

func (i ConnectionStatus) String() string {
    return []string{"UNKNOWN", "ATTEMPTED", "SUCCEEDED", "BLOCKED", "FAILED", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseConnectionStatus(v string) (interface{}, error) {
    result := UNKNOWN_CONNECTIONSTATUS
    switch strings.ToUpper(v) {
        case "UNKNOWN":
            result = UNKNOWN_CONNECTIONSTATUS
        case "ATTEMPTED":
            result = ATTEMPTED_CONNECTIONSTATUS
        case "SUCCEEDED":
            result = SUCCEEDED_CONNECTIONSTATUS
        case "BLOCKED":
            result = BLOCKED_CONNECTIONSTATUS
        case "FAILED":
            result = FAILED_CONNECTIONSTATUS
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_CONNECTIONSTATUS
        default:
            return 0, errors.New("Unknown ConnectionStatus value: " + v)
    }
    return &result, nil
}
func SerializeConnectionStatus(values []ConnectionStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
