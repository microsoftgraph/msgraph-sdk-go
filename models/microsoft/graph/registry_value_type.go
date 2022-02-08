package graph
import (
    "strings"
    "errors"
)
// 
type RegistryValueType int

const (
    UNKNOWN_REGISTRYVALUETYPE RegistryValueType = iota
    BINARY_REGISTRYVALUETYPE
    DWORD_REGISTRYVALUETYPE
    DWORDLITTLEENDIAN_REGISTRYVALUETYPE
    DWORDBIGENDIAN_REGISTRYVALUETYPE
    EXPANDSZ_REGISTRYVALUETYPE
    LINK_REGISTRYVALUETYPE
    MULTISZ_REGISTRYVALUETYPE
    NONE_REGISTRYVALUETYPE
    QWORD_REGISTRYVALUETYPE
    QWORDLITTLEENDIAN_REGISTRYVALUETYPE
    SZ_REGISTRYVALUETYPE
    UNKNOWNFUTUREVALUE_REGISTRYVALUETYPE
)

func (i RegistryValueType) String() string {
    return []string{"UNKNOWN", "BINARY", "DWORD", "DWORDLITTLEENDIAN", "DWORDBIGENDIAN", "EXPANDSZ", "LINK", "MULTISZ", "NONE", "QWORD", "QWORDLITTLEENDIAN", "SZ", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseRegistryValueType(v string) (interface{}, error) {
    result := UNKNOWN_REGISTRYVALUETYPE
    switch strings.ToUpper(v) {
        case "UNKNOWN":
            result = UNKNOWN_REGISTRYVALUETYPE
        case "BINARY":
            result = BINARY_REGISTRYVALUETYPE
        case "DWORD":
            result = DWORD_REGISTRYVALUETYPE
        case "DWORDLITTLEENDIAN":
            result = DWORDLITTLEENDIAN_REGISTRYVALUETYPE
        case "DWORDBIGENDIAN":
            result = DWORDBIGENDIAN_REGISTRYVALUETYPE
        case "EXPANDSZ":
            result = EXPANDSZ_REGISTRYVALUETYPE
        case "LINK":
            result = LINK_REGISTRYVALUETYPE
        case "MULTISZ":
            result = MULTISZ_REGISTRYVALUETYPE
        case "NONE":
            result = NONE_REGISTRYVALUETYPE
        case "QWORD":
            result = QWORD_REGISTRYVALUETYPE
        case "QWORDLITTLEENDIAN":
            result = QWORDLITTLEENDIAN_REGISTRYVALUETYPE
        case "SZ":
            result = SZ_REGISTRYVALUETYPE
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_REGISTRYVALUETYPE
        default:
            return 0, errors.New("Unknown RegistryValueType value: " + v)
    }
    return &result, nil
}
func SerializeRegistryValueType(values []RegistryValueType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
