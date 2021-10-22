package externalconnectors
import (
    "strings"
    "errors"
)
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
    switch strings.ToUpper(v) {
        case "UNSPECIFIED":
            return UNSPECIFIED_CONNECTIONOPERATIONSTATUS, nil
        case "INPROGRESS":
            return INPROGRESS_CONNECTIONOPERATIONSTATUS, nil
        case "COMPLETED":
            return COMPLETED_CONNECTIONOPERATIONSTATUS, nil
        case "FAILED":
            return FAILED_CONNECTIONOPERATIONSTATUS, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_CONNECTIONOPERATIONSTATUS, nil
    }
    return 0, errors.New("Unknown ConnectionOperationStatus value: " + v)
}
func SerializeConnectionOperationStatus(values []ConnectionOperationStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
