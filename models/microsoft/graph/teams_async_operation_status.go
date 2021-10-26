package graph
import (
    "strings"
    "errors"
)
// 
type TeamsAsyncOperationStatus int

const (
    INVALID_TEAMSASYNCOPERATIONSTATUS TeamsAsyncOperationStatus = iota
    NOTSTARTED_TEAMSASYNCOPERATIONSTATUS
    INPROGRESS_TEAMSASYNCOPERATIONSTATUS
    SUCCEEDED_TEAMSASYNCOPERATIONSTATUS
    FAILED_TEAMSASYNCOPERATIONSTATUS
    UNKNOWNFUTUREVALUE_TEAMSASYNCOPERATIONSTATUS
)

func (i TeamsAsyncOperationStatus) String() string {
    return []string{"INVALID", "NOTSTARTED", "INPROGRESS", "SUCCEEDED", "FAILED", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseTeamsAsyncOperationStatus(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "INVALID":
            return INVALID_TEAMSASYNCOPERATIONSTATUS, nil
        case "NOTSTARTED":
            return NOTSTARTED_TEAMSASYNCOPERATIONSTATUS, nil
        case "INPROGRESS":
            return INPROGRESS_TEAMSASYNCOPERATIONSTATUS, nil
        case "SUCCEEDED":
            return SUCCEEDED_TEAMSASYNCOPERATIONSTATUS, nil
        case "FAILED":
            return FAILED_TEAMSASYNCOPERATIONSTATUS, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_TEAMSASYNCOPERATIONSTATUS, nil
    }
    return 0, errors.New("Unknown TeamsAsyncOperationStatus value: " + v)
}
func SerializeTeamsAsyncOperationStatus(values []TeamsAsyncOperationStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
