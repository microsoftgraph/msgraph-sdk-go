package graph
import (
    "strings"
    "errors"
)
// Provides operations to manage the drive singleton.
type WorkbookOperationStatus int

const (
    NOTSTARTED_WORKBOOKOPERATIONSTATUS WorkbookOperationStatus = iota
    RUNNING_WORKBOOKOPERATIONSTATUS
    SUCCEEDED_WORKBOOKOPERATIONSTATUS
    FAILED_WORKBOOKOPERATIONSTATUS
)

func (i WorkbookOperationStatus) String() string {
    return []string{"NOTSTARTED", "RUNNING", "SUCCEEDED", "FAILED"}[i]
}
func ParseWorkbookOperationStatus(v string) (interface{}, error) {
    result := NOTSTARTED_WORKBOOKOPERATIONSTATUS
    switch strings.ToUpper(v) {
        case "NOTSTARTED":
            result = NOTSTARTED_WORKBOOKOPERATIONSTATUS
        case "RUNNING":
            result = RUNNING_WORKBOOKOPERATIONSTATUS
        case "SUCCEEDED":
            result = SUCCEEDED_WORKBOOKOPERATIONSTATUS
        case "FAILED":
            result = FAILED_WORKBOOKOPERATIONSTATUS
        default:
            return 0, errors.New("Unknown WorkbookOperationStatus value: " + v)
    }
    return &result, nil
}
func SerializeWorkbookOperationStatus(values []WorkbookOperationStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
