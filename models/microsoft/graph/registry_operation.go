package graph
import (
    "strings"
    "errors"
)
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
    switch strings.ToUpper(v) {
        case "UNKNOWN":
            return UNKNOWN_REGISTRYOPERATION, nil
        case "CREATE":
            return CREATE_REGISTRYOPERATION, nil
        case "MODIFY":
            return MODIFY_REGISTRYOPERATION, nil
        case "DELETE":
            return DELETE_REGISTRYOPERATION, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_REGISTRYOPERATION, nil
    }
    return 0, errors.New("Unknown RegistryOperation value: " + v)
}
func SerializeRegistryOperation(values []RegistryOperation) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
