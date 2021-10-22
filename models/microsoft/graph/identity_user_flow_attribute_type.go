package graph
import (
    "strings"
    "errors"
)
type IdentityUserFlowAttributeType int

const (
    BUILTIN_IDENTITYUSERFLOWATTRIBUTETYPE IdentityUserFlowAttributeType = iota
    CUSTOM_IDENTITYUSERFLOWATTRIBUTETYPE
    REQUIRED_IDENTITYUSERFLOWATTRIBUTETYPE
    UNKNOWNFUTUREVALUE_IDENTITYUSERFLOWATTRIBUTETYPE
)

func (i IdentityUserFlowAttributeType) String() string {
    return []string{"BUILTIN", "CUSTOM", "REQUIRED", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseIdentityUserFlowAttributeType(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "BUILTIN":
            return BUILTIN_IDENTITYUSERFLOWATTRIBUTETYPE, nil
        case "CUSTOM":
            return CUSTOM_IDENTITYUSERFLOWATTRIBUTETYPE, nil
        case "REQUIRED":
            return REQUIRED_IDENTITYUSERFLOWATTRIBUTETYPE, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_IDENTITYUSERFLOWATTRIBUTETYPE, nil
    }
    return 0, errors.New("Unknown IdentityUserFlowAttributeType value: " + v)
}
func SerializeIdentityUserFlowAttributeType(values []IdentityUserFlowAttributeType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
