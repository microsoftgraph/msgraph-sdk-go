package graph
import (
    "strings"
    "errors"
)
// 
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
    switch strings.ToUpper(v) {
        case "INVALID":
            return INVALID_TEAMSASYNCOPERATIONTYPE, nil
        case "CLONETEAM":
            return CLONETEAM_TEAMSASYNCOPERATIONTYPE, nil
        case "ARCHIVETEAM":
            return ARCHIVETEAM_TEAMSASYNCOPERATIONTYPE, nil
        case "UNARCHIVETEAM":
            return UNARCHIVETEAM_TEAMSASYNCOPERATIONTYPE, nil
        case "CREATETEAM":
            return CREATETEAM_TEAMSASYNCOPERATIONTYPE, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_TEAMSASYNCOPERATIONTYPE, nil
    }
    return 0, errors.New("Unknown TeamsAsyncOperationType value: " + v)
}
func SerializeTeamsAsyncOperationType(values []TeamsAsyncOperationType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
