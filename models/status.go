package models
import (
    "strings"
    "errors"
)
// Provides operations to manage the drive singleton.
type Status int

const (
    ACTIVE_STATUS Status = iota
    UPDATED_STATUS
    DELETED_STATUS
    IGNORED_STATUS
    UNKNOWNFUTUREVALUE_STATUS
)

func (i Status) String() string {
    return []string{"ACTIVE", "UPDATED", "DELETED", "IGNORED", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseStatus(v string) (interface{}, error) {
    result := ACTIVE_STATUS
    switch strings.ToUpper(v) {
        case "ACTIVE":
            result = ACTIVE_STATUS
        case "UPDATED":
            result = UPDATED_STATUS
        case "DELETED":
            result = DELETED_STATUS
        case "IGNORED":
            result = IGNORED_STATUS
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_STATUS
        default:
            return 0, errors.New("Unknown Status value: " + v)
    }
    return &result, nil
}
func SerializeStatus(values []Status) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
