package graph
import (
    "strings"
    "errors"
)
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
    switch strings.ToUpper(v) {
        case "STRING":
            return STRING_IDENTITYUSERFLOWATTRIBUTEDATATYPE, nil
        case "BOOLEAN":
            return BOOLEAN_IDENTITYUSERFLOWATTRIBUTEDATATYPE, nil
        case "INT64":
            return INT64_IDENTITYUSERFLOWATTRIBUTEDATATYPE, nil
        case "STRINGCOLLECTION":
            return STRINGCOLLECTION_IDENTITYUSERFLOWATTRIBUTEDATATYPE, nil
        case "DATETIME":
            return DATETIME_IDENTITYUSERFLOWATTRIBUTEDATATYPE, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_IDENTITYUSERFLOWATTRIBUTEDATATYPE, nil
    }
    return 0, errors.New("Unknown IdentityUserFlowAttributeDataType value: " + v)
}
func SerializeIdentityUserFlowAttributeDataType(values []IdentityUserFlowAttributeDataType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
