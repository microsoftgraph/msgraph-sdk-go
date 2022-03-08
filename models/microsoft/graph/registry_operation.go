package graph
import (
    "strings"
    "errors"
)
// Provides operations to manage the security singleton.
type RegistryOperation int

const (
    UNKNOWN_REGISTRYOPERATION RegistryOperation = iota
    CREATE_REGISTRYOPERATION
    MODIFY_REGISTRYOPERATION
    DELETE_REGISTRYOPERATION
    UNKNOWNFUTUREVALUE_REGISTRYOPERATION
)

func (i RegistryOperation) String() string {
    return []string{"UNKNOWN", "CREATE", "MODIFY", "DELETE", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseRegistryOperation(v string) (interface{}, error) {
    result := UNKNOWN_REGISTRYOPERATION
    switch strings.ToUpper(v) {
        case "UNKNOWN":
            result = UNKNOWN_REGISTRYOPERATION
        case "CREATE":
            result = CREATE_REGISTRYOPERATION
        case "MODIFY":
            result = MODIFY_REGISTRYOPERATION
        case "DELETE":
            result = DELETE_REGISTRYOPERATION
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_REGISTRYOPERATION
        default:
            return 0, errors.New("Unknown RegistryOperation value: " + v)
    }
    return &result, nil
}
func SerializeRegistryOperation(values []RegistryOperation) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
