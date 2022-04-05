package models
import (
    "strings"
    "errors"
)
// Provides operations to manage the auditLogRoot singleton.
type ConditionalAccessStatus int

const (
    SUCCESS_CONDITIONALACCESSSTATUS ConditionalAccessStatus = iota
    FAILURE_CONDITIONALACCESSSTATUS
    NOTAPPLIED_CONDITIONALACCESSSTATUS
    UNKNOWNFUTUREVALUE_CONDITIONALACCESSSTATUS
)

func (i ConditionalAccessStatus) String() string {
    return []string{"SUCCESS", "FAILURE", "NOTAPPLIED", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseConditionalAccessStatus(v string) (interface{}, error) {
    result := SUCCESS_CONDITIONALACCESSSTATUS
    switch strings.ToUpper(v) {
        case "SUCCESS":
            result = SUCCESS_CONDITIONALACCESSSTATUS
        case "FAILURE":
            result = FAILURE_CONDITIONALACCESSSTATUS
        case "NOTAPPLIED":
            result = NOTAPPLIED_CONDITIONALACCESSSTATUS
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_CONDITIONALACCESSSTATUS
        default:
            return 0, errors.New("Unknown ConditionalAccessStatus value: " + v)
    }
    return &result, nil
}
func SerializeConditionalAccessStatus(values []ConditionalAccessStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
