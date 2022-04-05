package models
import (
    "strings"
    "errors"
)
// Provides operations to manage the print singleton.
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
    result := PENDING_PRINTTASKPROCESSINGSTATE
    switch strings.ToUpper(v) {
        case "PENDING":
            result = PENDING_PRINTTASKPROCESSINGSTATE
        case "PROCESSING":
            result = PROCESSING_PRINTTASKPROCESSINGSTATE
        case "COMPLETED":
            result = COMPLETED_PRINTTASKPROCESSINGSTATE
        case "ABORTED":
            result = ABORTED_PRINTTASKPROCESSINGSTATE
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_PRINTTASKPROCESSINGSTATE
        default:
            return 0, errors.New("Unknown PrintTaskProcessingState value: " + v)
    }
    return &result, nil
}
func SerializePrintTaskProcessingState(values []PrintTaskProcessingState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
