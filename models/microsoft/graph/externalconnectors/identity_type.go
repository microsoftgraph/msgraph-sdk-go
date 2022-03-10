package externalconnectors
import (
    "strings"
    "errors"
)
// Provides operations to manage the collection of externalConnection entities.
type IdentityType int

const (
    USER_IDENTITYTYPE IdentityType = iota
    GROUP_IDENTITYTYPE
    EXTERNALGROUP_IDENTITYTYPE
    UNKNOWNFUTUREVALUE_IDENTITYTYPE
)

func (i IdentityType) String() string {
    return []string{"USER", "GROUP", "EXTERNALGROUP", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseIdentityType(v string) (interface{}, error) {
    result := USER_IDENTITYTYPE
    switch strings.ToUpper(v) {
        case "USER":
            result = USER_IDENTITYTYPE
        case "GROUP":
            result = GROUP_IDENTITYTYPE
        case "EXTERNALGROUP":
            result = EXTERNALGROUP_IDENTITYTYPE
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_IDENTITYTYPE
        default:
            return 0, errors.New("Unknown IdentityType value: " + v)
    }
    return &result, nil
}
func SerializeIdentityType(values []IdentityType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
