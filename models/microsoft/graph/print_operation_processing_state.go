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
    switch strings.ToUpper(v) {
        case "NOTSTARTED":
            return NOTSTARTED_PRINTOPERATIONPROCESSINGSTATE, nil
        case "RUNNING":
            return RUNNING_PRINTOPERATIONPROCESSINGSTATE, nil
        case "SUCCEEDED":
            return SUCCEEDED_PRINTOPERATIONPROCESSINGSTATE, nil
        case "FAILED":
            return FAILED_PRINTOPERATIONPROCESSINGSTATE, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_PRINTOPERATIONPROCESSINGSTATE, nil
    }
    return 0, errors.New("Unknown PrintOperationProcessingState value: " + v)
}
func SerializePrintOperationProcessingState(values []PrintOperationProcessingState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
