package graph
import (
    "strings"
    "errors"
)
type GroupType int

const (
    UNIFIEDGROUPS_GROUPTYPE GroupType = iota
    AZUREAD_GROUPTYPE
    UNKNOWNFUTUREVALUE_GROUPTYPE
)

func (i GroupType) String() string {
    return []string{"UNIFIEDGROUPS", "AZUREAD", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseGroupType(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "UNIFIEDGROUPS":
            return UNIFIEDGROUPS_GROUPTYPE, nil
        case "AZUREAD":
            return AZUREAD_GROUPTYPE, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_GROUPTYPE, nil
    }
    return 0, errors.New("Unknown GroupType value: " + v)
}
func SerializeGroupType(values []GroupType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
