package graph
import (
    "strings"
    "errors"
)
// Provides operations to manage the identityContainer singleton.
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
    result := BUILTIN_IDENTITYUSERFLOWATTRIBUTETYPE
    switch strings.ToUpper(v) {
        case "BUILTIN":
            result = BUILTIN_IDENTITYUSERFLOWATTRIBUTETYPE
        case "CUSTOM":
            result = CUSTOM_IDENTITYUSERFLOWATTRIBUTETYPE
        case "REQUIRED":
            result = REQUIRED_IDENTITYUSERFLOWATTRIBUTETYPE
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_IDENTITYUSERFLOWATTRIBUTETYPE
        default:
            return 0, errors.New("Unknown IdentityUserFlowAttributeType value: " + v)
    }
    return &result, nil
}
func SerializeIdentityUserFlowAttributeType(values []IdentityUserFlowAttributeType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
