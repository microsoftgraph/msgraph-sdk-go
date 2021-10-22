package graph
import (
    "strings"
    "errors"
)
type PermissionType int

const (
    APPLICATION_PERMISSIONTYPE PermissionType = iota
    DELEGATED_PERMISSIONTYPE
    DELEGATEDUSERCONSENTABLE_PERMISSIONTYPE
)

func (i PermissionType) String() string {
    return []string{"APPLICATION", "DELEGATED", "DELEGATEDUSERCONSENTABLE"}[i]
}
func ParsePermissionType(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "APPLICATION":
            return APPLICATION_PERMISSIONTYPE, nil
        case "DELEGATED":
            return DELEGATED_PERMISSIONTYPE, nil
        case "DELEGATEDUSERCONSENTABLE":
            return DELEGATEDUSERCONSENTABLE_PERMISSIONTYPE, nil
    }
    return 0, errors.New("Unknown PermissionType value: " + v)
}
func SerializePermissionType(values []PermissionType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
