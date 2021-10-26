package graph
import (
    "strings"
    "errors"
)
// 
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
    switch strings.ToUpper(v) {
        case "NONE":
            return NONE_REJECTREASON, nil
        case "BUSY":
            return BUSY_REJECTREASON, nil
        case "FORBIDDEN":
            return FORBIDDEN_REJECTREASON, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_REJECTREASON, nil
    }
    return 0, errors.New("Unknown RejectReason value: " + v)
}
func SerializeRejectReason(values []RejectReason) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
