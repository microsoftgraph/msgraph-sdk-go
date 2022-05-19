package models
import (
    "errors"
)
// Casts the previous resource to user.
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
    return []string{"invalid", "cloneTeam", "archiveTeam", "unarchiveTeam", "createTeam", "unknownFutureValue"}[i]
}
func ParseTeamsAsyncOperationType(v string) (interface{}, error) {
    result := INVALID_TEAMSASYNCOPERATIONTYPE
    switch v {
        case "invalid":
            result = INVALID_TEAMSASYNCOPERATIONTYPE
        case "cloneTeam":
            result = CLONETEAM_TEAMSASYNCOPERATIONTYPE
        case "archiveTeam":
            result = ARCHIVETEAM_TEAMSASYNCOPERATIONTYPE
        case "unarchiveTeam":
            result = UNARCHIVETEAM_TEAMSASYNCOPERATIONTYPE
        case "createTeam":
            result = CREATETEAM_TEAMSASYNCOPERATIONTYPE
        case "unknownFutureValue":
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
