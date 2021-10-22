package graph
import (
    "strings"
    "errors"
)
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
    switch strings.ToUpper(v) {
        case "NOTSTARTED":
            return NOTSTARTED_DATAPOLICYOPERATIONSTATUS, nil
        case "RUNNING":
            return RUNNING_DATAPOLICYOPERATIONSTATUS, nil
        case "COMPLETE":
            return COMPLETE_DATAPOLICYOPERATIONSTATUS, nil
        case "FAILED":
            return FAILED_DATAPOLICYOPERATIONSTATUS, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_DATAPOLICYOPERATIONSTATUS, nil
    }
    return 0, errors.New("Unknown DataPolicyOperationStatus value: " + v)
}
func SerializeDataPolicyOperationStatus(values []DataPolicyOperationStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
