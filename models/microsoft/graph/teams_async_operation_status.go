package graph
import (
    "strings"
    "errors"
)
// Provides operations to manage the educationRoot singleton.
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
    result := INVALID_TEAMSASYNCOPERATIONSTATUS
    switch strings.ToUpper(v) {
        case "INVALID":
            result = INVALID_TEAMSASYNCOPERATIONSTATUS
        case "NOTSTARTED":
            result = NOTSTARTED_TEAMSASYNCOPERATIONSTATUS
        case "INPROGRESS":
            result = INPROGRESS_TEAMSASYNCOPERATIONSTATUS
        case "SUCCEEDED":
            result = SUCCEEDED_TEAMSASYNCOPERATIONSTATUS
        case "FAILED":
            result = FAILED_TEAMSASYNCOPERATIONSTATUS
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_TEAMSASYNCOPERATIONSTATUS
        default:
            return 0, errors.New("Unknown TeamsAsyncOperationStatus value: " + v)
    }
    return &result, nil
}
func SerializeTeamsAsyncOperationStatus(values []TeamsAsyncOperationStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
