package graph
import (
    "strings"
    "errors"
)
// 
type OperationStatus int

const (
    NOTSTARTED_OPERATIONSTATUS OperationStatus = iota
    RUNNING_OPERATIONSTATUS
    COMPLETED_OPERATIONSTATUS
    FAILED_OPERATIONSTATUS
)

func (i OperationStatus) String() string {
    return []string{"NOTSTARTED", "RUNNING", "COMPLETED", "FAILED"}[i]
}
func ParseOperationStatus(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "NOTSTARTED":
            return NOTSTARTED_OPERATIONSTATUS, nil
        case "RUNNING":
            return RUNNING_OPERATIONSTATUS, nil
        case "COMPLETED":
            return COMPLETED_OPERATIONSTATUS, nil
        case "FAILED":
            return FAILED_OPERATIONSTATUS, nil
    }
    return 0, errors.New("Unknown OperationStatus value: " + v)
}
func SerializeOperationStatus(values []OperationStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
