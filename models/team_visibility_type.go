package models
import (
    "strings"
    "errors"
)
// Provides operations to manage the collection of drive entities.
type TeamVisibilityType int

const (
    PRIVATE_TEAMVISIBILITYTYPE TeamVisibilityType = iota
    PUBLIC_TEAMVISIBILITYTYPE
    HIDDENMEMBERSHIP_TEAMVISIBILITYTYPE
    UNKNOWNFUTUREVALUE_TEAMVISIBILITYTYPE
)

func (i TeamVisibilityType) String() string {
    return []string{"PRIVATE", "PUBLIC", "HIDDENMEMBERSHIP", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseTeamVisibilityType(v string) (interface{}, error) {
    result := PRIVATE_TEAMVISIBILITYTYPE
    switch strings.ToUpper(v) {
        case "PRIVATE":
            result = PRIVATE_TEAMVISIBILITYTYPE
        case "PUBLIC":
            result = PUBLIC_TEAMVISIBILITYTYPE
        case "HIDDENMEMBERSHIP":
            result = HIDDENMEMBERSHIP_TEAMVISIBILITYTYPE
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_TEAMVISIBILITYTYPE
        default:
            return 0, errors.New("Unknown TeamVisibilityType value: " + v)
    }
    return &result, nil
}
func SerializeTeamVisibilityType(values []TeamVisibilityType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
