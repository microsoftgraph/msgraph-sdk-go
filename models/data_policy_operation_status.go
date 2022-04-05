package models
import (
    "strings"
    "errors"
)
// Provides operations to manage the collection of dataPolicyOperation entities.
type DataPolicyOperationStatus int

const (
    NOTSTARTED_DATAPOLICYOPERATIONSTATUS DataPolicyOperationStatus = iota
    RUNNING_DATAPOLICYOPERATIONSTATUS
    COMPLETE_DATAPOLICYOPERATIONSTATUS
    FAILED_DATAPOLICYOPERATIONSTATUS
    UNKNOWNFUTUREVALUE_DATAPOLICYOPERATIONSTATUS
)

func (i DataPolicyOperationStatus) String() string {
    return []string{"NOTSTARTED", "RUNNING", "COMPLETE", "FAILED", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseDataPolicyOperationStatus(v string) (interface{}, error) {
    result := NOTSTARTED_DATAPOLICYOPERATIONSTATUS
    switch strings.ToUpper(v) {
        case "NOTSTARTED":
            result = NOTSTARTED_DATAPOLICYOPERATIONSTATUS
        case "RUNNING":
            result = RUNNING_DATAPOLICYOPERATIONSTATUS
        case "COMPLETE":
            result = COMPLETE_DATAPOLICYOPERATIONSTATUS
        case "FAILED":
            result = FAILED_DATAPOLICYOPERATIONSTATUS
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_DATAPOLICYOPERATIONSTATUS
        default:
            return 0, errors.New("Unknown DataPolicyOperationStatus value: " + v)
    }
    return &result, nil
}
func SerializeDataPolicyOperationStatus(values []DataPolicyOperationStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
