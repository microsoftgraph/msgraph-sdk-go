package graph
import (
    "strings"
    "errors"
)
type PrintTaskProcessingState int

const (
    PENDING_PRINTTASKPROCESSINGSTATE PrintTaskProcessingState = iota
    PROCESSING_PRINTTASKPROCESSINGSTATE
    COMPLETED_PRINTTASKPROCESSINGSTATE
    ABORTED_PRINTTASKPROCESSINGSTATE
    UNKNOWNFUTUREVALUE_PRINTTASKPROCESSINGSTATE
)

func (i PrintTaskProcessingState) String() string {
    return []string{"PENDING", "PROCESSING", "COMPLETED", "ABORTED", "UNKNOWNFUTUREVALUE"}[i]
}
func ParsePrintTaskProcessingState(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "PENDING":
            return PENDING_PRINTTASKPROCESSINGSTATE, nil
        case "PROCESSING":
            return PROCESSING_PRINTTASKPROCESSINGSTATE, nil
        case "COMPLETED":
            return COMPLETED_PRINTTASKPROCESSINGSTATE, nil
        case "ABORTED":
            return ABORTED_PRINTTASKPROCESSINGSTATE, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_PRINTTASKPROCESSINGSTATE, nil
    }
    return 0, errors.New("Unknown PrintTaskProcessingState value: " + v)
}
func SerializePrintTaskProcessingState(values []PrintTaskProcessingState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
