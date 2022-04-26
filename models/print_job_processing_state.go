package models
import (
    "strings"
    "errors"
)
// Provides operations to manage the print singleton.
type PrintJobProcessingState int

const (
    UNKNOWN_PRINTJOBPROCESSINGSTATE PrintJobProcessingState = iota
    PENDING_PRINTJOBPROCESSINGSTATE
    PROCESSING_PRINTJOBPROCESSINGSTATE
    PAUSED_PRINTJOBPROCESSINGSTATE
    STOPPED_PRINTJOBPROCESSINGSTATE
    COMPLETED_PRINTJOBPROCESSINGSTATE
    CANCELED_PRINTJOBPROCESSINGSTATE
    ABORTED_PRINTJOBPROCESSINGSTATE
    UNKNOWNFUTUREVALUE_PRINTJOBPROCESSINGSTATE
)

func (i PrintJobProcessingState) String() string {
    return []string{"UNKNOWN", "PENDING", "PROCESSING", "PAUSED", "STOPPED", "COMPLETED", "CANCELED", "ABORTED", "UNKNOWNFUTUREVALUE"}[i]
}
func ParsePrintJobProcessingState(v string) (interface{}, error) {
    result := UNKNOWN_PRINTJOBPROCESSINGSTATE
    switch strings.ToUpper(v) {
        case "UNKNOWN":
            result = UNKNOWN_PRINTJOBPROCESSINGSTATE
        case "PENDING":
            result = PENDING_PRINTJOBPROCESSINGSTATE
        case "PROCESSING":
            result = PROCESSING_PRINTJOBPROCESSINGSTATE
        case "PAUSED":
            result = PAUSED_PRINTJOBPROCESSINGSTATE
        case "STOPPED":
            result = STOPPED_PRINTJOBPROCESSINGSTATE
        case "COMPLETED":
            result = COMPLETED_PRINTJOBPROCESSINGSTATE
        case "CANCELED":
            result = CANCELED_PRINTJOBPROCESSINGSTATE
        case "ABORTED":
            result = ABORTED_PRINTJOBPROCESSINGSTATE
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_PRINTJOBPROCESSINGSTATE
        default:
            return 0, errors.New("Unknown PrintJobProcessingState value: " + v)
    }
    return &result, nil
}
func SerializePrintJobProcessingState(values []PrintJobProcessingState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
