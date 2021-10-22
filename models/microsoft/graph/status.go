package graph
import (
    "strings"
    "errors"
)
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
    switch strings.ToUpper(v) {
        case "ACTIVE":
            return ACTIVE_STATUS, nil
        case "UPDATED":
            return UPDATED_STATUS, nil
        case "DELETED":
            return DELETED_STATUS, nil
        case "IGNORED":
            return IGNORED_STATUS, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_STATUS, nil
    }
    return 0, errors.New("Unknown Status value: " + v)
}
func SerializeStatus(values []Status) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
