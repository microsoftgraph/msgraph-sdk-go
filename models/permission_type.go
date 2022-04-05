package models
import (
    "strings"
    "errors"
)
// Provides operations to manage the policyRoot singleton.
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
    result := APPLICATION_PERMISSIONTYPE
    switch strings.ToUpper(v) {
        case "APPLICATION":
            result = APPLICATION_PERMISSIONTYPE
        case "DELEGATED":
            result = DELEGATED_PERMISSIONTYPE
        case "DELEGATEDUSERCONSENTABLE":
            result = DELEGATEDUSERCONSENTABLE_PERMISSIONTYPE
        default:
            return 0, errors.New("Unknown PermissionType value: " + v)
    }
    return &result, nil
}
func SerializePermissionType(values []PermissionType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
