package models
import (
    "strings"
    "errors"
)
// Provides operations to manage the identityContainer singleton.
type IdentityUserFlowAttributeDataType int

const (
    STRING_IDENTITYUSERFLOWATTRIBUTEDATATYPE IdentityUserFlowAttributeDataType = iota
    BOOLEAN_IDENTITYUSERFLOWATTRIBUTEDATATYPE
    INT64_IDENTITYUSERFLOWATTRIBUTEDATATYPE
    STRINGCOLLECTION_IDENTITYUSERFLOWATTRIBUTEDATATYPE
    DATETIME_IDENTITYUSERFLOWATTRIBUTEDATATYPE
    UNKNOWNFUTUREVALUE_IDENTITYUSERFLOWATTRIBUTEDATATYPE
)

func (i IdentityUserFlowAttributeDataType) String() string {
    return []string{"STRING", "BOOLEAN", "INT64", "STRINGCOLLECTION", "DATETIME", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseIdentityUserFlowAttributeDataType(v string) (interface{}, error) {
    result := STRING_IDENTITYUSERFLOWATTRIBUTEDATATYPE
    switch strings.ToUpper(v) {
        case "STRING":
            result = STRING_IDENTITYUSERFLOWATTRIBUTEDATATYPE
        case "BOOLEAN":
            result = BOOLEAN_IDENTITYUSERFLOWATTRIBUTEDATATYPE
        case "INT64":
            result = INT64_IDENTITYUSERFLOWATTRIBUTEDATATYPE
        case "STRINGCOLLECTION":
            result = STRINGCOLLECTION_IDENTITYUSERFLOWATTRIBUTEDATATYPE
        case "DATETIME":
            result = DATETIME_IDENTITYUSERFLOWATTRIBUTEDATATYPE
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_IDENTITYUSERFLOWATTRIBUTEDATATYPE
        default:
            return 0, errors.New("Unknown IdentityUserFlowAttributeDataType value: " + v)
    }
    return &result, nil
}
func SerializeIdentityUserFlowAttributeDataType(values []IdentityUserFlowAttributeDataType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
