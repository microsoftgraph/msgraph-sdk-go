package graph
import (
    "strings"
    "errors"
)
// 
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
    switch strings.ToUpper(v) {
        case "STORE":
            return STORE_TEAMSAPPDISTRIBUTIONMETHOD, nil
        case "ORGANIZATION":
            return ORGANIZATION_TEAMSAPPDISTRIBUTIONMETHOD, nil
        case "SIDELOADED":
            return SIDELOADED_TEAMSAPPDISTRIBUTIONMETHOD, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_TEAMSAPPDISTRIBUTIONMETHOD, nil
    }
    return 0, errors.New("Unknown TeamsAppDistributionMethod value: " + v)
}
func SerializeTeamsAppDistributionMethod(values []TeamsAppDistributionMethod) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
