package models
import (
    "strings"
    "errors"
)
// Provides operations to manage the appCatalogs singleton.
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
    result := SUBMITTED_TEAMSAPPPUBLISHINGSTATE
    switch strings.ToUpper(v) {
        case "SUBMITTED":
            result = SUBMITTED_TEAMSAPPPUBLISHINGSTATE
        case "REJECTED":
            result = REJECTED_TEAMSAPPPUBLISHINGSTATE
        case "PUBLISHED":
            result = PUBLISHED_TEAMSAPPPUBLISHINGSTATE
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_TEAMSAPPPUBLISHINGSTATE
        default:
            return 0, errors.New("Unknown TeamsAppPublishingState value: " + v)
    }
    return &result, nil
}
func SerializeTeamsAppPublishingState(values []TeamsAppPublishingState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
