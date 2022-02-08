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
    result := NOTSTARTED_OPERATIONSTATUS
    switch strings.ToUpper(v) {
        case "NOTSTARTED":
            result = NOTSTARTED_OPERATIONSTATUS
        case "RUNNING":
            result = RUNNING_OPERATIONSTATUS
        case "COMPLETED":
            result = COMPLETED_OPERATIONSTATUS
        case "FAILED":
            result = FAILED_OPERATIONSTATUS
        default:
            return 0, errors.New("Unknown OperationStatus value: " + v)
    }
    return &result, nil
}
func SerializeOperationStatus(values []OperationStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
