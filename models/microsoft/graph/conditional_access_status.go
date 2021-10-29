package graph
import (
    "strings"
    "errors"
)
// 
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
    switch strings.ToUpper(v) {
        case "SUCCESS":
            return SUCCESS_CONDITIONALACCESSSTATUS, nil
        case "FAILURE":
            return FAILURE_CONDITIONALACCESSSTATUS, nil
        case "NOTAPPLIED":
            return NOTAPPLIED_CONDITIONALACCESSSTATUS, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_CONDITIONALACCESSSTATUS, nil
    }
    return 0, errors.New("Unknown ConditionalAccessStatus value: " + v)
}
func SerializeConditionalAccessStatus(values []ConditionalAccessStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
