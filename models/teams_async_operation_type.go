package models
import (
    "strings"
    "errors"
)
// Provides operations to manage the collection of drive entities.
type TeamsAsyncOperationType int

const (
    INVALID_TEAMSASYNCOPERATIONTYPE TeamsAsyncOperationType = iota
    CLONETEAM_TEAMSASYNCOPERATIONTYPE
    ARCHIVETEAM_TEAMSASYNCOPERATIONTYPE
    UNARCHIVETEAM_TEAMSASYNCOPERATIONTYPE
    CREATETEAM_TEAMSASYNCOPERATIONTYPE
    UNKNOWNFUTUREVALUE_TEAMSASYNCOPERATIONTYPE
)

func (i TeamsAsyncOperationType) String() string {
    return []string{"INVALID", "CLONETEAM", "ARCHIVETEAM", "UNARCHIVETEAM", "CREATETEAM", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseTeamsAsyncOperationType(v string) (interface{}, error) {
    result := INVALID_TEAMSASYNCOPERATIONTYPE
    switch strings.ToUpper(v) {
        case "INVALID":
            result = INVALID_TEAMSASYNCOPERATIONTYPE
        case "CLONETEAM":
            result = CLONETEAM_TEAMSASYNCOPERATIONTYPE
        case "ARCHIVETEAM":
            result = ARCHIVETEAM_TEAMSASYNCOPERATIONTYPE
        case "UNARCHIVETEAM":
            result = UNARCHIVETEAM_TEAMSASYNCOPERATIONTYPE
        case "CREATETEAM":
            result = CREATETEAM_TEAMSASYNCOPERATIONTYPE
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_TEAMSASYNCOPERATIONTYPE
        default:
            return 0, errors.New("Unknown TeamsAsyncOperationType value: " + v)
    }
    return &result, nil
}
func SerializeTeamsAsyncOperationType(values []TeamsAsyncOperationType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
