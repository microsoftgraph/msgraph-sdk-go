package graph
import (
    "strings"
    "errors"
)
// 
type PrintOperationProcessingState int

const (
    NOTSTARTED_PRINTOPERATIONPROCESSINGSTATE PrintOperationProcessingState = iota
    RUNNING_PRINTOPERATIONPROCESSINGSTATE
    SUCCEEDED_PRINTOPERATIONPROCESSINGSTATE
    FAILED_PRINTOPERATIONPROCESSINGSTATE
    UNKNOWNFUTUREVALUE_PRINTOPERATIONPROCESSINGSTATE
)

func (i PrintOperationProcessingState) String() string {
    return []string{"NOTSTARTED", "RUNNING", "SUCCEEDED", "FAILED", "UNKNOWNFUTUREVALUE"}[i]
}
func ParsePrintOperationProcessingState(v string) (interface{}, error) {
    result := NOTSTARTED_PRINTOPERATIONPROCESSINGSTATE
    switch strings.ToUpper(v) {
        case "NOTSTARTED":
            result = NOTSTARTED_PRINTOPERATIONPROCESSINGSTATE
        case "RUNNING":
            result = RUNNING_PRINTOPERATIONPROCESSINGSTATE
        case "SUCCEEDED":
            result = SUCCEEDED_PRINTOPERATIONPROCESSINGSTATE
        case "FAILED":
            result = FAILED_PRINTOPERATIONPROCESSINGSTATE
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_PRINTOPERATIONPROCESSINGSTATE
        default:
            return 0, errors.New("Unknown PrintOperationProcessingState value: " + v)
    }
    return &result, nil
}
func SerializePrintOperationProcessingState(values []PrintOperationProcessingState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
