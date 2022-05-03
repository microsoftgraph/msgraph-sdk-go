package models
import (
    "strings"
    "errors"
)
// Provides operations to manage the drive singleton.
type LongRunningOperationStatus int

const (
    NOTSTARTED_LONGRUNNINGOPERATIONSTATUS LongRunningOperationStatus = iota
    RUNNING_LONGRUNNINGOPERATIONSTATUS
    SUCCEEDED_LONGRUNNINGOPERATIONSTATUS
    FAILED_LONGRUNNINGOPERATIONSTATUS
    UNKNOWNFUTUREVALUE_LONGRUNNINGOPERATIONSTATUS
)

func (i LongRunningOperationStatus) String() string {
    return []string{"NOTSTARTED", "RUNNING", "SUCCEEDED", "FAILED", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseLongRunningOperationStatus(v string) (interface{}, error) {
    result := NOTSTARTED_LONGRUNNINGOPERATIONSTATUS
    switch strings.ToUpper(v) {
        case "NOTSTARTED":
            result = NOTSTARTED_LONGRUNNINGOPERATIONSTATUS
        case "RUNNING":
            result = RUNNING_LONGRUNNINGOPERATIONSTATUS
        case "SUCCEEDED":
            result = SUCCEEDED_LONGRUNNINGOPERATIONSTATUS
        case "FAILED":
            result = FAILED_LONGRUNNINGOPERATIONSTATUS
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_LONGRUNNINGOPERATIONSTATUS
        default:
            return 0, errors.New("Unknown LongRunningOperationStatus value: " + v)
    }
    return &result, nil
}
func SerializeLongRunningOperationStatus(values []LongRunningOperationStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
