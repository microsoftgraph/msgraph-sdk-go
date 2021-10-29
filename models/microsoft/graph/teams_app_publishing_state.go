package graph
import (
    "strings"
    "errors"
)
// 
type TeamsAppPublishingState int

const (
    SUBMITTED_TEAMSAPPPUBLISHINGSTATE TeamsAppPublishingState = iota
    REJECTED_TEAMSAPPPUBLISHINGSTATE
    PUBLISHED_TEAMSAPPPUBLISHINGSTATE
    UNKNOWNFUTUREVALUE_TEAMSAPPPUBLISHINGSTATE
)

func (i TeamsAppPublishingState) String() string {
    return []string{"SUBMITTED", "REJECTED", "PUBLISHED", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseTeamsAppPublishingState(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "SUBMITTED":
            return SUBMITTED_TEAMSAPPPUBLISHINGSTATE, nil
        case "REJECTED":
            return REJECTED_TEAMSAPPPUBLISHINGSTATE, nil
        case "PUBLISHED":
            return PUBLISHED_TEAMSAPPPUBLISHINGSTATE, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_TEAMSAPPPUBLISHINGSTATE, nil
    }
    return 0, errors.New("Unknown TeamsAppPublishingState value: " + v)
}
func SerializeTeamsAppPublishingState(values []TeamsAppPublishingState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
