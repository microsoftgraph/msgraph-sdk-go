package graph
import (
    "strings"
    "errors"
)
// 
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
    switch strings.ToUpper(v) {
        case "UNKNOWN":
            return UNKNOWN_PRINTJOBPROCESSINGSTATE, nil
        case "PENDING":
            return PENDING_PRINTJOBPROCESSINGSTATE, nil
        case "PROCESSING":
            return PROCESSING_PRINTJOBPROCESSINGSTATE, nil
        case "PAUSED":
            return PAUSED_PRINTJOBPROCESSINGSTATE, nil
        case "STOPPED":
            return STOPPED_PRINTJOBPROCESSINGSTATE, nil
        case "COMPLETED":
            return COMPLETED_PRINTJOBPROCESSINGSTATE, nil
        case "CANCELED":
            return CANCELED_PRINTJOBPROCESSINGSTATE, nil
        case "ABORTED":
            return ABORTED_PRINTJOBPROCESSINGSTATE, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_PRINTJOBPROCESSINGSTATE, nil
    }
    return 0, errors.New("Unknown PrintJobProcessingState value: " + v)
}
func SerializePrintJobProcessingState(values []PrintJobProcessingState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
