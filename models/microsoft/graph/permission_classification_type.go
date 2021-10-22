package graph
import (
    "strings"
    "errors"
)
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
    switch strings.ToUpper(v) {
        case "LOW":
            return LOW_PERMISSIONCLASSIFICATIONTYPE, nil
        case "MEDIUM":
            return MEDIUM_PERMISSIONCLASSIFICATIONTYPE, nil
        case "HIGH":
            return HIGH_PERMISSIONCLASSIFICATIONTYPE, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_PERMISSIONCLASSIFICATIONTYPE, nil
    }
    return 0, errors.New("Unknown PermissionClassificationType value: " + v)
}
func SerializePermissionClassificationType(values []PermissionClassificationType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
