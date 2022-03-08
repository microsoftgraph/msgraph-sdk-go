package graph
import (
    "strings"
    "errors"
)
// Provides operations to manage the appCatalogs singleton.
type TeamsAppDistributionMethod int

const (
    STORE_TEAMSAPPDISTRIBUTIONMETHOD TeamsAppDistributionMethod = iota
    ORGANIZATION_TEAMSAPPDISTRIBUTIONMETHOD
    SIDELOADED_TEAMSAPPDISTRIBUTIONMETHOD
    UNKNOWNFUTUREVALUE_TEAMSAPPDISTRIBUTIONMETHOD
)

func (i TeamsAppDistributionMethod) String() string {
    return []string{"STORE", "ORGANIZATION", "SIDELOADED", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseTeamsAppDistributionMethod(v string) (interface{}, error) {
    result := STORE_TEAMSAPPDISTRIBUTIONMETHOD
    switch strings.ToUpper(v) {
        case "STORE":
            result = STORE_TEAMSAPPDISTRIBUTIONMETHOD
        case "ORGANIZATION":
            result = ORGANIZATION_TEAMSAPPDISTRIBUTIONMETHOD
        case "SIDELOADED":
            result = SIDELOADED_TEAMSAPPDISTRIBUTIONMETHOD
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_TEAMSAPPDISTRIBUTIONMETHOD
        default:
            return 0, errors.New("Unknown TeamsAppDistributionMethod value: " + v)
    }
    return &result, nil
}
func SerializeTeamsAppDistributionMethod(values []TeamsAppDistributionMethod) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
