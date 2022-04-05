package models
import (
    "strings"
    "errors"
)
// Provides operations to manage the collection of drive entities.
type OnenoteUserRole int

const (
    NONE_ONENOTEUSERROLE OnenoteUserRole = iota
    OWNER_ONENOTEUSERROLE
    CONTRIBUTOR_ONENOTEUSERROLE
    READER_ONENOTEUSERROLE
)

func (i OnenoteUserRole) String() string {
    return []string{"NONE", "OWNER", "CONTRIBUTOR", "READER"}[i]
}
func ParseOnenoteUserRole(v string) (interface{}, error) {
    result := NONE_ONENOTEUSERROLE
    switch strings.ToUpper(v) {
        case "NONE":
            result = NONE_ONENOTEUSERROLE
        case "OWNER":
            result = OWNER_ONENOTEUSERROLE
        case "CONTRIBUTOR":
            result = CONTRIBUTOR_ONENOTEUSERROLE
        case "READER":
            result = READER_ONENOTEUSERROLE
        default:
            return 0, errors.New("Unknown OnenoteUserRole value: " + v)
    }
    return &result, nil
}
func SerializeOnenoteUserRole(values []OnenoteUserRole) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
