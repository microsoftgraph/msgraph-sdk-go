package externalconnectors
import (
    "strings"
    "errors"
)
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
    switch strings.ToUpper(v) {
        case "USER":
            return USER_ACLTYPE, nil
        case "GROUP":
            return GROUP_ACLTYPE, nil
        case "EVERYONE":
            return EVERYONE_ACLTYPE, nil
        case "EVERYONEEXCEPTGUESTS":
            return EVERYONEEXCEPTGUESTS_ACLTYPE, nil
        case "EXTERNALGROUP":
            return EXTERNALGROUP_ACLTYPE, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_ACLTYPE, nil
    }
    return 0, errors.New("Unknown AclType value: " + v)
}
func SerializeAclType(values []AclType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
