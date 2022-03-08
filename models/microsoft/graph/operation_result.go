package graph
import (
    "strings"
    "errors"
)
// Provides operations to manage the auditLogRoot singleton.
type OperationResult int

const (
    SUCCESS_OPERATIONRESULT OperationResult = iota
    FAILURE_OPERATIONRESULT
    TIMEOUT_OPERATIONRESULT
    UNKNOWNFUTUREVALUE_OPERATIONRESULT
)

func (i OperationResult) String() string {
    return []string{"SUCCESS", "FAILURE", "TIMEOUT", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseOperationResult(v string) (interface{}, error) {
    result := SUCCESS_OPERATIONRESULT
    switch strings.ToUpper(v) {
        case "SUCCESS":
            result = SUCCESS_OPERATIONRESULT
        case "FAILURE":
            result = FAILURE_OPERATIONRESULT
        case "TIMEOUT":
            result = TIMEOUT_OPERATIONRESULT
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_OPERATIONRESULT
        default:
            return 0, errors.New("Unknown OperationResult value: " + v)
    }
    return &result, nil
}
func SerializeOperationResult(values []OperationResult) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
