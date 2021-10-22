package graph
import (
    "strings"
    "errors"
)
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
    switch strings.ToUpper(v) {
        case "UNKNOWN":
            return UNKNOWN_CONNECTIONSTATUS, nil
        case "ATTEMPTED":
            return ATTEMPTED_CONNECTIONSTATUS, nil
        case "SUCCEEDED":
            return SUCCEEDED_CONNECTIONSTATUS, nil
        case "BLOCKED":
            return BLOCKED_CONNECTIONSTATUS, nil
        case "FAILED":
            return FAILED_CONNECTIONSTATUS, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_CONNECTIONSTATUS, nil
    }
    return 0, errors.New("Unknown ConnectionStatus value: " + v)
}
func SerializeConnectionStatus(values []ConnectionStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
