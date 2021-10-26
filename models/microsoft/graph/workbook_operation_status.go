package graph
import (
    "strings"
    "errors"
)
// 
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
    switch strings.ToUpper(v) {
        case "NOTSTARTED":
            return NOTSTARTED_WORKBOOKOPERATIONSTATUS, nil
        case "RUNNING":
            return RUNNING_WORKBOOKOPERATIONSTATUS, nil
        case "SUCCEEDED":
            return SUCCEEDED_WORKBOOKOPERATIONSTATUS, nil
        case "FAILED":
            return FAILED_WORKBOOKOPERATIONSTATUS, nil
    }
    return 0, errors.New("Unknown WorkbookOperationStatus value: " + v)
}
func SerializeWorkbookOperationStatus(values []WorkbookOperationStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
