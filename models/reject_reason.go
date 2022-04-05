package models
import (
    "strings"
    "errors"
)
// Provides operations to call the reject method.
type RejectReason int

const (
    NONE_REJECTREASON RejectReason = iota
    BUSY_REJECTREASON
    FORBIDDEN_REJECTREASON
    UNKNOWNFUTUREVALUE_REJECTREASON
)

func (i RejectReason) String() string {
    return []string{"NONE", "BUSY", "FORBIDDEN", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseRejectReason(v string) (interface{}, error) {
    result := NONE_REJECTREASON
    switch strings.ToUpper(v) {
        case "NONE":
            result = NONE_REJECTREASON
        case "BUSY":
            result = BUSY_REJECTREASON
        case "FORBIDDEN":
            result = FORBIDDEN_REJECTREASON
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_REJECTREASON
        default:
            return 0, errors.New("Unknown RejectReason value: " + v)
    }
    return &result, nil
}
func SerializeRejectReason(values []RejectReason) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
