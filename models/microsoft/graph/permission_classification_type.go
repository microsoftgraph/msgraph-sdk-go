package graph
import (
    "strings"
    "errors"
)
// Provides operations to call the instantiate method.
type PermissionClassificationType int

const (
    LOW_PERMISSIONCLASSIFICATIONTYPE PermissionClassificationType = iota
    MEDIUM_PERMISSIONCLASSIFICATIONTYPE
    HIGH_PERMISSIONCLASSIFICATIONTYPE
    UNKNOWNFUTUREVALUE_PERMISSIONCLASSIFICATIONTYPE
)

func (i PermissionClassificationType) String() string {
    return []string{"LOW", "MEDIUM", "HIGH", "UNKNOWNFUTUREVALUE"}[i]
}
func ParsePermissionClassificationType(v string) (interface{}, error) {
    result := LOW_PERMISSIONCLASSIFICATIONTYPE
    switch strings.ToUpper(v) {
        case "LOW":
            result = LOW_PERMISSIONCLASSIFICATIONTYPE
        case "MEDIUM":
            result = MEDIUM_PERMISSIONCLASSIFICATIONTYPE
        case "HIGH":
            result = HIGH_PERMISSIONCLASSIFICATIONTYPE
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_PERMISSIONCLASSIFICATIONTYPE
        default:
            return 0, errors.New("Unknown PermissionClassificationType value: " + v)
    }
    return &result, nil
}
func SerializePermissionClassificationType(values []PermissionClassificationType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
