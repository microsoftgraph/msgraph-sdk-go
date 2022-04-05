package models
import (
    "strings"
    "errors"
)
// Provides operations to manage the auditLogRoot singleton.
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
    result := UNIFIEDGROUPS_GROUPTYPE
    switch strings.ToUpper(v) {
        case "UNIFIEDGROUPS":
            result = UNIFIEDGROUPS_GROUPTYPE
        case "AZUREAD":
            result = AZUREAD_GROUPTYPE
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_GROUPTYPE
        default:
            return 0, errors.New("Unknown GroupType value: " + v)
    }
    return &result, nil
}
func SerializeGroupType(values []GroupType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
