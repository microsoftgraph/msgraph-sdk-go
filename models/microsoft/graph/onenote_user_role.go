package graph
import (
    "strings"
    "errors"
)
// 
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
    switch strings.ToUpper(v) {
        case "NONE":
            return NONE_ONENOTEUSERROLE, nil
        case "OWNER":
            return OWNER_ONENOTEUSERROLE, nil
        case "CONTRIBUTOR":
            return CONTRIBUTOR_ONENOTEUSERROLE, nil
        case "READER":
            return READER_ONENOTEUSERROLE, nil
    }
    return 0, errors.New("Unknown OnenoteUserRole value: " + v)
}
func SerializeOnenoteUserRole(values []OnenoteUserRole) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
