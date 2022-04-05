package externalconnectors
import (
    "strings"
    "errors"
)
// Provides operations to manage the collection of externalConnection entities.
type ConnectionOperationStatus int

const (
    UNSPECIFIED_CONNECTIONOPERATIONSTATUS ConnectionOperationStatus = iota
    INPROGRESS_CONNECTIONOPERATIONSTATUS
    COMPLETED_CONNECTIONOPERATIONSTATUS
    FAILED_CONNECTIONOPERATIONSTATUS
    UNKNOWNFUTUREVALUE_CONNECTIONOPERATIONSTATUS
)

func (i ConnectionOperationStatus) String() string {
    return []string{"UNSPECIFIED", "INPROGRESS", "COMPLETED", "FAILED", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseConnectionOperationStatus(v string) (interface{}, error) {
    result := UNSPECIFIED_CONNECTIONOPERATIONSTATUS
    switch strings.ToUpper(v) {
        case "UNSPECIFIED":
            result = UNSPECIFIED_CONNECTIONOPERATIONSTATUS
        case "INPROGRESS":
            result = INPROGRESS_CONNECTIONOPERATIONSTATUS
        case "COMPLETED":
            result = COMPLETED_CONNECTIONOPERATIONSTATUS
        case "FAILED":
            result = FAILED_CONNECTIONOPERATIONSTATUS
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_CONNECTIONOPERATIONSTATUS
        default:
            return 0, errors.New("Unknown ConnectionOperationStatus value: " + v)
    }
    return &result, nil
}
func SerializeConnectionOperationStatus(values []ConnectionOperationStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
