package graph
import (
    "strings"
    "errors"
)
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
    switch strings.ToUpper(v) {
        case "PRIVATE":
            return PRIVATE_TEAMVISIBILITYTYPE, nil
        case "PUBLIC":
            return PUBLIC_TEAMVISIBILITYTYPE, nil
        case "HIDDENMEMBERSHIP":
            return HIDDENMEMBERSHIP_TEAMVISIBILITYTYPE, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_TEAMVISIBILITYTYPE, nil
    }
    return 0, errors.New("Unknown TeamVisibilityType value: " + v)
}
func SerializeTeamVisibilityType(values []TeamVisibilityType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
