package externalconnectors
import (
    "strings"
    "errors"
)
// Provides operations to manage the collection of externalConnection entities.
type AclType int

const (
    USER_ACLTYPE AclType = iota
    GROUP_ACLTYPE
    EVERYONE_ACLTYPE
    EVERYONEEXCEPTGUESTS_ACLTYPE
    EXTERNALGROUP_ACLTYPE
    UNKNOWNFUTUREVALUE_ACLTYPE
)

func (i AclType) String() string {
    return []string{"USER", "GROUP", "EVERYONE", "EVERYONEEXCEPTGUESTS", "EXTERNALGROUP", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseAclType(v string) (interface{}, error) {
    result := USER_ACLTYPE
    switch strings.ToUpper(v) {
        case "USER":
            result = USER_ACLTYPE
        case "GROUP":
            result = GROUP_ACLTYPE
        case "EVERYONE":
            result = EVERYONE_ACLTYPE
        case "EVERYONEEXCEPTGUESTS":
            result = EVERYONEEXCEPTGUESTS_ACLTYPE
        case "EXTERNALGROUP":
            result = EXTERNALGROUP_ACLTYPE
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_ACLTYPE
        default:
            return 0, errors.New("Unknown AclType value: " + v)
    }
    return &result, nil
}
func SerializeAclType(values []AclType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
