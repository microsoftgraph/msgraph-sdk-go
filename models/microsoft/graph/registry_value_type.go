package graph
import (
    "strings"
    "errors"
)
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
    switch strings.ToUpper(v) {
        case "UNKNOWN":
            return UNKNOWN_REGISTRYVALUETYPE, nil
        case "BINARY":
            return BINARY_REGISTRYVALUETYPE, nil
        case "DWORD":
            return DWORD_REGISTRYVALUETYPE, nil
        case "DWORDLITTLEENDIAN":
            return DWORDLITTLEENDIAN_REGISTRYVALUETYPE, nil
        case "DWORDBIGENDIAN":
            return DWORDBIGENDIAN_REGISTRYVALUETYPE, nil
        case "EXPANDSZ":
            return EXPANDSZ_REGISTRYVALUETYPE, nil
        case "LINK":
            return LINK_REGISTRYVALUETYPE, nil
        case "MULTISZ":
            return MULTISZ_REGISTRYVALUETYPE, nil
        case "NONE":
            return NONE_REGISTRYVALUETYPE, nil
        case "QWORD":
            return QWORD_REGISTRYVALUETYPE, nil
        case "QWORDLITTLEENDIAN":
            return QWORDLITTLEENDIAN_REGISTRYVALUETYPE, nil
        case "SZ":
            return SZ_REGISTRYVALUETYPE, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_REGISTRYVALUETYPE, nil
    }
    return 0, errors.New("Unknown RegistryValueType value: " + v)
}
func SerializeRegistryValueType(values []RegistryValueType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
