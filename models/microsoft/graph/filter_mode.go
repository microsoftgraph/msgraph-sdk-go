package graph
import (
    "strings"
    "errors"
)
// Provides operations to manage the identityContainer singleton.
type FilterMode int

const (
    INCLUDE_FILTERMODE FilterMode = iota
    EXCLUDE_FILTERMODE
)

func (i FilterMode) String() string {
    return []string{"INCLUDE", "EXCLUDE"}[i]
}
func ParseFilterMode(v string) (interface{}, error) {
    result := INCLUDE_FILTERMODE
    switch strings.ToUpper(v) {
        case "INCLUDE":
            result = INCLUDE_FILTERMODE
        case "EXCLUDE":
            result = EXCLUDE_FILTERMODE
        default:
            return 0, errors.New("Unknown FilterMode value: " + v)
    }
    return &result, nil
}
func SerializeFilterMode(values []FilterMode) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
