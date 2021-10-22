package graph
import (
    "strings"
    "errors"
)
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
    switch strings.ToUpper(v) {
        case "SUCCESS":
            return SUCCESS_OPERATIONRESULT, nil
        case "FAILURE":
            return FAILURE_OPERATIONRESULT, nil
        case "TIMEOUT":
            return TIMEOUT_OPERATIONRESULT, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_OPERATIONRESULT, nil
    }
    return 0, errors.New("Unknown OperationResult value: " + v)
}
func SerializeOperationResult(values []OperationResult) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
